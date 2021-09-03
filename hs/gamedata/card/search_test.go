// Copyright (C) 2021 David Bundgaard
package card

import "testing"

func TestSearchRequest(t *testing.T) {

	tests := []struct {
		Input    *SearchRequest
		Expected string
	}{
		{Search(Set("standard"), Attack(1, 3)), "attack=1%2C3&set=standard"},
		{Search(Locale("en_US"), Attack(1, 2, 3, 4, 5, 6, 7, 8), Set("wild"), Manacost(5), Health(3)), "attack=1%2C2%2C3%2C4%2C5%2C6%2C7%2C8&health=3&locale=en_US&manacost=5&set=wild"},
	}

	for idx, test := range tests {
		got := test.Input.Encode()
		if got != test.Expected {
			t.Errorf("test[%04d] -- expected %q. got %q", idx, test.Expected, got)
		}
	}

}
