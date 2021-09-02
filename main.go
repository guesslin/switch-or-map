package main

import (
	"flag"
	"fmt"

	"gxn.tw/performance/switch-or-map/internal/plugs"
)

func main() {
	path := flag.String("path", "", "path to plugin file")
	key := flag.String("key", "aal", "key to lookup")
	flag.Parse()

	getter := plugs.LoadGetter(*path)
	fmt.Println(getter.Get(*key))
}
