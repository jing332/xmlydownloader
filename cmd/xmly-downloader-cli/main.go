package main

import (
	"fmt"
	xmlydown "github.com/jing332/xmlydownloader"
	"github.com/mattn/go-colorable"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var (
	albumID                                 = 4176595 //https://www.ximalaya.com/youshengshu/19383749/
	begin, end, maxTaskCount, maxRetryCount = 1, -1, 1, 3
	forceSave                               = false
	isAddNumber                             = false

	rmStr       = ""
	downloadDir = "download"
	cookie      = ""

	trackCountLen = ""
)

type DownloadURL struct {
	URL      string
	FilePath string
	Number   int
}

var (
	downloadFailed []DownloadURL
	client         = http.Client{}
	fileNameRegexp = regexp.MustCompile("[/\\/:*?\"<>|]")
)

func initCli() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	defaultDir := appPath + "/Downlaod"
	osName := runtime.GOOS
	if "android" == osName {
		defaultDir = "/storage/sdcard0/Downlaod"
	}

	app := cli.NewApp()
	app.Version = "1.0.1"
	app.Usage = "喜马拉雅FM专辑下载器"
	app.Author = "jing (https://github.com/jing332)"
	app.Flags = []cli.Flag{
		cli.IntFlag{Name: "id", Usage: "专辑ID", Value: 0, Destination: &albumID},
		cli.IntFlag{Name: "b, begin", Usage: "起始音频", Value: 1, Destination: &begin},
		cli.IntFlag{Name: "e, end", Usage: "结束音频", Value: -1, Destination: &end},
		cli.IntFlag{Name: "p, parallel", Usage: "最大并行下载任务数量(1-5)", Value: 3, Destination: &maxTaskCount},
		cli.IntFlag{Name: "r, retry", Usage: "下载失败后重试次数", Value: 3, Destination: &maxRetryCount},
		cli.StringFlag{Name: "s, save, saveDir", Usage: "专辑保存目录", Value: defaultDir, Destination: &downloadDir},
		cli.StringFlag{Name: "c, cookie", Usage: "Cookie (非免费专辑必须设置)", Value: "", Destination: &cookie},
		cli.StringFlag{Name: "rm, remove", Usage: "删除文件名中指定字符串", Value: "", Destination: &rmStr},
		cli.BoolFlag{Name: "an, addNum", Usage: "在文件名前添加序号", Destination: &isAddNumber},
		cli.BoolFlag{Name: "ow, overwrite", Usage: "不检查文件名与文件长度直接覆盖保存", Destination: &forceSave},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatalf("解析参数失败: %v", err)
	}

	if albumID == 0 {
		log.Fatal(AnsiRed("请使用 -id 参数设置专辑ID"))
	}

	if 0 <= maxTaskCount && 5 <= maxTaskCount {
		log.Fatal(AnsiRed("最大任务数量应大于0小于6"))
	}
}

func main() {
	log.SetFlags(log.Ltime)
	log.SetOutput(colorable.NewColorableStdout())
	initCli()

	err := Start()
	if err != nil {
		log.Fatal(err)
	}

	if len(downloadFailed) > 0 {
		log.Println("下载失败:", len(downloadFailed))
		for _, v := range downloadFailed {
			log.Printf("%s %s, %s", AnsiRed(fmt.Sprintf("[%v]", v.Number)),
				AnsiCyan(filepath.Base(v.FilePath)),
				AnsiYellow(v.URL))
		}
		var input string
		fmt.Print("\n重新下载请输入 Y 或 y 并回车:")
		fmt.Scanln(&input)
		if input == "Y" || input == "y" {
			fmt.Println("正在重新下载已失败文件")
			ch := make(chan int, maxTaskCount)
			var wg sync.WaitGroup

			for i, v := range downloadFailed {
				i++
				ch <- 1
				wg.Add(1)
				go func(url, dir, fileName string, index int, ch chan int) {
					log.Printf("%s 开始下载: %s, %s", AnsiGreen(fmt.Sprintf("[%v]", index)),
						AnsiCyan(fileName), AnsiYellow(url))

					//去除非法字符
					fileName = fileNameRegexp.ReplaceAllString(fileName, " ")
					//删除指定字符串
					fileName = strings.ReplaceAll(fileName, rmStr, "")

					err = downloadFile(url, fmt.Sprintf("%s/%s.m4a", dir, fileName), forceSave, index, 0, ch)
					if err != nil {
						log.Fatal(err)
					}
					wg.Add(-1)
				}(v.URL, downloadDir, filepath.Base(v.FilePath), i, ch)
			}
			wg.Wait()
		}
	}
}

