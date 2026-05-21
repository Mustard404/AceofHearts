package utils

import (
	"AceofHearts/config"
	"fmt"
)

// Title 打印一级标题
func Title(title string) {
	fmt.Println(config.Cyan(fmt.Sprintf(`
===========================================
## %s
===========================================`, title)))
	Log("\n## %s\n", title)
}

// Title2 打印二级标题
func Title2(title string) {
	fmt.Println(config.Cyan(fmt.Sprintf(`
-------------------------------------------
### %s
-------------------------------------------`, title)))
	Log("\n### %s\n", title)
}
