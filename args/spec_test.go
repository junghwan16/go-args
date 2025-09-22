package args_test

import (
	"testing"

	"github.com/junghwan16/go-args/args"
)

func TestBoolean_DefaultIsFalse(t *testing.T) {
	spec := args.New().Boolean("l")

	res, err := spec.Parse([]string{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := res.Bool("l"); got != false {
		t.Fatalf("l default = %v, want false", got)
	}
}
