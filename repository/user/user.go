package user

import (
	"database/sql"
	"fmt"
	"log"

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

func (r *UserRepository) CheckEmail(userChecked _models.User) (_models.User, error) {
	user := _models.User{}
	result, err := r.db.Query("SELECT email FROM users WHERE email=?", userChecked.Email)
	if err != nil {
		return user, err
	}
	defer result.Close()
	if isExist := result.Next(); isExist {
		return user, fmt.Errorf("user already exist")
	}

	if user.Email != userChecked.Email {
		// usernya belum ada
		return user, nil
	}

	return user, nil
}

func (r *UserRepository) Register(user _models.User) (_models.User, error) {
	_, err := r.db.Exec("INSERT INTO users(avatar, name, email, password, address, occupation, phone) VALUES(?, ?, ?, ?, ?, ?, ?)", "http://cdn.onlinewebfonts.com/svg/img_569204.png", user.Name, user.Email, user.Password, user.Address, user.Occupation, user.Phone)
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
	row := r.db.QueryRow(`SELECT id, name, email, password, address, occupation, phone FROM users WHERE id = ?;`, id)

	var user _models.User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Address, &user.Occupation, &user.Phone)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetUsers() ([]_models.User, error) {
	var users []_models.User
	rows, err := r.db.Query(`SELECT id, name, email, password, address, occupation, phone FROM users ORDER BY id ASC`)
	if err != nil {
		log.Fatalf("Error")
	}

	defer rows.Close()

	for rows.Next() {
		var user _models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Address, &user.Occupation, &user.Phone)
		if err != nil {
			log.Fatalf("Error")
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) UpdateUser(user _models.User) error {
	_, err := r.db.Exec(`UPDATE users 
						SET name = ?, email = ?, password = ?, address = ?, occupation =?, phone = ?, updated_at = CURRENT_TIMESTAMP
						WHERE id = ?`, user.Name, user.Email, user.Password, user.Address, user.Occupation, user.Phone, user.ID)
	return err
}

func (r *UserRepository) DeleteUser(user _models.User) error {
	_, err := r.db.Exec(`DELETE FROM users	WHERE id=?`, user.ID)
	return err
}
