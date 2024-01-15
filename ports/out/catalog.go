package out

type Catalog interface {
	NextRow() []string
	HasNext() bool
	Close()
}
