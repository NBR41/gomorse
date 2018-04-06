package morse

import (
	"reflect"
	"testing"
)

type testplayer struct {
	ch   chan bool
	buf  []int16
	buf2 []int16
}

func (p *testplayer) InitSoundDevice() error {
	return nil
}

func (p *testplayer) FlushSoundBuffer() {

}
func (p *testplayer) WaitLine() {
	<-p.ch
}
func (p *testplayer) Play(buf1, buf2 []int16) {
	p.buf = buf1
	p.buf2 = buf2
	p.ch <- true

}

func TestPlay(t *testing.T) {
	p := &testplayer{ch: make(chan bool)}
	ru := []rune{'E', ' ', 'T'}
	var duration int64 = 1
	var freq, bar float64 = 2, 3
	exp := getBuffer(ru, duration, freq, bar)
	Play(p, ru, duration, freq, bar)
	if !reflect.DeepEqual(p.buf, exp) {
		t.Error("unexpected buffer")
	}
	if !reflect.DeepEqual(p.buf2, exp) {
		t.Error("unexpected buffer")
	}
}

func TestGetSound(t *testing.T) {
	v := getSound(3, 2, 2)
	exp := []int16{0, 31}
	if !reflect.DeepEqual(exp, v) {
		t.Error("unexpected sound")
	}
}
