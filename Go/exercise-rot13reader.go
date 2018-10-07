/* https://tour.golang.org/methods/23
Make Implement io.Reader for the rot13 substitution cipher
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(b []byte) (int, error) {
	n, err := rot.r.Read(b) // reads in as integers from ascii codes
	for i := 0; i < n; i++ {
		// need to work in mod26 space, which means we have to subtract the
		// ascii offset, mod, then add the offset back depending on casing
		switch {
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
