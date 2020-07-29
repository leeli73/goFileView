package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"time"
)

func ComparePath(a string, b string) bool {
	if len(a) >= len(b) {
		if strings.Compare(a[0:len(b)], b) == 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ConvertToPDF(filePath string) string {
	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		commandName = "cmd"
		params = []string{"/c", "soffice", "--headless", "--invisible", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	} else if runtime.GOOS == "linux" {
		commandName = "libreoffice"
		params = []string{"--invisible", "--headless", "--convert-to", "pdf", "--outdir", "cache/pdf/", filePath}
	}
	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/pdf/" + strings.Split(path.Base(filePath), ".")[0] + ".pdf"
		if ok, _ := PathExists(resultPath); ok {
			log.Printf("Convert <%s> to pdf\n", path.Base(filePath))
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func ConvertToImg(filePath string) string {
	fileName := strings.Split(path.Base(filePath), ".")[0]
	//fileDir := path.Dir(filePath)
	fileExt := path.Ext(filePath)
	if fileExt != ".pdf" {
		return ""
	}
	os.Mkdir("cache/convert/"+fileName, os.ModePerm)
	commandName := ""
	var params []string
	if runtime.GOOS == "windows" {
		commandName = "cmd"
		params = []string{"/c", "magick", "convert", "-density", "130", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	} else if runtime.GOOS == "linux" {
		commandName = "convert"
		params = []string{"-density", "130", filePath, "cache/convert/" + fileName + "/%d.jpg"}
	}
	if _, ok := interactiveToexec(commandName, params); ok {
		resultPath := "cache/convert/" + strings.Split(path.Base(filePath), ".")[0]
		if ok, _ := PathExists(resultPath); ok {
			log.Printf("Convert <%s> to images\n", path.Base(filePath))
			return resultPath
		} else {
			return ""
		}
	} else {
		return ""
	}
}

func interactiveToexec(commandName string, params []string) (string, bool) {
	cmd := exec.Command(commandName, params...)
	buf, err := cmd.Output()
	w := bytes.NewBuffer(nil)
	cmd.Stderr = w
	if err != nil {
		log.Println("Error: <", err, "> when exec command read out buffer")
		return "", false
	} else {
		return string(buf), true
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetFileMD5(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Println("Error: <", err, "> when open file to get md5")
		return ""
	}
	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Error: <", err, "> when get md5")
		return ""
	}
	f.Close()
	return fmt.Sprintf("%x", md5hash.Sum(nil))
}

func randString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func IsInArr(key string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if key == arr[i] {
			return true
		}
	}
	return false
}
