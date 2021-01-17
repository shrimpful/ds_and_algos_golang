package main

var matrix1 = [2][2]int{
	{4, 5},
	{1, 2},
}
var matrix2 = [2][2]int{
	{6, 7},
	{3, 4},
}

func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var sum [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

func substract(matrix1, matrix2 [2][2]int) [2][2]int {
	var m, l int
	var difference [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			difference[l][m] = matrix1[l][m] - matrix2[l][m]
		}
	}
	return difference
}

func multiply(matrix1, matrix2 [2][2]int) [2][2]int {
	var l, n, m int
	var product [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			var productSum int = 0
			for n = 0; n < 2; n++ {
				productSum = productSum + matrix1[l][n]*matrix2[n][m]
			}
			product[l][m] = productSum
		}
	}
	return product
}

func transpose(matrix1 [2][2]int) [2][2]int {
	var m, l int
	var transMatrix [2][2]int
	for l = 0; l < 2; l++ {
		for m = 0; m < 2; m++ {
			transMatrix[l][m] = matrix1[m][1]
		}
	}
	return transMatrix
}

func determinant(matrix1 [2][2]int) float64 {
	var det float64
	det = det + float64((matrix1[0][0]*matrix1[1][1])-(matrix1[0][1]*matrix1[1][0]))
	return det
}

func inverse(matrix [2][2]int) [][]float64 {

	var det float64

	det = determinant(matrix)

	var invmatrix [][]float64
	invmatrix[0][0] = float64(float64(matrix[1][1]) / det)
	invmatrix[0][1] = float64(float64(-1*matrix[0][1]) / det)
	invmatrix[1][0] = float64(float64(-1*matrix[1][0]) / det)
	invmatrix[1][1] = float64(float64(matrix[0][0]) / det)

	return invmatrix

}
