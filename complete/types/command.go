package types

import (
	"github.com/atticus-sullivan/go-complete/complete/flag"
	"github.com/atticus-sullivan/go-complete/complete/positional"
)

type Completer struct {
	Flags       []flag.Flag
	Positionals []positional.Positional
	Name        string
}
