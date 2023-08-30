package greet

/*
AÇIKLAMAAAA

*/
import (
	"log"
	"strings"
)

// SayHi greets given names.
func SayHi(names ...string) string {
	if len(names) == 0 {
		log.Println("This is a log message")
		return "hi everybody!"

	}
	out := make([]string, len(names))
	for i, name := range names {
		out[i] = "hi " + name + "!"
	}

	return strings.Join(out, "\n")
}
