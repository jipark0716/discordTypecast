package handler

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

var FileLogger *log.Logger

func init() {
	os.MkdirAll("log", os.ModePerm)

	file, err := os.OpenFile("log/discord_typecast.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	FileLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Report(err error) {
	err = errors.Wrap(err, "report error")
	FileLogger.Println(fmt.Sprintf("%+v\n", err))
}
