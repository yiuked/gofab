package gofab

import "log"

type ErrLevel int

const (
	ErrorLog ErrLevel = iota
	ErrorIgnore
	ErrorFinal
)

func Error(err error, level ErrLevel) {
	if err == nil {
		return
	}

	switch level {
	case ErrorFinal:
		log.Panic(err)
		break
	case ErrorIgnore:
		break
	case ErrorLog:
	default:
		log.Println(err)
	}
}

func ErrLog(err error) {
	Error(err, ErrorLog)
}

func ErrIgnore(err error) {
	Error(err, ErrorIgnore)
}

func ErrFinal(err error) {
	Error(err, ErrorFinal)
}
