// +build linux

package beep

/*
#cgo LDFLAGS: -lasound

#include <alsa/asoundlib.h>

void *cbuf[2];
*/
import "C"

import (
	"fmt"
	"os"
	"unsafe"
)

var (
	pcmHandle   *C.snd_pcm_t
	pcmHwParams *C.snd_pcm_hw_params_t
)

const (
	bitsPerSample = 16
	sample16bit   = bitsPerSample == 16
)

func strerror(code C.int) string {
	return C.GoString(C.snd_strerror(code))
}

// OpenSoundDevice opens hardware sound device
func OpenSoundDevice(device string) error {
	code := C.snd_pcm_open(&pcmHandle, C.CString(device), C.SND_PCM_STREAM_PLAYBACK, 0)
	if code < 0 {
		err := fmt.Errorf("snd_pcm_open: %v", strerror(code))
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	C.snd_pcm_drop(pcmHandle)

	return nil
}

// InitSoundDevice initialize sound device
func InitSoundDevice() error {
	var sampleFormat C.snd_pcm_format_t = C.SND_PCM_FORMAT_S8
	if sample16bit {
		sampleFormat = C.SND_PCM_FORMAT_S16
	}

	if code := C.snd_pcm_hw_params_malloc(&pcmHwParams); code < 0 {
		err := fmt.Errorf("snd_pcm_hw_params_malloc: %v", strerror(code))
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	if code := C.snd_pcm_hw_params_any(pcmHandle, pcmHwParams); code < 0 {
		err := fmt.Errorf("snd_pcm_hw_params_any: %v", strerror(code))
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	// C.SND_PCM_ACCESS_RW_NONINTERLEAVED - is not working with PulseAudio
	code := C.snd_pcm_set_params(pcmHandle, sampleFormat, C.SND_PCM_ACCESS_RW_INTERLEAVED, 1, 44100, 1, 500000)
	if code < 0 {
		err := fmt.Errorf("snd_pcm_set_params: %v", strerror(code))
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	code = C.snd_pcm_prepare(pcmHandle)
	if code < 0 {
		err := fmt.Errorf("snd_pcm_prepare: %v", strerror(code))
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}

// CloseSoundDevice closes sound device
func CloseSoundDevice() {
	if pcmHandle != nil {
		C.snd_pcm_close(pcmHandle)
		C.snd_pcm_hw_free(pcmHandle)
	}
}

// FlushSoundBuffer flushes sound buffer
func FlushSoundBuffer() {
	if pcmHandle != nil {
		C.snd_pcm_drain(pcmHandle)
	}
}

// Play sends stereo wave buffer to sound device
func (m *Player) Play(buf1, buf2 []int16) {
	bufsize := len(buf1)
	if bufsize < SampleRate {
		// prevent buffer underrun
		rest := make([]int16, SampleRate)
		buf1 = append(buf1, rest...)
		//buf2 = append(buf2, rest...)
	}

	// Changing to single channel interleaved buffer format for PulseAudio
	buf := unsafe.Pointer(&buf1[0])

	for {
		n := C.snd_pcm_writei(pcmHandle, buf, C.snd_pcm_uframes_t(bufsize))
		written := int(n)
		if written < 0 {
			if m.stopping {
				break
			}
			// error
			code := C.int(written)
			written = 0
			_ = written
			fmt.Fprintln(os.Stderr, "snd_pcm_writei:", code, strerror(code))
			code = C.snd_pcm_recover(pcmHandle, code, 0)
			if code < 0 {
				fmt.Fprintln(os.Stderr, "snd_pcm_recover:", strerror(code))
				break
			}
		}
		break // don't retry, breaks timing
	}
	m.linePlayed <- true // notify that playback is done
}
