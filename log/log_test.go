package log

import "testing"

func TestBasic(t *testing.T) {
	InitZapSugared(true, false)
	Infof("Logging test ! %s !", "Let's see")
}
