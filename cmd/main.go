package main

import (
	"log"

	"github.com/ditwrd/wedinv/internal"
	"github.com/ditwrd/wedinv/internal/handlers"
	_ "github.com/ditwrd/wedinv/internal/migrations"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	invitationHandler := handlers.NewInvitationHandler(app)

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/{path...}", apis.Static(internal.StaticFiles, false))
		se.Router.GET("/home", handlers.Home)
		se.Router.GET("/inv/{id}", invitationHandler.GetInviteeData)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
