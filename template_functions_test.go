package view_test

import (
	"io/fs"
	"strings"
	"testing"
	"time"

	"github.com/maddiesch/go-view"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInGroupsOf(t *testing.T) {
	t.Run("when there are more elements than the group size", func(t *testing.T) {
		v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		vv := view.InGroupsOf(v, 3)

		assert.Equal(t, [][]any{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10}}, vv)
	})

	t.Run("when there are less elements than the group size", func(t *testing.T) {
		v := []int{1, 2, 3, 4}

		vv := view.InGroupsOf(v, 5)

		assert.Equal(t, [][]any{{1, 2, 3, 4}}, vv)
	})
}

func TestTemplateFunction(t *testing.T) {
	templates, _ := fs.Sub(templateFS, "templates")
	r := view.NewTemplateRenderer(templates)

	data := struct {
		NumberList  []int
		UpdatedTime time.Time
		SourceTime  time.Time
	}{
		NumberList:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		SourceTime:  time.Unix(1694718675, 0),
		UpdatedTime: time.Unix(1694718600, 0),
	}

	var buf strings.Builder

	err := r.RenderWithLayout(&buf, "pages/functions.html", data)
	require.NoError(t, err)

	t.Run("inGroupsOf", func(t *testing.T) {
		assert.Contains(t, buf.String(), `<p>Group:1,2,3,</p>`)
		assert.Contains(t, buf.String(), `<p>Group:4,5,6,</p>`)
		assert.Contains(t, buf.String(), `<p>Group:7,8,9,</p>`)
		assert.Contains(t, buf.String(), `<p>Group:10,</p>`)
	})

	t.Run("timeSince", func(t *testing.T) {
		assert.Contains(t, buf.String(), `<h1>timeSince</h1>1m15s`)
	})
}
