package pbReader

import (
	"github.com/dh1tw/remoteAudio/audio"
	"github.com/dh1tw/remoteAudio/audiocodec"
)

// Option is the type for a function option
type Option func(*Options)

type Options struct {
	DeviceName string
	Decoder    audiocodec.Decoder
	Channels   int
	Samplerate float64
	Callback   audio.OnDataCb
}

// Channels is a functional option to set the amount of channels to be used
// with the audio device. Typically this is either Mono (1) or Stereo (2).
// Make sure that your audio device supports the specified amount of channels.
func Channels(chs int) Option {
	return func(args *Options) {
		args.Channels = chs
	}
}

// Samplerate is a functional option to set the sampling rate of the
// audio device. Make sure your audio device supports the specified sampling
// rate.
func Samplerate(s float64) Option {
	return func(args *Options) {
		args.Samplerate = s
	}
}

// Decoder is a functional option to set a specific decoder. By default,
// the opus decoder is used.
func Decoder(dec audiocodec.Decoder) Option {
	return func(args *Options) {
		args.Decoder = dec
	}
}
