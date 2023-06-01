package utils

import (
	"AceofHearts/config"
	"fmt"
)

func Title2(title string) {
	fmt.Println(config.Cyan(`
===========================================
## ` + title + `
===========================================
	`))
	Log("### %s\t\n", title)
}
