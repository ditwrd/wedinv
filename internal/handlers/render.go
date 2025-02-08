package handlers

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(c templ.Component, ctx context.Context, w http.ResponseWriter) error {
	return c.Render(ctx, w)
}
