package main

const (
	counter = 5
	nonCount = 0
)
func main() {
	Repeat("a", 6)
}

func Repeat(character string, loopCount int) string {
	var repeated string
	if loopCount == nonCount {
		loopCount = counter
	}
	for i := 0; i < loopCount; i++ {
		repeated += character
	}
	return repeated
}
