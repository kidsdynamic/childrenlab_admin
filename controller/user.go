package controller

import (
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetAllUser(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	searchEmail := c.Query("searchEmail")
	var searchQuery string
	if searchEmail != "" {
		searchQuery = fmt.Sprintf("and email like '%%%s%%'", searchEmail)
	}

	max, page, err := GetMaxAndPage(*c)
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on parse string to int"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on parse string to int",
			"error":   err,
		})
	}

	var userList []User
	query := fmt.Sprintf("select user.id, email, first_name, last_name, profile, registration_id, android_registration_token, sign_up_country_code, date_created"+
		" from user join role ON role.id = user.role_id where role.authority = 'ROLE_USER' %s order by date_created desc LIMIT ?,?", searchQuery)
	if err := db.Select(&userList, query, (max*page)-max, max); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retriving user list from admin"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving user list",
			"error":   err,
		})
		return
	}

	var totalNumber int
	if searchEmail != "" {
		searchQuery = fmt.Sprintf("where email like '%%%s%%'", searchEmail)
	}
	query = fmt.Sprintf("select count(*) from user %s", searchQuery)
	if err := db.Get(&totalNumber, query); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on count users"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on count users",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  userList,
		"total": totalNumber,
	})
}

func CheckAuth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
