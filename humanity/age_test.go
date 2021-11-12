package humanity

import (
	"testing"
	"time"
)

func TestHowOld(t *testing.T) {

	test := time.Date(1992, 19, 12, 0, 0, 0, 0, time.UTC)

	x, err := HowOld(test)

	if err != nil {
		t.Error(err)
	}
	if x != 28 {
		t.Error("Expected", 19, "Got", x)
	}
}
