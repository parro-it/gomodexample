package gomodexample

import (
	"embed"
	"io/fs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

func TestPeriod(t *testing.T) {
	content,err := fs.ReadFile(fixtureFS, "placeholder")
	if assert.NoError(t,err) {
		assert.Equal(t, "this is a placeholder", string(content))
	}
}

func TestFunc(t *testing.T) {
	assert.Equal(t, 42, Func())
}
