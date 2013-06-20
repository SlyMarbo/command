package command

import (
	"bufio"
	"io"
	"errors"
	"strings"
)

// Scanner provides a convenient interface for reading data such as
// a file of newline-delimited lines of text. Successive calls to
// the Scan method will step through the 'tokens' of a file, skipping
// the bytes between the tokens. The specification of a token is
// defined by a split function of type SplitFunc; the default split
// function breaks the input into lines with line termination stripped. Split
// functions are defined in this package for scanning a file into
// lines, bytes, UTF-8-encoded runes, and space-delimited words. The
// client may instead provide a custom split function.
//
// Scanning stops unrecoverably at EOF, the first I/O error, or a token too
// large to fit in the buffer. When a scan stops, the reader may have
// advanced arbitrarily far past the last token. Programs that need more
// control over error handling or large tokens, or must run sequential scans
// on a reader, should use bufio.Reader instead.
type Scanner struct {
	s *bufio.Scanner
	t *Token
	c bool
}

// NewScanner returns a new Scanner to read from r.
// The split function defaults to ScanLines.
func NewScanner(r io.Reader, caseInsensitive bool) *Scanner {
	out := new(Scanner)
	out.s = bufio.NewScanner(r)
	out.c = caseInsensitive
	return out
}

// Err returns the first non-EOF error that was encountered by the Scanner.
func (s *Scanner) Err() error {
	if s.s != nil {
		return s.s.Err()
	} else {
		return errors.New("Error: Scanner not initialised.")
	}
}

// Scan advances the Scanner to the next token, which will then be
// available through the Bytes or Text method. It returns false when the
// scan stops, either by reaching the end of the input or an error.
// After Scan returns false, the Err method will return any error that
// occurred during scanning, except that if it was io.EOF, Err
// will return nil.
func (s *Scanner) Scan() bool {
	if s.s != nil {
		b := s.s.Scan()
		if b {
			str := s.s.Text()
			if s.c {
				str = strings.ToLower(str)
			}
			s.t = &Token{str, str, s.c}
		}
		return b
	}
	return false
}

// Split sets the split function for the Scanner. If called, it must be
// called before Scan. The default split function is ScanLines.
func (s *Scanner) Split(split bufio.SplitFunc) {
	if s.s != nil {
		s.s.Split(split)
	}
}
