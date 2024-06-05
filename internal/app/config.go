package app

import (
	"fmt"
	"github.com/go-viper/mapstructure/v2"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/v2"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"os"
	"time"
)

type Config struct {
	PostgresDSN     string     `koanf:"postgres.dsn"`
	ParquetPath     string     `koanf:"parquet.file.path" `
	DateStarted     *time.Time `koanf:"date.start"`
	DateEnded       *time.Time `koanf:"date.end"`
	WindowDays      int        `koanf:"date.window_days"`
	TargetType      string     `koanf:"target"`
	MilvusServer    string     `koanf:"milvus.server"`
	LimitBatchSize  int        `koanf:"limits.batch_size"`
	LimitMaxRecords int        `koanf:"limits.max_records"`
}

func MustLoadConfig(log *zap.Logger, configYamlFile string) *Config {
	k := koanf.New(".")
	// Use the POSIX compliant pflag lib instead of Go's flag lib.
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	f.String("target", "xxx", "target upload type: 'postgres' or 'milvus'")
	f.Parse(os.Args[1:])

	// Defaults
	k.Load(confmap.Provider(map[string]interface{}{
		"postgres.dsn":       "postgresql://postgres:postgres@localhost:5432/test?sslmode=disable",
		"milvus.server":      "localhost:19530",
		"limits.batch_size":  "1",
		"limits.max_records": "2",
	}, "."), nil)

	if err := k.Load(file.Provider(configYamlFile), yaml.Parser()); err != nil {
		log.Fatal("error loading config file", zap.Error(err))
	}

	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Fatal("error loading config", zap.Error(err))
	}

	config := &Config{}
	if err := k.UnmarshalWithConf("", config,
		koanf.UnmarshalConf{
			Tag:       "koanf",
			FlatPaths: true,
			DecoderConfig: &mapstructure.DecoderConfig{
				DecodeHook: mapstructure.ComposeDecodeHookFunc(
					mapstructure.StringToTimeHookFunc("2006-01-02")),
				Metadata:         nil,
				Result:           config,
				WeaklyTypedInput: true,
			},
		}); err != nil {
		log.Fatal("error parsing config", zap.Error(err))
	}

	return config
}
