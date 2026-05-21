package models

import (
	"AceofHearts/config"
	"fmt"
)

// Banner 打印程序 Logo
func Banner() {
	fmt.Println(`
	` + config.Red("┌──────────┐") + `┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐
	` + config.Red("│ A        │") + `│││││││││││││││││││││
	` + config.Red("│          │") + `│││││││││││││││││││││
	` + config.Red("│    ❤     │") + `│││││││││││││││││││││
	` + config.Red("│          │") + `│││││││││││││││││││││
	` + config.Red("│        A │") + `│││││││││││││││││││││
	` + config.Red("└──────────┘") + `┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘
			-- Version: ` + config.Version + `
			-- By: Mustard404
	`)
}
