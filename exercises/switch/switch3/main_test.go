// switch3
// Make me compile!

package main_test

import "testing"

func weekDay(day int) string {
	// Return the day of the week based on the
	// integer. Use a switch case to satisfy all test cases below
	weekDay := "Saturday"
	switch day {
		case 0: weekDay = "Sunday" 
		case 1: weekDay = "Monday" 
		case 2: weekDay = "Tuesday" 
		case 3: weekDay = "Wednesday" 
		case 4: weekDay = "Thursday" 
		case 5: weekDay = "Friday" 
	}
	return weekDay
}

func TestWeekDay(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{input: 0, want: "Sunday"},
		{input: 1, want: "Monday"},
		{input: 2, want: "Tuesday"},
		{input: 3, want: "Wednesday"},
		{input: 4, want: "Thursday"},
		{input: 5, want: "Friday"},
		{input: 6, want: "Saturday"},
	}

	for _, tc := range tests {
		day := weekDay(tc.input)
		if day != tc.want {
			t.Errorf("expected %s but got %s", tc.want, day)
		}
	}
}
