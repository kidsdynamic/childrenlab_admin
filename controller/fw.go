package controller

import (
	"net/http"

	"time"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kidsdynamic/childrenlab_admin/config"
	"github.com/pkg/errors"
)

type FwFile struct {
	ID           int64     `db:"id",json:"id"`
	Version      string    `db:"version",json:"version"`
	FileAURL     string    `db:"file_a_url",json:"fileAUrl"`
	FileBURL     string    `db:"file_b_url",json:"fileBUrl"`
	UploadedDate time.Time `db:"uploaded_date",json:"uploadedDate"`
}

func GetFwFileList(c *gin.Context) {
	db := NewDB()
	defer db.Close()

	var fwFiles []FwFile
	if err := db.Select(&fwFiles, "select * from fw_file order by id desc"); err != nil {
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
