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

package invitee

import (
	"net/http"

	m "github.com/ditwrd/wedinv/internal/models"
	t "github.com/ditwrd/wedinv/template"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/daos"
)

func FindInvitation(pb *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		invitationID := c.PathParam("invitationID")

		invitee := m.Invitee{}

		err := pb.Dao().DB().
			NewQuery("select id, name, status, invited_by from invitee where id={:invitationID} limit 1").
			Bind(dbx.Params{"invitationID": invitationID}).
			One(&invitee)
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/missing")
			return err
		}

		return t.Render(c, http.StatusOK, t.Index(invitee))
	}
}

func AcceptInvitation(pb *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		invitationID := c.PathParam("invitationID")

		err := pb.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			_, err := txDao.DB().
				NewQuery("update invitee set status='accepted' where invitee.id={:invitationID}").
				Bind(dbx.Params{"invitationID": invitationID}).Execute()
			return err
		})
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/missing")
			return err
		}
		return t.Render(c, http.StatusOK, t.StatusCheck("accepted"))
	}
}

func DeclineInvitation(pb *pocketbase.PocketBase) func(echo.Context) error {
	return func(c echo.Context) error {
		invitationID := c.PathParam("invitationID")

		err := pb.Dao().RunInTransaction(func(txDao *daos.Dao) error {
			_, err := txDao.DB().
				NewQuery("update invitee set status='declined' where invitee.id={:invitationID}").
				Bind(dbx.Params{"invitationID": invitationID}).Execute()
			return err
		})
		if err != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/missing")
			return err
		}
		return t.Render(c, http.StatusOK, t.StatusCheck("declined"))
	}
}
