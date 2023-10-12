package converter

import "strings"

func GenerateForKanaConverter(in string) <-chan KanaConverterRune {
	out := make(chan KanaConverterRune)
	go func() {
		defer close(out)
		for _, r := range in {
			out <- KanaConverterRune{Rune: r}
		}
	}()
	return out
}

func StringForKanaConverter(in <-chan KanaConverterRune) string {
	var b strings.Builder
	for r := range in {
		b.WriteRune(r.Rune)
	}
	return b.String()
}
