package humanity

import (
	"errors"
	"math"
	"time"
)

var currentYear = time.Now().Year()
var currentMonth = time.Now().Month()
var currentDay = time.Now().Day()
var currentDate = time.Date(currentYear, currentMonth, currentDay, 0, 0, 0, 0, time.UTC)

// HowOld You can get human age
// Time Human birthday
func HowOld(t time.Time) (int, error) {
	r := int((currentDate.Sub(t).Hours() / 24) / 365)

	if r <= 0 {
		return r, errors.New("please check human age params")
	}
	return int((currentDate.Sub(t).Hours() / 24) / 365), nil
}

// HowOldComma You can find out your exact age
// Time Human birthday
func HowOldComma(t time.Time) (float64, error) {
	r := (currentDate.Sub(t).Hours() / 24) / 365

	if r <= 0 {
		return r, errors.New("please check human age params")
	}

	res := (currentDate.Sub(t).Hours() / 24) / 365

	return math.Round(res*100) / 100, nil
}
