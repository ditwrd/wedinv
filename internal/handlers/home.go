package handlers

import (
	"github.com/ditwrd/wedinv/internal/components"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	return components.Render(components.Home(), e.Response)
}
