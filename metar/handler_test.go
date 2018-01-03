package function

import (
	"errors"
	"fmt"
	"testing"
)

func Test_correct_parsing_of_input_string(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      error
	}{
		{"Echo Golf Aplha Charlie", "EGAC", nil},
		{"Aplha Bravo", "AB", nil},
		{"Aplha Bravo Charlie Delta Echo Foxtrot", "ABCDEF", nil},
		{"", "", errors.New("Oops, no weather station was provided, ensure you provide a station using the Phonetic Alphabet, ie Echo Golf Alpha Charlie")},
	}

	for testNum, test := range tests {
		fmt.Println(fmt.Sprintf("Running test #%v", testNum))
		icao, err := parseInput([]byte(test.input))
		if test.err != nil {
			if test.err.Error() != err.Error() {
				t.Errorf("Actual error %s did not match expected %s", err.Error(), test.err.Error())
			}
		}
		if string(icao) != test.expected {
			t.Errorf("Actual icao %s did not match expected %s", icao, test.expected)
		}
	}
}
