package beep

import (
	"math"
	"testing"
)

func TestHertzToFreq(t *testing.T) {
	v := HertzToFreq(SampleRate64)
	exp := 2.0 * math.Pi
	if v != exp {
		t.Errorf("unexpected value: exp %f got %f", exp, v)
	}
}

func TestVolumeToBar(t *testing.T) {
	v := VolumeToBar(100)
	if v != SampleAmp16bit {
		t.Errorf("unexpected value: exp %f got %f", SampleAmp16bit, v)
	}
}

func TestDurationToSampleSize(t *testing.T) {
	v := DurationToSampleSize(1000)
	if v != int(SampleRate64) {
		t.Errorf("unexpected value: exp %d got %d", int(SampleRate64), v)
	}
}
