package twelve

const (
	testVersion = 1
)

func Verse(n int) string {
	otherDays := [...]string{"a Partridge in a Pear Tree.", "two Turtle Doves, and ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}
	weekDays := [...]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

	song := "On the " + weekDays[n-1] + " day of Christmas my true love gave to me, "
	i := n - 1
	for i >= 0 {
		song += otherDays[i]
		i--
	}
	return song
}

func Song() string {
	allSong := ""
	for j := 1; j <= 12; j++ {
		allSong += Verse(j) + "\n"
	}
	return allSong
}
