package morse

import (
	"testing"
)

func TestCheckInput(t *testing.T) {
	err := CheckInput([]rune{'é'})
	switch {
	case err == nil:
		t.Error("expecting error")
	case err != errInvalidChar:
		t.Errorf("unexpected error: exp %v , got %v", errInvalidChar, err)
	}
	err = CheckInput([]rune{'A'})
	if err != nil {
		t.Errorf("unexpected error: exp nil , got %v", err)
	}
}
