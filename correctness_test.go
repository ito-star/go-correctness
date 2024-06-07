package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYourFunc(t *testing.T) {
	// First string is expected, second string is input
	testCases := [][2]string{
		{"StellarMLS #U8216253;DOM 60", "StellarMLS #U8216253;DON 60"},
		{"MARIS#23061532;DOM 39", "MARIS#23061532;DG 39"},
		{"MARIS#23029660;DOM 42", "MARIS#23029660;DOR 42"},
		{"TN Property Data;DOM 1", "TN Property Data: DOM 1"},
		{"MLS # 4057001;DOM 0", "MLS # 4057001 DOM 0"},
		{"NCR #100412013;DOM 39", "NCR #100412013:DOM 39"},
		{"Bldr/PR/FMLS#7236386;DOM 99", "Bldr/PR/FMLS#7236386;D 99"},
		{"MLS #202302649;DOM 101", "MLS #202302649; 101"},
		{"RMLS #6409190;DOM 19", "RMLS #6409190 DOM 19"},
		{"SEFMLX #A11468828;DOM 41", "SEFMLX #A11468828;D DOM 41"},
		{"GFKMLS#23-1561;DOM 48", "GFKMLS#23-1561; 48"},
		{"Bldr/PR/FMLS#7264757;DOM 7", "Bldr/PR/FMLS#7264757;0 7"},
	}

	for k, tc := range testCases {
		corrected := yourFunc(tc[1])
		assert.Equalf(t, tc[0], corrected, "Failed test case %d: %s != %s", k, tc[0], corrected)
	}
}
