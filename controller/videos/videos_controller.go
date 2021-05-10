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

	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://3.93.43.122:3000/")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PUT")

	c.JSON(http.StatusCreated, string(videoMarshalled))

}

func Get(c *gin.Context) {

	var video videos.VideoFromUser

	video.Category = c.Param("video_category")

	videoStruct := videos.Get(&video)

	sliceOfVideoStruct := videoStruct[len(videoStruct)-5:]

	fmt.Println("slice of video struct: ", sliceOfVideoStruct)

	videoMarshalled, _ := json.Marshal(sliceOfVideoStruct)

	// videoMarshalled, _ := json.Marshal(videoStruct)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.JSON(http.StatusOK, string(videoMarshalled))

	// fmt.Println("video marshalled: ", string(videoMarshalled))
	fmt.Println("video marshalled: ", videoStruct)
	fmt.Println("length of video marshalled: ", len(videoStruct))
	// sliceOfVideoStruct := videoStruct[len(videoStruct)-5:]
	fmt.Println("slice of video struct: ", sliceOfVideoStruct)
}
