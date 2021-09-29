package timex

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	MONDAY Day = iota
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

type Day int

func (d Day) Duration() time.Duration {
	return time.Duration(d) * 24 * time.Hour
}

func LookupDay(s string) (Day, error) {
	switch strings.ToLower(s) {
	case "monday":
		return MONDAY, nil
	case "tuesday":
		return TUESDAY, nil
	case "wednesday":
		return WEDNESDAY, nil
	case "thursday":
		return THURSDAY, nil
	case "friday":
		return FRIDAY, nil
	case "saturday":
		return SATURDAY, nil
	case "sunday":
		return SUNDAY, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	switch i {
	case 0:
		return MONDAY, nil
	case 1:
		return TUESDAY, nil
	case 2:
		return WEDNESDAY, nil
	case 3:
		return THURSDAY, nil
	case 4:
		return FRIDAY, nil
	case 5:
		return SATURDAY, nil
	case 6:
		return SUNDAY, nil
	}

	return 0, fmt.Errorf("unknown day %q", s)
}
