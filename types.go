package xmlydownloader

//TrackInfo 音频信息(MobileV1API)
//
//TrackList
type TrackInfo struct {
	TrackID                int     `json:"trackId"`
	TrackRecordID          int     `json:"trackRecordId"`
	UID                    int     `json:"uid"`
	PlayURL64              string  `json:"playUrl64"`
	PlayURL32              string  `json:"playUrl32"`
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
	IsTrailer              bool    `json:"isTrailer"`
}

//TrackList 音频列表(MobileV1API)
//
//https://mobile.ximalaya.com/mobile/v1/album/track/ts-%d?ac=WIFI&albumId=%d&device=android&isAsc=%t&pageId=%d&pageSize=200
type TrackList struct {
	Msg  string `json:"msg"`
	Ret  int    `json:"ret"`
	Data struct {
		List       []*TrackInfo `json:"list"`
		PageID     int          `json:"pageId"`
		PageSize   int          `json:"pageSize"`
		MaxPageID  int          `json:"maxPageId"`
		TotalCount int          `json:"totalCount"`
	} `json:"data"`
}

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

//Playlist 播放列表
//
//https://mobwsa.ximalaya.com/mobile/playlist/album/page?albumId=%d&pageId=%d
type Playlist struct {
	Msg        string       `json:"msg"`
	Ret        int          `json:"ret"`
	MaxPageID  int          `json:"maxPageId"`
	PageSize   int          `json:"pageSize"`
	List       []*AudioInfo `json:"list"`
	PageID     int          `json:"pageId"`
	TotalCount int          `json:"totalCount"`
}

//VipTrackInfo VIP音频信息(MobileAPI)
//
//需使用 Java unidbg库 解密音频URL
//unidbg: https://github.com/zhkl0228/unidbg
type ChargeTrackInfo struct {
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
}

//VipAudioInfo VIP音频信息(WebAPI)
//
//需使用 DecryptFileName(send, fileId) 与 DecryptUrlParams(ep) 解密音频URL
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

//QRCode 二维码
type QRCode struct {
	Ret  int    `json:"ret"`
	Msg  string `json:"msg"`
	QrID string `json:"qrId"`
	Img  string `json:"img"`
}

//QRCodeStatus 二维码状态
type QRCodeStatus struct {
	Ret                int         `json:"ret"`
	Msg                string      `json:"msg"`
	BizKey             interface{} `json:"bizKey"`
	UID                int         `json:"uid"`
	Token              interface{} `json:"token"`
	UserType           interface{} `json:"userType"`
	IsFirst            bool        `json:"isFirst"`
	ToSetPwd           bool        `json:"toSetPwd"`
	LoginType          string      `json:"loginType"`
	MobileMask         string      `json:"mobileMask"`
	MobileCipher       string      `json:"mobileCipher"`
	CaptchaInfo        interface{} `json:"captchaInfo"`
	Avatar             interface{} `json:"avatar"`
	ThirdpartyAvatar   interface{} `json:"thirdpartyAvatar"`
	ThirdpartyNickname interface{} `json:"thirdpartyNickname"`
	SmsKey             interface{} `json:"smsKey"`
	ThirdpartyID       interface{} `json:"thirdpartyId"`
}
