package xmlydownloader

import (
	"log"
	"strconv"
	"strings"
)

//DecryptFileName 解密文件名
func DecryptFileName(send int, s string) string {
	x := New(float64(send))

	uri := x.CgFun(s)
	if "/" != string(uri[0]) {
		uri = "/" + uri
	}
	return uri
}

type Xmly struct {
	CgStr      string
	RandomSeed float64
}

func New(send float64) *Xmly {
	x := &Xmly{}
	x.RandomSeed = send
	x.CgHun()
	return x
}

func (x *Xmly) Ran() float64 {
	j := 211*x.RandomSeed + 30031
	x.RandomSeed = float64(int(j) % 65536)
	return x.RandomSeed / float64(65536)
}

func (x *Xmly) CgHun() {
	x.CgStr = ""
	key :=
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ/\\:._-1234567890"

	kLen := len(key)
	for i := 0; i < kLen; i++ {
		ran := x.Ran()
		r := int(ran * float64(len(key)))

		x.CgStr += string([]byte(key)[r])
		key = strings.ReplaceAll(key, string([]byte(key)[r]), "")
	}
}

func (x *Xmly) CgFun(t string) string {
	strs := strings.Split(t, "*")

	e := ""
	for n := 0; n < len(strs)-1; n++ {
		if strs[n] != "" {
			index, err := strconv.Atoi(strs[n])
			if err != nil {
				log.Fatal(err)
			}
			e += string([]byte(x.CgStr)[index])
		}
	}
	return e
}

//DecryptUrlParams 解密URL参数
func DecryptUrlParams(s string) (sign string, buyKey, token, timestamp int) {
	s1 := encrypt3(s)
	s2 := encrypt(encrypt2("d"+o+"9", u), s1)
	ss := strings.Split(s2, "-")

	buyKey, _ = strconv.Atoi(ss[0])
	token, _ = strconv.Atoi(ss[2])
	timestamp, _ = strconv.Atoi(ss[3])

	return ss[1], buyKey, token, timestamp
}

func charCodeAt(s string, n int) rune {
	i := 0
	for _, r := range s {
		if i == n {
			return r
		}
		i++
	}
	return 0
}

const o = "g3utf1k6yxdwi0"

var u = []int{
	19, 1, 4, 7, 30, 14, 28, 8, 24, 17, 6, 35, 34, 16, 9, 10, 13, 22, 32, 29, 31,
	21, 18, 3, 2, 23, 25, 27, 11, 20, 5, 15, 12, 0, 33, 26,
}

func encrypt(e string, t []int16) string {
	var (
		n = 0
		r []int
		a = 0
		s string
		o = 0
	)

	for o = 0; o < 256; o++ {
		r = append(r, o)
	}

	for o = 0; o < 256; o++ {
		a = (a + r[o] + int(charCodeAt(e, o%len(e)))) % 256
		n = r[o]
		r[o] = r[a]
		r[a] = n
	}

	a, o = 0, 0
	for u := 0; u < len(t); u++ {
		o = (o + 1) % 256
		a = (a + r[o]) % 256
		n = r[o]
		r[o] = r[a]
		r[a] = n
		s += string(int(t[u]) ^ (r[(r[o]+r[a])%256]))
	}

	return s
}

func encrypt2(key string, key2 []int) string {
	var n []string
	for r := 0; r < len(key); r++ {
		a := []rune("a")[0]
		if int16(a) <= int16(key[r]) && int16([]rune("z")[0]) >= int16(key[r]) {
			a = rune(int16(charCodeAt(key, r)) - 97)
		} else {
			a = rune(int16(charCodeAt(key, r)) - 48 + 26)
		}

		n = append(n, "")
		for i := 0; 36 > i; i++ {
			if int16(key2[i]) == int16(a) {
				a = rune(i)
				break
			}
		}

		if 25 < a {
			n[r] = string(int16(a) - 26 + 48)
		} else {
			n[r] = string(int16(a) + 97)
		}
	}

	var str string
	if len(n) != 0 {
		for _, s := range n {
			str += s
		}
	} else {
		str = ""
	}

	return str
}

func encrypt3(s string) []int16 {
	var (
		t    = 0
		n    = 0
		r    = 0
		sLen = len(s)
		i    []int16
		o    = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 62, -1, -1, -1, 63, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, -1, -1, -1, -1, -1, -1, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, -1, -1, -1, -1, -1}
	)

	for r < sLen {
		t = o[255&charCodeAt(s, r)]
		r++
		for r < sLen && -1 == t {
			t = o[255&charCodeAt(s, r)]
			r++
		}
		if -1 == t {
			break
		}

		n = o[255&charCodeAt(s, r)]
		r++
		for r < sLen && -1 == n {
			n = o[255&charCodeAt(s, r)]
			r++
		}
		if -1 == t {
			break
		}

		i = append(i, int16((t<<2)|((48&n)>>4)))

		t = 255 & int(charCodeAt(s, r))
		r++
		if 61 == t {
			return i
		}
		t = o[t]
		for r < sLen && -1 == t {
			t = 255 & int(charCodeAt(s, r))
			r++

			if 61 == t {
				return i
			}
			t = o[t]
		}
		if -1 == t {
			break
		}

		i = append(i, int16(((15&n)<<4)|((60&t)>>2)))

		n = 255 & int(charCodeAt(s, r))
		r++
		if 61 == n {
			return i
		}
		n = o[n]
		for r < sLen && -1 == n {
			n = 255 & int(charCodeAt(s, r))
			r++
			if 61 == n {
				return i
			}
			n = o[n]
		}
		if -1 == n {
			break
		}
		i = append(i, int16(((3&t)<<6)|n))
	}

	return i
}
