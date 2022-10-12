package user

import (
	"go_rest_api/authentication"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type RepositoryPg struct {
	DBConnection *sqlx.DB
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (r RepositoryPg) UserRegistration(user User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	_, err = r.DBConnection.NamedExec(`
		INSERT INTO users (
			email, password, first_name, last_name, id_card_file,
			id_card_number, id_card_laser_number, date_of_birth
		) 
		VALUES (
			:email, :password, :first_name, :last_name, :id_card_file,
			:id_card_number, :id_card_laser_number, :date_of_birth
		)`, user)
	return err
}

func (r RepositoryPg) UserLogin(loginInfo LoginInfo) (LoginResponse, error) {
	user := User{}
	loginResponse := LoginResponse{}
	err := r.DBConnection.Get(&user, "SELECT * FROM users WHERE email=$1 LIMIT 1", loginInfo.Email)
	if err != nil {
		return loginResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
	if err != nil {
		return loginResponse, err
	}
	jwtService := authentication.JWTAuthService()
	loginResponse.AccessToken = jwtService.GenerateToken(loginInfo.Email, true)
	return loginResponse, nil
}
