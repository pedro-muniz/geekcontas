package userDataSql

func Insert() string {
	return `insert into geekcontas_user (id, name, email, password, created_at, updated_at, created_by, updated_by)
	values (nextval('sq_geekcontas_user'), $1, $2, $3, $4, $5, $6, $7) RETURNING id`
}
