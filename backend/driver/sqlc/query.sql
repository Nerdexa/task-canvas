-- name: FindTodo :many
SELECT
  id,
  content,
  completed
FROM
  task_canvas.todo
;

-- name: InsertTodo :exec
INSERT INTO
task_canvas.todo (
  id,
  content,
  completed
)
VALUES (
  $1,
  $2,
  $3
);

-- name: UpdateTodo :exec
UPDATE task_canvas.todo
SET
  content = $2,
  completed = $3
WHERE
  id = $1
;