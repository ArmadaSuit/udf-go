package converter

import "strings"

func Generate(in string) <-chan rune {
	out := make(chan rune)
	go func() {
		defer close(out)
		for _, i := range in {
			out <- i
		}
	}()
	return out
}

func String(in <-chan rune) string {
	var b strings.Builder
	for r := range in {
		b.WriteRune(r)
	}
	return b.String()
}
