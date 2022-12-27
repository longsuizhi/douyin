package model

type FeedVideoReq struct {
	LatestTime string `json:"latest_time" schema:"latest_time"`
	Token      string `json:"token" schema:"token"`
}
