package gigasecond

import "time"

const testVersion = 4
const gigaSecond = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond * time.Second)
}
