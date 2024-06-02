package fetchticks

type ObjectWriter interface {
	Write(data any) error
}
