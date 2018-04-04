// content of reverse_test.go
package reverse

import "testing"

func TestReverseToReturnReversedInputString(t *testing.T) {
	for value, expected := range reverseDataProvider() {
		actual := Reverse(value)
		if actual != expected {
			t.Fatalf("Expected %s but got %s", expected, actual)
		}
	}
}

func reverseDataProvider() map[string]string {
	return map[string]string{
		"Hello": "olleH",
		"Привіт": "тівирП",
		"": "",
		"1": "1",
	}
}