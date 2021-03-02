package log

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	l := InitZapSugared(false, true, 1)
	l.Infow("k1",1, "k2",2)
}

func TestGetLogID(t *testing.T) {
	id := GetLogID("12.322.333.33")
	fmt.Println(id)
}