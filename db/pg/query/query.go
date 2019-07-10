// Package query attempts to implement an intuitive and "fluent" interface to
// querying data out of postgres (via go-pg's orm).
package query

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Scope func(*orm.Query) *orm.Query

type Query struct {
	q *orm.Query
}

func New(db *pg.DB, dst interface{}) *Query {
	return &Query{
		q: db.Model(dst),
	}
}

func (q *Query) Scope(fn Scope) *Query {
	fn(q.q)

	return q
}

func (q *Query) Limit(n int) *Query {
	q.q.Limit(n)

	return q
}

func (q *Query) Column(columns ...string) *Query {
	q.q.Column(columns...)

	return q
}

func (q *Query) Select(values ...interface{}) error {
	return q.q.Select(values...)
}

func (q *Query) Update(scan ...interface{}) (orm.Result, error) {
	return q.q.Update(scan...)
}
