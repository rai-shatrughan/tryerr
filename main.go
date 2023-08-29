package tryerr

import (
	"fmt"
	"log"
)

func HandleErrWithPanic(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}

func HandleErrWithPrint(err error, msg string) {
	if err != nil {
		fmt.Printf(msg, err)
	}
}

func HandleErrWithLogFatal(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func HandleErrWithLogPanic(err error, msg string) {
	if err != nil {
		log.Panicf(msg, err)
	}
}
