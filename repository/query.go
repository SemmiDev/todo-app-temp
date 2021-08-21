package repository

const (
	SaveTodoCmd         = `insert into todos(id, task, starting_at, ends_at, duration, is_expired, done) values (?,?,?,?,?,?,?)`
	UpdateTodoCmd       = `update todos set task = ?, starting_at = ?, ends_at = ?, duration = ?, is_expired = ?, done = ? where id = ?`
	UpdateStatusTodoCmd = `update todos set done = ? where id = ?`
	DeleteTodoCmd       = "delete from todos where id = ?"
	FindByIdQuery       = `select id, task, starting_at, ends_at, duration, is_expired, done from todos where id = ?`
	FindAllQuery        = `select id, task, starting_at, ends_at, duration, is_expired, done from todos`
)
