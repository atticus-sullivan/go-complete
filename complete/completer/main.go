package completer

import "complete/complete/types"

type MainCompleter struct {
	types.Completer
	Opts CompleterOpts
}

type CompleterOpts struct {
	OptionStacking     bool
	OptionStackingArgs bool
	DashSep            bool
}
