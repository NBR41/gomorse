package morse

import (
	"math"

	"github.com/NBR41/gomorse/beep"
)

type player interface {
	InitSoundDevice() error
	FlushSoundBuffer()
	WaitLine()
	Play(buf1, buf2 []int16)
}

// Play play morse symbols according to given params
func Play(pl player, ru []rune, duration int64, freq, bar float64) {
	buf := getBuffer(ru, duration, freq, bar)
	_ = pl.InitSoundDevice()
	go pl.Play(buf, buf)
	pl.WaitLine()
	pl.FlushSoundBuffer()
}

func getBuffer(ru []rune, duration int64, freq, bar float64) []int16 {
	ti := beep.DurationToSampleSize(duration)
	ta := 3 * ti
	word := make([]int16, 6*ti)
	letter := make([]int16, ti)
	buf := make([]int16, 0)
	for i := range ru {
		if ru[i] == ' ' {
			buf = append(buf, word...)
			continue
		}
		for j := range Alphabet[ru[i]] {
			switch Alphabet[ru[i]][j] {
			case Ti:
				buf = append(buf, getSound(freq, bar, ti)...)
			case Ta:
				buf = append(buf, getSound(freq, bar, ta)...)
			}
		}

		if i != len(ru)-1 {
			buf = append(buf, letter...)
		}
	}
	return buf
}

func getSound(freq, bar float64, samples int) []int16 {
	buf := make([]int16, samples)
	var last int16
	var fade = 1024
	if samples < fade {
		fade = 1
	}
	for i := range buf {
		if i < samples-fade {
			buf[i] = int16(bar * math.Sin(float64(i)*freq))
			last = buf[i]
		} else {
			if last > 0 {
				last -= 31
			} else {
				last += 31
			}
			buf[i] = last
		}
	}
	return buf
}
