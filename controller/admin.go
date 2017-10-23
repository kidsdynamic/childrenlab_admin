package controller

import (
	"net/http"

	"fmt"

	"time"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type AdminSignUpRequest struct {
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type Role struct {
	ID        int64  `db:"id",json:"-"`
	Authority string `db:"authority",json:"authority"`
}

type User struct {
	ID                       int64     `db:"id",json:"id"`
	Email                    string    `db:"email",json:"email"`
	Password                 string    `db:"password",json:"-"`
	FirstName                string    `db:"first_name",json:"firstName"`
	LastName                 string    `db:"last_name",json:"lastName"`
	LastUpdated              time.Time `db:"last_updated",json:"lastUpdate"`
	DateCreated              time.Time `db:"date_created",json:"dateCreated"`
	ZipCode                  string    `db:"zip_code",json:"zipCode"`
	PhoneNumber              string    `db:"phone_number",json:"phoneNumber"`
	Profile                  *string   `db:"profile",json:"profile"`
	Language                 string    `db:"language",json:"language"`
	RegistrationID           string    `db:"registration_id",json:"ios_registration_id"`
	AndroidRegistrationToken *string   `db:"android_registration_token",json:"android_registration_id"`
	Role                     Role      `db:"role",json:"-"`
	RoleID                   int64     `db:"role_id",json:"-"`
	ResetPasswordToken       *string   `db:"reset_password_token",json:"-"`
	SignUpIP                 *string   `db:"sign_up_ip",json:"-"`
	SignUpCountryCode        *string   `db:"sign_up_country_code",json:"country"`
}

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
