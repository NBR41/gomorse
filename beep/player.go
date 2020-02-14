// Package beep play beep sound
package beep

import (
	"math"
)

const (
	// SampleAmp16bit - 16-bit sample amplitude
	SampleAmp16bit = 32767.0

	// SampleRate - sample rate
	SampleRate = 44100

	// SampleRate64 - float64 sample rate
	SampleRate64 = float64(SampleRate)
)

// Player is a music player
type Player struct {
	stopping   bool
	played     chan bool // for syncing player
	linePlayed chan bool // for syncing lines
}

// NewPlayer returns new player
func NewPlayer() *Player {
	return &Player{
		played:     make(chan bool),
		linePlayed: make(chan bool),
	}
}

// Wait until sheet is played
func (m *Player) Wait() {
	<-m.played
}

// WaitLine waits until line is played
func (m *Player) WaitLine() {
	<-m.linePlayed
}

// HertzToFreq converts Hertz to frequency unit
func HertzToFreq(hertz float64) float64 {
	// 1 second = 44100 samples
	// 1 hertz = freq * 2Pi
	// freq = 2Pi / 44100 * hertz
	freq := 2.0 * math.Pi / SampleRate64 * hertz
	return freq
}

// VolumeToBar converts volume to bar
func VolumeToBar(volume int64) float64 {
	return SampleAmp16bit * (float64(volume) / 100.0)
}

// DurationToSampleSize converts duration to sample size
// SampleRate64 = 1s
func DurationToSampleSize(duration int64) int {
	return int(SampleRate64 * (float64(duration) / 1000.0))
}
