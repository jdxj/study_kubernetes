package db

import "testing"

func TestInit(t *testing.T) {
	err := Init("root", "123456", "127.0.0.1", "3306", "sys")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
