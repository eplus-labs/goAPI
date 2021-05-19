package ethereum

import (
	"fmt"
	"illuminate_crypto_api/datasources/mysql/ethereum_mappings_db"
	"time"
)

const (
	apiDateLayout       = "2006-01-02T15:04:05Z"
	apiDbLayout         = "2006-01-02 15:04:05"
	queryInsertCampaign = "INSERT into kickstart_campaigns(address, name, description, minimumContribution, dateTime) VALUES(?,?,?,?,?);"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDbLayout)
}

func Save(Campaign CampaignFromUser) CampaignFromUser {
	dbTime := GetNowDBFormat()
	stmt, err := ethereum_mappings_db.Client.Prepare(queryInsertCampaign)
	if err != nil {
		fmt.Println(err)
	}

	insertResult, saveErr := stmt.Exec(Campaign.CampaignAddress, Campaign.CampaignName, Campaign.CampaignDescription, Campaign.MinimumContribution, dbTime)
	if saveErr != nil {
		fmt.Println(saveErr)
	}

	_, err = insertResult.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	return Campaign
}

func Get(campaign *CampaignFromUser) []CampaignFromUser {

	rows, err := ethereum_mappings_db.Client.Query("SELECT address, minimumContribution, name, description FROM kickstart_campaigns WHERE address='" + campaign.CampaignAddress + "'")

	if err != nil {
		fmt.Println(err)
	}

	sliceOfCampaign := []CampaignFromUser{}
	for rows.Next() {
		err := rows.Scan(&campaign.CampaignAddress, &campaign.MinimumContribution, &campaign.CampaignName, &campaign.CampaignDescription)
		if err != nil {
			fmt.Println(err)
		}

		sliceOfCampaign = append(sliceOfCampaign, *campaign)
	}

	return sliceOfCampaign
}

func GetName(campaign *CampaignFromUser) []CampaignFromUser {

	rows, err := ethereum_mappings_db.Client.Query("SELECT address, minimumContribution, name, description FROM kickstart_campaigns WHERE name='" + campaign.CampaignName + "'")

	if err != nil {
		fmt.Println(err)
	}

	sliceOfCampaign := []CampaignFromUser{}
	for rows.Next() {
		err := rows.Scan(&campaign.CampaignAddress, &campaign.MinimumContribution, &campaign.CampaignName, &campaign.CampaignDescription)
		if err != nil {
			fmt.Println(err)
		}

		sliceOfCampaign = append(sliceOfCampaign, *campaign)
	}

	return sliceOfCampaign
}

func GetAllCampaigns() []CampaignFromUser {

	rows, err := ethereum_mappings_db.Client.Query("SELECT address, minimumContribution, name, description FROM kickstart_campaigns")

	if err != nil {
		fmt.Println(err)
	}

	campaign := CampaignFromUser{}

	sliceOfCampaign := []CampaignFromUser{}
	for rows.Next() {
		err := rows.Scan(&campaign.CampaignAddress, &campaign.MinimumContribution, &campaign.CampaignName, &campaign.CampaignDescription)
		if err != nil {
			fmt.Println(err)
		}

		sliceOfCampaign = append(sliceOfCampaign, campaign)
	}

	return sliceOfCampaign
}
