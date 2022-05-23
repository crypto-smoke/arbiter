package driver

type Driver interface {
	New() Conn
}
type Conn interface {
	Close()
}
