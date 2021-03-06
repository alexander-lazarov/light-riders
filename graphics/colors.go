package graphics

import (
	"image/color"

	"github.com/alexander-lazarov/lightriders/models"
)

func getColor(board *models.Board, p models.Position) color.RGBA {
	cell := board.GetCell(p)

	switch cell {
	case models.P1Head:
		return color.RGBA{R: 0, G: 255, B: 0, A: 255}
	case models.P1Tail:
		return color.RGBA{R: 0, G: 200, B: 0, A: 255}
	case models.P2Head:
		return color.RGBA{R: 0, G: 0, B: 255, A: 255}
	case models.P2Tail:
		return color.RGBA{R: 0, G: 0, B: 200, A: 255}
	case models.Crash:
		return color.RGBA{R: 255, G: 0, B: 0, A: 255}
	default:
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}
}
