/*
Package scylladb provides an interface and implementation for a ScyllaDB session
wrapper to simplify database interactions using gocql and gocqlx libraries.
*/
package scylladb

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type SessionxInterface interface {
	ContextQuery(ctx context.Context, stmt string, names []string) gocqlx.Queryx
	Query(stmt string, names []string) gocqlx.Queryx
	ExecStmt(stmt string) error
	AwaitSchemaAgreement(ctx context.Context) error
	Close()
}

type Session struct {
	S *gocqlx.Session
}

func (s *Session) ContextQuery(ctx context.Context, stmt string, names []string) gocqlx.Queryx {
	return *(s.S.ContextQuery(ctx, stmt, names))
}

func (s *Session) Query(stmt string, names []string) gocqlx.Queryx {
	return *(s.S.Query(stmt, names))
}

func (s *Session) ExecStmt(stmt string) error {
	return s.S.ExecStmt(stmt)
}

func (s *Session) AwaitSchemaAgreement(ctx context.Context) error {
	return s.S.AwaitSchemaAgreement(ctx)
}

func (s *Session) Close() {
	s.S.Close()
}

func NewSession(session *gocql.Session) SessionxInterface {
	gocqlxSession := gocqlx.NewSession(session)

	return &Session{S: &gocqlxSession}
}

func WrapSession(session *gocql.Session, err error) (SessionxInterface, error) {
	gocqlxSession, wrapErr := gocqlx.WrapSession(session, err)

	return &Session{&gocqlxSession}, wrapErr
}
