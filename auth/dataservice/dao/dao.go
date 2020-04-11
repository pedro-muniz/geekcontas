package dao

type PostgreSqlInterface interface {
	Read(query string, params ...interface{}) ([]map[string]interface{}, error)
	Delete(query string, params ...interface{}) error
	Write(query string, params ...interface{}) (id int, err error)
}
