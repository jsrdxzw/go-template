package extensions

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jmoiron/sqlx"
	"yuanzhoucehui.com/be/gen/entschema"
)

type EntConf struct {
	Driver       string `json:"driver"`
	URL          string `json:"url"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type EntClient interface {
	Close() error
}

type EntExt struct {
	EntConf
	drv    *dialect.Driver
	Client EntClient
	RowSQL *sqlx.DB
}

func NewEntExt(c EntConf) *EntExt {
	return &EntExt{EntConf: c}
}

func (e *EntExt) Init() error {
	db, err := sql.Open(e.EntConf.Driver, e.URL)
	if err != nil {
		return err
	}

	e.RowSQL = sqlx.NewDb(db, e.EntConf.Driver)
	db.SetMaxOpenConns(e.MaxOpenConns)
	db.SetMaxIdleConns(e.MaxIdleConns)
	drv := entsql.OpenDB(e.EntConf.Driver, db)
	e.Client = entschema.NewClient(entschema.Driver(drv))
	return nil
}

func (e *EntExt) Close() error {
	return e.Client.Close()
}
