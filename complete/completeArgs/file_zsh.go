package completeargs

import (
	"complete/internal"
	"strings"
)

func (cf CTfile) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*internal.AddFuncZsh {
	builderArguments.WriteString(":path:")
	if cf.OnlyDirs {
		builderArguments.WriteString("_files -/")
	} else {
		builderArguments.WriteString("_files -f")
	}
	if cf.Glob != "" {
		builderArguments.WriteString(" -g ")
		builderArguments.WriteString(cf.Glob)
	}
	return nil
}
