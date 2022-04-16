package interation

const defaultTimes = 5

// Repeat takes a string and how many it should repeat and return the string
// over and over again according to the times. If the time should <= 0 we use
// the default time (5)
func Repeat(ch string, times int) string {
	var repeated string

	if times <= 0 {
		times = defaultTimes
	}

	for i := 0; i < times; i++ {
		repeated += ch
	}
	return repeated
}
