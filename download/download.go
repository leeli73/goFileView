package download

//https://blog.csdn.net/shixingya/article/details/88951782

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/leeli73/goFileView/utils"
)

func IsFileExist(filename string, filesize int64) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	if filesize == info.Size() {
		return true
	}
	os.Remove(filename)
	return false
}
func DownloadFile(url string, localPath string) (string, error) {
	var (
		fsize   int64
		buf     = make([]byte, 32*1024)
		written int64
	)
	tmpFilePath := localPath + ".download"
	client := new(http.Client)
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	fsize, err = strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 32)
	if err != nil {
		log.Println("Error: <", err, "> when get file remote size")
		return "", err
	}
	if IsFileExist(localPath, fsize) {
		return "had", nil
	}
	file, err := os.Create(tmpFilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	if resp.Body == nil {
		return "", errors.New("body is null")
	}
	defer resp.Body.Close()
	for {
		nr, er := resp.Body.Read(buf)
		if nr > 0 {
			nw, ew := file.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	if err == nil {
		file.Close()
		fileMd5 := utils.GetFileMD5(tmpFilePath)
		newPath := "cache/download/" + fileMd5 + path.Ext(localPath)
		os.Rename(tmpFilePath, newPath)
		log.Printf("Download file <filename:%s, md5:%s> success\n", path.Base(localPath), fileMd5)
		return newPath, nil
	}
	return "", err
}
