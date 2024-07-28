package completeargs

import (
	"strings"

	"github.com/atticus-sullivan/go-complete/internal"
	sh "mvdan.cc/sh/v3/syntax"
)

func (cf CTfile) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(":path:")
	if cf.OnlyDirs {
		builderArguments.WriteString("_files -/")
	} else {
		builderArguments.WriteString("_files -f")
	}
	s,err := sh.Quote(cf.Glob, sh.LangBash)
	if cf.Glob != "" && err != nil {
		builderArguments.WriteString(" -g '")
		builderArguments.WriteString(s)
		builderArguments.WriteString("'")
	}
	return nil
}
