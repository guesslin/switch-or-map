package helpers

import "log"

func PanicOn(err error, msg ...string) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
