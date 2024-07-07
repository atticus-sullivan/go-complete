package positional

import (
	"complete/internal"
	"fmt"
	"strings"
)

func (p *Positional) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	add := make([]*internal.AddFuncZsh, 0)

	builderArguments.WriteString(" \\\n")
	builderArguments.WriteString(indent)
	builderArguments.WriteRune('"')

	if p.Optional {
		fmt.Fprintf(builderArguments, "%d::", p.Idx+1)
	} else {
		fmt.Fprintf(builderArguments, "%d:", p.Idx+1)
	}
	// help must default to " " so that :: vs : can be differentiated
	if p.Help == "" {
		p.Help = " "
	}
	fmt.Fprintf(builderArguments, "%s:", p.Help)

	add = append(add, p.Arg.GenerateZsh(builderArguments, indent, progName)...)
	builderArguments.WriteRune('"')

	return add
}
