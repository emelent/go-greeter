package greeter

import (
	"testing"
)

func TestHello(t *testing.T) {
	var tests = []struct {
		name     string
		expected string
	}{
		{"John", "Hello John."},
		{"Jake", "Hello Jake."},
	}

	for _, tt := range tests {
		testName := tt.name
		t.Run(testName, func(t *testing.T) {
			actual := Hello(tt.name)
			if actual != tt.expected {
				t.Errorf("Expected '%s', actual '%s'", tt.expected, actual)
			}
		})
	}
}
