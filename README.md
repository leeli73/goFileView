go File View
============

Go File View是受kkFileView( *https://gitee.com/kekingcn/file-online-preview.git* )启发并基于其网站前端开发的。目前goFileView处于最原始的起步状态，相对简陋，相信随着不断完善成为一套强壮的系统。本人代码风格相对较”狂”，欢迎大家一起来提出建议和完善Go File View。

特别要感谢kkfileview的开源，让我可以使用它的前端页面直接开发。调用方式也在很大程度上参考了kkfileview。

从我有想法到写出这个beta版，只有半天时间，所以可能有很多问题，目前我也即将毕业，所以很少有时间能维护goFileView，如果您对goFileView有想法或者建议欢迎在issue中提问，我看到后会尽快完善的。

![](media/mainshow.png)

上面是goFileView的预览效果(顺便给我自己打打广告,手动滑稽)

免费预览测试服务
====

本人刚好有一个空闲的服务器，所以打算将其贡献出来，以供大家测试和预览goFileView

这个服务可能随时取消，请大家不要将其引入自己的项目，以免造成意外的错误

接口 http://gofileview.onlinecode.cn/perview/onlinePreview?url= 被预览文件的url

你可以直接访问 http://gofileview.onlinecode.cn/ 使用预置的文件查看goFileView的效果

更新
====

2019年8月

    1.发布第一个版本
    2.完成在Windows10 WSL中的运行

2019年9月

    1.完成Ubuntu的完美兼容

2020年4月

    1.新增了对Windows系统的支持

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

Windows版
----

    准备

        1.安装Libreoffice,下载官方msi包,傻瓜式安装即可 (https://zh-cn.libreoffice.org)

        2.将Libreoffice安装路径下的program文件夹加入PATH中
![](media/win_path.png)

        3.安装ImageMagick,官方包,傻瓜式安装即可,安装7.0以上版本 (https://ghostscript.com/download/gsdnld.html)

        4.安装GhostScript,官方包,傻瓜式安装即可 (https://ghostscript.com/download/gsdnld.html)

        5.git clone <https://github.com/leeli73/goFileView.git>
    
    编译

        1.cd goFileView
        2.go build

    运行

        1. goFileView.exe
        2. 访问 http://127.0.0.1:8089/perview/onlinePreview?url=被预览文件的url (例如 http://127.0.0.1:8089/perview/onlinePreview?url=http://127.0.0.1:88/test.docx)

    你可以在代码中修改监听的URL、端口等信息。

Linux版
----

    准备

        1.安装Libreoffice:sudo apt install libreoffice

        2.安装ImageMagick:sudo apt install imagemagick

        4.修改ImageMagick的配置,vi etc/ImageMagick-6/policy.xml

            修改
            <policy domain="coder" rights="none" pattern="PDF" />
            为
            <policy domain="coder" rights="read|write" pattern="PDF" />
            下方新增一行
            <policy domain="coder" rights="read|write" pattern="LABEL" />

            wq退出保存

        5.安装字体(如果出现乱码)

            打包一台Windows机器的C:\Windows\Fonts下的所有文件
            发送到Linux机器上
            解压后进入Fonts文件夹，依次执行mkfontscale,mkfontdir,fc-cache

        5.git clone <https://github.com/leeli73/goFileView.git>
    
    编译

        1.cd goFileView
        2.go build

    运行

        1. ./goFileView
        2. 访问 http://127.0.0.1:8089/perview/onlinePreview?url=被预览文件的url (例如 http://127.0.0.1:8089/perview/onlinePreview?url=http://127.0.0.1:88/test.docx)

    你可以再代码中修改监听的URL、端口等信息。

在自己的项目中集成
==================

准备
----

    go get github.com/leeli73/goFileView

demo
----
```
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
```