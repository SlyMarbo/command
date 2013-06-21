package command

import (
	"strconv"
	"strings"
)

// Token represents a token scanned from a Scanner's input.
type Token struct {
	full            string
	modified        string
	caseInsensitive bool
}

// Blank returns true when the token is the empty string.
func (t *Token) Blank() bool {
	return t.full == ""
}

// Body returns the token, possibly after modification by
// HasPrefix and HasSuffix.
func (t *Token) Body() string {
	return t.modified
}

// Bool attempts to parse the token (possibly modified by
// HasPrefix or HasSuffix) as a bool.
// It accepts 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False.
// Any other value returns an error.
func (t *Token) Bool() (bool, error) {
	return strconv.ParseBool(t.modified)
}

// Bytes returns the raw bytes of the token generated by a call
// to Scan.
func (t *Token) Bytes() []byte {
	return []byte(t.full)
}

// Equals returns true if the token is the same as any of the
// provided strings.
func (t *Token) Equals(tokens ...string) bool {
	if t.caseInsensitive {
		t.modified = strings.ToLower(t.modified)
	}
	for _, token := range tokens {
		if t.caseInsensitive {
			token = strings.ToLower(token)
		}
		if token == t.modified {
			return true
		}
	}
	return false
}

// Float attempts to parse the token (possibly modified by
// HasPrefix or HasSuffix) as a floating-point number with
// the precision specified by bitSize: 32 for float32, or
// 64 for float64. When bitSize=32, the result still has
// type float64, but it will be convertible to float32
// without changing its value.
//
// If the token is well-formed and near a valid floating
// point number, Float returns the nearest floating point
// number rounded using using IEEE754 unbiased rounding.
//
// The errors that Float returns have concrete type
// *strconv.NumError and include err.Num = s.
//
// If s is not syntactically well-formed, Float returns
// err.Error = ErrSyntax.
//
// If s is syntactically well-formed but is more than 1/2 ULP
// away from the largest floating point number of the given size,
// ParseFloat returns f = ±Inf, err.Error = ErrRange.
func (t *Token) Float(bitSize int) (float64, error) {
	return strconv.ParseFloat(t.modified, bitSize)
}

// HasPrefix returns true if the token has the given prefix. If
// the prefix matches, the string will be accessible with Body,
// with the prefix removed. The full string (with prefix) is
// still accessible with String.
func (t *Token) HasPrefix(prefix ...string) bool {
	if t.caseInsensitive {
		t.modified = strings.ToLower(t.modified)
	}
	for _, p := range prefix {
		if t.caseInsensitive {
			p = strings.ToLower(p)
		}
		if strings.HasPrefix(t.modified, p) {
			t.modified = t.modified[len(p):]
			return true
		}
	}
	return false
}

// HasSuffix returns true if the token has the given suffix. If
// the suffix matches, the string will be accessible with Body,
// with the suffix removed. The full string (with suffix) is
// still accessible with String.
func (t *Token) HasSuffix(suffix ...string) bool {
	if t.caseInsensitive {
		t.modified = strings.ToLower(t.modified)
	}
	for _, s := range suffix {
		if t.caseInsensitive {
			s = strings.ToLower(s)
		}
		if strings.HasSuffix(t.modified, s) {
			t.modified = t.modified[:len(t.modified)-len(s)]
			return true
		}
	}
	return false
}

// Int interprets the token in the given base (2 to 36) and
// returns the corresponding value i. If base == 0, the base is
// implied by the string's prefix: base 16 for "0x", base 8 for
// "0", and base 10 otherwise.
//
// The bitSize argument specifies the integer type
// that the result must fit into.  Bit sizes 0, 8, 16, 32, and 64
// correspond to int, int8, int16, int32, and int64.
//
// The errors that Int returns have concrete type *strconv.NumError
// and include err.Num = s.  If s is empty or contains invalid
// digits, err.Error = ErrSyntax; if the value corresponding
// to s cannot be represented by a signed integer of the
// given size, err.Error = ErrRange.
func (t *Token) Int(base int, bitSize int) (int64, error) {
	return strconv.ParseInt(t.modified, base, bitSize)
}

// String returns the token generated by a call to Scan
// as a newly allocated string holding its bytes.
func (t *Token) String() string {
	return t.full
}

// Uint is like Int but for unsigned numbers.
func (t *Token) Uint(base int, bitSize int) (uint64, error) {
	return strconv.ParseUint(t.modified, base, bitSize)
}
