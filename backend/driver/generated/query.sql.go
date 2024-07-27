// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package db_driver

import (
	"context"
)

const findTodo = `-- name: FindTodo :many
SELECT
  id,
  content,
  completed
FROM
  task_canvas.todo
`

func (q *Queries) FindTodo(ctx context.Context) ([]TaskCanvasTodo, error) {
	rows, err := q.db.Query(ctx, findTodo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskCanvasTodo
	for rows.Next() {
		var i TaskCanvasTodo
		if err := rows.Scan(&i.ID, &i.Content, &i.Completed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
