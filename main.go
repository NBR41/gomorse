package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/NBR41/gomorse/beep"
	"github.com/NBR41/gomorse/morse"
)

var (
	volume, duration int64
	freqH, freq, bar float64
)

func main() {
	var cmd = &cobra.Command{
		Use:   "gomorse",
		Short: "gomorse is a tool to transcript phrases to morse code",
		Long: `gomorse is a tool to transcript phrases to morse code.
3 parameters:
- volume
- beep duration
- beep frequence`,
		PreRun: func(cmd *cobra.Command, args []string) {
			bar = beep.VolumeToBar(volume)
			freq = beep.HertzToFreq(freqH)
		},
		Run: func(cmd *cobra.Command, args []string) {
			player := beep.NewPlayer()
			scanner := bufio.NewScanner(os.Stdin)
			beep.OpenSoundDevice("default")
			defer beep.CloseSoundDevice()
			for {
				fmt.Print("Type your phrase to code> ")
				scanner.Scan()
				s := strings.ToUpper(scanner.Text())
				ru := bytes.Runes([]byte(s))
				if len(ru) == 0 {
					fmt.Println("Empty input")
					continue
				}
				fmt.Print("Checking input: ")
				if err := morse.CheckInput(ru); err != nil {
					fmt.Println(err.Error())
					continue
				}
				fmt.Println("OK")
				fmt.Print("Type enter to play code>")
				scanner.Scan()
				morse.Play(player, ru, duration, freq, bar)
			}
		},
	}
	cmd.Flags().Int64VarP(&volume, "volume", "v", 100, "beep volume (0 to 100)")
	cmd.Flags().Int64VarP(&duration, "duration", "d", 250, "beep duration in ms")
	cmd.Flags().Float64VarP(&freqH, "frequence", "f", 523.25, "frequency in Hertz (1-22050)")
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
