package controllers

import (
	"encoding/json"
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

type Track struct {
	Key        string
	Title      string
	Subtitle   string
	Plays      int
	ArtistName string
	WebUrl     string
	Artists    []struct {
		Adamid string
	} `json:"artists,omitempty"`
}

type Artist struct {
	Name   string
	WebUrl string
	Adamid string
}

type ResponseSerializer struct {
	Tracks struct {
		Hits []struct {
			Track Track
		}
	}

	Artists struct {
		Hits []struct {
			Artist Artist
		}
	}
}

type SongStreamSerializer struct {
	Id    string `json:"id"`
	Total int    `json:"total"`
	Type  string `json:"type"`
}

func (ac *MainController) MusicSearchView(ctx *gin.Context) {
	var payload SearchSerializer

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	songTitle := payload.SongTitle

	rapidResponse, err := utils.MusicSearch(songTitle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var responseObject ResponseSerializer
	json.Unmarshal(rapidResponse, &responseObject)

	songLayout := responseObject.Tracks.Hits
	artistLayout := responseObject.Artists.Hits

	for i, v := range songLayout {

		for _, artist := range artistLayout {
			if v.Track.Artists[0].Adamid == artist.Artist.Adamid {
				songLayout[i].Track.ArtistName = artist.Artist.Name
				songLayout[i].Track.WebUrl = artist.Artist.WebUrl
				songLayout[i].Track.Artists = nil

				// get total play
				rapidApiTotalStreamResponse, err := utils.GetTotalStream(v.Track.Key)

				if err != nil {
					songLayout[i].Track.Plays = 0
				}

				var totalStreamResponse SongStreamSerializer
				json.Unmarshal(rapidApiTotalStreamResponse, &totalStreamResponse)

				songLayout[i].Track.Plays = totalStreamResponse.Total
			}
		}

	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Song title received", "data": songLayout})

}
