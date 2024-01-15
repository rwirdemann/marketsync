package out

import "time"

type Catalog interface {
	NextRow() []string
	HasNext() bool
	Close()
	MarkSynced(row []string, t time.Time)
}
