package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"memorabilia/entity"
	"memorabilia/model"
	"strings"

	"github.com/google/uuid"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, error)
	LoginUser(user *model.UserLogin) (*entity.User, error)
	GetUserById(id uuid.UUID) (*entity.User, error)
	UpdateUser(id uuid.UUID, userReq *model.UpdateUser) (*model.UpdateUser, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepository *UserRepository) CreateUser(user *entity.User) (*entity.User, error) {
	stmt := `INSERT INTO users (id, name, username, password, email, roleId, profilePicture, createdAt) 
	VALUES (?, ?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(stmt, user.ID, user.Name, user.Username, user.Password, user.Email, user.RoleID, user.ProfilePicture)
	if err != nil {
		tx.Rollback()
		return user, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) LoginUser(userReq *model.UserLogin) (*entity.User, error) {
	stmt := `SELECT * FROM users WHERE username = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, userReq.Username)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.RoleID, &user.ProfilePicture, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) GetUserById(id uuid.UUID) (*entity.User, error) {
	stmt := `SELECT * FROM users WHERE id = ?`
	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	user := &entity.User{}

	row := tx.QueryRow(stmt, id)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Email, &user.RoleID, &user.ProfilePicture, &user.CreatedAt)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return user, err
}

func (userRepository *UserRepository) UpdateUser(id uuid.UUID, userReq *model.UpdateUser) (*model.UpdateUser, error) {
	var column []string
	var values []interface{}

	if userReq.Name != "" {
		column = append(column, "name = ?")
		values = append(values, userReq.Name)
	}
	if userReq.Username != "" {
		column = append(column, "username = ?")
		values = append(values, userReq.Username)
	}
	if userReq.Password != "" {
		column = append(column, "password = ?")
		values = append(values, userReq.Password)
	}
	if userReq.Email != "" {
		column = append(column, "email = ?")
		values = append(values, userReq.Email)
	}
	if userReq.ProfilePicture != "" {
		column = append(column, "profilePicture = ?")
		values = append(values, userReq.ProfilePicture)
	}
	values = append(values, id)

	stmt := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(column, ", "))

	tx, err := userRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	result, err := tx.Exec(stmt, values...)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	fmt.Println(rowsAffected)
	if rowsAffected <= 0 {
		tx.Rollback()
		return nil, errors.New("no row updated")
	}

	err = tx.Commit()
	return userReq, err
}
