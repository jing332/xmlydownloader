package xmlydownloader

import (
	"log"
	"testing"
)

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
