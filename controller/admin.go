package controller

import (
	"net/http"

	"fmt"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Only for internal use
func CreateAdminUser(c *gin.Context) {
	var request AdminSignUpRequest

	if c.BindJSON(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := NewDB()
	defer db.Close()
	var admin User

	request.Password = EncryptPassword(request.Password)
	if err := db.Get(&admin, "select * from user a join role r on u.role_id = r.id WHERE email = ? AND password = ? AND authority = ?", request.Name, request.Password, "ROLE_ADMIN"); err != nil {
		if err != sql.ErrNoRows {
			fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve getting admin user"))
			c.JSON(http.StatusInternalServerError, err)
			return
		}

	}

	if admin.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{
			"message": "The admin name exists",
		})
		return
	}

	var role Role
	if err := db.Get(&role, "select * from role where authority = ?", "ROLE_ADMIN"); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on getting admin role"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on getting admin role",
			"error":   err,
		})
		return
	}

	adminUser := User{
		Email:     request.Name,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Role:      role,
	}

	if _, err := db.NamedExec("INSERT INTO user (email, password, first_name, last_name, role_id) VALUES (:email, :password, :first_name, :last_name, :role_id)", &adminUser); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on creating admin user"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error on creating admin user",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
