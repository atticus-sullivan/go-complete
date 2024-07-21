package completeargs

import (
	"complete/internal"
	"strings"
)

func (cf CTfile) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	builderArguments.WriteString(indent)
	if cf.OnlyDirs {
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -d`)
	} else {
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -f`)
	}

	if cf.Glob != "" {
		builderArguments.WriteString(" -X '!")
		builderArguments.WriteString(cf.Glob)
		builderArguments.WriteRune('\'')
	}

	builderArguments.WriteString(` -- "${cur}") )`)
	builderArguments.WriteRune('\n')

	if cf.Glob != "" {
		builderArguments.WriteString(indent)
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -S / -d -- "${cur}") )`)
		builderArguments.WriteRune('\n')
	}

	return nil,nil
}
