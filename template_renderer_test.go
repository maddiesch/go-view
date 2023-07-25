package view_test

import (
	"bytes"
	"embed"
	"io/fs"
	"testing"

	"github.com/maddiesch/go-view"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed templates
var templateFS embed.FS

func TestTemplateRenderer(t *testing.T) {
	templates, _ := fs.Sub(templateFS, "templates")
	r := view.NewTemplateRenderer(templates)

	t.Run("RenderWithLayout", func(t *testing.T) {
		type content struct {
			Greeting string
		}

		var buf bytes.Buffer

		err := r.RenderWithLayout(&buf, "pages/landing.html", content{
			Greeting: "World",
		})

		require.NoError(t, err)

		assert.Contains(t, buf.String(), `<html lang="en">`)
		assert.Contains(t, buf.String(), "<h1>Hello World!</h1>")
	})

	t.Run("Render", func(t *testing.T) {
		type content struct {
			Message string
		}

		var buf bytes.Buffer

		err := r.Render(&buf, "404.html", content{
			Message: "The requested resource could not be found.",
		})

		require.NoError(t, err)

		assert.Contains(t, buf.String(), "<h1>Not Found</h1>")
		assert.Contains(t, buf.String(), "<p>The requested resource could not be found.</p>")
	})
}
