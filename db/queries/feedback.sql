-- name: CreateFeedback :one
insert into feedbacks (userid, feedback)
values ($1, $2)
returning *;