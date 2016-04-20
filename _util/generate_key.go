package main

import (
	"fmt"

	"github.com/wcharczuk/instabot/server/core"
)

func main() {
	key := core.CreateKey(32)
	fmt.Println(core.Base64Encode(key))
}
