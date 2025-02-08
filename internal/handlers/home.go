package handlers

import (
	"github.com/ditwrd/wedinv/internal/components"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	Req := e.Request.Context()
	Res := e.Response
	return Render(components.Home(), Req, Res)
}
