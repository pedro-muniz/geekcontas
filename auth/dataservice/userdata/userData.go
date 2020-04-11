package userdata

import (
	"log"

	pgsql "github.com/pedro-muniz/geekcontas/auth/container/dataservicefactory/daofactory/pgsqlfactory"
	dao "github.com/pedro-muniz/geekcontas/auth/dataservice/dao"
	sql "github.com/pedro-muniz/geekcontas/auth/dataservice/userdata/sql"
	"github.com/pedro-muniz/geekcontas/auth/model"
)

type UserData struct {
	Postgres dao.PostgreSqlInterface
}

func (userData *UserData) Insert(user *model.User) (*model.User, error) {
	var query string = sql.Insert()
	var conn dao.PostgreSqlInterface = pgsql.NewPgsql()

	id, err := conn.Write(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.BaseProperties.CreatedAt,
		user.BaseProperties.UpdatedAt,
		user.BaseProperties.CreatedBy,
		user.BaseProperties.UpdatedBy,
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	user.Id = id
	return user, nil
}
