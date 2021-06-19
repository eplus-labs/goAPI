package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	fmt.Println("about to start application")

	router.Use(cors.Default())

	// router.Run(":8080")
	router.RunTLS(":8080", "/Users/scarruthers/go/src/illuminate_crypto_api/app/example.crt", "/Users/scarruthers/go/src/illuminate_crypto_api/app/example.key")

}
