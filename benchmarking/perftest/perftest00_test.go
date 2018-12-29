package perftest

import (
	"regexp"
	"strings"
	"testing"
)

var haystack = `Lorem ipsum dolor sit amet, consectetur adipiscing 
[...]
Vivamus vitae nulla posuere, pellentesque quam posuere`
var pattern = regexp.MustCompile("auctor")

func BenchmarkSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains(haystack, "auctor")
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pattern.MatchString(haystack)
	}
}