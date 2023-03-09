package models

import (
	"database/sql"
	"errors"
	"html"
	"log"
	"strings"
	"time"
)

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        int       `form:"id" json:"id"`
	Name      string    `form:"name" json:"name"`
	Email     string    `form:"email" json:"email"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	UpdatedAt time.Time `form:"updated_at" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("Required Name")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		return nil
	}
}

func (u *User) FindAllUsers(db *sql.DB) (*[]User, error) {
	users := []User{}
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM user")
	if err != nil {
		log.Print(err)
		return &[]User{}, err
	}

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			users = append(users, *u)
		}
	}
	return &users, nil
}

func (u *User) FindUserByID(db *sql.DB, uid uint32) (*User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM user WHERE id=?", uid)

	if !rows.Next() {
		return &User{}, errors.New("Id not exist")
	}

	err = rows.Scan(&u.Id, &u.Name, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		log.Fatal(err.Error())
		return &User{}, err
	}
	return u, nil
}

func (u *User) SaveUser(db *sql.DB) (*User, error) {
	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec("INSERT INTO user(name, email, password, created_at, updated_at) "+
		"VALUES(?, ?, ?, ?, ?)", u.Name, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	id, _ := res.LastInsertId()
	u.Id = int(id)
	return u, nil
}

func (u *User) UpdateAUser(db *sql.DB, uid uint32) (*User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM user WHERE id=?", uid)
	if !rows.Next() {
		return &User{}, errors.New("Id not exist")
	}

	_, err = db.Exec("UPDATE user SET name=?, email = ?, updated_at=? WHERE id=?", u.Name, u.Email, time.Now(), uid)
	if err != nil {
		log.Print(err)
	}
	u.Id = int(uid)
	return u, nil
}

func (u *User) DeleteAUser(db *sql.DB, uid uint32) (int, error) {
	_, err := db.Exec("DELETE FROM user WHERE id=?", uid)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return u.Id, nil
}

func (u *User) FindUserByEmail(db *sql.DB, email string) (*User, error) {
	rows, err := db.Query("SELECT id, name, password, email, created_at, updated_at FROM user WHERE email=?", email)

	if !rows.Next() {
		return &User{}, errors.New("Email not exist")
	}

	err = rows.Scan(&u.Id, &u.Name, &u.Password, &u.Email, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		log.Fatal(err.Error())
		return &User{}, err
	}
	return u, nil
}
