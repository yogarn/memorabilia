package repository

import (
	"database/sql"
	"memorabilia/entity"
)

type IPeopleRepository interface {
	CreatePeople(people *entity.People) (*entity.People, error)
}

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) IPeopleRepository {
	return &PeopleRepository{db}
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
