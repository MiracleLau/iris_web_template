package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/utahta/go-cronowriter"
	"os"
	"time"
	"web_template/config"
	"web_template/db"
)

var App *iris.Application

func InitApp() {

	App = iris.New()
	// 一些目录设置
	App.HandleDir("/static", "static")

	// 设置模板引擎
	App.RegisterView(iris.Django("./views", ".html").Reload(true))

	// 中间件设置
	App.Use(recover.New())

	session := sessions.New(sessions.Config{
		Expires: time.Hour * 2,
		Cookie:  config.Config.SessionName,
	})
	session.UseDatabase(db.NewRedis())
	App.Use(session.Handler())
}

func SetApplicationLog(path string) {
	w1 := cronowriter.MustNew(path +"/iris.log.%Y%m%d")
	App.Logger().SetOutput(w1)
	App.Logger().AddOutput(os.Stdout)
}

func InitRoute() {

}

func Start() {
	App.Listen(config.Config.SiteAddress)
}
