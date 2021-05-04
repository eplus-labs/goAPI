package videos

import (
	"encoding/json"
	"fmt"
	videos "illuminate_crypto_api/domain/videos_domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var video videos.VideoFromUser

	if err := c.ShouldBindJSON(&video); err != nil {
		fmt.Println("Error from comtroller: ", err)
		return
	}

	videoStruct := videos.Save(video)

	videoMarshalled, _ := json.Marshal(videoStruct)

	fmt.Println("Struct from create: ", videoStruct)

	c.JSON(http.StatusCreated, string(videoMarshalled))

}

func Get(c *gin.Context) {

	var video videos.VideoFromUser

	video.Category = c.Param("video_category")

	videoStruct := videos.Get(&video)

	videoMarshalled, _ := json.Marshal(videoStruct)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.JSON(http.StatusOK, string(videoMarshalled))

}
