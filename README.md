go File View
============

Go File
View是受kkFileView(https://gitee.com/kekingcn/file-online-preview.git)启发并基于其网站前端开发的。目前goFileView处于最原始的起步状态，相对简陋，相信随着不断完善成为一套强壮的系统。本人代码风格相对较”狂”，欢迎大家一起来提出建议和完善Go
File View。

特别要感谢kkfileview的开源，让我可以使用它的前端页面直接开发。调用方式也在很大程度上参考了kkfileview。

目前仅支持Linux系统，因为时间仓促，从我有想法到写出这个beta版，只有半天时间，所以可能有很多问题，比如url不支持中文等问题。请大家见谅，我目前也只是在Win10
WSL Ubuntu里完成了测试。

![](media/356e96c952b74e31f28357899547ba49.png)

![](media/e9ac0e8245cbca32fcc8da292f9f935e.png)

上面是预览效果(顺便给我自己打打广告,手动滑稽)

目前已经完成
============

Word、Excel、PPT转码为PDF

PDF转码为图片

对Word,Excel,PPT和PDF的图片式在线预览

未来
====

PDF文件直接在线预览

PDF转SVG矢量图形

多文件的接受

ftp、xftp、scp等文件传输形式的兼容

内置Fire Server

本地路径指定，省去下载步骤

部署编译
========

准备
----

安装Libreoffice

安装convert

确保Libreoffice和conver都在path目录下

编译
----

git clone <https://github.com/leeli73/goFileView.git>

cd goFileView

go build main.go

在自己的项目中集成
==================

准备
----

go get github.com/leeli73/goFileView

demo
----

package main

import(

"net/http"

"github.com/leeli73/goFileView/perview"

)

func index(w http.ResponseWriter, r \*http.Request) {

w.Write([]byte("I'm Index"))

}

func main(){

perview.Init("/perview/","no") //初始化

http.HandleFunc("/index",index)

http.HandleFunc("/perview/",perview.Handle) //绑定到preview的Handle

http.ListenAndServe(":80", nil)

}
