package db

import (
	"io"
)

type Connector interface {
	Exec(request Request, extr RecordExtractor) (Result, error)
	io.Closer
}
