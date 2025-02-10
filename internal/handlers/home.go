package handlers

import (
	"github.com/ditwrd/wedinv/internal/components"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	req := e.Request
	res := e.Response
	return Render(components.Home(), req.Context(), res)
}
