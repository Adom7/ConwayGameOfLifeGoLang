package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 600
	screenHeight = 600
	gridSize     = 100 // Taille de la grille (20x20)
	cellSize     = 10
)

const (
	buttonWidth  = 100
	buttonHeight = 50
)

type Game struct {
	grid      [][]bool
	isStarted bool
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		for i := range g.grid {
			for j := range g.grid[i] {
				x := i * cellSize
				y := j * cellSize
				if mx >= x && mx <= x+cellSize && my >= y && my <= y+cellSize {
					g.grid[i][j] = !g.grid[i][j]
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	// Dessine les cellules
	for i := range g.grid {
		for j := range g.grid[i] {
			if g.grid[i][j] {
				ebitenutil.DrawRect(screen, float64(i*cellSize), float64(j*cellSize), float64(cellSize), float64(cellSize), color.White)
			} else {
				ebitenutil.DrawRect(screen, float64(i*cellSize), float64(j*cellSize), float64(cellSize), float64(cellSize), color.Black)
			}
		}
	}

	// Dessine les lignes entre les cellules
	for i := 1; i < gridSize; i++ {
		ebitenutil.DrawLine(screen, float64(i*cellSize), 0, float64(i*cellSize), float64(screenHeight), color.Gray16{0x3333})
	}
	for j := 1; j < gridSize; j++ {
		ebitenutil.DrawLine(screen, 0, float64(j*cellSize), float64(screenWidth), float64(j*cellSize), color.Gray16{0x3333})
	}

	buttonX := (screenWidth - buttonWidth) / 2
	buttonY := screenHeight - buttonHeight - 20
	ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), buttonWidth, buttonHeight, color.RGBA{187, 187, 187, 255})
	ebitenutil.DebugPrintAt(screen, "Start Game", buttonX, buttonY)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func GameLogic() {

}

func main() {
	grid := make([][]bool, gridSize)
	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	game := &Game{
		grid: grid,
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway Game of life")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
