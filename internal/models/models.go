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
	Level               int
	Vocation            string
	Hitpoints           int
	Manapoints          int
	Experience          int
	Cap                 int
	BlessingRegularOne  int
	BlessingTwist       int
	BlessingRegularFive int
	BlessingSeven       int
	BlessingFull        int
}

type MidnightShard struct {
	Level      int
	Experience int
	Quantity   int
}

type StoneOfInsight struct {
	Level      int
	NewLevel   int
	Experience int
}

func (f *Character) Load() {
	f.Hitpoints = calcHitpoints(f.Level, f.Vocation)
	f.Manapoints = calcManapoints(f.Level, f.Vocation)
	f.Cap = calcCap(f.Level, f.Vocation)
	f.Experience = CalcExp(f.Level)
	f.BlessingRegularOne = calcOneBless(f.Level)
	f.BlessingTwist = calcTwistBless(f.Level)
	f.BlessingRegularFive = calcFiveBless(f.Level)
	f.BlessingSeven = calcSevenBless(f.Level)
	f.BlessingFull = calcFullBless(f.Level)
}

func (soi *StoneOfInsight) Load() {
	soiCalc := CalcSOI(soi.Level)
	soi.Experience = soiCalc
	soi.NewLevel = FindNewLevel(soiCalc)
}

func (m *MidnightShard) Load() {
	m.Experience = CalcMidnightShard(m.Level, m.Quantity)
}

func CalcMidnightShard(lvl, qty int) int {
	return (300 * qty) * lvl
}

func CalcSOI(lvl int) int {
	return int(100 * math.Pow(float64(lvl), 2))
}

func FindNewLevel(newExp int) int {
	x := math.Pow((3*float64(newExp)/100-3), 2) + 125/27
	y := 3*newExp/100 - 3

	a := math.Cbrt(math.Sqrt(x) + float64(y))
	b := math.Cbrt(-math.Sqrt(x) + float64(y))

	return int(math.Floor(a + b + 2))
}

func CalcExp(lvl int) int {
	l := float64(lvl)
	return int((50*math.Pow(l-1, 3) - 150*math.Pow(l-1, 2) + 400*(l-1)) / 3)
}

func calcOneBless(lvl int) int {
	if lvl <= 30 {
		return 2000
	}

	if lvl < 120 {
		return 200 * (lvl - 20)
	}

	return 20000 + 75*(lvl-120)
}

func calcFiveBless(lvl int) int {
	if lvl <= 30 {
		return 2000 * 5
	}

	return calcOneBless(lvl) * 5
}

func calcOneEspecialBless(lvl int) int {
	if lvl <= 30 {
		return 2600
	}

	if lvl < 120 {
		return 260 * (lvl - 20)
	}

	return 26000 + 100*(lvl-120)
}

func calcTwistBless(lvl int) int {
	if lvl <= 30 {
		return 2000
	}

	if lvl < 120 {
		return 200 * (lvl - 20)
	}

	return 50000
}

func calcSevenBless(lvl int) int {
	return calcFiveBless(lvl) + (calcOneEspecialBless(lvl) * 2)
}

func calcFullBless(lvl int) int {
	return calcSevenBless(lvl) + calcTwistBless(lvl)
}

func calcHitpoints(level int, vocation string) int {
	if level < 9 || vocation == "mage" || vocation == "none" {
		return 5 * (level + 29)
	}

	if vocation == "paladin" {
		return 5 * (2*level + 21)
	} else {
		return 5 * (3*level + 13)
	}
}

func calcManapoints(level int, vocation string) int {
	if level < 9 || vocation == "knight" || vocation == "none" {
		return 5 * (level + 10)
	}

	if vocation == "paladin" {
		return 5 * (3*level - 6)
	} else {
		return 5 * (6*level - 30)
	}
}

func calcCap(level int, vocation string) int {
	if level < 9 || vocation == "mage" || vocation == "none" {
		return 10 * (level + 39)
	}

	if vocation == "paladin" {
		return 10 * (2*level + 31)
	} else {
		return 5 * (5*level + 54)
	}
}
