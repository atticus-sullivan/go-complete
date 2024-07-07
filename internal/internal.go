package internal

import "strings"

type AddFuncZsh struct {
	Cas func(*strings.Builder, string)
	Fun func(*strings.Builder, string) []*AddFuncZsh
}

type AddIfBash struct {
	Fun func(*strings.Builder, string)
}

type AddFuncBash struct {
	Fun func(*strings.Builder, string) []*AddFuncBash
}
