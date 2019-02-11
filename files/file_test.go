package files

import (
	"os"
	"testing"
)

func TestListingDirectory(t *testing.T) {
	f := os.Getenv("HOME")
	c := Ctx{RootKey: f}
	_, e := ListingDirectory(c, "/")
	if e != nil {
		t.Fatal(e)
	}
}
