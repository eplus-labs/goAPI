package videos_domain

import (
	"encoding/json"
	"fmt"
	"illuminate_crypto_api/datasources/mysql/videos_mappings_db"
	"io/ioutil"
	"net/http"

	"time"
)

type YoutubeVideoResponse struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind    string `json:"kind"`
		Etag    string `json:"etag"`
		ID      string `json:"id"`
		Snippet struct {
			PublishedAt time.Time `json:"publishedAt"`
			ChannelID   string    `json:"channelId"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Thumbnails  struct {
				Default struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"default"`
				Medium struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"medium"`
				High struct {
					URL    string `json:"url"`
					Width  int    `json:"width"`
					Height int    `json:"height"`
				} `json:"high"`
			} `json:"thumbnails"`
			ChannelTitle         string   `json:"channelTitle"`
			Tags                 []string `json:"tags"`
			CategoryID           string   `json:"categoryId"`
			LiveBroadcastContent string   `json:"liveBroadcastContent"`
			DefaultLanguage      string   `json:"defaultLanguage"`
			Localized            struct {
				Title       string `json:"title"`
				Description string `json:"description"`
			} `json:"localized"`
			DefaultAudioLanguage string `json:"defaultAudioLanguage"`
		} `json:"snippet"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

const (
	apiDateLayout    = "2006-01-02T15:04:05Z"
	apiDbLayout      = "2006-01-02 15:04:05"
	queryInsertVideo = "INSERT into videos(Name, DateTime, Category, EmbedID, Thumbnail, Title) VALUES(?,?,?,?,?,?);"
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

func Save(video VideoFromUser) VideoToDB {
	dbTime := GetNowDBFormat()
	stmt, err := videos_mappings_db.Client.Prepare(queryInsertVideo)
	if err != nil {
		fmt.Println(err)
	}

	var videoToDB VideoToDB

	videoToDB = *CreateVideo(video)

	fmt.Println("Video to DB variable in Save function: ", videoToDB)

	insertResult, saveErr := stmt.Exec(videoToDB.Name, dbTime, videoToDB.Category, videoToDB.EmbedID, videoToDB.Thumbnail, videoToDB.Title)
	if saveErr != nil {
		fmt.Println(saveErr)
	}

	_, err = insertResult.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	return videoToDB

}

func Get(video *VideoFromUser) []VideoFromUser {

	rows, err := videos_mappings_db.Client.Query("SELECT ID, Name, DateTime, Category, EmbedID, Title, Thumbnail FROM videos WHERE Category='" + video.Category + "'")

	if err != nil {
		fmt.Println(err)
	}

	sliceOfVideos := []VideoFromUser{}
	for rows.Next() {
		err := rows.Scan(&video.ID, &video.Name, &video.DateTime, &video.Category, &video.EmbedID, &video.Title, &video.Thumbnail)
		if err != nil {
			fmt.Println(err)
		}

		sliceOfVideos = append(sliceOfVideos, *video)
	}

	return sliceOfVideos
}

func CreateVideo(videoFromUser VideoFromUser) *VideoToDB {

	url := "https://www.googleapis.com/youtube/v3/videos"
	method := "GET"

	const KEY = "AIzaSyAya5REfQqq96eOMmTXk3yXUYuLqsVFYXE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	q := req.URL.Query()
	q.Add("part", "snippet")
	q.Add("maxResults", "5")
	q.Add("key", KEY)
	q.Add("type", "video")
	q.Add("id", videoFromUser.EmbedID)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	YoutubeVideoResponse1 := YoutubeVideoResponse{}
	videoToDB := VideoToDB{}

	json.Unmarshal(body, &YoutubeVideoResponse1)

	fmt.Println("Within CreateVideo - YoutubeVideoResponse1: ", YoutubeVideoResponse1)

	// retrurn to videoFromUser.EmbedID (ie) when testing complete
	videoToDB.EmbedID = videoFromUser.EmbedID
	videoToDB.Name = videoFromUser.Name
	videoToDB.Category = videoFromUser.Category
	videoToDB.Title = YoutubeVideoResponse1.Items[0].Snippet.Title
	videoToDB.Thumbnail = YoutubeVideoResponse1.Items[0].Snippet.Thumbnails.Medium.URL

	return &videoToDB
}
