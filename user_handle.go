package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Login     string `json:"login" db:"login"`
	Timestamp string `json:"timestamp" db:"timestamp"`
	Password  string `json:"password,omitempty" db:"password"`
}

func (u *User) ValidateUser() error {
	password := u.Password

	if err := db.Get(u, getPasswordUserSql, u.Login); err != nil {
		return err
	}

	ok := CheckPasswordHash(password, u.Password)
	if !ok {
		return fmt.Errorf("Senha inválida para o login %s", u.Login)
	}

	return nil
}

func (u *User) GenerateToken() (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["login"] = u.Login
	claims["name"] = u.Name
	claims["id"] = strconv.Itoa(u.ID)
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, err := t.SignedString([]byte(jwtSecret))
	if err != nil {
		return token, err
	}

	return token, nil
}

func (users *Users) GetAllUsers() error {
	if err := db.Select(&users.Users, getAllUsersSql); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}

func (user *User) GerUser() error {
	if err := db.Get(user, getUserSql, user.ID); err != nil {
		if err != sql.ErrNoRows {
			return err
		}
	}

	return nil
}

func (user *User) CreateNewUser() error {
	msgErro := "Falha ao gerar novo usuário"

	if ok := user.ValidateLogin(); !ok {
		return fmt.Errorf("%s: Login inválido ou indefinido", msgErro)
	}
	if ok := user.ValidateName(); !ok {
		return fmt.Errorf("%s: Nome inválido ou indefinido", msgErro)
	}
	if ok := user.ValidatePassword(); !ok {
		return fmt.Errorf("%s: Senha inválida ou indefinida", msgErro)
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	r, err := db.Exec(createUserSql, user.Name, user.Login, hash)
	if err != nil {
		return err
	}

	if ra, _ := r.RowsAffected(); ra == 0 {
		return fmt.Errorf("%s", msgErro)
	}

	return nil
}

func (user *User) ValidateName() bool {
	// Começa com letra maiscula seguido por ao menos 2 letras
	re := regexp.MustCompile(`^[A-Z][a-zA-Z]{2,}.*`)
	return re.MatchString(user.Name)
}

func (user *User) ValidateLogin() bool {
	grupos := []string{`[@!#+-]+`, `[0-9]+`, `[A-Z]+`}

	for _, g := range grupos {
		re := regexp.MustCompile(g)
		if !re.MatchString(user.Login) {
			return false
		}
	}

	return len(user.Login) > 4
}

// ValidatePassword deve ser maior que
func (user *User) ValidatePassword() bool {
	grupos := []string{`[@!#+-]+`, `[0-9]+`, `[A-Z]+`}

	for _, g := range grupos {
		re := regexp.MustCompile(g)
		if !re.MatchString(user.Password) {
			return false
		}
	}

	return len(user.Password) > 4
}
