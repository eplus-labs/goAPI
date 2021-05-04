// package services

// import (
// 	"encoding/json"
// 	"fmt"
// 	"illuminate_crypto_api/domain/videos_domain"
// 	"io/ioutil"
// 	"net/http"
// 	"time"
// )

// type VideoFromUser struct {
// 	Name     string
// 	Category string
// 	EmbedID  string
// 	ID       string
// 	DateTime string
// }

// type VideoToDB struct {
// 	Name      string
// 	Category  string
// 	DateTime  string
// 	EmbedID   string
// 	Thumbnail string
// 	Title     string
// }

// type YoutubeVideoResponse struct {
// 	Kind  string `json:"kind"`
// 	Etag  string `json:"etag"`
// 	Items []struct {
// 		Kind    string `json:"kind"`
// 		Etag    string `json:"etag"`
// 		ID      string `json:"id"`
// 		Snippet struct {
// 			PublishedAt time.Time `json:"publishedAt"`
// 			ChannelID   string    `json:"channelId"`
// 			Title       string    `json:"title"`
// 			Description string    `json:"description"`
// 			Thumbnails  struct {
// 				Default struct {
// 					URL    string `json:"url"`
// 					Width  int    `json:"width"`
// 					Height int    `json:"height"`
// 				} `json:"default"`
// 				Medium struct {
// 					URL    string `json:"url"`
// 					Width  int    `json:"width"`
// 					Height int    `json:"height"`
// 				} `json:"medium"`
// 				High struct {
// 					URL    string `json:"url"`
// 					Width  int    `json:"width"`
// 					Height int    `json:"height"`
// 				} `json:"high"`
// 			} `json:"thumbnails"`
// 			ChannelTitle         string   `json:"channelTitle"`
// 			Tags                 []string `json:"tags"`
// 			CategoryID           string   `json:"categoryId"`
// 			LiveBroadcastContent string   `json:"liveBroadcastContent"`
// 			DefaultLanguage      string   `json:"defaultLanguage"`
// 			Localized            struct {
// 				Title       string `json:"title"`
// 				Description string `json:"description"`
// 			} `json:"localized"`
// 			DefaultAudioLanguage string `json:"defaultAudioLanguage"`
// 		} `json:"snippet"`
// 	} `json:"items"`
// 	PageInfo struct {
// 		TotalResults   int `json:"totalResults"`
// 		ResultsPerPage int `json:"resultsPerPage"`
// 	} `json:"pageInfo"`
// }

// func CreateVideo(videoFromUser videos_domain.VideoFromUser) (videoToDomain *VideoToDB) {

// 	url := "https://www.googleapis.com/youtube/v3/videos"
// 	method := "GET"

// 	const KEY = "AIzaSyAya5REfQqq96eOMmTXk3yXUYuLqsVFYXE"

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, nil)

// 	q := req.URL.Query()
// 	q.Add("part", "snippet")
// 	q.Add("maxResults", "5")
// 	q.Add("key", KEY)
// 	q.Add("type", "video")
// 	q.Add("id", videoFromUser.EmbedID)
// 	req.URL.RawQuery = q.Encode()

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	res, err := client.Do(req)
// 	defer res.Body.Close()
// 	body, err := ioutil.ReadAll(res.Body)

// 	YoutubeVideoResponse1 := YoutubeVideoResponse{}
// 	videoToDB := VideoToDB{}

// 	json.Unmarshal(body, &YoutubeVideoResponse1)

// 	// retrurn to videoFromUser.EmbedID (ie) when testing complete
// 	videoToDB.EmbedID = videoFromUser.EmbedID
// 	videoToDB.Name = videoFromUser.Name
// 	videoToDB.Category = videoFromUser.Category
// 	videoToDB.Title = YoutubeVideoResponse1.Items[0].Snippet.Title
// 	videoToDB.Thumbnail = YoutubeVideoResponse1.Items[0].Snippet.Thumbnails.Medium.URL

// 	fmt.Println("video struct from video service: ", videoToDB)
// 	return &videoToDB
// }
