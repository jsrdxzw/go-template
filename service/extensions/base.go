package extensions

type Extension interface {
	Init() error
	Close() error
}
