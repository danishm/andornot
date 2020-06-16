package identity

import "testing"

func TestGet(t *testing.T) {
	cases := []struct {
		name     string
		expected string
	}{
		{"and", "and1"},
		{"and", "and2"},
		{"or", "or1"},
		{"board", "board1"},
		{"and", "and3"},
		{"board", "board2"},
	}

	for _, c := range cases {
		actual := Get(c.name)
		if actual != c.expected {
			t.Logf("Get(%s) expected:%s got:%s", c.name, c.expected, actual)
			t.Fail()
		}
	}
}
