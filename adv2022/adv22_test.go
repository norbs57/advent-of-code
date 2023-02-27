package adv2022

import "testing"

func TestSnafu(t *testing.T) {
	for i := 0; i < 10000; i++ {
		s := ToSnafu(i)
		j := FromSnafu(s)
		if i != j {
			t.Errorf("Expected %d, got %d, toSnafu= %v", i, j, s)
		}
	}
}
