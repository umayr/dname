package dname_test

import (
	"strings"
	"testing"

	"github.com/umayr/dname"
)

func TestGenerateSingle(t *testing.T) {
	n := dname.Generate()
	if n == "" {
		t.Error("generated name shouldn't be empty")
	}

	if !strings.Contains(n, "_") {
		t.Error("generated name should consist of an underscore")
	}

	s := strings.Split(n, "_")

	if s[0] == "" || s[1] == "" {
		t.Error("generated name should consist of both parts, i.e. adjective and subject")
	}
}

func TestGenerateMultiple(t *testing.T) {
	for i := 0; i < 1000; i++ {
		n := dname.Generate()
		if n == "" {
			t.Error("generated name shouldn't be empty")
		}

		if !strings.Contains(n, "_") {
			t.Error("generated name should consist of an underscore")
		}

		s := strings.Split(n, "_")

		if s[0] == "" || s[1] == "" {
			t.Error("generated name should consist of both parts, i.e. adjective and subject")
		}
	}
}
