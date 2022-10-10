package services

import (
	"fmt"

	"github.com/pkg/errors"
)

func report(err error) {
	err = errors.Wrap(err, "report error")
	fmt.Printf("%+v\n", err)
}
