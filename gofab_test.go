package gofab

import (
	"fmt"
	"testing"
)

func TestErrLog(t *testing.T) {
	ErrLog(fmt.Errorf("hello"))
}