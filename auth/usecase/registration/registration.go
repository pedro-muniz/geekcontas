package registration

import (
	ds "github.com/pedro-muniz/geekcontas/auth/dataservice/userdata"
	"github.com/pedro-muniz/geekcontas/auth/model"
	"github.com/pkg/errors"
)

type RegistrationUseCase struct {
	UserDataInterface ds.UserData //TODO: Use UserDataInterface as type
}

/*
func (ruc *RegistrationUseCase) isDuplicate(email string) (bool, error) {
	user, err := ruc.UserDataInterface.FindByEmail(email)
	if err != nil {
		return true, errors.Wrap(err, "")
	}

	return user != nil, nil
}
*/

func (ruc *RegistrationUseCase) insert(user *model.User) (*model.User, error) {
	err := user.ValidateNew()
	if err != nil {
		return nil, errors.Wrap(err, "User validation failed")
	}

	/*
		isDup, err := ruc.isDuplicate(user.Email)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}

		if isDup {
			return nil, errors.New("This email is already in use: " + user.Email)
		}
	*/

	resultUser, err := ruc.UserDataInterface.Insert(user)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return resultUser, nil
}

func (ruc *RegistrationUseCase) update(user *model.User) (*model.User, error) {
	err := user.ValidateModify()
	if err != nil {
		return nil, errors.Wrap(err, "User validation failed")
	}

	//_, err = ruc.UserDataInterface.Update(user)
	return user, nil
}

func (ruc *RegistrationUseCase) Save(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("Invalid user")
	}

	if user.BaseProperties.Id == 0 {
		return ruc.insert(user)
	}

	return ruc.update(user)
}

func (ruc *RegistrationUseCase) Delete(user *model.User) (*model.User, error) {
	if user == nil {
		return nil, errors.New("Invalid user")
	}

	user.IsDeleted = true
	return ruc.update(user)
}
