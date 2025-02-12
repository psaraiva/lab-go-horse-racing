package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	DELAY_REFRESH_SCREEN = time.Duration(300 * time.Millisecond)
	SCORE_LIMIT          = 50
	STEP_LIMIT           = 5
	QUANTITY_HORSES      = 7
)

type Horse struct {
	Label string
	Score int
}

var horses = []*Horse{}

func main() {
	load_horses()
	run()
}

func load_horses() {
	qtd_horses := 2
	if QUANTITY_HORSES > 1 || QUANTITY_HORSES < 11 {
		qtd_horses = QUANTITY_HORSES
	}

	for i := range qtd_horses {
		horses = append(horses, &Horse{Label: "H" + strconv.Itoa(i+1), Score: 0})
	}
}

func run() {
	end_race := make(chan bool)

	for _, horse := range horses {
		go func(target *Horse) {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			for {
				target.Score += r.Intn(STEP_LIMIT) + 1
				if target.Score >= SCORE_LIMIT {
					end_race <- true
					break
				}
				time.Sleep(DELAY_REFRESH_SCREEN)
			}
		}(horse)
	}

	go func() {
		for {
			clear_terminal()
			print_race()
			time.Sleep(time.Duration(100 * time.Millisecond))
		}
	}()

	select {
	case <-end_race:
		clear_terminal()
		print_race()
	case <-time.After(10 * time.Second):
		fmt.Println("The Horses Are Tired!")
	}
}

func clear_terminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func print_race() {
	println(generateTrackLimit())
	for _, horse := range horses {
		println(generateHorseTrack(horse.Label, horse.Score))
	}
	println(generateTrackLimit())
}

func generateHorseTrack(horse string, score int) string {
	more := ""
	if SCORE_LIMIT-score > 0 {
		more = strings.Repeat(" ", SCORE_LIMIT-score)
	}

	less := strings.Repeat(".", score)
	return fmt.Sprintf("|%v%v%v|", less, horse, more)
}

func generateTrackLimit() string {
	limit := "+"
	limit += strings.Repeat("-", SCORE_LIMIT+2)
	limit += "+"
	return limit
}
