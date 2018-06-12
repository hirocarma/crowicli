// Package client for Crowi and Crowi-plus.
package client

import (
	"flag"

	a "github.com/hirocarma/crowicli/mode"
)

// Opt has cmdline argments.
var Opt struct {
	Mode string
	Path string
	File string
}

func flagparse() {
	flag.StringVar(&Opt.Mode, "m", "list", "mode:create,update,get,list")
	flag.StringVar(&Opt.Path, "p", "", "wiki path")
	flag.StringVar(&Opt.File, "f", "./file.md", "markdown file")
	flag.Parse()
}

func help() string {
	return "mode is invalid. [create|update|get|list]\n"
}

var modes = map[string]a.Mode{}

// Register sets command to commands.
// http://blog.engineer.adways.net/entry/2017/07/14/120000
func Register(key string, mode a.Mode) {
	modes[key] = mode
}

// Run calls commands.
func Run() (string, error) {
	flagparse()
	err := Env()
	if err != nil {
		return "", err
	}

	mode, ok := modes[Opt.Mode]
	if ok {
		return mode.API()
	}
	return help(), nil
}
