package logo

import (
	"bytes"
	"strings"
	"testing"
)

func TestNormalOutput(t *testing.T) {
	var buf = bytes.NewBuffer(nil)
	AddLogger("buffer", buf, ALL)
	defer ReleaseLogger("buffer")
	str := "Buffer one"
	Debug(str)

	if !strings.Contains(buf.String(), str) {
		t.Error("Error output something in this buffer")
	}
}

func TestMaskedOutput(t *testing.T) {
	var buf = bytes.NewBuffer(nil)
	AddLogger("buffer", buf, CRITICAL|ERROR)
	defer ReleaseLogger("buffer")
	str := "Buffer two"
	Debug(str)

	if len(buf.String()) > 0 {
		t.Error("No log need to write to this logger")
	}
}
