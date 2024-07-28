package completeargs

import (
	"github.com/atticus-sullivan/go-complete/internal"
	"strings"
	sh "mvdan.cc/sh/v3/syntax"
)

func (cf CTfile) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*internal.AddFuncBash, []*internal.AddIfBash) {
	builderArguments.WriteString(indent)
	if cf.OnlyDirs {
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -d`)
	} else {
		builderArguments.WriteString(`    COMPREPLY+=( $(compgen -f`)
	}

	s,err := sh.Quote(cf.Glob, sh.LangBash)
	if cf.Glob != "" && err != nil {
		builderArguments.WriteString(" -X '!")
		builderArguments.WriteString(s)
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
