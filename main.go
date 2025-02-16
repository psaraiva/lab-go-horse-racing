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
	DELAY_HORSE_STEP     = time.Duration(300 * time.Millisecond)
	DELAY_REFRESH_SCREEN = time.Duration(100 * time.Millisecond)
	HORSE_LABEL          = "H"
	HORSE_MAX_SPEED      = 5
	HORSE_MIN_SPEED      = 1
	HORSES_QUANTITY      = 10
	HORSES_MAX_QUANTITY  = 99
	SCORE_TARGET         = 75
	TIMEOUT_GAME         = time.Duration(10 * time.Second)
)

type Horse struct {
	Label string
	Score int
}

var horses = []*Horse{}

func (h *Horse) Winner() string {
	return fmt.Sprintf("The horse winner is: %s - Score %d", h.Label, h.Score)
}

func main() {
	run()
}

func run() {
	load_horses()
	end_race := make(chan bool)
	for _, horse := range horses {
		go goHorse(horse, end_race)
	}

	go display()

	select {
	case <-end_race:
		clear_terminal()
		print_race()
	case <-time.After(TIMEOUT_GAME):
		fmt.Println("The Horses Are Tired!")
	}

	println(getHorseWinner().Winner())
}

func load_horses() {
	qtd_horses := 2
	if HORSES_QUANTITY > 1 || HORSES_QUANTITY <= HORSES_MAX_QUANTITY {
		qtd_horses = HORSES_QUANTITY
	}

	for i := range qtd_horses {
		index := i + 1
		prefix := ""
		if index < 10 {
			prefix = "0"
		}

		horses = append(horses, &Horse{Label: HORSE_LABEL + prefix + strconv.Itoa(index), Score: 0})
	}
}

func goHorse(target *Horse, end_race chan (bool)) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		target.Score += r.Intn(HORSE_MAX_SPEED) + HORSE_MIN_SPEED
		if target.Score >= SCORE_TARGET {
			end_race <- true
			break
		}
		time.Sleep(DELAY_HORSE_STEP)
	}
}

func display() {
	for {
		clear_terminal()
		print_race()
		time.Sleep(DELAY_REFRESH_SCREEN)
	}
}

func getHorseWinner() *Horse {
	horse_winner := &Horse{Label: "", Score: 0}
	for _, horse := range horses {
		if horse_winner.Label == "" {
			horse_winner = horse
			continue
		}

		if horse.Score > horse_winner.Score {
			horse_winner = horse
		}
	}
	return horse_winner
}

func clear_terminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func print_race() {
	println(generateTrackLimit())
	for _, horse := range horses {
		println(generateHorseTrack(horse, horse.Score))
	}
	println(generateTrackLimit())
}

func generateHorseTrack(horse *Horse, score int) string {
	more := ""
	if SCORE_TARGET-score > 0 {
		more = strings.Repeat(" ", SCORE_TARGET-score-1)
	}

	less := strings.Repeat(".", score)
	return fmt.Sprintf("%s|%v%v%v|", horse.Label, less, horse.Label, more)
}

func generateTrackLimit() string {
	return "   +" + strings.Repeat("-", SCORE_TARGET+2) + "+"
}
