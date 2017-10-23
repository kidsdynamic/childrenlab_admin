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

	var userList []User
	if err := db.Select(&userList, "select user.id, email, first_name, last_name, profile, registration_id, android_registration_token, sign_up_country_code, date_created"+
		" from user join role ON role.id = user.role_id where role.authority = 'ROLE_USER' order by date_created desc"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retriving user list from admin"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on retriving user list",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, userList)
}
