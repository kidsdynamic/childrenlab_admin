package controller

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_admin/model"
	"github.com/pkg/errors"
)

const (
	SignedUserKey = "SignedUser"
	timeLayout    = "2006-01-02T15:04:05Z"
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

var SuperAdminToken string

func storeToken(db *sqlx.DB, accessToken model.AccessToken) error {
	var existToken model.AccessToken
	if err := db.Get(&existToken, "SELECT * FROM authentication_token WHERE email = ?", accessToken.Email); err != nil {
		if err != sql.ErrNoRows {
			return errors.Wrap(err, "Error on retrieve auth")
		}
	}
	existToken.LastUpdated = time.Now()

	if existToken.ID != 0 {
		existToken.Token = accessToken.Token
		existToken.AccessCount += 1
		if _, err := db.NamedExec("UPDATE authentication_token SET token = :token, access_count = :access_count, last_updated = :last_updated WHERE email = :email", &existToken); err != nil {
			return errors.Wrap(err, "Error on update token")
		}
	} else {
		if _, err := db.NamedExec("INSERT INTO authentication_token (email, token, access_count, last_updated) VALUES (:email, :token, :access_count, :last_updated)", &existToken); err != nil {
			return errors.Wrap(err, "Error on insert token")
		}
	}
	return nil

}

func randToken() string {
	b := make([]byte, 8)
	rand.Read(b)
	h := md5.New()
	io.WriteString(h, fmt.Sprintf("%x", b))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func GetNowTime() (time.Time, error) {
	now := time.Now()

	s := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	return time.Parse(timeLayout, s)

}

func SuperAdminAuth(c *gin.Context) {
	authToken := c.Request.Header.Get("x-auth-token")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	if authToken != SuperAdminToken {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}
	c.Next()
}

func AdminAuth(c *gin.Context) {
	authToken := c.Request.Header.Get("x-auth-token")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	db := NewDB()
	defer db.Close()

	var user User

	if err := db.Get(&user, "select u.* from user u join authentication_token a ON u.email = a.email JOIN role ON u.role_id = role.id "+
		"where a.token = ? and (authority = ? or authority = ?)", authToken, "ROLE_ADMIN", "ROLE_SUPER_ADMIN"); err != nil {
		fmt.Println(err)
		fmt.Println("test")
		c.JSON(http.StatusForbidden, gin.H{})
		c.Abort()
		return
	}

	c.Set(SignedUserKey, user)

	c.Next()

}

func GetMaxAndPage(c gin.Context) (int, int, error) {
	maxString := c.Query("max")
	pageString := c.Query("page")

	max, err := strconv.Atoi(maxString)
	page, err := strconv.Atoi(pageString)

	return max, page, err
}
