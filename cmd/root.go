/*
Copyright © 2024 Aditya Wardianto

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/ditwrd/wedinv/internal/services"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

func Execute() {
	pb := pocketbase.New()

	migratecmd.MustRegister(pb, pb.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Static("/public", "public")
		e.Router.GET("/notfound", func(c echo.Context) error {
			c.HTML(404, "You shouldn't be here")
			return nil
		})
		e.Router.RouteNotFound("/*", func(c echo.Context) error {
			c.Redirect(http.StatusTemporaryRedirect, "/missing")
			return nil
		})

		inviteeService := services.NewInviteeService(pb.Dao())

		baseRouter := e.Router.Group("/inv")
		baseRouter.GET("/:invitationID", inviteeService.FindInvitation)
		baseRouter.PUT("/:invitationID/accept", inviteeService.AcceptInvitation)
		baseRouter.PUT("/:invitationID/decline", inviteeService.DeclineInvitation)

		return nil
	})

	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}
