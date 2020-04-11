package pgsqlfactory

import (
	pgsqlDao "github.com/pedro-muniz/geekcontas/auth/dataservice/dao/pgsql"
)

func NewPgsql() *pgsqlDao.Connection {
	return new(pgsqlDao.Connection)
}
