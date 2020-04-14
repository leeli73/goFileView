package main

import(
	"github.com/leeli73/goFileView/perview"
)

func main() {
	perview.Init("/perview/","0.0.0.0:8089")
	perview.StartServer()
}