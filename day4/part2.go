package day4

import (
	"os"
	"strings"
)

var (
	top    = -1
	bottom = 1
	left   = -1
	right  = 1
)

func Part2(filePath string) int {
	byteInput, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(byteInput)
	rows := strings.Split(input, "\r\n")

	count := 0
	for x := 1; x < len(rows)-1; x++ {
		for y := 1; y < len(rows[x])-1; y++ {
			letter := string(rows[x][y])

			if letter != "A" {
				continue
			}

			{ //top-left && bottom-right
				tlx := x + left
				tly := y + top

				TL := string(rows[tlx][tly])

				brx := x + right
				bry := y + bottom

				BR := string(rows[brx][bry])

				if (TL != "M" && TL != "S") ||
					(BR != "M" && BR != "S") ||
					(TL == BR) {
					continue
				}
			}

			{ //top-right && bottom-left
				trx := x + right
				try := y + top
				TR := string(rows[trx][try])

				blx := x + left
				bly := y + bottom

				BL := string(rows[blx][bly])

				if (TR != "M" && TR != "S") ||
					(BL != "M" && BL != "S") ||
					(TR == BL) {
					continue
				}
			}

			count += 1
		}

	}

	return count
}
