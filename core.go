package xmlydownloader

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

//GetAlbumInfo 获取专辑信息
func GetAlbumInfo(albumID int) (title string, audioCount, pageCount int, err error) {
	resp, err := http.Get(fmt.Sprintf("https://www.ximalaya.com/youshengshu/%d/", albumID))
	if err != nil {
		return "", 0, 0, fmt.Errorf("获取专辑信息失败: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", -1, -1, fmt.Errorf("获取专辑信息失败: StatusCode != 200")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", -1, -1, fmt.Errorf("new goquery document fail:" + err.Error())
	}

	//标题
	title = doc.Find("h1.title").Text()

	//音频数量
	r, _ := regexp.Compile("\\d+\\.?\\d*")
	num := r.FindString(doc.Find("div.head").Text())
	audioCount, _ = strconv.Atoi(num)

	//页面数量
	list := doc.Find("ul.pagination-page").Children()
	size := list.Size()
	if size > 6 { //超过5页
		i, _ := strconv.Atoi(list.Eq(list.Size() - 2).Text())
		pageCount = i
	} else if size == 0 { //仅一页
		pageCount = 1
	} else { //大于0页 && 小于等于5页
		pageCount = size - 1 //1为下一页按钮
	}
	return
}

//GetVipAudioInfo 获取VIP音频信息
func GetVipAudioInfo(trackId int, cookie string) (ai *AudioInfo, err error) {
	ts := time.Now().Unix()
	url := fmt.Sprintf(
		"https://mpay.ximalaya.com/mobile/track/pay/%d/%d?device=pc&isBackend=true&_=%d",
		trackId, ts, ts)

	resp, err := HttpGetByCookie(url, cookie, Android)
	if err != nil {
		return ai, fmt.Errorf("获取音频信息失败: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ai, err
	}

	var vipAi VipAudioInfo
	err = jsoniter.Unmarshal(body, &vipAi)
	if err != nil {
		return ai, fmt.Errorf("无法解析JSON: %v", err)
	}

	if vipAi.Ret != 0 {
		return ai, fmt.Errorf("无法获取VIP音频信息: %v", jsoniter.Get(body, "msg").ToString())
	}

	fileName := DecryptFileName(vipAi.Seed, vipAi.FileID)
	sign, _, token, timestamp := DecryptUrlParams(vipAi.Ep)

	args := fmt.Sprintf("?sign=%s&buy_key=%s&token=%d&timestamp=%d&duration=%d",
		sign, vipAi.BuyKey, token, timestamp, vipAi.Duration)

	ai = &AudioInfo{TrackID: trackId, Title: vipAi.Title}

	ai.PlayPathAacv164 = vipAi.Domain + "/download/" + vipAi.APIVersion + fileName + args
	return ai, nil
}

//GetAudioInfo 获取音频信息
func GetAudioInfo(albumID, page, pageSize int) (audioList []AudioInfo, err error) {
	format := fmt.Sprintf("https://m.ximalaya.com/m-revision/common/album/queryAlbumTrackRecordsByPage?albumId=%d&page=%d&pageSize=%d&asc=true", albumID, page, pageSize)

	resp, err := client.Get(format)
	if err != nil {
		return nil, fmt.Errorf("http get %v fail:%v", format, err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	list := jsoniter.Get(body, "data").Get("trackDetailInfos")
	for i2 := 0; i2 < list.Size(); i2++ {
		v := list.Get(i2).Get("trackInfo")
		audioList = append(audioList, AudioInfo{TrackID: v.Get("id").ToInt(),
			PlayPathAacv164: v.Get("playPath").ToString(), Title: v.Get("title").ToString()})
	}

	return audioList, nil
}

//GetAllAudioInfo 获取所有音频信息
func GetAllAudioInfo(albumID int) (list []*AudioInfo, err error) {
	firstPlayList, err := GetAudioInfoListByPageID(albumID, 0)
	if err != nil {
		return nil, fmt.Errorf("无法获取播放列表: 0, %s", err)
	}
	for _, v := range firstPlayList.List {
		list = append(list, v)
	}

	for i := 1; i < firstPlayList.MaxPageID; i++ {
		playList, err := GetAudioInfoListByPageID(albumID, i)
		if err != nil {
			return nil, fmt.Errorf("无法获取播放列表: %d, %s", i, err)
		}
		for _, v := range playList.List {
			list = append(list, v)
		}
	}
	return list, nil
}

//GetAudioInfoListByPageID 使用PageID获取音频信息列表
func GetAudioInfoListByPageID(albumID, pageID int) (playlist *Playlist, err error) {
	url := fmt.Sprintf("http://mobwsa.ximalaya.com/mobile/playlist/album/page?albumId=%d&pageId=%d",
		albumID, pageID)
	resp, err := HttpGet(url, Android)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	playlist = &Playlist{}
	err = jsoniter.Unmarshal(data, playlist)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

//GetUserInfo 使用Cookie获取用户信息
func GetUserInfo(cookie string) (*UserInfo, error) {
	resp, err := HttpGetByCookie("https://www.ximalaya.com/revision/main/getCurrentUser", cookie, PC)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ui := &UserInfo{}
	err = jsoniter.Unmarshal(data, ui)
	if err != nil {
		return nil, err
	}

	return ui, nil
}

//GetAlbumInfoByMobileAPI 使用MobileV1API获取音频列表
//
//isAsc: true为升序(默认), false为降序
func GetTrackListByMobile(albumID, pageID int, isAsc bool) (tracks *TrackList, err error) {
	url := fmt.Sprintf(
		"https://mobile.ximalaya.com/mobile/v1/album/track/ts-%d?ac=WIFI&albumId=%d&device=android&isAsc=%t&pageId=%d&pageSize=200",
		time.Now().Unix(), albumID, isAsc, pageID)
	resp, err := HttpGet(url, Android)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tracks = &TrackList{}
	err = jsoniter.Unmarshal(data, tracks)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

//GetQRCode 获取登录二维码
func GetQRCode() (qrCode *QRCode, err error) {
	resp, err := HttpGet("https://passport.ximalaya.com/web/qrCode/gen?level=L", PC)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	qrCode = &QRCode{}
	err = jsoniter.Unmarshal(data, qrCode)
	if err != nil {
		return nil, err
	}
	return qrCode, err
}

//CheckQRCodeStatus 检查二维码的状态
func CheckQRCodeStatus(qrID string) (status *QRCodeStatus, cookie string, err error) {
	url := fmt.Sprintf("https://passport.ximalaya.com/web/qrCode/check/%s/%d", qrID, time.Now().Unix())
	resp, err := HttpGet(url, PC)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	status = &QRCodeStatus{}
	err = jsoniter.Unmarshal(data, status)
	if err != nil {
		return nil, "", err
	}
	if status.Ret == 0 {
		return status, resp.Header.Values("Set-Cookie")[1], nil
	}

	return status, "", nil
}

////GetVipTrackInfoByMobile 使用Cookie获取VIP音频信息
//func GetTrackInfoByMobile(trackID, cookie string) error {
//	url := fmt.Sprintf("https://mpay.ximalaya.com/mobile/track/pay/%d/%d?device=android",
//		time.Now().Unix(), trackID)
//	resp, err := HttpGet(url)
//	if err != nil {
//		return err
//	}
//	defer resp.Body.Close()
//
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//}

//func GetAudioInfoByMobile(trackID int, cookie string) error {
//	url := fmt.Sprintf("https://mpay.ximalaya.com/mobile/track/pay/%d?device=pc&isBackend=true&_=%d",
//		trackID, math.Round(float64(time.Now().UnixNano()/1e6)))
//	req, err := http.NewRequest("GET", url, nil)
//	if err != nil {
//		return err
//	}
//
//	sign, err := GetXmMd5()
//	if err != nil {
//		return err
//	}
//
//	req.Header.Set("Cookie", cookie)
//	req.Header.Set("xm-sign", sign)
//	req.Header.Set("Connection", "keep-alive")
//	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
//	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36")
//
//	resp, err := client.Do(req)
//	if err != nil {
//		return err
//	}
//
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	log.Println(string(data))
//
//	return nil
//}
//
//func GetXmMd5() (string, error) {
//	resp, err := HttpGet("https://www.ximalaya.com/revision/time")
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	data, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		return "", err
//	}
//
//	ts := time.Now().UnixNano() / 1e6
//	hash := md5.Sum([]byte("himalaya-" + string(data)))
//	str := fmt.Sprintf("%x", hash)
//	sign := fmt.Sprintf("%s(%d)%s(%d)%d", str, int(math.Round(float64(rand.Float32()*100))), string(data),
//		int(math.Round(float64(rand.Float32()*100))), ts)
//
//	return sign, err
//}
