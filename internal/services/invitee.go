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
package services

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/a-h/templ"
	"github.com/ditwrd/wedinv/internal/models"
	"github.com/ditwrd/wedinv/template"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

type inviteeService struct {
	dao *daos.Dao
}

func NewInviteeService(dao *daos.Dao) inviteeService {
	return inviteeService{dao: dao}
}

func (i inviteeService) render(ctx echo.Context, status int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(status)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (i inviteeService) validateInvitationID(id string) bool {
	r := regexp.MustCompile("^[a-zA-Z0-9]*$")
	return r.MatchString(id)
}

func (i inviteeService) FindInvitation(ctx echo.Context) error {
	invitationID := ctx.PathParam("invitationID")
	if !i.validateInvitationID(invitationID) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return errors.New("invalid invitation ID")
	}

	invitee := models.Invitee{}
	err := i.dao.DB().
		NewQuery("select id, name, status, invited_by from invitee where id={:invitationID} limit 1").
		Bind(dbx.Params{"invitationID": invitationID}).
		One(&invitee)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	return i.render(ctx, http.StatusOK, template.Index(invitee))
}

func (i inviteeService) updateStatusHandler(ctx echo.Context, status string) error {
	invitationID := ctx.PathParam("invitationID")
	if !i.validateInvitationID(invitationID) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return errors.New("invalid invitation ID")
	}
	err := i.dao.RunInTransaction(func(txDao *daos.Dao) error {
		_, err := txDao.DB().
			NewQuery("update invitee set status={:status} where invitee.id={:invitationID}").
			Bind(dbx.Params{"invitationID": invitationID, "status": status}).Execute()
		return err
	})
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}
	return i.render(ctx, http.StatusOK, template.StatusResponseBox(status, invitationID))
}

func (i inviteeService) AcceptInvitation(ctx echo.Context) error {
	err := i.updateStatusHandler(ctx, "accepted")
	return err
}

func (i inviteeService) DeclineInvitation(ctx echo.Context) error {
	err := i.updateStatusHandler(ctx, "declined")
	return err
}
