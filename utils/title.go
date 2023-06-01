package utils

import (
	"AceofHearts/config"
	"fmt"
)

func Title(title string) {
	fmt.Println(config.Cyan(`
===========================================
## ` + title + `
===========================================
	`))
	Log("## %s\t\n", title)
}
