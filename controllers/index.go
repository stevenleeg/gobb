package controllers

import (
	"fmt"
	"github.com/stevenleeg/gobb/models"
	"github.com/stevenleeg/gobb/utils"
	"net/http"
)

func Index(w http.ResponseWriter, request *http.Request) {
	current_user := utils.GetCurrentUser(request)
	boards, err := models.GetBoardsUnread(current_user)

	if err != nil {
		fmt.Printf("[error] Could not get boards (%s)\n", err.Error())
	}

	user_count, _ := models.GetUserCount()
	latest_user, _ := models.GetLatestUser()
	total_posts, _ := models.GetPostCount()

	utils.RenderTemplate(w, request, "index.html", map[string]interface{}{
		"boards":       boards,
		"user_count":   user_count,
		"online_users": models.GetOnlineUsers(),
		"latest_user":  latest_user,
		"total_posts":  total_posts,
	}, map[string]interface{}{
		"IsUnread": func(join *models.JoinBoardView) bool {
			latest_post := join.Board.GetLatestPost()

			if current_user != nil && !current_user.LastUnreadAll.Time.Before(latest_post.Op.LatestReply) {
				return false
			}

			return !join.ViewedOn.Valid || join.ViewedOn.Time.Before(latest_post.Op.LatestReply)
		},
	})
}
