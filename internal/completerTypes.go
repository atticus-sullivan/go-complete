package internal

import (
	"strings"
)

// defines what can be completed (usually used for arguments)
type CompleteType interface {
	// generate the code for the argument (for the _arguments function)
	// returns a slice of functions which generate stuff that also needs to be written (but at another position)
	GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*AddFuncZsh
	GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*AddFuncBash, []*AddIfBash)
}

type CTfile struct{
	Glob string
	OnlyDirs bool
}
func (cf CTfile) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*AddFuncZsh {
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

func (cf CTfile) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*AddFuncBash, []*AddIfBash) {
	return nil,nil
}

type CTalternatives struct {
	Alts []string
}

func (ca CTalternatives) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*AddFuncZsh {
	builderArguments.WriteString(": :(")
	first := true
	for _,a := range ca.Alts {
		if !first {
			builderArguments.WriteRune(' ')
		}
		first = false
		builderArguments.WriteString(a)
	}
	builderArguments.WriteRune(')')
	return nil
}

func (ca CTalternatives) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*AddFuncBash, []*AddIfBash) {
	return nil,nil
}

type CTempty struct{}

func (ce CTempty) GenerateZsh(builderArguments *strings.Builder, indent string, progName string) []*AddFuncZsh {
	builderArguments.WriteString(": : ")
	return nil
}

func (ce CTempty) GenerateBash(builderArguments *strings.Builder, indent string, progName string) ([]*AddFuncBash, []*AddIfBash) {
	return nil,nil
}
