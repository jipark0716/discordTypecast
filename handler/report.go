package handler

import (
	"fmt"

	"github.com/pkg/errors"
)

func Report(err error) {
	err = errors.Wrap(err, "report error")
	fmt.Printf("%+v\n", err)
}
