package main

import (
	"fmt"
)

/*
 *  write const declarations for KB, MB, up through YB as compactly as you can
 */

const (
	KB = 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000
)

func main() {
	fmt.Println(EB)
}
