package types

import (
	"complete/complete/flag"
	"complete/complete/positional"
)

type Completer struct {
	Flags       []flag.Flag
	Positionals []positional.Positional
	Name        string
}
