package layout

import (
	"kuiper/config"
	"kuiper/fonts"
	"kuiper/utils"
	"math"
	"strings"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Layout struct {
	Data string
}

type KeyPosition struct {
	X int
	Y int
}

type DualChar struct {
	Display string
	Backend string
}

const Chars = 10

var Layouts = map[string]Layout{
	"azerty": {"azertyuiop"},
}

var Positions = "QWERTYUIOP"

func StringToPosition(str string) (float64, float64) {
	w := utils.PowInts(Chars, config.Config.WidthLevels)
	h := utils.PowInts(Chars, config.Config.HeightLevels)
	i := 0
	x := 0.0
	y := 0.0
	k_pow := 1.0
	l_pow := 1.0
	for w > 1 {
		if len(str) > i {
			x += float64(strings.Index(Positions, string(str[i]))) / math.Pow(float64(Chars), k_pow)
		}
		w /= Chars
		i++
		k_pow++

	}
	for h > 1 {
		if len(str) > i {
			y += float64(strings.Index(Positions, string(str[i]))) / math.Pow(float64(Chars), l_pow)
		}
		h /= Chars
		i++
		l_pow++
	}
	return x, y
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PositionToString(layout Layout, x, y int) DualChar {
	w := utils.PowInts(Chars, config.Config.WidthLevels)
	h := utils.PowInts(Chars, config.Config.HeightLevels)
	s := ""
	b := ""
	for h > 1 {
		s += string(layout.Data[y%Chars])
		b += string(Positions[y%Chars])
		h /= Chars
		y /= Chars
	}
	for w > 1 {
		s += string(layout.Data[x%Chars])
		b += string(Positions[x%Chars])
		w /= Chars
		x /= Chars
	}
	return DualChar{Reverse(s), Reverse(b)}
}

var TFace *text.GoTextFace

func Init() {
	TFace = &text.GoTextFace{
		Source: fonts.Font,
		Size:   config.Config.FontSize,
	}
	text.CacheGlyphs(Layouts[config.Config.SelectedLayout].Data, TFace)
}
