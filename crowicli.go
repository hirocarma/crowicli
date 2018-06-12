package main

import (
	"fmt"
	"os"

	"github.com/hirocarma/crowicli/client"
	_ "github.com/hirocarma/crowicli/mode/create"
	_ "github.com/hirocarma/crowicli/mode/get"
	_ "github.com/hirocarma/crowicli/mode/list"
	_ "github.com/hirocarma/crowicli/mode/update"
)

func main() {
	res, err := client.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(res)
}
