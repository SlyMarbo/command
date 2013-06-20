/*
Package command is a simple Go package for implementing a CLI.

Example usage:
		package main

		import (
			"github.com/SlyMarbo/command"
			"os"
		)

		func main() {
			scanner := command.NewScanner(os.Stdin, true) // Read from stdin.

			for scanner.Scan() {
				t := scanner.Token()
				switch {

					// Check for an empty token.
					case token.Empty():
						// ...

					// Compare against one string.
					case token.Equals("example"):
						// ...

					// Compare against multiple strings.
					case token.Equals("more", "other"):
						// ...

					// Check for a prefix.
					case token.HasPrefix("start"):
						token.String() // Full string, starting with "start".
						token.Body()   // Modified string, starting after "start".

					// HasSuffix works identically.
				}
			}

			if err := scanner.Err(); err != nil {
				// handle error.
			}
		}
*/
package command
