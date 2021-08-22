package playcricket

import "time"

func GetCurrentSeason() string {
	return time.Now().Format("2006")
}
