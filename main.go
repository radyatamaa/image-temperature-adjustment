package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/radyatamaa/image-temperature-adjustment/internal"

	beego "github.com/beego/beego/v2/server/web"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/beego/i18n"
	"github.com/radyatamaa/image-temperature-adjustment/internal/middlewares"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/response"
	"github.com/radyatamaa/image-temperature-adjustment/pkg/zaplogger"

	imageAdjustmentHandler "github.com/radyatamaa/image-temperature-adjustment/internal/image_adjustment/delivery/http/v1"
	imageAdjustmentUsecase "github.com/radyatamaa/image-temperature-adjustment/internal/image_adjustment/usecase"
)

// @title Api Gateway V1
// @version v1
// @contact.name radyatama
// @contact.email mohradyatama24@gmail.com
// @description api "API Gateway v1"
// @BasePath /api
// @query.collection.format multi

func main() {
	err := beego.LoadAppConfig("ini", "conf/app.ini")
	if err != nil {
		panic(err)
	}
	// global execution timeout
	serverTimeout := beego.AppConfig.DefaultInt64("serverTimeout", 60)
	// global execution timeout
	requestTimeout := beego.AppConfig.DefaultInt("executionTimeout", 5)
	// global execution timeout to second
	timeoutContext := time.Duration(requestTimeout) * time.Second
	// web hook to slack error log
	slackWebHookUrl := beego.AppConfig.DefaultString("slackWebhookUrlLog", "")
	// app version
	appVersion := beego.AppConfig.DefaultString("version", "1")
	// log path
	logPath := beego.AppConfig.DefaultString("logPath", "./logs/api.log")


	// language
	lang := beego.AppConfig.DefaultString("lang", "en|id")
	languages := strings.Split(lang, "|")
	for _, value := range languages {
		if err := i18n.SetMessage(value, "./conf/"+value+".ini"); err != nil {
			panic("Failed to set message file for l10n")
		}
	}

	// beego config
	beego.BConfig.Log.AccessLogs = false
	beego.BConfig.Log.EnableStaticLogs = false
	beego.BConfig.Listen.ServerTimeOut = serverTimeout

	// zap logger
	zapLog := zaplogger.NewZapLogger(logPath, slackWebHookUrl)

	err = os.Mkdir("external/storage", 0755)
	if err != nil && !strings.Contains(err.Error(),"Cannot create a file when that file already exists."){
		fmt.Println(err.Error())
		panic(err)
	}

	if beego.BConfig.RunMode == "dev" {
		// static files swagger
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}


	if beego.BConfig.RunMode != "prod" {
		// static files swagger
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.BConfig.WebConfig.StaticDir["/external/storage"] = "external/storage"

	// middleware init
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:    []string{http.MethodGet, http.MethodPost},
		AllowAllOrigins: true,
	}))


	beego.InsertFilterChain("*", middlewares.RequestID())
	beego.InsertFilterChain("/api/*", middlewares.BodyDumpWithConfig(middlewares.NewAccessLogMiddleware(zapLog, appVersion).Logger()))

	// health check
	beego.Get("/health", func(ctx *beegoContext.Context) {
		ctx.Output.SetStatus(http.StatusOK)
		ctx.Output.JSON(beego.M{"status": "alive"}, beego.BConfig.RunMode != "prod", false)
	})

	// default error handler
	beego.ErrorController(&response.ErrorController{})

	// init repository


	// init usecase
	imageAdjustmentUseCase := imageAdjustmentUsecase.NewImageAdjustmentUseCase(timeoutContext, zapLog)

	// init handler
	imageAdjustmentHandler.NewImageAdjustmentHandler(imageAdjustmentUseCase, zapLog)

	// default error handler
	beego.ErrorController(&internal.BaseController{})

	beego.Run()
}
