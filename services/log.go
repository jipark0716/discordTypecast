package services

import "fmt"

func report(err error) {
	fmt.Printf("%+v\n", err)
}
