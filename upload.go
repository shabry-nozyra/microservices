package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
			return
		}

		filename := filepath.Base(file.Filename)

		if err := c.SaveUploadedFile(file, "file/" + filename); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Unable to save the file",
			})
			return
		}

		// File saved successfully. Return proper result
		c.JSON(http.StatusOK,  gin.H{
			"message": fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email),
		})

		//if err := c.SaveUploadedFile(file, "file/"+ filename); err != nil {
		//	c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		//	return
		//}
		//
		//c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email))

	})
	router.Run(":8085")
}