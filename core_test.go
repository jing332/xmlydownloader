package xmlydownloader

import (
	"log"
	"testing"
)

func TestGetAllTrackList(t *testing.T) {
	var trackInfoList []*TrackInfo
	albumID := 2780581 //https://www.ximalaya.com/youshengshu/2780581/
	tracks, err := GetTrackList(albumID, 1, false)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range tracks.Data.List {
		trackInfoList = append(trackInfoList, v)
	}

	for i := 2; i <= tracks.Data.MaxPageID; i++ {
		tracks, err = GetTrackList(albumID, i, false)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range tracks.Data.List {
			trackInfoList = append(trackInfoList, v)
		}
	}
	for i, v := range trackInfoList {
		t.Log(i, v.Title)
	}
}

func TestQRCode(t *testing.T) {
	qrCode, err := GetQRCode()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(qrCode.Img)
	log.Println(qrCode.QrID)
	//
	status, cookie, err := CheckQRCodeStatus("FEDECD84A3014713B396B4B6ED4F3483")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status, cookie)
}
