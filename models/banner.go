package models

import (
	"AceofHearts/config"
	"fmt"
)

func Banner() {
	fmt.Println(`
	┌──────────┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐┐
	│ ` + config.Red("A") + `        ││││││││││││││││││││││
	│          ││││││││││││││││││││││
	│    ` + config.Red("❤") + `     ││││││││││││││││││││││
	│          ││││││││││││││││││││││
	│        ` + config.Red("A") + ` ││││││││││││││││││││││	
	└──────────┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘┘
			-- Version: 1.0
			-- By: Mustard404
	`)
}
