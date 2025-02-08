package components

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func Render(c templ.Component, w http.ResponseWriter) error {
	return c.Render(context.Background(), w)
}
