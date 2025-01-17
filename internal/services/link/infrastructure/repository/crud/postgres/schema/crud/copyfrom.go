// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: copyfrom.go

package crud

import (
	"context"
)

// iteratorForCreateLinks implements pgx.CopyFromSource.
type iteratorForCreateLinks struct {
	rows                 []CreateLinksParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateLinks) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateLinks) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Url,
		r.rows[0].Hash,
		r.rows[0].Describe,
		r.rows[0].Json,
	}, nil
}

func (r iteratorForCreateLinks) Err() error {
	return nil
}

func (q *Queries) CreateLinks(ctx context.Context, arg []CreateLinksParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"link", "links"}, []string{"url", "hash", "describe", "json"}, &iteratorForCreateLinks{rows: arg})
}
