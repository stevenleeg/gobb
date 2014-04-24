package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

type View struct {
	Id     string    `db:"id"`
	Post   *Post     `db:"-"`
	PostId int64     `db:"post_id"`
	User   *User     `db:"-"`
	UserId int64     `db:"user_id"`
	Time   time.Time `db:"time"`
}

func AddView(user *User, post *Post) *View {
	db := GetDbSession()

	// Generate the hash of userid and post id
	h := md5.New()
	hash := fmt.Sprintf("%d_%d", user.Id, post.Id)
	h.Write([]byte(hash))
	hash = hex.EncodeToString(h.Sum(nil))

	var view *View
	obj, _ := db.Get(&View{}, hash)
	if obj == nil {
		view = &View{
			Id:     hash,
			Post:   post,
			PostId: post.Id,
			User:   user,
			UserId: user.Id,
			Time:   time.Now(),
		}

		db.Insert(view)
	} else {
		view = obj.(*View)
		view.User = user
		view.Post = post
		view.Time = time.Now()

		db.Update(view)
	}

	return view
}
