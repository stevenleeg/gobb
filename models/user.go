package models

import (
    "crypto/sha1"
    "crypto/rand"
    "encoding/base64"
    "encoding/binary"
    "strconv"
    "io"
    "time"
    "errors"
    "fmt"
)


type User struct {
    Id       int64      "db:id"
    Created  int64      "db:created"
    Username string     "db:username"
    Password string     "db:password"
    Salt     string     "db:salt"
    Sid      string     "db:sid"
}

func NewUser(username, password string) *User {
    var salt string
    binary.Read(rand.Reader, binary.LittleEndian, &salt)
    
    hasher := sha1.New()
    io.WriteString(hasher, password)
    io.WriteString(hasher, salt)
    password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))

    return &User {
        Created: time.Now().Unix(),
        Username: username,
        Password: password,
        Salt: salt,
    }
}

func AuthenticateUser(username, password string) (error, *User) {
    db := GetDbSession()
    user := &User{}
    err := db.SelectOne(user, "SELECT * FROM users WHERE username=?", username)
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
        fmt.Printf("[warn] Invalid auth (%s != %s)\n", password, user.Password)
        return errors.New("Invalid username/password"), nil
    }

    return nil, user
}

func (user *User) GenerateSid() {
    var random string
    binary.Read(rand.Reader, binary.LittleEndian, &random)

    hasher := sha1.New()
    io.WriteString(hasher, strconv.FormatInt(time.Now().UnixNano(), 10))
    io.WriteString(hasher, random)
    sid := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

    user.Sid = sid
}
