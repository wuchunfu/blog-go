package model

type WebsiteConfig struct {
	Universal
	Config string `json:"config"`
}

type WebsiteConfigVO struct {
	WebsiteAvatar     string   `json:"websiteAvatar"`
	WebsiteName       string   `json:"websiteName"`
	WebsiteAuthor     string   `json:"websiteAuthor"`
	WebsiteIntro      string   `json:"websiteIntro"`
	WebsiteNotice     string   `json:"websiteNotice"`
	WebsiteCreateTime string   `json:"websiteCreateTime"`
	WebsiteRecordNo   string   `json:"websiteRecordNo"`
	SocialLoginList   []string `json:"socialLoginList"`
	SocialUrlList     []string `json:"socialUrlList"`
	Qq                string   `json:"qq"`
	Github            string   `json:"github"`
	Gitee             string   `json:"gitee"`
	TouristAvatar     string   `json:"touristAvatar"`
	UserAvatar        string   `json:"userAvatar"`
	IsCommentReview   int      `json:"isCommentReview"`
	IsMessageReview   int      `json:"isMessageReview"`
	IsEmailNotice     int      `json:"isEmailNotice"`
	IsReward          int      `json:"isReward"`
	WeiXinQRCode      string   `json:"weiXinQRCode"`
	AlipayQRCode      string   `json:"alipayQRCode"`
	ArticleCover      string   `json:"articleCover"`
	IsChatRoom        int      `json:"isChatRoom"`
	WebsocketUrl      string   `json:"websocketUrl"`
	IsMusicPlayer     int      `json:"isMusicPlayer"`
}
