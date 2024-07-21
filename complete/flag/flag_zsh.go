package flag

import (
	"complete/internal"
	"strings"
)

func (f *Flag) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	var addFuncs []*internal.AddFuncZsh

	builderArguments.WriteString(" \\\n")
	builderArguments.WriteString(indent)
	builderArguments.WriteString("+ '(")
	builderArguments.WriteString(f.Long)
	builderArguments.WriteString(")'")

	if f.Short != 0 {
		builderArguments.WriteString(" \\\n")
		builderArguments.WriteString(indent)
		builderArguments.WriteRune('"')

		builderArguments.WriteRune('-')
		builderArguments.WriteRune(f.Short)
		if f.Help != "" {
			builderArguments.WriteRune('[')
			builderArguments.WriteString(f.Help)
			builderArguments.WriteRune(']')
		}
		for _, arg := range f.Args {
			addFuncs = append(addFuncs, arg.GenerateZsh(builderArguments, indent, progName)...)
		}

		builderArguments.WriteRune('"')
	}

	if f.Long != "" {
		builderArguments.WriteString(" \\\n")
		builderArguments.WriteString(indent)
		builderArguments.WriteRune('"')

		builderArguments.WriteString("--")
		builderArguments.WriteString(f.Long)
		if f.Help != "" {
			builderArguments.WriteRune('[')
			builderArguments.WriteString(f.Help)
			builderArguments.WriteRune(']')
		}
		for _, arg := range f.Args {
			addFuncs = append(addFuncs, arg.GenerateZsh(builderArguments, indent, progName)...)
		}

		builderArguments.WriteRune('"')
	}
	return addFuncs
}
