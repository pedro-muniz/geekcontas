package registration

import (
	"testing"
	"time"

	pgsqlFactory "github.com/pedro-muniz/geekcontas/auth/container/dataservicefactory/daofactory/pgsqlfactory"
	userData "github.com/pedro-muniz/geekcontas/auth/dataservice/userdata"
	"github.com/pedro-muniz/geekcontas/auth/model"
)

func TestInsert_shouldPass(t *testing.T) {
	registrationUseCase := RegistrationUseCase{
		UserDataInterface: userData.UserData{
			Postgres: pgsqlFactory.NewPgsql(),
		},
	}
	baseProperties := model.BaseProperties{
		CreatedBy: "Testing suite",
		UpdatedBy: "Testing suite",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user := &model.User{
		BaseProperties: baseProperties,
		Name:           "Pedro",
		Email:          "pmuniz09@gmail.com",
		Password:       "18d50aac-7c39-11ea-bc55-0242ac130003",
	}

	registrationUseCase.Save(user)
}
