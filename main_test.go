package vtune

import "testing"

var thresholdTests = map[string]struct {
	baseThreshold, scaleFactor, tuples, want float64
}{
	"defaults":            {50, 0.2, 1000, 250.0},
	"zero base threshold": {0, 0.2, 1000, 200.0},
	"zero scale factor":   {50, 0, 1000, 50.0},
	"zero tuples":         {50, 0.2, 0, 50.0},
}

func TestGetAutovacuumThreshold(t *testing.T) {
	for name, test := range thresholdTests {
		t.Run(name, func(t *testing.T) {
			got := getAutovacuumThreshold(test.baseThreshold, test.scaleFactor, test.tuples)
			if got != test.want {
				t.Errorf("got %f, want %f", got, test.want)
			}
		})
	}
}
