package command

import (
	"strings"
)

// Token returns the most recent token generated by a call to Scan.
func (s *Scanner) Token() *Token {
	return s.t
}

// Token represents a token scanned from a Scanner's input.
type Token struct {
	s string
	b string
	c bool
}

// Blank returns true when the token is the empty string.
func (t *Token) Blank() bool {
	return t.s == ""
}

// Body returns the token, possibly after modification by
// HasPrefix and HasSuffix.
func (t *Token) Body() string {
	return t.b
}

// Bytes returns the raw bytes of the token generated by a call
// to Scan.
func (t *Token) Bytes() []byte {
	return []byte(t.s)
}

// Equals returns true if the token is the same as any of the
// provided strings.
func (t *Token) Equals(tokens ...string) bool {
	for _, token := range tokens {
		if token == t.s {
			return true
		}
	}
	return false
}

// HasPrefix returns true if the token has the given prefix. If
// the prefix matches, the string will be accessible with Body,
// with the prefix removed. The full string (with prefix) is
// still accessible with String.
func (t *Token) HasPrefix(prefix string) bool {
	if t.c {
		prefix = strings.ToLower(prefix)
	}
	if strings.HasPrefix(t.b, prefix) {
		t.b = t.b[len(prefix):]
		return true
	}
	return false
}

// HasSuffix returns true if the token has the given suffix. If
// the suffix matches, the string will be accessible with Body,
// with the suffix removed. The full string (with suffix) is
// still accessible with String.
func (t *Token) HasSuffix(suffix string) bool {
	if t.c {
		suffix = strings.ToLower(suffix)
	}
	if strings.HasSuffix(t.b, suffix) {
		t.b = t.b[:len(t.b)-len(suffix)]
		return true
	}
	return false
}

// String returns the token generated by a call to Scan
// as a newly allocated string holding its bytes.
func (t *Token) String() string {
	return t.s
}