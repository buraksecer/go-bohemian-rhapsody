package humanity

import (
	"errors"
	"time"
)

// HowOld You can get human age
func HowOld(t time.Time) (int, error) {
	currentMonth := time.Now().Month()

	if currentMonth < t.Month() {
		return 0, errors.New("please check human age params")
	}

	return int((currentMonth - t.Month()) / 12), nil
}
