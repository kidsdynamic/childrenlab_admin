package controller

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_admin/model"
)

func Login(c *gin.Context) {
	var adminLogin model.AdminLogin

	if err := c.Bind(&adminLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return

	}

	db := NewDB()
	defer db.Close()

	var admin model.User

	adminLogin.Password = EncryptPassword(adminLogin.Password)

	if err := db.Get(&admin, "SELECT id, email, first_name, last_name FROM user WHERE email = ? AND PASSWORD = ?", adminLogin.Name, adminLogin.Password); err != nil {
		if err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

	}

	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	t, err := GetNowTime()
	if err != nil {
		fmt.Printf("%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Store token failed",
		})
		return
	}
	accessToken := model.AccessToken{
		Email:       admin.Email,
		Token:       randToken(),
		LastUpdated: t,
	}

	err = storeToken(db, accessToken)

	if err != nil {
		fmt.Printf("%+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Store token failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"username":     accessToken.Email,
		"access_token": accessToken.Token,
	})

}

func EncryptPassword(password string) string {
	h := sha256.New()
	io.WriteString(h, password)
	fmt.Printf("\n%x\n", h.Sum(nil))

	return fmt.Sprintf("%x", h.Sum(nil))

}
