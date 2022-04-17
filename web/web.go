package web

import (
	"context"
	_ "embed"
	"html/template"

	"github.com/gobuffalo/plush/v4"
)

//go:embed layout.plush.html
var layout string

func Render(ctx context.Context, content string, data map[string]interface{}) (string, error) {
	pctx := plush.NewContextWithContext(ctx)
	for k, v := range data {
		pctx.Set(k, v)
	}

	content, err := plush.Render(content, pctx)
	if err != nil {
		return "", err
	}

	pctx.Set("yield", template.HTML(content))

	return plush.Render(layout, pctx)
}
