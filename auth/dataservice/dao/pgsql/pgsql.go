package pgsql

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"time"

	_ "github.com/lib/pq"
	config "github.com/pedro-muniz/geekcontas/auth/config"
)

var dbConf *config.AppConfig = config.NewAppConfig()
var (
	host     = dbConf.PostgreSql.Host
	port     = dbConf.PostgreSql.Port
	user     = dbConf.PostgreSql.User
	password = dbConf.PostgreSql.Password
	dbname   = dbConf.PostgreSql.DbName
)

type Connection struct {
	Conn *sql.DB
}

type DbData struct {
	valid bool
	value interface{}
}

func (scanner *DbData) getBytes(src interface{}) []byte {
	if a, ok := src.([]uint8); ok {
		return a
	}
	return nil
}

func (scanner *DbData) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		if value, ok := src.(int64); ok {
			scanner.value = int(value)
			scanner.valid = true
		}

	case float64:
		if value, ok := src.(float64); ok {
			scanner.value = value
			scanner.valid = true
		}
	case bool:
		if value, ok := src.(bool); ok {
			scanner.value = value
			scanner.valid = true
		}
	case string:
		scanner.value = src.(string)
		scanner.valid = true
	case time.Time:
		if value, ok := src.(time.Time); ok {
			scanner.value = value
			scanner.valid = true
		}

	case []byte:
		scanner.value = string(src.([]byte))
		scanner.valid = true

	case nil:
		scanner.value = nil
		scanner.valid = true
	default:
		log.Println("Type not implemented ")
		log.Println(reflect.TypeOf(src))
		log.Println(src)
	}

	return nil
}

func getConnection() (conn *Connection, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	this := Connection{
		Conn: db,
	}

	return &this, nil
}

func (conn *Connection) Read(query string, params ...interface{}) ([]map[string]interface{}, error) {
	db, err := getConnection()
	if err != nil {
		return nil, err
	}
	defer db.Conn.Close()

	dbRows, queryErr := db.Conn.Query(query, params...)
	defer dbRows.Close()
	if queryErr != nil {
		log.Println(err)
		return nil, err
	}

	//Para não fazer outro sql para buscar a quantidade de linhas
	var rows []map[string]interface{}

	columns, _ := dbRows.Columns()
	for dbRows.Next() {
		row := make([]interface{}, len(columns))

		for idx := range columns {
			row[idx] = new(DbData)
		}

		err = dbRows.Scan(row...)
		if err != nil {
			return nil, err
		}

		mapRow := make(map[string]interface{})
		for idx, column := range columns {
			var scanner = row[idx].(*DbData)
			mapRow[column] = scanner.value
		}
		rows = append(rows, mapRow)
	}

	return rows, nil
}

func (conn *Connection) Write(query string, params ...interface{}) (id int, err error) {
	db, err := getConnection()
	if err != nil {
		return 0, err
	}
	defer db.Conn.Close()

	writeErr := db.Conn.QueryRow(query, params...).Scan(&id)
	if writeErr != nil && writeErr != sql.ErrNoRows {
		return 0, writeErr
	}

	return id, err
}

func (conn *Connection) Delete(query string, params ...interface{}) error {
	db, err := getConnection()
	if err != nil {
		return err
	}
	defer db.Conn.Close()

	_, err = db.Conn.Exec(query, params...)
	if err != nil {
		return err
	}

	return err
}
