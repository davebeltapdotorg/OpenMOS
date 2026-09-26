package service

import "testing"

func TestResolveItemTiming(t *testing.T) {
	for _, tc := range []struct {
		name                   string
		item, object, rate     string
		samples, seconds, base int
		haveSamples, haveRate  bool
	}{
		{"editorial wins", "120", "300", "59.94", 120, 2, 60, true, true},
		{"object fallback", "", "0x96", "50", 150, 3, 50, true, true},
		{"uppercase object hex", "", "0X96", "50", 150, 3, 50, true, true},
		{"rate missing", "150", "", "", 150, 0, 0, true, false},
		{"invalid rate", "150", "", "NaN", 150, 0, 0, true, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			samples, haveSamples, seconds, base, haveRate := resolveItemTiming(tc.item, tc.object, tc.rate)
			if samples != tc.samples || haveSamples != tc.haveSamples || seconds != tc.seconds || base != tc.base || haveRate != tc.haveRate {
				t.Fatalf("got (%d, %t, %d, %d, %t)", samples, haveSamples, seconds, base, haveRate)
			}
		})
	}
}
