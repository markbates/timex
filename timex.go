package timex

import "time"

var (
	DAY  time.Duration = 24 * time.Hour
	TERM               = 13 * WEEK
	WEEK               = 7 * DAY
)
