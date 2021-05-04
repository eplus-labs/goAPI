package app

import (
	"illuminate_crypto_api/controller/ping"
	"illuminate_crypto_api/controller/videos"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/postvideo", videos.Create)
	router.GET("/video/:video_category", videos.Get)
}
