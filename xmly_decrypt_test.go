package xmlydownloader

import (
	"testing"
)

func TestDecryptFileName(t *testing.T) {
	send := 2047
	fileId := "51*10*20*60*53*9*7*34*16*13*7*13*19*7*44*35*7*31*58*51*34*0*8*65*38*60*21*65*31*23*14*42*18*62*56*25*43*25*13*59*58*29*53*55*59*33*33*64*61*21*46*"
	t.Log(DecryptFileName(send, fileId))
}

func TestDecryptUrlParams(t *testing.T) {
	ep := "3kNrPox/Sn5Sj6gKPokctQtfTU52gnKTStYYeA+0XXn9y+nciv2AmOoN2/fegvBlDLVxznoAf6B82/T2wQYQ074aPQ=="
	t.Log(DecryptUrlParams(ep))
}
