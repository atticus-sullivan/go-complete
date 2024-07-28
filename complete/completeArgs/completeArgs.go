package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
)

// defines what can be completed (usually used for arguments) includes subcommands
type CompleteTypeSub interface {
	// generate the code for the argument (for the _arguments function)
	// returns a slice of functions which generate stuff that also needs to be written (but at another position)
	GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh
	GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash)
	// if the argument is a positional one, it might need to know its index
	SetIdxS(idx uint) CompleteTypeSub
}

// defines what can be completed (usually used for arguments), excludes subcommands
type CompleteTypeNoSub interface {
	// generate the code for the argument (for the _arguments function)
	// returns a slice of functions which generate stuff that also needs to be written (but at another position)
	GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh
	GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash)
	// if the argument is a positional one, it might need to know its index
	SetIdxNS(idx uint) CompleteTypeNoSub
	// just to restrict what type implements this interface
	isNoSub()
}
