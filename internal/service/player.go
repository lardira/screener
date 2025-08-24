package service

import (
	"bytes"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const (
	sampleRate = 41000
)

//go:embed assets/scream.mp3
var soundBytes []byte

func PlayScream() error {
	audioContext := audio.NewContext(sampleRate)

	s, err := mp3.DecodeF32(bytes.NewReader(soundBytes))
	if err != nil {
		return err
	}

	p, err := audioContext.NewPlayerF32(s)
	if err != nil {
		return err
	}

	p.SetVolume(0.5)
	p.Play()

	return nil
}
