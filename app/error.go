package main

import "github.com/importfmt/logger"

func checkErr(err error, message string) {
	if err != nil {
		logger.Error.Println(message)
	}
}

func checkWarn(err error, message string) {
	if err != nil {
		logger.Warning.Println(message)
	}
}

func checkInfo(err error, message string) {
	if err != nil {
		logger.Info.Println(message)
	}
}

func catchFalse(b bool, message string) {
	if b != true {
		logger.Warning.Println(message)
	}
}

func catchTrue(b bool, message string) {
	if b != false {
		logger.Warning.Println(message)
	}
}
