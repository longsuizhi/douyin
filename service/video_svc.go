package service

import (
	"douyin/dao"
	"douyin/model"
	"github.com/gin-gonic/gin"
)

func FeedVideo(c *gin.Context, req model.FeedVideoReq) (*dao.FeedVideo, error) {
	res := &dao.FeedVideo{}
	/*var latestTime time.Time
	intTime, err := strconv.ParseInt(req.LatestTime, 10, 64)
	if err != nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	//无登录状态
	if req.Token == "" {

	} else {
		//登录状态

	}*/
	return res, nil
}
