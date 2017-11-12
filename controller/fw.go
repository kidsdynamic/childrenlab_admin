package controller

import (
	"net/http"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kidsdynamic/childrenlab_admin/config"
	"github.com/pkg/errors"
)

type FwFile struct {
	ID           int64     `db:"id",json:"id"`
	Version      string    `db:"version",json:"version"`
	FileAURL     string    `db:"file_a_url",json:"fileAUrl"`
	FileBURL     string    `db:"file_b_url",json:"fileBUrl"`
	UploadedDate time.Time `db:"uploaded_date",json:"uploadedDate"`
	Active       bool      `db:"active",json:"active"`
}

type updateFWRequest struct {
	ID     int64 `json:"id"`
	Active bool  `json:"active"`
}

func GetFwFileList(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	fwFiles, err := getFwList(db)

	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on retrieve FW list"))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something wrong when retrieve FW list from database",
			"error":   err,
		})
		return
	}

	for index, fw := range fwFiles {
		fwFiles[index].FileAURL = fmt.Sprintf("%s/%s", config.ConfigSettings.S3BasedURL, fw.FileAURL)
		fwFiles[index].FileBURL = fmt.Sprintf("%s/%s", config.ConfigSettings.S3BasedURL, fw.FileBURL)
	}

	c.JSON(http.StatusOK, fwFiles)

}

func UpdateActivation(c *gin.Context) {
	var request updateFWRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	db := NewDB()
	defer db.Close()

	if _, err := db.Exec("update fw_file set active = ? where id = ?", request.Active, request.ID); err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on update fw active"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	fwList, err := getFwList(db)
	if err != nil {
		fmt.Printf("%+v", errors.Wrap(err, "Error on update fw active"))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, fwList)
}

func getFwList(db *sqlx.DB) ([]FwFile, error) {
	var fwFiles []FwFile
	if err := db.Select(&fwFiles, "select * from fw_file order by id desc"); err != nil {
		return fwFiles, err
	}
	return fwFiles, nil
}
