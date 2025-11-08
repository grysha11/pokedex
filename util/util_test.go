package util

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input	string
		want	[]string
	}{
		"whitespaces":			{input: "charmander bulbasaur pikachu", want: []string{"charmander", "bulbasaur", "pikachu"}},
		"multiple_whitespaces":	{input: "    charizard  charmeleon     weedle   ", want: []string{"charizard", "charmeleon", "weedle"}},
		"upper_case":			{input: "JIGGLYpuff WigglyTuff ZuBaT", want: []string{"jigglypuff", "wigglytuff", "zubat"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}