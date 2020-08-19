package xmlydownloader

//AudioInfoList AudioInfo的指针数组
type AudioInfoList []*AudioInfo

//AudioInfo Mobile接口的音频信息
type AudioInfo struct {
	TrackID         int    `json:"trackId"`
	TrackRecordID   int    `json:"trackRecordId"`
	UID             int    `json:"uid"`
	PlayURL64       string `json:"playUrl64"` //mp3 64kbps
	PlayURL32       string `json:"playUrl32"` //mp3 32kbps
	PlayPathHq      string `json:"playPathHq"`
	PlayPathAacv164 string `json:"playPathAacv164"` //m4a 64kbps
	PlayPathAacv224 string `json:"playPathAacv224"` //m4a 24kbps
	Title           string `json:"title"`
	Duration        int    `json:"duration"`
	AlbumID         int    `json:"albumId"`
	AlbumTitle      string `json:"albumTitle"`
	AlbumImage      string `json:"albumImage"`
	IsPaid          bool   `json:"isPaid"`
	IsFree          bool   `json:"isFree"`
	IsVideo         bool   `json:"isVideo"`
	IsDraft         bool   `json:"isDraft"`
	IsRichAudio     bool   `json:"isRichAudio"`
	IsAuthorized    bool   `json:"isAuthorized"`
	PriceTypeID     int    `json:"priceTypeId"`
	PriceTypeEnum   int    `json:"priceTypeEnum"`
	Type            int    `json:"type"`
	RelatedID       int    `json:"relatedId"`
	OrderNo         int    `json:"orderNo"`
	VipFirstStatus  int    `json:"vipFirstStatus"`
	PaidType        int    `json:"paidType"`
}

//Playlist Mobwsa接口: https://mobwsa.ximalaya.com/mobile/playlist/album/page?albumId=%d&pageId=%d
type Playlist struct {
	Msg        string        `json:"msg"`
	Ret        int           `json:"ret"`
	MaxPageID  int           `json:"maxPageId"`
	PageSize   int           `json:"pageSize"`
	List       AudioInfoList `json:"list"`
	PageID     int           `json:"pageId"`
	TotalCount int           `json:"totalCount"`
}

//VipAudioInfo VIP音频信息
type VipAudioInfo struct {
	Ret                  int    `json:"ret"`
	Msg                  string `json:"msg"`
	TrackID              int    `json:"trackId"`
	UID                  int    `json:"uid"`
	AlbumID              int    `json:"albumId"`
	Title                string `json:"title"`
	Domain               string `json:"domain"`
	TotalLength          int    `json:"totalLength"`
	SampleDuration       int    `json:"sampleDuration"`
	SampleLength         int    `json:"sampleLength"`
	IsAuthorized         bool   `json:"isAuthorized"`
	APIVersion           string `json:"apiVersion"`
	Seed                 int    `json:"seed"`
	FileID               string `json:"fileId"`
	BuyKey               string `json:"buyKey"`
	Duration             int    `json:"duration"`
	Ep                   string `json:"ep"`
	HighestQualityLevel  int    `json:"highestQualityLevel"`
	DownloadQualityLevel int    `json:"downloadQualityLevel"`
	AuthorizedType       int    `json:"authorizedType"`
}

//UserInfo 用户信息
type UserInfo struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	Data struct {
		UID           int         `json:"uid"`
		RealUID       int         `json:"realUid"`
		Nickname      string      `json:"nickname"`
		LogoPic       string      `json:"logoPic"`
		IsLoginBan    bool        `json:"isLoginBan"`
		IsVerified    bool        `json:"isVerified"`
		Ptitle        interface{} `json:"ptitle"`
		Mobile        string      `json:"mobile"`
		IsRobot       bool        `json:"isRobot"`
		VerifyType    int         `json:"verifyType"`
		IsVip         bool        `json:"isVip"`
		VipExpireTime int         `json:"vipExpireTime"`
		AnchorGrade   int         `json:"anchorGrade"`
		UserGrade     int         `json:"userGrade"`
		UserTitle     interface{} `json:"userTitle"`
		LogoType      int         `json:"logoType"`
	} `json:"data"`
}

//AudioInfoMobile Mobile接口
type AudioInfoMobile struct {
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Data struct {
		List []struct {
			TrackID                int     `json:"trackId"`
			TrackRecordID          int     `json:"trackRecordId"`
			UID                    int     `json:"uid"`
			PlayURL64              string  `json:"playUrl64"`
			PlayURL32              string  `json:"playUrl32"`
			PlayPathHq             string  `json:"playPathHq"`
			PlayPathAacv164        string  `json:"playPathAacv164"`
			PlayPathAacv224        string  `json:"playPathAacv224"`
			Title                  string  `json:"title"`
			Duration               int     `json:"duration"`
			AlbumID                int     `json:"albumId"`
			IsPaid                 bool    `json:"isPaid"`
			IsFree                 bool    `json:"isFree"`
			IsVideo                bool    `json:"isVideo"`
			IsDraft                bool    `json:"isDraft"`
			IsRichAudio            bool    `json:"isRichAudio"`
			IsAuthorized           bool    `json:"isAuthorized"`
			Price                  float64 `json:"price"`
			DiscountedPrice        float64 `json:"discountedPrice"`
			PriceTypeID            int     `json:"priceTypeId"`
			SampleDuration         int     `json:"sampleDuration"`
			PriceTypeEnum          int     `json:"priceTypeEnum"`
			DisplayPrice           string  `json:"displayPrice"`
			DisplayDiscountedPrice string  `json:"displayDiscountedPrice"`
			VipPrice               float64 `json:"vipPrice"`
			DisplayVipPrice        string  `json:"displayVipPrice"`
			Type                   int     `json:"type"`
			RelatedID              int     `json:"relatedId"`
			OrderNo                int     `json:"orderNo"`
			IsHoldCopyright        bool    `json:"isHoldCopyright"`
			VipFirstStatus         int     `json:"vipFirstStatus"`
			PaidType               int     `json:"paidType"`
			IsSample               bool    `json:"isSample"`
			ProcessState           int     `json:"processState"`
			CreatedAt              int64   `json:"createdAt"`
			CoverSmall             string  `json:"coverSmall"`
			CoverMiddle            string  `json:"coverMiddle"`
			CoverLarge             string  `json:"coverLarge"`
			Nickname               string  `json:"nickname"`
			SmallLogo              string  `json:"smallLogo"`
			UserSource             int     `json:"userSource"`
			OpType                 int     `json:"opType"`
			IsPublic               bool    `json:"isPublic"`
			Likes                  int     `json:"likes"`
			Playtimes              int     `json:"playtimes"`
			Comments               int     `json:"comments"`
			Shares                 int     `json:"shares"`
			Status                 int     `json:"status"`
			ExpireTime             int64   `json:"expireTime"`
			IsTrailer              bool    `json:"isTrailer"`
		} `json:"list"`
		PageID     int `json:"pageId"`
		PageSize   int `json:"pageSize"`
		MaxPageID  int `json:"maxPageId"`
		TotalCount int `json:"totalCount"`
	} `json:"data"`
}
