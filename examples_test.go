package gomodexample_test

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"
	"time"

	"github.com/parro-it/gomodexample"
)

//go:embed fixtures
var fixtureRootFS embed.FS
var fixtureFS, _ = fs.Sub(fixtureRootFS, "fixtures")

// This example show how to use gomodexample.Func()
func ExampleFunc() {
	fmt.Println(gomodexample.Func())
	// Output: 42
}
