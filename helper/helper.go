package helper

import "log"

func CheckErrorNill(err error) {
	log.Panic(err)
}
