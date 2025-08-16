package parser

// isAsciiLetter checks whether r is an ASCII letter (A-Z, a-z).
func isAsciiLetter(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}

// isAsciiDigit checks whether r is an ASCII digit (0-9).
func isAsciiDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
