package function

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type capitalizedReader struct {
	r io.Reader
}

func CapReader(r io.Reader) io.Reader {
	return &capitalizedReader{r: r}
}

func (r *capitalizedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p) // r有r这个变量，有read这个方法
	if err != nil {
		return 0, err
	}
	q := bytes.ToUpper(p)
	for i, v := range q {
		p[i] = v
	}
	return n, err
}

func TestCapReader(t *testing.T) {
	r := strings.NewReader("hello,gopher!\n")
	r1 := CapReader(io.LimitReader(r, 4))
	if _, err := io.Copy(os.Stdout, r1); err != nil {
		log.Fatal(err)
	}
}
