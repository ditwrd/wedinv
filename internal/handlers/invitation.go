package handlers

import (
	"github.com/ditwrd/wedinv/internal/components"
	"github.com/ditwrd/wedinv/internal/models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type invitationHandler struct {
	app *pocketbase.PocketBase
	db  func() dbx.Builder
}

func NewInvitationHandler(app *pocketbase.PocketBase) invitationHandler {
	db := app.DB
	return invitationHandler{app: app, db: db}
}

func (i invitationHandler) GetInviteeData(e *core.RequestEvent) error {
	req := e.Request
	res := e.Response

	id := req.PathValue("id")

	invitee := models.Invitee{}

	err := i.db().NewQuery("select id, name from invitee where id = {:id}").
		Bind(dbx.Params{"id": id}).
		One(&invitee)

	println(err)
	if err != nil {
		println(err)
		return err
	}

	return Render(components.Invitation(invitee), req.Context(), res)
}
