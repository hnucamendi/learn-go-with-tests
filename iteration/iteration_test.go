package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 25)
	expected := "aaaaaaaaaaaaaaaaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("Elaris Wilson", 2)
	fmt.Println(repeated)
	// Output: Elaris WilsonElaris Wilson
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 13)
	}
}
