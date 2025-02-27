package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	DELAY_HORSE_STEP           = time.Duration(300 * time.Millisecond)
	DELAY_REFRESH_SCREEN       = time.Duration(100 * time.Millisecond)
	HORSE_LABEL_DEFAULT        = "H"
	HORSE_MAX_SPEED            = 5
	HORSE_MIN_SPEED            = 1
	HORSE_QUANTITY_DEFAULT     = 2
	HORSE_QUANTITY_MAX         = 99
	HORSE_QUANTITY_MIN         = 2
	SCORE_TARGET_DEFAULT       = 75
	SCORE_TARGET_MIN           = 15
	SCORE_TARGET_MAX           = 100
	GAME_TIMEOUT_DEFAULT       = "10s"
	GAME_TIMEOUT_REGEX_PATTERN = `^\d{1,2}s$`
)

var (
	horseLabel          = HORSE_LABEL_DEFAULT
	horseQuantity       = HORSE_QUANTITY_DEFAULT
	scoreTarget         = SCORE_TARGET_DEFAULT
	gameTimeout         = GAME_TIMEOUT_DEFAULT
	gameTimeoutDuration = 10 * time.Second
)

type Horse struct {
	Label string
	Score int
}

var horses = []*Horse{}

func (h *Horse) Winner() string {
	return fmt.Sprintf("The horse winner is: %s - Score %d", h.Label, h.Score)
}

// Load config enf file
// If file not found, the default config is apply
func loadConfig(fileEnv string) {
	godotenv.Load(fileEnv)
	intEnv()
}

func intEnv() {
	setHorseLabel()
	setHorseQuantity()
	setScoreTarget()
	setGameTimeout()
	setGameTimeoutDuration()
}

func setHorseLabel() {
	horseLabel = HORSE_LABEL_DEFAULT
	temp := os.Getenv("HORSE_LABEL")
	if len(temp) == 1 {
		horseLabel = temp
	}
}

func setScoreTarget() {
	scoreTarget = SCORE_TARGET_DEFAULT
	tmp, err := strconv.Atoi(os.Getenv("SCORE_TARGET"))
	if err == nil && isValidScoreTarget(tmp) {
		scoreTarget = tmp
	}
}

func setHorseQuantity() {
	horseQuantity = HORSE_QUANTITY_DEFAULT
	tmp, err := strconv.Atoi(os.Getenv("HORSE_QUANTITY"))
	if err == nil && tmp >= HORSE_QUANTITY_MIN && tmp <= HORSE_QUANTITY_MAX {
		horseQuantity = tmp
	}
}

func setGameTimeout() {
	gameTimeout = GAME_TIMEOUT_DEFAULT
	tmp := os.Getenv("GAME_TIMEOUT")
	r, err := regexp.Compile(GAME_TIMEOUT_REGEX_PATTERN)
	if err == nil && r.MatchString(tmp) {
		gameTimeout = tmp
	}
}

func setGameTimeoutDuration() {
	gameTimeoutDuration = 10 * time.Second
	tmp, err := time.ParseDuration(gameTimeout)
	if err == nil {
		gameTimeoutDuration = tmp
	}
}

func main() {
	run()
}

func run() {
	loadConfig(".env")
	loadHorses(horseQuantity)
	end_race := make(chan bool)
	for _, horse := range horses {
		go goHorse(horse, end_race)
	}

	go display()

	select {
	case <-end_race:
		clear_terminal()
		println(getRaceStr())
	case <-time.After(gameTimeoutDuration):
		fmt.Println("The Horses Are Tired!")
	}

	println(getHorseWinner().Winner())
}

func loadHorses(horsesQuantity int) {
	clearHorses()
	for i := range horsesQuantity {
		index := i + 1
		prefix := ""
		if index < 10 {
			prefix = "0"
		}

		horses = append(horses, &Horse{Label: horseLabel + prefix + strconv.Itoa(index)})
	}
}

func goHorse(target *Horse, end_race chan<- (bool)) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		target.Score += r.Intn(HORSE_MAX_SPEED) + HORSE_MIN_SPEED
		if target.Score >= scoreTarget {
			end_race <- true
			break
		}
		time.Sleep(DELAY_HORSE_STEP)
	}
}

func display() {
	for {
		clear_terminal()
		println(getRaceStr())
		time.Sleep(DELAY_REFRESH_SCREEN)
	}
}

func getHorseWinner() *Horse {
	horse_winner := &Horse{}
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

func getRaceStr() string {
	msg := ""
	msg += generateTrackLimit(scoreTarget) + "\n"
	for _, horse := range horses {
		msg += generateHorseTrack(horse, scoreTarget) + "\n"
	}
	msg += generateTrackLimit(scoreTarget) + "\n"
	return msg
}

func generateHorseTrack(horse *Horse, scoreTarget int) string {
	more := ""
	if !isValidScoreTarget(scoreTarget) {
		scoreTarget = SCORE_TARGET_DEFAULT
	}

	if scoreTarget-horse.Score > 0 {
		more = strings.Repeat(" ", scoreTarget-horse.Score-1)
	}

	less := strings.Repeat(".", horse.Score)
	return fmt.Sprintf("%s|%v%v%v|", horse.Label, less, horse.Label, more)
}

func generateTrackLimit(scoreTarget int) string {
	if !isValidScoreTarget(scoreTarget) {
		scoreTarget = SCORE_TARGET_DEFAULT
	}

	return "   +" + strings.Repeat("-", scoreTarget+2) + "+"
}

func clearHorses() {
	horses = []*Horse{}
}

func isValidScoreTarget(scoreTarget int) bool {
	return scoreTarget >= SCORE_TARGET_MIN && scoreTarget <= SCORE_TARGET_MAX
}
