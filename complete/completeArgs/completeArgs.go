package completeargs

import (
	"complete/internal"
	"strings"
)

// defines what can be completed (usually used for arguments)
type CompleteType interface {
	// generate the code for the argument (for the _arguments function)
	// returns a slice of functions which generate stuff that also needs to be written (but at another position)
	GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh
	GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash)
}
