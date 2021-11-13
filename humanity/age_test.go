package humanity

import (
	"testing"
	"time"
)

func TestHowOld(t *testing.T) {

	test := time.Now().AddDate(-28,0,0)

	x, err := HowOld(test)

	if err != nil {
		t.Error(err)
	}
	if x != 28 {
		t.Error("Expected", 19, "Got", x)
	}
}

func TestHowOldComma(t *testing.T) {
	test := time.Now().AddDate(-28,0,0)

	x, err := HowOldComma(test)

	if err != nil {
		t.Error(err)
	}
	if x != 28.02 {
		t.Error("Expected", 19, "Got", x)
	}
}
