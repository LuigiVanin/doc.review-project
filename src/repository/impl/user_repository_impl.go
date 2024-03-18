package impl

import (
	"doc-review/src/dto"
	"doc-review/src/entity"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryImpl struct {
	database *sqlx.DB
}

func NewUserRepositoryImpl(database *sqlx.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{database: database}

}

func (repository *UserRepositoryImpl) FindById(id string) (dto.GetUserDto, error) {
	query := "SELECT id, type, name, email, created_at, updated_at FROM users WHERE id = $1"
	var user dto.GetUserDto

	err := repository.database.Get(user, query, id)

	return user, err
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (entity.User, error) {

	query := "SELECT * updated_at FROM users WHERE email = $1"
	var user entity.User

	err := repository.database.Get(&user, query, email)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Create(user dto.CreateUserDto) (entity.User, error) {

	query := "INSERT INTO users (name, email, password, type) VALUES ($1, $2, $3, $4) RETURNING *"
	var responseUser entity.User

	err := repository.database.Get(&responseUser, query, user.Name, user.Email, user.Password, user.Type)

	if err != nil {
		return responseUser, err
	}

	return responseUser, nil
}
