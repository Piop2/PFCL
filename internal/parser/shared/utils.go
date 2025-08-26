package shared

// IsAsciiLetter checks whether r is an ASCII letter (A-Z, a-z).
func IsAsciiLetter(r rune) bool {
	return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}

// IsAsciiDigit checks whether r is an ASCII digit (0-9).
func IsAsciiDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// IsSpace checks whether r is a space or tab character (' ', '\t').
func IsSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// IsNewline checks whether r is a newline character ('\n' or '\r').
func IsNewline(r rune) bool {
	return r == '\n' || r == '\r'
}

// IsWhitespace checks whether r is a whitespace character
// (space, tab, newline).
func IsWhitespace(r rune) bool {
	return IsSpace(r) || IsNewline(r)
}
