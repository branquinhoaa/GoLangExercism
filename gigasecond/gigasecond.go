package gigasecond

import "time"

const testVersion = 4 // find the value in gigasecond_test.go
const gigaSecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond * time.Second)
}
