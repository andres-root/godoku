package main

import "fmt"

const (
	cols string = "123456789"
	rows string = "ABCDEFGHI"
	grid string = "..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3.."
)

var (
	squareCols = []string{"123", "456", "789"}
	squareRows = []string{"ABC", "DEF", "GHI"}
)

func displayGrid(values map[string]string, boxes []string) {
	// TODO
	return
}

func cross(a string, b string) []string {
	var result []string
	for _, s := range a {
		for _, t := range b {
			result = append(result, string(s)+string(t))
		}
	}
	return result
}

func getRowUnits() [][]string {
	var rowUnits [][]string
	for _, r := range rows {
		result := cross(string(r), cols)
		rowUnits = append(rowUnits, result)
	}
	return rowUnits
}

func getColumnUnits() [][]string {
	var columnUnits [][]string
	for _, c := range cols {
		result := cross(rows, string(c))
		columnUnits = append(columnUnits, result)
	}
	return columnUnits
}

func getSquareUnitsOld() [][]string {
	var squareUnits [][]string
	n := 1
	i := 0
	for n <= 9 {
		for _, r := range squareRows {
			for _, c := range squareCols {
				squareUnits = append(squareUnits, cross(r, c))
				i++
			}
		}
		n++
	}
	return squareUnits
}

func getSquareUnits() [][]string {
	var squareUnits [][]string
	for _, r := range squareRows {
		for _, c := range squareCols {
			squareUnits = append(squareUnits, cross(r, c))
		}
	}
	return squareUnits
}

func getGridValues(boxes []string) map[string]string {
	gridValues := make(map[string]string)
	unsolved := make(map[string]string)
	for i, box := range boxes {
		gridValues[box] = string(grid[i])
	}

	for k, v := range gridValues {
		if v == "." {
			unsolved[k] = cols
		} else {
			unsolved[k] = v
		}
	}
	return unsolved
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func getUnits(unitlist [][]string, boxes []string) map[string][][]string {
	units := make(map[string][][]string)

	for _, box := range boxes {
		var tempList [][]string

		for _, unit := range unitlist {
			if stringInSlice(box, unit) {
				tempList = append(tempList, unit)
			}
		}
		units[box] = tempList
		if box == "A1" {
			break
		}
	}
	return units
}

func getPeers(boxes []string, units map[string][][]string) map[string]map[string]bool {
	// peers = dict((s, set(sum(units[s],[]))-set([s])) for s in boxes)
	var peers = make(map[string]map[string]bool)
	for _, box := range boxes {
		var tempMap = make(map[string]bool)
		for _, unit := range units[box] {
			for _, unitBox := range unit {
				if unitBox != box {
					tempMap[unitBox] = true
				}
			}
		}
		peers[box] = tempMap
	}

	return peers
}

// func elimitate(values map[string]string) map[string]string {
// 	//
// 	var solvedBoxes []string
// 	for k, v := range values {
// 		if len(v) == 1 {
// 			solvedBoxes = append(solvedBoxes, k)
// 		}
// 	}

// 	for _, box := range solvedBoxes {
// 		n := box
// 		// for k, v := range values

// 	}

// 	return values
// }

func main() {
	var unitList [][]string

	boxes := cross(rows, cols)
	rowUnits := getRowUnits()
	columnUnits := getColumnUnits()
	squareUnits := getSquareUnits()
	unitList = append(rowUnits, columnUnits...)
	unitList = append(unitList, squareUnits...)
	// gridValues := getGridValues(boxes)
	units := getUnits(unitList, boxes)
	// peers = dict((s, set(sum(units[s],[]))-set([s])) for s in boxes)
	peers := getPeers(boxes, units)
	fmt.Println("---------------------------")
	fmt.Println(peers["A1"])
	// fmt.Println(gridValues)
	// fmt.Println(rowUnits)
	// fmt.Println(units["A1"])
	// fmt.Println(len(units["A1"]))
}
