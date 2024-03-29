package levenshtein

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Min(array [3]int) int {
	var min int = array[0]
	for _, value := range array {
		if min > value {
			min = value
		}
	}
	return min
}

func printLevMatrix(source string, target string, a [][]int) {
	fmt.Print("    ")
	for i, _ := range target {
		fmt.Printf(" %v", string(target[i]))
	}
	fmt.Println()
	fmt.Printf("  %v\n", a[0])
	for i, line := range a[1:] {
		fmt.Printf("%+v", string([]rune(source)[i]))
		fmt.Printf(" %v\n", line)
	}
}

func GetLevMatrix(source, target string) [][]int {
	// Levenshtein distance
	// https://en.wikipedia.org/wiki/Levenshtein_distance#Iterative_with_full_matrix
	a := make([][]int, len(source)+1)
	for i := range a {
		a[i] = make([]int, len(target)+1)
	}

	for i := range a {
		a[i][0] = i
	}

	for j := range a[0] {
		a[0][j] = j
	}
	substitutionCost := 0

	for j := 1; j <= len(target); j++ {
		for i := 1; i <= len(source); i++ {
			if source[i-1] == target[j-1] {
				substitutionCost = 0
			} else {
				substitutionCost = 1
			}

			changes := [...]int{
				a[i-1][j] + 1,                  // deletion
				a[i][j-1] + 1,                  // insertion
				a[i-1][j-1] + substitutionCost} // substitution

			a[i][j] = Min(changes)

		}
	}

	return a

}

func GetDistance(levMatrix [][]int, source string, target string) int {
	distance := levMatrix[len(source)][len(target)]
	return distance
}

func CalculateDistance(source string, target string) int {
	mtx := GetLevMatrix(source, target)
	distance := GetDistance(mtx, source, target)
	return distance
}
