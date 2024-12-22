package main

import (
	"kuiper/config"
	"kuiper/fonts"
	"kuiper/layout"
	"kuiper/utils"
	"log"
	"math"
	"os"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	keys     []ebiten.Key
	filtered string
}

func (g *Game) Update() error {

	g.keys = inpututil.AppendJustPressedKeys(g.keys[:0])
	for _, key := range g.keys {
		if strings.Contains(layout.Positions, key.String()) {
			g.filtered += key.String()
		}
	}
	if len(g.filtered) >= config.Config.WidthLevels+config.Config.HeightLevels {
		x, y := layout.StringToPosition(g.filtered)
		log.Printf("%d %d", int((x+1/float64(utils.PowInts(layout.Chars, config.Config.WidthLevels))/2)*float64(config.Config.WindowWidth)), int((y+1/float64(utils.PowInts(layout.Chars, config.Config.HeightLevels))/2)*float64(config.Config.WindowHeight)))
		//robotgo.Move(int((x+1/float64(utils.PowInts(layout.Chars, config.Config.WidthLevels))/2)*float64(config.Config.WindowWidth)), int((y+1/float64(utils.PowInts(layout.Chars, config.Config.HeightLevels))/2)*float64(config.Config.WindowHeight)))
		os.Exit(1)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	//ebitenutil.DebugPrint(screen, g.filtered)
	for i := 0; i < layout.Chars; i++ {
		op := &text.DrawOptions{}
		var box_width float64
		var box_height float64
		var characters string
		var t_width float64
		var t_height float64
		var bx float32
		var by float32
		if len(g.filtered) < config.Config.WidthLevels {
			box_width = float64(config.Config.WindowWidth) / math.Pow(layout.Chars, float64(len(g.filtered)+1))
			box_height = float64(config.Config.WindowHeight)
			characters = string(layout.Layouts[config.Config.SelectedLayout].Data[i])
			t_width, t_height = text.Measure(characters, layout.TFace, 0)

			x, y := layout.StringToPosition(g.filtered)
			bx = float32(x*float64(config.Config.WindowWidth)) + float32(i)*float32(box_width)
			by = float32(y * float64(config.Config.WindowHeight))
		} else {
			box_width = float64(config.Config.WindowWidth) / math.Pow(layout.Chars, float64(config.Config.WidthLevels))
			box_height = float64(config.Config.WindowHeight) / math.Pow(layout.Chars, float64(len(g.filtered)-config.Config.WidthLevels+1))
			characters = string(layout.Layouts[config.Config.SelectedLayout].Data[i])
			t_width, t_height = text.Measure(characters, layout.TFace, 0)
			x, y := layout.StringToPosition(g.filtered)
			bx = float32(x * float64(config.Config.WindowWidth))
			by = float32(y*float64(config.Config.WindowHeight)) + float32(i)*float32(box_height)
		}
		op.GeoM.Translate(float64(bx)+float64(box_width)/2-t_width/2, float64(by)+float64(box_height)/2-t_height/2)
		op.ColorScale.ScaleWithColor(config.AccentColor)
		text.Draw(screen, characters, layout.TFace, op)
		vector.StrokeRect(screen, bx, by, float32(box_width), float32(box_height), 1, config.FirstColor, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(config.Config.WindowWidth), int(config.Config.WindowHeight)
}

func main() {
	config.Init()
	fonts.Init()
	layout.Init()
	//ebiten.SetWindowIcon([]image.Image{icons.WindowIcon48, icons.WindowIcon32, icons.WindowIcon16})
	ebiten.SetWindowSize(int(config.Config.WindowWidth), int(config.Config.WindowHeight))
	ebiten.SetWindowTitle("kuiper")
	ebiten.SetWindowMousePassthrough(false)
	ebiten.SetTPS(int(config.Config.TargetFPS))
	ebiten.SetWindowDecorated(false)
	screenW, screenH := ebiten.Monitor().Size()
	ebiten.SetWindowPosition(screenW/2-int(config.Config.WindowWidth)/2, screenH/2-int(config.Config.WindowHeight)/2)
	ebiten.SetVsyncEnabled(true)
	if err := ebiten.RunGameWithOptions(&Game{}, &ebiten.RunGameOptions{ScreenTransparent: true}); err != nil {
		log.Fatal(err)
	}
}
