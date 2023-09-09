package model

import (
	"testing"
)

func TestReflect(t *testing.T) {
	a := Action{
		Name: "CreteAlgo",
	}

	a.Validate()

}
