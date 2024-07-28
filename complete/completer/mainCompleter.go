package completer

import "github.com/atticus-sullivan/go-complete/complete/types"

type MainCompleter struct {
	types.Completer
	Opts CompleterOpts
}

type CompleterOpts struct {
	OptionStacking     bool
	OptionStackingArgs bool
	DashSep            bool
	ToFile             bool
}
