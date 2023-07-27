package main

import (
	"strings"
	"testing"
)

func TestReadLines(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Test 1",
			input:    "n0001\nn0002\nn0003\nn0004\nn0005",
			expected: []string{"n0001", "n0002", "n0003", "n0004", "n0005"},
		},
		{
			name:     "Test 2",
			input:    "n0006\nn0007\nn0008\nn0009\nn0010",
			expected: []string{"n0006", "n0007", "n0008", "n0009", "n0010"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readLines(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("ReadLines() error = %v", err)
			}

			for i, line := range got {
				if line != tt.expected[i] {
					t.Errorf("Line %d: got %v, want %v", i, line, tt.expected[i])
				}
			}
		})
	}
}
