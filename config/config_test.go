package config

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

func TestMarshalConfig(t *testing.T) {
	cfg := &Ins
	data, err := yaml.Marshal(cfg)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}
	fmt.Printf("%s\n", data)
}
