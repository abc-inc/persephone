package relational

import (
	"context"

	"github.com/abc-inc/merovingian/db"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Conn struct {
	logger zerolog.Logger
	conn   sqlx.Conn
}

func NewConn(conn sqlx.Conn) *Conn {
	l := zerolog.New(zerolog.NewConsoleWriter())
	return &Conn{l, conn}
}

func (c Conn) Close() error {
	return c.conn.Close()
}

func (c Conn) Exec(r db.Request, m db.RecordExtractor) (db.Result, error) {
	c.logger.Info().
		Str("query", r.Query).
		Str("format", r.Format).
		Str("template", r.Template).
		Interface("params", r.Params).
		Msg("Executing query")

	res, err := c.conn.QueryxContext(context.Background(), r.Query, r.Params)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	recs := db.Result{}
	for res.Next() {
		valByName := map[string]interface{}{}
		err := res.MapScan(valByName)
		if err != nil {
			return nil, err
		}

		keys, _ := res.Columns()
		rec := m(keys, func(key string) (interface{}, bool) {
			v, ok := valByName[key]
			return v, ok
		})
		recs = append(recs, rec)
	}
	return recs, nil
}

func (c Conn) Metadata() (es []db.Entity, err error) {
	return nil, nil
}