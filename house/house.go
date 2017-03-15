package house

const testVersion = 1

func Verse(n int) string {
	subject := []string{"house", "malt", "rat", "cat", "dog", "cow with the crumpled horn", "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn", "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"}
	firstVerse := "This is the " + subject[n-1]
	if n > 1 {
		firstVerse += "\n"
	}
	return firstVerse + Actions(n-1)
}

func Actions(n int) string {
	actions := []string{" that Jack built.", "that lay in the house", "that ate the malt", "that killed the rat", "that worried the cat", "that tossed the dog", "that milked the cow with the crumpled horn", "that kissed the maiden all forlorn", "that married the man all tattered and torn", "that woke the priest all shaven and shorn", "that kept the rooster that crowed in the morn", "that belonged to the farmer sowing his corn"}
	text := actions[n]
	if n == 0 {
		return text
	}
	if n > 1 {
		text += "\n"
	}
	return text + Actions(n-1)
}

func Song() string {
	allSong := ""
	i := 1
	for i < 13 {
		allSong += Verse(i)
		if i < 12 {
			allSong += "\n\n"
		}
		i++
	}
	return allSong
}
