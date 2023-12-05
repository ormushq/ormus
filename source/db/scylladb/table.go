package scylladb

import (
	"context"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
)

type TableInterface interface {
	Metadata() Metadata
	PrimaryKeyCmp() []qb.Cmp
	Name() string
	Get(columns ...string) (stmt string, names []string)
	GetQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx
	GetQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx
	Select(columns ...string) (stmt string, names []string)
	SelectQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx
	SelectQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx

	SelectAll() (stmt string, names []string)
	Insert() (stmt string, names []string)
	InsertQuery(session *gocqlx.Session) gocqlx.Queryx
	InsertQueryContext(ctx context.Context, session *gocqlx.Session) gocqlx.Queryx

	Update(columns ...string) (stmt string, names []string)
	UpdateQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx
	UpdateQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx

	Delete(columns ...string) (stmt string, names []string)
	DeleteQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx
	DeleteQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx
}

type Metadata struct {
	M *table.Metadata
}

type Table struct {
	T *table.Table
}

func (t *Table) Metadata() Metadata {
	gocqlxmetadata := t.T.Metadata()

	return Metadata{M: &gocqlxmetadata}
}

func (t *Table) PrimaryKeyCmp() []qb.Cmp {
	return t.T.PrimaryKeyCmp()
}

func (t *Table) Name() string {
	return t.T.Name()
}

func (t *Table) Get(columns ...string) (stmt string, names []string) {
	return t.T.Get(columns...)
}

func (t *Table) GetQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.GetQuery(*session, columns...))
}

func (t *Table) GetQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.GetQueryContext(ctx, *session, columns...))
}

func (t *Table) Select(columns ...string) (stmt string, names []string) {
	return t.T.Select(columns...)
}

func (t *Table) SelectQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.SelectQuery(*session, columns...))
}

func (t *Table) SelectQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.SelectQueryContext(ctx, *session, columns...))
}

func (t *Table) SelectAll() (stmt string, names []string) {
	return t.T.SelectAll()
}

func (t *Table) Insert() (stmt string, names []string) {
	return t.T.Insert()
}

func (t *Table) InsertQuery(session *gocqlx.Session) gocqlx.Queryx {
	return *(t.T.InsertQuery(*session))
}

func (t *Table) InsertQueryContext(ctx context.Context, session *gocqlx.Session) gocqlx.Queryx {
	return *(t.T.InsertQueryContext(ctx, *session))
}

func (t *Table) Update(columns ...string) (stmt string, names []string) {
	return t.T.Update(columns...)
}

func (t *Table) UpdateQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.UpdateQuery(*session, columns...))
}

func (t *Table) UpdateQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.UpdateQueryContext(ctx, *session, columns...))
}

func (t *Table) Delete(columns ...string) (stmt string, names []string) {
	return t.T.Delete(columns...)
}

func (t *Table) DeleteQuery(session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.DeleteQuery(*session, columns...))
}

func (t *Table) DeleteQueryContext(ctx context.Context, session *gocqlx.Session, columns ...string) gocqlx.Queryx {
	return *(t.T.DeleteQueryContext(ctx, *session, columns...))
}

func New(m table.Metadata) TableInterface {
	return &Table{T: table.New(m)}
}
