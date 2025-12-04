package utils

/*
* No, I'm not crazy.
* I created this test file because I already had this boilerplate test function
* somewhere else and it was quicker than creating a dedicated main.go to
* just try out a bunch of test cases for this function :)
 */

import (
	"reflect"
	"testing"
)

func TestSplitStringEqual(t *testing.T) {
	run := func(name string, s string, parts int, want []string, wantPanic bool) {
		t.Run(name, func(t *testing.T) {
			var got []string
			if wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("expected panic, but none occurred")
					}
				}()
				got = SplitStringEqual(s, parts)
			} else {
				got = SplitStringEqual(s, parts)
			}
			if !wantPanic && !reflect.DeepEqual(got, want) {
				t.Fatalf("got %v, want %v", got, want)
			}
		})
	}

	run("normal case",
		"hello world", 3,
		[]string{"hel", "lo ", "wor"},
		false,
	)

	run("even split",
		"abcd", 2,
		[]string{"ab", "cd"},
		false,
	)

	run("empty string",
		"", 3,
		[]string{"", "", ""},
		false,
	)

	run("parts greater than length",
		"abcd", 5,
		[]string{"", "", "", "", ""},
		false,
	)

	run("zero parts (panic)",
		"abcd", 0,
		nil,
		true,
	)

	run("123123123",
		"123123123", 3,
		[]string{"123", "123", "123"},
		false)
}
