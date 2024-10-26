package toolkit

import "testing"

// Name will consist of:
// - Test - this is mandatory to be a test
// - Tools - is the receiver name
// - RandomString - the name of the function
func TestTools_RandomString(t *testing.T) {
	var testTools Tools

	s := testTools.RandomString(10)

	if len(s) != 10 {
		t.Error("wrong length random string returned")
	}
}
