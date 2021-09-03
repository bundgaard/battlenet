package cardback

import "testing"

func TestSearchRequest(t *testing.T) {

	tests := []struct {
		Input    *SearchRequest
		Expected string
	}{
		{Search(Locale("en_US")), "locale=en_US"},
		{Search(Locale("en_US"), Sort(NameDesc)), "locale=en_US&sort=name%3Adesc"},
	}

	for idx, test := range tests {
		got := test.Input.Encode()
		if got != test.Expected {
			t.Errorf("test[%04d] -- expected %q. got %q", idx, test.Expected, got)
		}
	}

}
