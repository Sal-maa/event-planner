package user

import (
	"database/sql"
	"fmt"

	_models "github.com/justjundana/event-planner/models"
)

type UserRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Register(user _models.User) (_models.User, error) {
	_, err := r.db.Exec("INSERT INTO users(name,email,password,address,occupation,phone) VALUES(?,?,?,?,?,?)", user.Name, user.Email, user.Password, user.Address, user.Occupation, user.Phone)
	return user, err
}

func (r *UserRepository) Login(email string) (_models.User, error) {
	row := r.db.QueryRow(`SELECT id, email, password FROM users WHERE email = ?;`, email)

	var user _models.User

	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Profile(id int) (_models.User, error) {
	fmt.Println(id)
	row := r.db.QueryRow(`SELECT id, name, email, password, address, occupation, phone FROM users WHERE id = ?;`, id)

	var user _models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Address, &user.Occupation, &user.Phone)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUser(user _models.User) error {
	_, err := r.db.Exec(`UPDATE users 
						SET name=?, email=?, password=?, address=?, occupation=?, phone=?
						WHERE id=?`, user.Name, user.Email, user.Password, user.Address, user.Occupation, user.Phone, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(user _models.User) error {
	_, err := r.db.Exec(`DELETE FROM users	WHERE id=?`, user.ID)
	return err
}
