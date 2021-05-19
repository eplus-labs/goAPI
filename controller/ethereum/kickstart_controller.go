package ethereum

import (
	"encoding/json"
	"fmt"
	"net/http"

	campaigns "illuminate_crypto_api/domain/ethereum_domain"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var campaign campaigns.CampaignFromUser

	if err := c.ShouldBindJSON(&campaign); err != nil {
		fmt.Println("Error from comtroller: ", err)
		return
	}

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD, OPTIONS, PUT, DELETE")

	campaignStruct := campaigns.Save(campaign)

	campaignMarshalled, _ := json.Marshal(campaignStruct)

	c.JSON(http.StatusCreated, string(campaignMarshalled))

	fmt.Println("C from within GIN post: ", c)

}

func Get(c *gin.Context) {

	var campaign campaigns.CampaignFromUser

	campaign.CampaignAddress = c.Param("campaignAddress")

	videoStruct := campaigns.Get(&campaign)

	fmt.Println("video struct: ", videoStruct)

	campaignMarshalled, _ := json.Marshal(videoStruct)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.JSON(http.StatusOK, string(campaignMarshalled))

	// fmt.Println("campaign marshalled: ", string(campaignMarshalled))

}

func GetName(c *gin.Context) {

	var campaign campaigns.CampaignFromUser

	campaign.CampaignName = c.Param("campaignName")

	campaignStruct := campaigns.GetName(&campaign)

	fmt.Println("campaign struct: ", campaignStruct)

	campaignMarshalled, _ := json.Marshal(campaignStruct)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.JSON(http.StatusOK, string(campaignMarshalled))

	// fmt.Println("campaign marshalled: ", string(campaignMarshalled))

}

func GetAllCampaigns(c *gin.Context) {

	campaignStruct := campaigns.GetAllCampaigns()

	fmt.Println("campaign struct: ", campaignStruct)

	campaignMarshalled, _ := json.Marshal(campaignStruct)

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	c.JSON(http.StatusOK, string(campaignMarshalled))

	// fmt.Println("campaign marshalled: ", string(campaignMarshalled))

}
