package cli

import (
	"fmt"
	"os"
)

func Run() {
	err := rootCommand.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
