package Services

import (
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func UploadAvatar(c *gin.Context) {
	cld, _ := cloudinary.NewFromURL("")
	fileName := c.PostForm("name")

	// Add tags
	fileTags := c.PostForm("tags")
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "Failed to upload",
		})
	}

	result, err := cld.Upload.Upload(c, file, uploader.UploadParams{
		PublicID: fileName,
		// Split the tags by comma
		Tags: strings.Split(",", fileTags),
	})

	if err != nil {
		c.String(http.StatusConflict, "Upload to cloudinary failed")
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":   "Successfully uploaded the file",
		"secureURL": result.SecureURL,
		"publicURL": result.URL,
	})
}
