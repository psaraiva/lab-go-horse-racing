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
	// DelayHorseStep is the time that the horse will wait to move
	DelayHorseStep time.Duration = time.Duration(300 * time.Millisecond)
	// DelayRefreshScreen is the time that the screen will wait to refresh
	DelayRefreshScreen time.Duration = time.Duration(100 * time.Millisecond)
	// HorseLabelDefault is the default label for the horse
	HorseLabelDefault string = "H"
	// HorseMaxSpeed is the maximum speed that the horse can reach
	HorseMaxSpeed int = 5
	// HorseMinSpeed is the minimum speed that the horse can reach
	HorseMinSpeed int = 1
	// HorseQuantityDefault is the default quantity of horses
	HorseQuantityDefault int = 2
	// HorseQuantityMax is the maximum quantity of horses
	HorseQuantityMax int = 99
	// HorseQuantityMin is the minimum quantity of horses
	HorseQuantityMin int = 2
	// ScoreTargetDefault is the default score target
	ScoreTargetDefault int = 75
	// ScoreTargetMin is the minimum score target
	ScoreTargetMin int = 15
	// ScoreTargetMax is the maximum score target
	ScoreTargetMax int = 100
	// GameTimeoutDefault is the default game timeout
	GameTimeoutDefault string = "10s"
	// GameTimeoutRegexPattern is the regex pattern for the game timeout
	GameTimeoutRegexPattern string = `^\d{1,2}s$`
)

var (
	horseLabel          = HorseLabelDefault
	horseQuantity       = HorseQuantityDefault
	scoreTarget         = ScoreTargetDefault
	gameTimeout         = GameTimeoutDefault
	gameTimeoutDuration = 10 * time.Second
)

// Horse struct
type Horse struct {
	Label string
	Score int
}

var horses = []*Horse{}

// Winner returns the winner phrase
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
	horseLabel = HorseLabelDefault
	temp := os.Getenv("HORSE_LABEL")
	if len(temp) == 1 {
		horseLabel = temp
	}
}

func setScoreTarget() {
	scoreTarget = ScoreTargetDefault
	tmp, err := strconv.Atoi(os.Getenv("SCORE_TARGET"))
	if err == nil && isValidScoreTarget(tmp) {
		scoreTarget = tmp
	}
}

func setHorseQuantity() {
	horseQuantity = HorseQuantityDefault
	tmp, err := strconv.Atoi(os.Getenv("HORSE_QUANTITY"))
	if err == nil && tmp >= HorseQuantityMin && tmp <= HorseQuantityMax {
		horseQuantity = tmp
	}
}

func setGameTimeout() {
	gameTimeout = GameTimeoutDefault
	tmp := os.Getenv("GAME_TIMEOUT")
	r, err := regexp.Compile(GameTimeoutRegexPattern)
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
	endRace := make(chan bool)
	for _, horse := range horses {
		go goHorse(horse, endRace)
	}

	go display()

	select {
	case <-endRace:
		clearTerminal()
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

func goHorse(target *Horse, endRace chan<- (bool)) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		target.Score += r.Intn(HorseMaxSpeed) + HorseMinSpeed
		if target.Score >= scoreTarget {
			endRace <- true
			break
		}
		time.Sleep(DelayHorseStep)
	}
}

func display() {
	for {
		clearTerminal()
		println(getRaceStr())
		time.Sleep(DelayRefreshScreen)
	}
}

func getHorseWinner() *Horse {
	horseWinner := &Horse{}
	for _, horse := range horses {
		if horseWinner.Label == "" {
			horseWinner = horse
			continue
		}

		if horse.Score > horseWinner.Score {
			horseWinner = horse
		}
	}
	return horseWinner
}

func clearTerminal() {
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
		scoreTarget = ScoreTargetDefault
	}

	if scoreTarget-horse.Score > 0 {
		more = strings.Repeat(" ", scoreTarget-horse.Score-1)
	}

	less := strings.Repeat(".", horse.Score)
	return fmt.Sprintf("%s|%v%v%v|", horse.Label, less, horse.Label, more)
}

func generateTrackLimit(scoreTarget int) string {
	if !isValidScoreTarget(scoreTarget) {
		scoreTarget = ScoreTargetDefault
	}

	return "   +" + strings.Repeat("-", scoreTarget+2) + "+"
}

func clearHorses() {
	horses = []*Horse{}
}

func isValidScoreTarget(scoreTarget int) bool {
	return scoreTarget >= ScoreTargetMin && scoreTarget <= ScoreTargetMax
}
