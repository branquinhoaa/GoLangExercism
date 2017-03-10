package clock

import "fmt"

const testVersion = 4
const minutesInDay = 24 * 60

type Clock int

func New(hour, min int) Clock {
	minutes := min + 60*hour
	return Minutize(Clock(minutes))
}

func (c Clock) String() string {
	hours := c / 60
	minutes := c % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func (c Clock) Add(min int) Clock {
	return Minutize(c + Clock(min))
}

func Minutize(c Clock) Clock {
	for c < 0 {
		c += minutesInDay
	}
	return Clock(c % minutesInDay)
}
