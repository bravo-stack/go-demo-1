package mascot_test

import (
	"testing"

	"github.com/bravo-stack/go-demo-1/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}