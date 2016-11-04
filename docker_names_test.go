package docker_names_test

import (
	"testing"
	"strings"
	
	"github.com/umayr/docker-names"
)

func TestGenerateSingle(t *testing.T) {
	n := docker_names.Generate()
	if n == "" {
		t.Error("generated name shouldn't be empty")
	}

	if !strings.Contains(n, "-") {
		t.Error("generated name should consist of an hypen")
	}

	s := strings.Split(n, "-")

	if s[0] == "" || s[1] == "" {
		t.Error("generated name should consist of both parts, i.e. adjective and subject")
	}
}

func TestGenerateMultiple(t *testing.T) {
	for i := 0; i < 100; i++ {
		n := docker_names.Generate()
		if n == "" {
			t.Error("generated name shouldn't be empty")
		}

		if !strings.Contains(n, "-") {
			t.Error("generated name should consist of an hypen")
		}

		s := strings.Split(n, "-")

		if s[0] == "" || s[1] == "" {
			t.Error("generated name should consist of both parts, i.e. adjective and subject")
		}
	}
}
