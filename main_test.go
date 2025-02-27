package main

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	t.Setenv("HORSE_LABEL", "C")
	t.Setenv("HORSE_QUANTITY", "10")
	t.Setenv("SCORE_TARGET", "30")
	t.Setenv("GAME_TIMEOUT", "15s")

	intEnv()

	assert.Equal(t, "C", horseLabel)
	assert.Equal(t, 10, horseQuantity)
	assert.Equal(t, 30, scoreTarget)
	assert.Equal(t, "15s", gameTimeout)
	assert.Equal(t, 15*time.Second, gameTimeoutDuration)
}

func TestLoadConfigWithInvalidValures(t *testing.T) {
	t.Setenv("HORSE_LABEL", "ABACATE")
	t.Setenv("HORSE_QUANTITY", "999")
	t.Setenv("SCORE_TARGET", "150")
	t.Setenv("GAME_TIMEOUT", "120s")

	intEnv()

	assert.Equal(t, HORSE_LABEL_DEFAULT, horseLabel)
	assert.Equal(t, HORSE_QUANTITY_DEFAULT, horseQuantity)
	assert.Equal(t, SCORE_TARGET_DEFAULT, scoreTarget)
	assert.Equal(t, GAME_TIMEOUT_DEFAULT, gameTimeout)
	assert.Equal(t, 10*time.Second, gameTimeoutDuration)
}

func TestIsScoreTargetValid(t *testing.T) {
	assert.True(t, isValidScoreTarget(55))
}

func TestIsScoreTargetValidWithValueInvalid(t *testing.T) {
	assert.False(t, isValidScoreTarget(999))
}

func TestGenerateTrackLimit(t *testing.T) {
	expected := "   +-----------------+"
	found := generateTrackLimit(15)
	assert.Equal(t, expected, found)
}

func TestGenerateTrackLimitWithInvalidVal(t *testing.T) {
	expected := "   +" + strings.Repeat("-", scoreTarget) + "--+"
	found := generateTrackLimit(-10)
	assert.Equal(t, expected, found)
}

func TestGenerateHorseTrack(t *testing.T) {
	horse := Horse{Label: "H01", Score: 21}
	expected := "H01|.....................H01|"
	found := generateHorseTrack(&horse, 20)
	assert.Equal(t, expected, found)
}

func TestGenerateHorseTrackWithInvalidValue(t *testing.T) {
	horse := Horse{Label: "H01", Score: 70}
	expected := "H01|......................................................................H01    |"
	found := generateHorseTrack(&horse, -10)
	assert.Equal(t, expected, found)
}

func TestGetHorseWinner(t *testing.T) {
	horses = []*Horse{}
	horses = append(horses, &Horse{Label: "H01", Score: 21})
	horses = append(horses, &Horse{Label: "H02", Score: 25})
	horses = append(horses, &Horse{Label: "H03", Score: 30})

	expected := horses[2]
	found := getHorseWinner()
	assert.Equal(t, expected, found)
}

func TestClearHorses(t *testing.T) {
	horses = []*Horse{}
	horses = append(horses, &Horse{Label: "H01", Score: 21})
	horses = append(horses, &Horse{Label: "H02", Score: 25})
	horses = append(horses, &Horse{Label: "H03", Score: 30})
	assert.Equal(t, 3, len(horses))
	clearHorses()
	assert.Equal(t, 0, len(horses))
}

func TestLoadHorses(t *testing.T) {
	loadHorses(5)
	assert.Equal(t, 5, len(horses))
	assert.Equal(t, horses[0].Label, "H01")
	assert.Equal(t, horses[1].Label, "H02")
	assert.Equal(t, horses[2].Label, "H03")
	assert.Equal(t, horses[3].Label, "H04")
	assert.Equal(t, horses[4].Label, "H05")
}

func TestWinner(t *testing.T) {
	horse := Horse{Label: "H01", Score: 21}
	assert.Equal(t, "The horse winner is: H01 - Score 21", horse.Winner())
}

func TestGetRaceStr(t *testing.T) {
	loadHorses(2)
	horses[0].Score = 5
	horses[1].Score = 7

	scoreTarget = SCORE_TARGET_DEFAULT
	expected := "   +-----------------------------------------------------------------------------+\n"
	expected += "H01|.....H01                                                                     |\n"
	expected += "H02|.......H02                                                                   |\n"
	expected += "   +-----------------------------------------------------------------------------+\n"

	assert.Equal(t, expected, getRaceStr())
}
