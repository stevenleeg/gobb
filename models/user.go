package models

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"time"
)

type User struct {
	Id       int64  `db:"id"`
	GroupId  int64  `db:"group_id"`
	Created  int64  `db:"created"`
	Username string `db:"username"`
	Password string `db:"password"`
	Avatar   string `db:"avatar"`
	Salt     string `db:"salt"`
}

func NewUser(username, password string) *User {
	var salt string
	binary.Read(rand.Reader, binary.LittleEndian, &salt)

	hasher := sha1.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, salt)
	password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return &User{
		Created:  time.Now().Unix(),
		Username: username,
		Password: password,
		Salt:     salt,
	}
}

func AuthenticateUser(username, password string) (error, *User) {
	db := GetDbSession()
	user := &User{}
	err := db.SelectOne(user, "SELECT * FROM users WHERE username=$1", username)
	if err != nil {
		fmt.Printf("[error] Cannot select user (%s)\n", err.Error())
		return err, nil
	}

	if user.Id == 0 {
		return errors.New("Invalid username/password"), nil
	}

	hasher := sha1.New()
	io.WriteString(hasher, password)
	io.WriteString(hasher, user.Salt)
	password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if password != user.Password {
		return errors.New("Invalid username/password"), nil
	}

	return nil, user
}

func GetUserCount() (int64, error) {
    db := GetDbSession()

    count, err := db.SelectInt("SELECT COUNT(*) FROM users")
    if err != nil {
        fmt.Printf("[error] Error selecting user count (%s)\n", err.Error())
        return 0, errors.New("Database error: " + err.Error())
    }

    return count, nil
}

func GetLatestUser() (*User, error) {
    db := GetDbSession()

    user := &User{}
    err := db.SelectOne(user, "SELECT * FROM users ORDER BY created DESC LIMIT 1")
    
    if err != nil {
        fmt.Printf("[error] Error selecting latest user (%s)\n", err.Error())
        return nil, errors.New("Database error: " + err.Error())
    }

    if user.Username == "" {
        return nil, nil
    }

    return user, nil
}

func (user *User) IsAdmin() bool {
	if user.GroupId == 1 {
		return true
	}

	return false
}

func (user *User) CanModerate() bool {
    if user.GroupId > 0 {
        return true
    }

    return false
}