func Start() error {
	//获取有声小说信息
	title, trackCount, _, err := xmlydown.GetAlbumInfo(albumID)
	if err != nil {
		return err
	}

	trackCountLen = strconv.Itoa(GetIntLength(trackCount) + 1)
	downloadDir += "/" + fileNameRegexp.ReplaceAllString(title, "")

	ch := make(chan int, maxTaskCount)
	var wg sync.WaitGroup

	//获取所有音频
	list := xmlydown.GetAudioInfoList(albumID, trackCount)
	for i, v := range list {
		if v.URL == "" {
			log.Printf("%s %s", AnsiRed(fmt.Sprintf("- [%v]", i+1)), AnsiYellow(v.Title))
		} else {
			log.Printf("%s %s", AnsiGreen(fmt.Sprintf("+ [%v]", i+1)), AnsiYellow(v.Title))
		}
	}

	fmt.Print("开始下载请输入 Y/y 并回车: ")
	var input string
	_, err = fmt.Scanln(&input)
	if err != nil || !("Y" == input || "y" == input) {
		os.Exit(0)
	}

	for i, v := range list {
		i++
		if i >= begin && (i <= end || end == -1) {
			ch <- 1
			wg.Add(1)

			go func(url, dir, fileName string, trackId, index int, ch chan int) {
				if v.URL == "" {
					if cookie != "" {
						albumInfo, err := xmlydown.GetVipAudioInfo(trackId, cookie)
						if err != nil {
							log.Fatalf("%s %s %v", AnsiRed(fmt.Sprintf("[%v]", index)),
								AnsiCyan(fileName), err)
						}

						url = albumInfo.URL
						fileName = albumInfo.Title
					} else {
						log.Fatalf("非免费音频无法下载, 请设置参数 -cookie !")
					}
				}

				log.Printf("%s 开始下载: %s, %s", AnsiGreen(fmt.Sprintf("[%d]", index)),
					AnsiCyan(fileName), AnsiYellow(url))
				//去除非法字符
				fileName = fileNameRegexp.ReplaceAllString(fileName, "")
				//删除指定字符串
				fileName = strings.ReplaceAll(fileName, rmStr, "")
				err = downloadFile(url, fmt.Sprintf("%s/%s.m4a", dir, fileName), false, index, 0, ch)
				if err != nil {
					log.Fatal(err)
				}
				wg.Add(-1)
			}(v.URL, downloadDir, v.Title, v.TrackId, i, ch)
		}
	}

	wg.Wait()
	return nil
}

func GetIntLength(n int) int {
	var i = 0

	if n == 0 {
		i = 1
	}

	for i = 0; n != 0; i++ {
		n /= 10
	}
	return i
}

var errFormat = "\u001B[91m[%v]\u001B[0m 下载失败(第\u001B[32m%v\u001B[0m次重试) \u001B[36m%s\u001B[0m: %s"

func downloadFile(url, filePath string, forceSave bool, number, retried int, ch chan int) error {
	var tmpFilePath string
	if isAddNumber {
		fileName := filepath.Base(filePath)
		tmpFilePath = fmt.Sprintf("%s/%0"+trackCountLen+"d %s", downloadDir, number, fileName)
	} else {
		tmpFilePath = filePath
	}

	if !forceSave {
		if fileInfo, err := os.Stat(tmpFilePath); err == nil {
			resp, err := http.Head(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if fileInfo.Size() == resp.ContentLength {
				log.Printf("%s 文件已存在: %s", AnsiRed(fmt.Sprintf("[%v]", number)),
					AnsiCyan(filepath.Base(tmpFilePath)))
				<-ch
				return nil
			}
		}
	}

	resp, err := httpGet(url)
	if err != nil {
		log.Printf(errFormat, number, retried, filepath.Base(tmpFilePath), err.Error())
		return retry(maxRetryCount, retried, number, ch, url, filePath, forceSave)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Printf(errFormat, number, retried, filepath.Base(tmpFilePath), "HTTP状态码 != 200")
		return retry(maxRetryCount, retried, number, ch, url, filePath, forceSave)
	}

	//目录不存在则创建
	err = os.MkdirAll(filepath.Dir(tmpFilePath), 0777)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("make dir fail: %s", err.Error())
	}

	//创建并写入文件
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return fmt.Errorf("create file %s fail: %s", filePath, err.Error())
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Printf("%s 保存文件到本地时失败: %v", AnsiRed(fmt.Sprintf("[%v]", number)), err)
		return retry(maxRetryCount, retried, number, ch, url, filePath, forceSave)
	}

	<-ch
	return nil
}

func retry(maxRetry, retried, number int, ch chan int, url, filePath string, forceSave bool) error {
	if retried < maxRetry {
		return downloadFile(url, filePath, forceSave, number, retried+1, ch)
	} else {
		downloadFailed = append(downloadFailed, DownloadURL{URL: url, FilePath: filePath, Number: number})
		<-ch
		return nil
	}
}

func httpGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4170.0 Safari/537.36 Edg/85.0.552.1")

	return client.Do(req)
}

func AnsiGreen(v interface{}) string {
	return fmt.Sprintf("\u001B[92m%v\u001B[0m", v)
}

func AnsiRed(v interface{}) string {
	return fmt.Sprintf("\u001B[91m%v\u001B[0m", v)
}

func AnsiYellow(v interface{}) string {
	return fmt.Sprintf("\u001B[1;33m%v\u001B[0m", v)
}

func AnsiCyan(v interface{}) string {
	return fmt.Sprintf("\u001B[33m%v\u001B[0m", v)
}
