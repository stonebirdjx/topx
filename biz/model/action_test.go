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

func TestObjectID(t *testing.T) {
	//primitive.ObjectID("64fd232a23e5f7e306115ed5")
}
