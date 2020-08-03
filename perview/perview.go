package perview

// sudo convert -density 300 test.pdf %d.jpg
// libreoffice  --invisible --convert-to pdf test.docx
import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/leeli73/goFileView/download"
	"github.com/leeli73/goFileView/utils"
)

type NowFile struct {
	Md5            string
	Ext            string
	LastActiveTime int64
}

var (
	Pattern      string
	Address      string
	AllFile      map[string]*NowFile
	ExpireTime   int64
	AllOfficeEtx = []string{".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx", ".txt"}
	AllImageEtx  = []string{".jpg", ".png", ".gif"}
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if Pattern == "" || Address == "" {
		log.Fatal("Error: Please Init Go File View First!")
	}
	requestUrl := r.URL.String()
	filePath := requestUrl[len(Pattern):]
	if utils.ComparePath(filePath, "onlinePreview") {
		r.ParseForm()
		submitUrl := r.Form.Get("url")
		fileUrl, err := url.QueryUnescape(submitUrl)
		if err != nil {
			w.Write([]byte("URL解码出现错误"))
			return
		}
		submitType := r.Form.Get("type")
		if filePath, err := download.DownloadFile(fileUrl, "cache/download/"+path.Base(fileUrl)); err == nil {
			if submitType == "pdf" && (path.Ext(filePath) == ".pdf" || utils.IsInArr(path.Ext(filePath), AllOfficeEtx)) { //预留的PDF预览接口
				if path.Ext(filePath) == ".pdf" {
					dataByte := pdfPageDownload("cache/download/" + path.Base(filePath))
					w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
					w.Header().Set("content-type", "text/html;charset=UTF-8")
					w.Write([]byte(dataByte))
					setFileMap(path.Base(filePath))
				} else {
					if pdfPath := utils.ConvertToPDF(filePath); pdfPath != "" {
						dataByte := pdfPage("cache/pdf/" + path.Base(pdfPath))
						w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
						w.Header().Set("content-type", "text/html;charset=UTF-8")
						w.Write([]byte(dataByte))
						setFileMap(path.Base(filePath))
					} else {
						w.Write([]byte("转换为PDF时出现错误!"))
					}
				}
			} else if utils.IsInArr(path.Ext(filePath), AllImageEtx) {
				dataByte := imagePage(filePath)
				w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
				w.Header().Set("content-type", "text/html;charset=UTF-8")
				w.Write([]byte(dataByte))
			} else if utils.IsInArr(path.Ext(filePath), AllOfficeEtx) {
				if isHave(path.Base(filePath)) {
					dataByte := officePage("cache/convert/" + strings.Split(path.Base(filePath), ".")[0])
					w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
					w.Header().Set("content-type", "text/html;charset=UTF-8")
					w.Write([]byte(dataByte))
					return
				}
				if pdfPath := utils.ConvertToPDF(filePath); pdfPath != "" {
					if imgPath := utils.ConvertToImg(pdfPath); imgPath != "" {
						dataByte := officePage(imgPath)
						w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
						w.Header().Set("content-type", "text/html;charset=UTF-8")
						w.Write([]byte(dataByte))
						setFileMap(path.Base(filePath))
					} else {
						w.Write([]byte("转换为图片时出现错误!"))
					}
				} else {
					w.Write([]byte("转换为PDF时出现错误!"))
				}
			} else if path.Ext(filePath) == ".pdf" {
				if isHave(path.Base(filePath)) {
					dataByte := officePage("cache/convert/" + strings.Split(path.Base(filePath), ".")[0])
					w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
					w.Header().Set("content-type", "text/html;charset=UTF-8")
					w.Write([]byte(dataByte))
					return
				}
				if imgPath := utils.ConvertToImg(filePath); imgPath != "" {
					dataByte := officePage(imgPath)
					w.Header().Set("content-length", strconv.Itoa(len(dataByte)))
					w.Header().Set("content-type", "text/html;charset=UTF-8")
					w.Write([]byte(dataByte))
					setFileMap(path.Base(filePath))
				} else {
					w.Write([]byte("转换为图片时出现错误!"))
				}
			}
		} else {
			log.Println("Error: <", err, "> when download file")
			w.Write([]byte("获取文件失败...请检查你的路径是否正确!"))
		}
	} else if utils.ComparePath(filePath, "img_asset") {
		imgPath := requestUrl[len(Pattern+"img_asset"):]
		DataByte, err := ioutil.ReadFile("cache/download/" + imgPath)
		if err != nil {
			w.Header().Set("content-length", strconv.Itoa(len("404")))
			w.Header().Set("content-type", "text/html;charset=UTF-8")
			w.Write([]byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
		w.Header().Set("content-length", strconv.Itoa(len(DataByte)))
		w.Write(DataByte)
	} else if utils.ComparePath(filePath, "office_asset") {
		imgPath := requestUrl[len(Pattern+"office_asset"):]
		DataByte, err := ioutil.ReadFile("cache/convert/" + imgPath)
		if err != nil {
			w.Header().Set("content-length", strconv.Itoa(len("404")))
			w.Header().Set("content-type", "text/html;charset=UTF-8")
			w.Write([]byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
		w.Header().Set("content-length", strconv.Itoa(len(DataByte)))
		w.Write(DataByte)
	} else if utils.ComparePath(filePath, "pdf_asset") {
		pdfPath := requestUrl[len(Pattern+"pdf_asset"):]
		DataByte, err := ioutil.ReadFile("cache/pdf/" + pdfPath)
		if err != nil {
			w.Header().Set("content-length", strconv.Itoa(len("404")))
			w.Header().Set("content-type", "text/html;charset=UTF-8")
			w.Write([]byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
		w.Header().Set("content-length", strconv.Itoa(len(DataByte)))
		w.Header().Set("content-type", "application/pdf;charset=UTF-8")
		w.Write(DataByte)
	} else {
		DataByte, err := ioutil.ReadFile("html/" + filePath)
		if err != nil {
			w.Header().Set("content-length", strconv.Itoa(len("404")))
			w.Header().Set("content-type", "text/html;charset=UTF-8")
			w.Write([]byte("出现了一些问题,导致File View无法获取您的数据!"))
			return
		}
		w.Header().Set("content-length", strconv.Itoa(len(DataByte)))
		if path.Ext(filePath) == ".css" {
			w.Header().Set("content-type", "text/css;charset=UTF-8")
		} else if path.Ext(filePath) == ".js" {
			w.Header().Set("content-type", "application/x-javascript;charset=UTF-8")
		}
		w.Write(DataByte)
	}
}

func officePage(imgPath string) []byte {
	rd, _ := ioutil.ReadDir(imgPath)
	dataByte, _ := ioutil.ReadFile("html/office.html")
	dataStr := string(dataByte)
	htmlCode := ""
	for _, fi := range rd {
		if !fi.IsDir() {
			htmlCode = htmlCode + `<img class="my-photo" alt="loading" title="查看大图" style="cursor: pointer;"
									data-src="office_asset/` + path.Base(imgPath) + "/" + fi.Name() + `" src="images/loading.gif"
									">`
		}
	}
	dataStr = strings.Replace(dataStr, "{{AllImages}}", htmlCode, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func imagePage(filePath string) []byte {
	dataByte, _ := ioutil.ReadFile("html/image.html")
	dataStr := string(dataByte)
	imageUrl := "img_asset/" + path.Base(filePath)
	htmlCode := `<li>
					<img id="` + imageUrl + `" url="` + imageUrl + `"
						src="` + imageUrl + `" width="1px" height="1px">
				 </li>`
	dataStr = strings.Replace(dataStr, "{{AllImages}}", htmlCode, -1)
	dataStr = strings.Replace(dataStr, "{{FirstPath}}", imageUrl, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func pdfPage(filePath string) []byte {
	dataByte, _ := ioutil.ReadFile("html/pdf.html")
	dataStr := string(dataByte)
	pdfUrl := "pdf_asset/" + path.Base(filePath)
	dataStr = strings.Replace(dataStr, "{{url}}", pdfUrl, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func pdfPageDownload(filePath string) []byte {
	dataByte, _ := ioutil.ReadFile("html/pdf.html")
	dataStr := string(dataByte)
	pdfUrl := "img_asset/" + path.Base(filePath)
	dataStr = strings.Replace(dataStr, "{{url}}", pdfUrl, -1)
	dataByte = []byte(dataStr)
	return dataByte
}

func isHave(fileName string) bool {
	fileName = strings.Split(fileName, ".")[0]
	if _, ok := AllFile[fileName]; ok {
		AllFile[fileName].LastActiveTime = time.Now().Unix()
		return true
	} else {
		return false
	}
}

func setFileMap(fileName string) {
	ext := path.Ext(fileName)
	fileName = strings.Split(fileName, ".")[0]
	if _, ok := AllFile[fileName]; ok {
		AllFile[fileName].LastActiveTime = time.Now().Unix()
		return
	} else {
		temp := &NowFile{
			Md5:            fileName,
			Ext:            ext,
			LastActiveTime: time.Now().Unix(),
		}
		AllFile[fileName] = temp
	}
}

func Monitor() {
	log.Println("Info: Starting Monitor Thread")
	for {
		for _, v := range AllFile {
			if time.Now().Unix()-v.LastActiveTime > ExpireTime {
				if v.Md5 != "" {
					os.RemoveAll("cache/convert/" + v.Md5)
					os.Remove("cache/download/" + v.Md5 + v.Ext)
					os.Remove("cache/pdf/" + v.Md5 + ".pdf")
					log.Println("Cache file ", v.Md5, " delete")
					delete(AllFile, v.Md5)
				} else {
					delete(AllFile, v.Md5)
					log.Println("Cache file ", v.Md5, " delete with error")
				}
			}
		}
		time.Sleep(time.Second * 60)
	}
}

func StartServer() {
	http.HandleFunc(Pattern, Handle)
	log.Println("Info: Go File View Listening Address: " + Address + " on " + Pattern)
	if err := http.ListenAndServe(Address, nil); err != nil {
		log.Fatal("Error: <", err, "> when StartServer")
	}
}

func Init(pattern string, address string) {
	Pattern = pattern
	Address = address
	AllFile = make(map[string]*NowFile)
	ExpireTime = 3600
	go Monitor()
}
