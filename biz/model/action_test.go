package model

import (
	"context"
	"testing"
)

func TestReflect(t *testing.T) {
	a := Action{
		Name: "CreteAlgo",
	}

	a.Validate(context.Background())
}

func TestObjectID(t *testing.T) {
	//primitive.ObjectID("64fd232a23e5f7e306115ed5")
}
