package astra

import "testing"

var optionSets = [][]Option{
	{
		IgnoreVariables, IgnoreMethods,
	},
	{
		IgnoreFunctions, IgnoreInterfaces, IgnoreMethods, IgnoreConstants,
	},
}

func TestOptions(t *testing.T) {
	for _, s := range optionSets {
		concated := concatOptions(s)
		for i := range s {
			if !concated.check(s[i]) {
				t.Error(concated, s[i])
			}
		}
	}
}
