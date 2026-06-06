package storage

import "testing"

func TestCompile(t *testing.T) {
	var s StorageEngine
	if &s == nil {
		t.Fatal()
	}
}
