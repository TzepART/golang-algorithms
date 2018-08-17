package methods_and_interfaces

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13byte(sb byte) byte {
	s := rune(sb)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		sb += 13
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		sb -= 13
	}

	return sb
}

func (rot13 rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot13.r.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13byte(p[i])
	}
	return
}

func ReadersExercise23() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
