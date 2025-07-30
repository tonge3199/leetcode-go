package string

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		want        int
	}{
		{
			name:        "twoSubstring Case",
			inputString: "abcdabcdabcdeabcdefaebcd",
			want:        6,
		},
		{
			name:        "simple Case",
			inputString: "bbbbb",
			want:        1,
		},
		{
			name:        "gap Case",
			inputString: "pwwkew",
			want:        3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.inputString)

			if got != tt.want {
				t.Errorf("lengthOfLongestSubstring(%v) got = %v, want %v", tt.inputString, got, tt.want)
			}
		})
	}
}
