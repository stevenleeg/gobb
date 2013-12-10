package controllers

import (
	"fmt"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Index(w http.ResponseWriter, request *http.Request) {

	boards, err := models.GetBoards()

	if err != nil {
		fmt.Printf("[error] Could not get boards (%s)\n", err.Error())
	}

	user_count, _ := models.GetUserCount()
	latest_user, _ := models.GetLatestUser()
	total_posts, _ := models.GetPostCount()

	utils.RenderTemplate(w, request, "index.html", map[string]interface{}{
		"boards":      boards,
		"user_count":  user_count,
		"latest_user": latest_user,
		"total_posts": total_posts,
	}, nil)
}
