package token_bucket

import (
	"testing"
)

func Test_Get_One_Token(t *testing.T) {
	tb := DefaultTokenBucket()

	if tb.Get() != nil {
		t.Error("Cannot get token from token bucket.")
	}

	if tb.Remain != tb.Total-1 {
		t.Errorf("Expect '%d' but got '%d'", tb.Total-1, tb.Remain)
	}
}

func Test_Get_Part_Of_Tokens(t *testing.T) {
	// TODO: concurrent get
}
