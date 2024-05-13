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
	"fmt"
	"net/http"

	"github.com/ditwrd/wedinv/internal/models"
	"github.com/ditwrd/wedinv/template"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
)

var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

type inviteeService struct {
	dao *daos.Dao
}

func NewInviteeService(dao *daos.Dao) inviteeService {
	return inviteeService{dao: dao}
}

func (i inviteeService) FindInvitation(ctx echo.Context) error {
	findRequest := models.FindInvitationRequest{}
	err := ctx.Bind(&findRequest)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	err = validate.Struct(&findRequest)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	invitee := models.Invitee{}
	err = i.dao.DB().
		NewQuery("select id, name, status, invited_by from invitee where id={:invitationID} limit 1").
		Bind(dbx.Params{"invitationID": findRequest.ID}).
		One(&invitee)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	return template.Render(ctx, http.StatusOK, template.Index(invitee))
}

func (i inviteeService) ConfirmInvitation(ctx echo.Context) error {
	confirmRequest := models.ConfirmInvitationRequest{}
	err := ctx.Bind(&confirmRequest)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	err = validate.Struct(&confirmRequest)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	err = i.dao.RunInTransaction(func(txDao *daos.Dao) error {
		_, err = txDao.DB().
			NewQuery("update invitee set status={:status} where invitee.id={:invitationID}").
			Bind(dbx.Params{"invitationID": confirmRequest.ID, "status": confirmRequest.Answer}).Execute()
		return err
	})
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/notfound")
		return err
	}

	template.Render(ctx, -1, template.AnswerButton(confirmRequest.ID, "accept", confirmRequest.Answer))
	template.Render(ctx, -1, template.AnswerButton(confirmRequest.ID, "decline", confirmRequest.Answer))
	return template.Render(ctx, http.StatusOK, template.InvitationStatus(confirmRequest.Answer))
}
