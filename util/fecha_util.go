package util

import "time"

func GenerateFecha() time.Time {
	return time.Now().Add(-time.Hour * 3)
}