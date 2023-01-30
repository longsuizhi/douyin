package define

// 每次最多返回的视频流数量
const (
	MaxVideoNum = 30
)

// 用户名和密码长度
const (
	MaxUsernameLength = 50 //最大用户名长度
	MinUsernameLength = 2  //最小用户名长度
	MaxPasswordLength = 20 //最大密码长度
	MinPasswordLength = 6  //最小密码长度
)

// 用户信息
const (
	//UserFollowNum   = "user_follow_num"   //用户关注数量
	//UserFollowerNum = "user_follower_num" //用户粉丝数量
	UserFollow      = "user_follow:"       //用户关注列表
	UserFollower    = "user_follower:"     //用户粉丝列表
)
