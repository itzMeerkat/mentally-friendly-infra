package log

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	InitZapSugared(true, false)
	Infof("Logging test ! %s !", "Let's see")
}

func TestGetLogID(t *testing.T) {
	id := GetLogID("12.322.333.33")
	fmt.Println(id)
}