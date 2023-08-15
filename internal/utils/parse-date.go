package parsers

import "time"

func ParseDate(layoutDate string, dateString string) time.Time {
	t, err := time.Parse(layoutDate, dateString)
	if err != nil {
		panic(err)
	}
	return t
}
