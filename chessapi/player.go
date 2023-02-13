package chessapi

import (
	"encoding/json"
	"log"
	"net/http"
)

type Player struct {
	Avatar     string `json:"avatar"`
	Country    string `json:"country"`
	Followers  int    `json:"followers"`
	Id         string `json:"@id"`
	IsStreamer string `json:"is_streamer"`
	Joined     int    `json:"joined"`
	LastOnline int    `json:"last_online"`
	League     string `json:"league"`
	Location   string `json:"location"`
	Name       string `json:"name"`
	PlayerId   int    `json:"player_id"`
	Url        string `json:"url"`
	Username   string `json:"username"`
	Status     string `json:"status"`
	Title      string `json:"title"`
	TwitchUrl  string `json:"twitch_url"`
	Verified   bool   `json:"verified"`
}

type DatedRating struct {
	Rating int `json:"rating"`
	Date   int `json:"date"`
	Rd     int `json:"rd"`
}

type BestRating struct {
	Rating  int    `json:"rating"`
	Date    int    `json:"date"`
	GameUrl string `json:"game"`
}

type Record struct {
	Wins   int `json:"win"`
	Losses int `json:"losses"`
	Draws  int `json:"draw"`
}

type Stats struct {
	Last DatedRating `json:"last"`
	Best BestRating  `json:"best"`
}

type PlayerStats struct {
	DailyStats    Stats `json:"chess_daily"`
	Daily960Stats Stats `json:"chess960_daily"`
	RapidStats    Stats `json:"chess_rapid"`
	BulletStats   Stats `json:"chess_bullet"`
	BlitzStats    Stats `json:"chess_blitz"`
	FideRating    int   `json:"fide"`
}

func (api ChessAPI) FindPlayerByUsername(username string) (*Player, error) {
	resp, err := http.Get(api.baseUrl + "/player/" + username)

	if err != nil {
		log.Printf("unable to retrieve player %s - %s", username, err)
		return nil, err
	}

	var player Player
	err = json.NewDecoder(resp.Body).Decode(&player)

	if err != nil {
		log.Printf("unable to decode payload for player %s - %s", username, err)
	}

	return &player, nil
}

func (api ChessAPI) FindPlayerStatsByUsername(username string) (*PlayerStats, error) {
	resp, err := http.Get(api.baseUrl + "/player/" + username + "/stats/")

	if err != nil {
		log.Printf("unable to retrieve stats for player %s - %s", username, err)
		return nil, err
	}

	var playerStats PlayerStats
	err = json.NewDecoder(resp.Body).Decode(&playerStats)

	if err != nil {
		log.Printf("unable to decode payload for stats for player %s - %s", username, err)
	}

	return &playerStats, nil
}
