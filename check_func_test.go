package done

import (
	"testing"

	"github.com/pkg/errors"
)

func TestDone(t *testing.T) {
	if false {
		Done(errors.New("wrong"))
	}
}

func TestMust(t *testing.T) {
	if false {
		Must(errors.New("wrong"))
	}
}

func TestSoft(t *testing.T) {
	Soft(errors.New("wrong"))
}

func TestFata(t *testing.T) {
	if false {
		Fata(errors.New("123"))
	}
}
