package dao

type Author struct {
	ID            int64  `json:"id"`             //用户id
	Name          string `json:"name"`           //用户名称
	FollowCount   int64  `json:"follow_count"`   //关注总数
	FollowerCount int64  `json:"follower_count"` //粉丝总数
	IsFollow      bool   `json:"is_follow"`      // 是否关注 true 已关注	false 未关注
}

type VideoLiset struct {
	ID            int64    `json:"id"`             //视频唯一标识
	PlayUrl       string   `json:"play_url"`       //视频播放地址
	CoverUrl      string   `json:"cover_url"`      //视频封页地址
	FavoriteCount int64    `json:"favorite_count"` //视频点赞数
	CommentCount  int64    `json:"comment_count"`  //视频评论数
	IsFavorite    bool     `json:"is_favorite"`    //是否点赞	true 已点赞	false 未点赞
	Title         string   `json:"title"`          //视频标题
	List          []Author `json:"list"`
}

type FeedVideo struct {
	NextTime int64        `json:"next_time"` //本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	List     []VideoLiset `json:"list"`
}

func GetFeedVideoList(lastestTime string, token string) (FeedVideo, error) {
	res := FeedVideo{}
	err := SvDB.Model(FeedVideo{}).Find(&res).Error
	return res, err
}
