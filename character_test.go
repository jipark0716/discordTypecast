package main_test

import (
	"fmt"
	"testing"
	."github.com/jipark0716/discordTypecast"
)

func TestSearch(t *testing.T) {
	t.Error("1234")
	char, err := SearchCharacter("머박")
	fmt.Printf("%#v\n%#v\n", char, err)
}
