package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// Documents the function in "godoc"
func ExampleRepeat() {
	repeated := Repeat("b", 5)
	fmt.Println(repeated)
	// Output: bbbbb
}

// runs b.N times to measure how long it takes to run
// To run the benchmark do "go test =bench=. -v"
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ { //New version is b.Loop()
		Repeat("a", 5)
	}
}
