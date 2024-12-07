package commons

func RotateMatrix45[t any](m [][]t) [][]t {
	maxIndex := len(m) + len(m[0]) - 1
	newMatrix := make([][]t, maxIndex)

	for i := 0; i < maxIndex; i++ {
		newMatrix[i] = make([]t, 0)
	}

	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			newMatrix[i+j] = append(newMatrix[i+j], m[j][i])
		}
	}

	return newMatrix
}

func RotateMatrix90[t any](m [][]t) [][]t {
	maxIndex := len(m) - 1
	newMatrix := make([][]t, len(m))
	for i := 0; i < len(m); i++ {
		newMatrix[i] = make([]t, len(m))
	}

	for i := 0; i < len(m); i++ { //rows
		for j := 0; j < len(m); j++ { //cols
			newMatrix[j][maxIndex-i] = m[i][j]
		}
	}

	return newMatrix
}
