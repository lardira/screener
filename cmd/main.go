package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/lardira/screener/internal/service"
)

const (
	maxDuration = 1 * time.Second

	tryInterval = 1 * time.Second

	targetChance = 0.005

	runCommand string = "scream"
)

func main() {

	// Check for game mode flag.
	if len(os.Args) > 1 && os.Args[1] == runCommand {
		screenerConfig := &service.ScreenerConfig{
			SetFullscreen: true,
			MaxDuration:   maxDuration,
		}

		screener := service.NewScreener(screenerConfig)
		screener.Run()
	}

	// Run main loop for creating subprocesses
	for {
		try := rand.Float64()
		fmt.Printf("try: %v, target: %v\n", try, targetChance)

		if try <= targetChance {
			fp, _ := filepath.Abs(os.Args[0])

			cmd := exec.Command(fp, runCommand)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			err := cmd.Run()
			if err != nil {
				if !strings.Contains(err.Error(), strconv.Itoa(service.SuccessfulRunExitCode)) {
					fmt.Println("Error running screener subprocess:", err)
				}
			}

			fmt.Println("Screener session completed")
		}
		time.Sleep(tryInterval)
	}
}
