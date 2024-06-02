package common

import (
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// Closer contains all service closers to handle 'em gently on signal received from OS
type Closer struct {
	sync.Mutex
	once       sync.Once
	closers    []io.Closer
	closeFuncs []func()
	sigCh      chan os.Signal
	sigs       []os.Signal
	timeout    time.Duration
	done       chan struct{}
	logger     *zap.Logger
}

type Option func(*Closer)

// New ...
func New(
	timeout time.Duration,
	logger *zap.Logger,
	opts ...Option,
) *Closer {
	c := Closer{
		sigs:    []os.Signal{syscall.SIGINT, syscall.SIGTERM},
		sigCh:   make(chan os.Signal, 1),
		timeout: timeout,
		done:    make(chan struct{}, 1),
		logger:  logger.Named("closer"),
	}

	for _, opt := range opts {
		opt(&c)
	}

	go func() {
		signal.Notify(c.sigCh, c.sigs...)

		sig := <-c.sigCh
		c.logger.Info("received syscall signal", zap.String("signal", sig.String()))
		c.drop()
	}()

	return &c
}

// AddCloser any io.Closer
func (c *Closer) AddCloser(cl io.Closer) {
	c.Lock()
	defer c.Unlock()
	c.closers = append(c.closers, cl)
}

// AddCloseFunc any func() that will be run on exit
func (c *Closer) AddCloseFunc(f func()) {
	c.Lock()
	defer c.Unlock()
	c.closeFuncs = append(c.closeFuncs, f)
}

// Close jobs forcibly, bypassing system signals
func (c *Closer) Close() {
	c.drop()
}

// Wait for closing all the Jobs
func (c *Closer) Wait() {
	<-c.done
}

func (c *Closer) drop() {
	c.once.Do(func() {
		c.Lock()
		defer c.Unlock()

		// prepare sync mechanism
		wg := sync.WaitGroup{}
		chGracefulJobClose := make(chan struct{})
		chTimeoutJobClose := time.After(c.timeout)

		for _, cl := range c.closers {
			wg.Add(1)
			go func(cl io.Closer) {
				defer wg.Done()
				err := cl.Close()
				if err != nil {
					c.logger.Error("failed to close", zap.Error(err))
				}
			}(cl)
		}

		for _, cf := range c.closeFuncs {
			wg.Add(1)
			go func(cf func()) {
				defer wg.Done()
				cf()
			}(cf)
		}

		// run wait after all the closers startup, preventing empty wg
		go func() {
			wg.Wait()
			close(chGracefulJobClose)
		}()

		c.logger.Info(
			"waiting before terminate or end up earlier if funcs ready",
			zap.Float64("timeout seconds", c.timeout.Seconds()),
			zap.Int("jobs to close amount", len(c.closers)+len(c.closeFuncs)),
		)

		select {
		case <-chGracefulJobClose:
			c.logger.Info("jobs are gracefully closed")
			c.done <- struct{}{}
		case <-chTimeoutJobClose:
			c.logger.Info("jobs are not closed due to timeout")
			c.done <- struct{}{}
		}
	})
}
