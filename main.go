/*package main

import(
	"github.com/leeli73/goFileView/perview"
)

func main() {
	perview.Init("/perview/","0.0.0.0:8089")
	perview.StartServer()
}*/

package main

import (
	"net/http"

	"github.com/leeli73/goFileView/perview"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<!DOCTYPE html>
	<html>
	
	<head>
	  <title>goFileView测试服务</title>
	  <meta charset="utf-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1">
	  <link rel="stylesheet" href="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/css/bootstrap.min.css">
	  <script src="https://cdn.staticfile.org/jquery/3.2.1/jquery.min.js"></script>
	  <script src="https://cdn.staticfile.org/popper.js/1.15.0/umd/popper.min.js"></script>
	  <script src="https://cdn.staticfile.org/twitter-bootstrap/4.3.1/js/bootstrap.min.js"></script>
	</head>
	
	<body>
	
	  <div class="container">
		<h2>goFileView</h2>
		<p>测试的预览服务</p>
		<table class="table">
		  <thead>
			<tr>
			  <th>文件名</th>
			  <th>文件类型</th>
			  <th>调用接口</th>
			  <th>预览</th>
			</tr>
		  </thead>
		  <tbody>
			<tr>
			  <td>test.docx</td>
			  <td>docx</td>
			  <td>http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.docx</td>
			  <td>
				<a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.docx">普通预览</a>
				<a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.docx&type=pdf">PDF预览</a>
			  </td>
			</tr>
			<tr>
			  <td>test.xlsx</td>
			  <td>xlsx</td>
			  <td>http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.xlsx</td>
			  <td><a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.xlsx">普通预览</a>
				  <a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.xlsx&type=pdf">PDF预览</a>
			  </td>
			</tr>
			<tr>
			  <td>test.pptx</td>
			  <td>pptx</td>
			  <td>http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pptx</td>
			  <td><a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pptx">普通预览</a>
				  <a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pptx&type=pdf">PDF预览</a>
			  </td>
			</tr>
			<tr>
			  <td>test.pdf</td>
			  <td>pdf</td>
			  <td>http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pdf</td>
			  <td><a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pdf">普通预览</a>
				  <a
				  href="http://gofileview.onlinecode.cn/perview/onlinePreview?url=http://onlinecode.cn/test.pdf&type=pdf">PDF预览</a>
			  </td>
			</tr>
		  </tbody>
		</table>
	  </div>
	
	</body>
	
	</html>`))
}

func main() {

	perview.Init("/perview/", "no") //初始化

	http.HandleFunc("/", index)

	http.HandleFunc("/perview/", perview.Handle) //绑定到preview的Handle

	http.ListenAndServe(":80", nil)

}
