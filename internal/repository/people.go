package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"memorabilia/entity"
	"memorabilia/model"
	"strings"
)

type IPeopleRepository interface {
	GetPeople(id string) (*entity.People, error)
	CreatePeople(people *entity.People) (*entity.People, error)
	UpdatePeople(id string, people *model.UpdatePeople) (*model.UpdatePeople, error)
	DeletePeople(id string) error
}

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) IPeopleRepository {
	return &PeopleRepository{db}
}

func (peopleRepository *PeopleRepository) GetPeople(id string) (*entity.People, error) {
	people := &entity.People{}
	stmt := `SELECT * FROM peoples WHERE id = ?`

	tx, err := peopleRepository.db.Begin()
	if err != nil {
		return nil, err
	}

	result := tx.QueryRow(stmt, id)
	err = result.Scan(&people.ID, &people.UserID, &people.Name, &people.Description, &people.Relation, &people.Picture, &people.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return people, nil
}

func (peopleRepository *PeopleRepository) CreatePeople(people *entity.People) (*entity.People, error) {
	stmt := `INSERT INTO peoples(id, userId, name, description, relation, picture, createdAt)
	VALUES(?, ?, ?, ?, ?, ?, ?)`

	tx, err := peopleRepository.db.Begin()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(stmt, people.ID, people.UserID, people.Name, people.Description, people.Relation, people.Picture, people.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	return people, err
}

func (peopleRepository *PeopleRepository) UpdatePeople(id string, people *model.UpdatePeople) (*model.UpdatePeople, error) {
	var column []string
	var values []interface{}

	if people.Name != "" {
		column = append(column, "name = ?")
		values = append(values, people.Name)
	}
	if people.Description != "" {
		column = append(column, "description = ?")
		values = append(values, people.Description)
	}
	if people.Relation != "" {
		column = append(column, "relation = ?")
		values = append(values, people.Relation)
	}

	values = append(values, id)

	stmt := fmt.Sprintf("UPDATE peoples SET %s WHERE id = ?", strings.Join(column, ", "))
	tx, err := peopleRepository.db.Begin()
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

	if rowsAffected <= 0 {
		tx.Rollback()
		return nil, errors.New("no row updated")
	}

	err = tx.Commit()
	return people, err
}

func (peopleRepository *PeopleRepository) DeletePeople(id string) error {
	stmt := `DELETE FROM peoples WHERE id = ?`
	tx, err := peopleRepository.db.Begin()
	if err != nil {
		return err
	}

	result, err := tx.Exec(stmt, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected <= 0 {
		tx.Rollback()
		return errors.New("no row updated")
	}

	err = tx.Commit()
	return err
}
