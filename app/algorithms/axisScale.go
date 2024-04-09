package algorithms

import "fmt"

type AxisScale struct {
	MinX, MaxX int
	MinY, MaxY int
}

func NewAxisScale(minX, maxX, minY, maxY int) AxisScale {
	return AxisScale{
		MinX: minX,
		MaxX: maxX,
		MinY: minY,
		MaxY: maxY,
	}
}

func DrawXAxis(scale AxisScale, totalProcessTime, totalProcessNumber int) {
	marginWidth := 5
	fmt.Printf("%*s", marginWidth, "")

	for i := scale.MinX; i <= scale.MaxX; i++ {
		fmt.Print("-")
	}
	fmt.Println()

	for i := 0; i <= totalProcessTime; i++ {
		fmt.Printf("%*d", totalProcessNumber+1, i)
	}
	fmt.Println()
}

func DrawYAxis(scale AxisScale, p []Process) {
	idCounter := len(p) - 1
	marginWidth := 5

	space := scale.MaxY / len(p)

	for i := scale.MinY; i <= scale.MaxY; i++ {
		fmt.Printf("%*s", marginWidth, "")
		fmt.Print("|")
		fmt.Println()

		if i == (space - 1) {
			if idCounter >= 0 {
				fmt.Printf("%s - |\n", p[idCounter].ProcessId)
			}
			idCounter--
			space += scale.MaxY / len(p)
		}

	}
}

func DrawValues() {

}
