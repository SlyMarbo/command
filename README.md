command
=======

command is a simple Go package for implementing a CLI.

Example usage:
```go
package main

import (
	"github.com/SlyMarbo/command"
	"os"
)

func main() {
	scanner := command.NewScanner(os.Stdin, true) // Read from stdin.
	
	for scanner.Scan() {
		token := scanner.Token()
		switch {
			
			// Check for an empty token.
			case token.Blank():
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
```

If the entered command was "stop pid 1234", then the command package could be used as follows:
```go
t := scanner.Token()
switch {
	
	case token.HasPrefix("stop", "kill"):
		if token.HasPrefix("pid") {
			pid, err := token.Int(0, 0) // Automatic base, type int.
			
			// ...
		}
	
	// ...
}
```
