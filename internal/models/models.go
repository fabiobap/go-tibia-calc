package models

import (
	"math"

	"github.com/fabiobap/go-tibia-calc/internal/forms"
)

const BASE_HITPOINTS = 150
const HP_LVL_MAGE = 5
const HP_LVL_NONE = 5
const HP_LVL_PALADIN = 10
const HP_LVL_KNIGHT = 15
const BASE_MANAPOINTS = 55
const MANA_LVL_MAGE = 30
const MANA_LVL_NONE = 5
const MANA_LVL_PALADIN = 15
const MANA_LVL_KNIGHT = 5
const BASE_CAP = 400
const CAP_LVL_NONE = 10
const CAP_LVL_MAGE = 10
const CAP_LVL_PALADIN = 20
const CAP_LVL_KNIGHT = 25
const BASE_EXP = 0
const BASE_LEVEL = 1

type TemplateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	FlashMessage    string
	Warning         string
	Error           string
	Form            *forms.Form
	IsAuthenticated int
}

type Character struct {
	Level             int
	Vocation          string
	Hitpoints         int
	Manapoints        int
	Experience        int
	Cap               int
	BlessingCostOne   int
	BlessingCostFive  int
	BlessingCostSeven int
	BlessingCostFull  int
}

func (f *Character) Load() {
	f.Hitpoints = calcHitpoints(f.Level, f.Vocation)
	f.Manapoints = calcManapoints(f.Level, f.Vocation)
	f.Cap = calcCap(f.Level, f.Vocation)
	f.Experience = CalcExp(f.Level)
	f.BlessingCostOne = calcOneBless(f.Level)
	f.BlessingCostFive = calcFiveBless(f.Level)
	f.BlessingCostSeven = calcSevenBless(f.Level)
	f.BlessingCostFull = calcFullBless(f.Level)
}

func CalcExp(lvl int) int {
	l := float64(lvl)
	return int((50*math.Pow(l-1, 3) - 150*math.Pow(l-1, 2) + 400*(l-1)) / 3)
}

func calcOneBless(lvl int) int {
	if lvl <= 30 {
		return 2000
	}

	return 200 * (lvl - 20)
}

func calcFiveBless(lvl int) int {
	if lvl <= 30 {
		return 2000 * 5
	}

	return lvl - 20
}

func calcSevenBless(lvl int) int {
	if lvl <= 30 {
		return calcFiveBless(lvl) + 5200
	}

	return int((float64(lvl) - 20) * 1.52)
}

func calcFullBless(lvl int) int {
	if lvl <= 20 {
		return calcSevenBless(lvl)
	}

	return int((float64(lvl) - 20) * 1.72)
}

func calcHitpoints(level int, vocation string) int {
	if level == BASE_LEVEL {
		return BASE_HITPOINTS
	}

	if level < 9 {
		return BASE_HITPOINTS + (level * HP_LVL_NONE) - HP_LVL_NONE
	}

	base := BASE_HITPOINTS + (7 * HP_LVL_NONE)
	var hitpoints int

	switch vocation {
	case "mage":
		hitpoints = base + (level-8)*HP_LVL_MAGE
	case "paladin":
		hitpoints = base + (level-8)*HP_LVL_PALADIN
	case "knight":
		hitpoints = base + (level-8)*HP_LVL_KNIGHT
	default:
		hitpoints = base + (level-8)*HP_LVL_NONE
	}

	return hitpoints
}

func calcManapoints(level int, vocation string) int {
	if level == BASE_LEVEL {
		return BASE_MANAPOINTS
	}

	if level < 9 {
		return BASE_MANAPOINTS + (level * MANA_LVL_NONE) - MANA_LVL_NONE
	}

	base := BASE_MANAPOINTS + (7 * MANA_LVL_NONE)
	var manapoints int

	switch vocation {
	case "mage":
		manapoints = base + (level-8)*MANA_LVL_MAGE
	case "paladin":
		manapoints = base + (level-8)*MANA_LVL_PALADIN
	case "knight":
		manapoints = base + (level-8)*MANA_LVL_KNIGHT
	default:
		manapoints = base + (level-8)*MANA_LVL_NONE
	}

	return manapoints
}

func calcCap(level int, vocation string) int {
	if level == BASE_LEVEL {
		return BASE_CAP
	}

	if level < 9 {
		return BASE_CAP + (level * CAP_LVL_NONE) - CAP_LVL_NONE
	}

	base := BASE_CAP + (7 * CAP_LVL_NONE)
	var cap int

	switch vocation {
	case "mage":
		cap = base + (level-8)*CAP_LVL_MAGE
	case "paladin":
		cap = base + (level-8)*CAP_LVL_PALADIN
	case "knight":
		cap = base + (level-8)*CAP_LVL_KNIGHT
	default:
		cap = base + (level-8)*CAP_LVL_NONE

	}

	return cap
}
