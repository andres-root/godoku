package main

const (
	rows string = "ABCDEFGHI"
	cols string = "123456789"
	grid string = "..3.2.6..9..3.5..1..18.64....81.29..7.......8..67.82....26.95..8..2.3..9..5.1.3.."
)

func displayGrid(values map[string]string, boxes []string) {
	// TODO
}

func cross(a string, b string) []string {
	var s []string
	for _, ac := range a {
		for _, bc := range b {
			s = append(s, string(ac)+string(bc))
		}
	}
	return s
}

func getRowUnits() []string {
	var rowUnits []string
	for _, r := range rows {
		rowUnits = append(rowUnits, cross(string(r), cols)...)
	}
	return rowUnits
}

func getColumnUnits() []string {
	var columnUnits []string
	for _, c := range cols {
		columnUnits = append(columnUnits, cross(rows, string(c))...)
	}
	return columnUnits
}

func getSquareUnits() []string {
	var squareUnits []string
	for _, r := range "ABC" {
		for _, c := range "123" {
			squareUnits = append(squareUnits, cross(string(r), string(c))...)
		}
	}
	return squareUnits
}

func getGridValues(grid string, boxes []string) map[string]string {
	gridValues := make(map[string]string)
	for i, box := range boxes {
		gridValues[box] = string(grid[i])
	}
	return gridValues
}

func main() {
	boxes := cross(rows, cols)
	rowUnits := getRowUnits()
	columnUnits := getColumnUnits()
	squareUnits := getSquareUnits()
	var unitList [][]string
	unitList = append(unitList, rowUnits)
	unitList = append(unitList, columnUnits)
	unitList = append(unitList, squareUnits)
	gridValues := getGridValues(grid, boxes)
	displayGrid(gridValues, boxes)
}
