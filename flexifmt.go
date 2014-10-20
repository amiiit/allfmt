package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/voxelbrain/goptions"
)

var (
	options = struct {
		Dartfmt string `goptions:"-d, --dart, description='Path to dartfmt'"`
		Gofmt   string `goptions:"-g, --go, description='Path to gofmt'"`
		Jsfmt   string `goptions:"-j, --js, description='Path to jsfmt'"`
		Input   string `goptions:"-i, --input, description='Input source', obligatory"`
	}{
		Dartfmt: "dartfmt",
		Gofmt:   "gofmt",
		Jsfmt:   "jsfmt",
	}
)

func standardFmt(file string, formatter string) {
	exec.Command(formatter, "-w", file).Run()
}

func main() {
	goptions.ParseAndFail(&options)

	split := strings.Split(options.Input, ".")

	filetype := split[len(split)-1]

	switch filetype {
	case "js":
		standardFmt(options.Input, options.Jsfmt)
	case "dart":
		standardFmt(options.Input, options.Dartfmt)
	case "go":
		standardFmt(options.Input, options.Gofmt)
	default:
		fmt.Println("unknown type")
	}
}
