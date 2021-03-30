package redis

import "testing"

func TestInitClient(t *testing.T) {
	err := InitClient("127.0.0.1", "6379", "", 0)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
