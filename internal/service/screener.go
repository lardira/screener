package service

import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lardira/screener/internal/errors"
)

const (
	tps = 60

	SuccessfulRunExitCode = 125
)

type ScreenerConfig struct {
	SetFullscreen bool
	MaxDuration   time.Duration
}

type Screener struct {
	fullscreen  bool
	maxDuration time.Duration

	currentRuntimeTicks uint
}

func NewScreener(config *ScreenerConfig) *Screener {
	var s Screener
	s.fullscreen = config.SetFullscreen
	s.maxDuration = config.MaxDuration

	return &s
}

func (s *Screener) Update() error {
	s.currentRuntimeTicks++

	runtimeSeconds := time.Duration(s.currentRuntimeTicks/tps) * time.Second
	if runtimeSeconds >= s.maxDuration {
		return errors.ErrGameStop
	}

	return nil
}

func (s *Screener) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (s *Screener) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 100
}

func (s *Screener) Run() {
	ebiten.SetWindowSize(1, 1)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetTPS(tps)
	ebiten.SetWindowFloating(true) // above others
	ebiten.SetFullscreen(true)

	options := &ebiten.RunGameOptions{
		InitUnfocused:     false,
		ScreenTransparent: true,
		SkipTaskbar:       true,
	}

	err := PlayScream()
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGameWithOptions(s, options); err != nil {
		log.Println(err)
		os.Exit(SuccessfulRunExitCode)
	}
}
