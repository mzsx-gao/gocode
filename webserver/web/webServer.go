/**
  author: gaoshudian
*/
package web

import (
	"fmt"
	"gocode/webserver/web/controller"
	"net/http"
)

func WebStart(app *controller.Application) {

	fs := http.FileServer(http.Dir("webserver/web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", app.IndexView)
	http.HandleFunc("/index.html", app.IndexView)
	http.HandleFunc("/setInfo.html", app.SetInfoView)
	http.HandleFunc("/setReq", app.SetInfo)
	http.HandleFunc("/queryReq", app.QueryInfo)

	fmt.Println("启动Web服务, 监听端口号: 9000")

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("启动Web服务错误")
	}

}
