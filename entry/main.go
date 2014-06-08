package danielqsjblog

import (
	"Controllers"
	"github.com/astaxie/beegae"
	"github.com/astaxie/beego/middleware"
	"net/http"
)

func init() {
	beegae.Router("/", &Controllers.MainController{})
	beegae.Router("/index", &Controllers.UserController{})

	//beegae.SetStaticPath("/css", "views/css")
	//beegae.SetStaticPath("/js", "views/js")
	//beegae.SetStaticPath("/img", "views/img")
	//beegae.SetStaticPath("/font", "views/font")
	//beegae.SetStaticPath("/fonts", "views/fonts")
	//beegae.SetStaticPath("/images", "views/images")

	err := beegae.BuildTemplate(beegae.ViewsPath)
	if err != nil {
		panic(err)
	}

	middleware.VERSION = beegae.VERSION
	middleware.AppName = beegae.AppName
	middleware.RegisterErrorHandler()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		beegae.BeeApp.Handlers.ServeHTTP(w, r)
	})
}
