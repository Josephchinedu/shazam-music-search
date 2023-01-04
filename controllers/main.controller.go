package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"shazam_music_query/utils"
)

type MainController struct {
	gin.RouterGroup
}

type SearchSerializer struct {
	SongTitle string `json:"song_title" binding:"required"`
}

type ResponseSerializer struct {
	Tracks struct {
		Hits []struct {
			Track struct {
				Layout   string `json:"layout"`
				Type     string `json:"type"`
				Key      string `json:"key"`
				Title    string `json:"title"`
				Subtitle string `json:"subtitle"`
				Share    struct {
					Subject  string `json:"subject"`
					Text     string `json:"text"`
					Href     string `json:"href"`
					Image    string `json:"image"`
					Twitter  string `json:"twitter"`
					Google   string `json:"google"`
					Html     string `json:"html"`
					Avatar   string `json:"avatar"`
					Snapchat string `json:"snapchat"`
				} `json:"share"`
				Images struct {
					Background string `json:"background"`
					CoverArt   string `json:"coverart"`
					CoverArtHq string `json:"coverarthq"`
					JoeColor   string `json:"joecolor"`
				} `json:"images"`
				Hub struct {
					Type    string `json:"type"`
					Image   string `json:"image"`
					Actions []struct {
						Name string `json:"name"`
						Type string `json:"type"`
						ID   string `json:"id"`
					} `json:"actions"`
				} `json:"hub"`
			} `json:"track"`
		} `json:"hits"`
	} `json:"tracks"`

	Artists struct {
		Hits []struct {
			Artist struct {
				Avatar   string `json:"avatar"`
				Name     string `json:"name"`
				Verified bool   `json:"verified"`
				WebUel   string `json:"webUrl"`
				Adamid   string `json:"adamid"`
			} `json:"artist"`
		} `json:"hits"`
	} `json:"artists"`
}

// func getArtistName(track Track) string {
// 	// get the artist name
// 	artistName := track.Artist.Name

// 	return artistName
// }

func (ac *MainController) MusicSearchView(ctx *gin.Context) {
	var payload SearchSerializer

	// fmt.Println("MusicSearchView")
	fmt.Println(payload)

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	// get the song title from the payload
	songTitle := payload.SongTitle

	rapidResponse, err := utils.MusicSearch(songTitle)

	// print("Rapid Response: ", string(rapidResponse))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var responseObject ResponseSerializer
	json.Unmarshal(rapidResponse, &responseObject)

	// fmt.Println("Response Object: ", responseObject)

	// get the first track
	// track := responseObject.Tracks.Hits[0].Track
	songLayout := responseObject.Tracks.Hits
	// artistLayout := responseObject.Artists.Hits

	for i, v := range songLayout {
		fmt.Println("Index: ", i, "Value: ", v)

	}
	// fmt.Println("Track: ", track)

	// dataResponse := map[string]interface{}{
	// 	"layout":   track.Layout,
	// 	"type":     track.Type,
	// 	"key":      track.Key,
	// 	"title":    track.Title,
	// 	"subtitle": track.Subtitle,
	// }

	// return all the data
	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Song title received", "layout": layout, "type": tracType, "key": key, "title": title, "subtitle": subtitle, "shareSubject": shareSubject})

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Song title received"})

	// ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Song title received", "data": dataResponse})

}
