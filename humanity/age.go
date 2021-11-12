package humanity

import (
	"errors"
	"time"
)

// HowOld You can get human age
func HowOld(t time.Time) (int, error) {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	currentDay := time.Now().Day()

	currentDate := time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.UTC)

	r := int((currentDate.Sub(t).Hours() / 24) / 365)

	if r <= 0 {
		return r, errors.New("please check human age params")
	}
	return int((currentDate.Sub(t).Hours() / 24) / 365), nil
}
