package chess

import (
	"strconv"
	"testing"
)

func TestFlipDiagA1H8(t *testing.T) {
	for i, tc := range []struct {
		input      uint64
		wantOutput uint64
	}{
		{0, 0},
		{fileA, rank1},
		{rank2, fileB},
		{rank3 | fileG, fileC | rank7},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			gotOutput := flipDiagA1H8(tc.input)
			if gotOutput != tc.wantOutput {
				t.Logf("input: %v", tc.input)
				t.Logf("want: %v", tc.wantOutput)
				t.Logf("got: %v", gotOutput)
				t.Errorf("mismatch")
			}
		})
	}
}
