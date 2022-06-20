// +build tests
package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW
func Test_PeopleLen(t *testing.T) {
	testCases := []struct{
		persons []Person
	}{
		{persons: make([]Person, 0)},
		{persons: []Person{Person{}, Person{}}},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_PeopleLen_%d", i), func (t *testing.T)  {
			people := People{}

			for _, person := range testCase.persons {
				people = append(people, person)
			}

			if people.Len() != len(testCase.persons) {
				t.Errorf("Wrong people len")
			}
		})
	}
}

func Test_PeopleLess(t *testing.T) {
	testCases := []struct{
		persons [2]Person
		firstBigger bool
	}{
		{
			persons: [2]Person{
				Person{
					birthDay: time.Date(2001, time.January, 1, 23, 0, 0, 0, time.UTC),
				},
				Person{
					birthDay: time.Date(2000, time.January, 1, 23, 0, 0, 0, time.UTC),
				},
			}, 
			firstBigger: true,
		},
		{
			persons: [2]Person{
				Person{
					firstName: "A",
				},
				Person{
					firstName: "B",
				},
			}, 
			firstBigger: true,
		},
		{
			persons: [2]Person{
				Person{
					lastName: "A",
				},
				Person{
					lastName: "B",
				},
			}, 
			firstBigger: true,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_PeopleLess_%d", i), func (t *testing.T)  {
			people := People{}

			for _, person := range testCase.persons {
				people = append(people, person)
			}

			if people.Less(0, 1) != testCase.firstBigger {
				t.Errorf("Error comparison")
			}
		})
	}
}

func Test_PeopleSwap(t *testing.T) {
	testCases := []struct{
		persons [2]Person
	}{
		{
			persons: [2]Person{
				Person{
					firstName: "A",
				},
				Person{
					firstName: "B",
				},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_PeopleSwap_%d", i), func (t *testing.T)  {
			people := People{}

			for _, person := range testCase.persons {
				people = append(people, person)
			}

			people.Swap(0, 1)

			if people[0].firstName != testCase.persons[1].firstName && people[1].firstName != testCase.persons[0].firstName {
				t.Errorf("Error swap")
			}
		})
	}
}

func Test_NewMatrix(t *testing.T) {
	testCases := []struct{
		matrix string
		err error
		data []int
		rows int
		cols int
	}{
		{
			matrix: "1 2 3",
			data: []int{1, 2, 3},
			rows: 1,
			cols: 3,
		},
		{
			matrix: "a",
			err: fmt.Errorf("strconv.Atoi: parsing \"a\": invalid syntax"),
		},
		{
			matrix: "1 2 3\n1 2",
			err: fmt.Errorf("Rows need to be the same length"),
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_NewMatrix_%d", i), func (t *testing.T)  {
			m, err := New(testCase.matrix)
			if testCase.err != nil {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}

				if testCase.err.Error() != err.Error() {
					t.Errorf("Expected error `%s`, got `%s`", testCase.err.Error(), err.Error())
				}

				return
			}

			if err != nil {
				t.Errorf("Unexpected error `%s`", err.Error())

				return
			}

			if m.cols != testCase.cols {
				t.Errorf("Expected %d cols, got %d", testCase.cols, m.cols)
			}
			
			if m.rows != testCase.rows {
				t.Errorf("Expected %d rows, got %d", testCase.rows, m.rows)
			}

			for i, expected := range testCase.data {
				if m.data[i] != expected {
					t.Errorf("Expected %d at position %d, got %d", expected, i, m.data[i])
				}
			}
		})
	}
}

func Test_MatrixRows(t *testing.T) {
	testCases := []struct{
		matrix string
		data [][]int
	}{
		{
			matrix: "1 2 3",
			data: [][]int{{1, 2, 3}},
		},
		{
			matrix: "1 2 3\n1 2 3",
			data: [][]int{{1, 2, 3}, {1, 2, 3}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_MatrixRows_%d", i), func (t *testing.T)  {
			m, err := New(testCase.matrix)
			
			if err != nil {
				t.Errorf("Unexpected error `%s`", err.Error())
				
				return
			}

			rows := m.Rows()

			for i, cols := range testCase.data {
				for j, expected := range cols {
					if rows[i][j] != expected {
						t.Errorf("Expected %d at position [%d][%d], got %d", expected, i, j, rows[i][j])
					}
				}
			}
		})
	}
}

func Test_MatrixCols(t *testing.T) {
	testCases := []struct{
		matrix string
		data [][]int
	}{
		{
			matrix: "1 2 3",
			data: [][]int{{1}, {2}, {3}},
		},
		{
			matrix: "1 2 3\n1 2 3",
			data: [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_MatrixCols_%d", i), func (t *testing.T)  {
			m, err := New(testCase.matrix)
			
			if err != nil {
				t.Errorf("Unexpected error `%s`", err.Error())
				
				return
			}

			rows := m.Cols()

			for i, cols := range testCase.data {
				for j, expected := range cols {
					if rows[i][j] != expected {
						t.Errorf("Expected %d at position [%d][%d], got %d", expected, i, j, rows[i][j])
					}
				}
			}
		})
	}
}

func Test_MatrixSet(t *testing.T) {
	testCases := []struct{
		matrix string
		row int 
		col int 
		value int
		result bool
	}{
		{
			matrix: "1 2 3\n1 2 3",
			row: 1,
			col: 1,
			value: 0,
			result: true,
		},
		{
			matrix: "1 2 3\n1 2 3",
			row: -1,
			result: false,
		},
		{
			matrix: "1 2 3\n1 2 3",
			col: -1,
			result: false,
		},
		{
			matrix: "1 2 3\n1 2 3",
			row: 5,
			result: false,
		},
		{
			matrix: "1 2 3\n1 2 3",
			col: 5,
			result: false,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test_MatrixSet_%d", i), func (t *testing.T)  {
			m, err := New(testCase.matrix)
			
			if err != nil {
				t.Errorf("Unexpected error `%s`", err.Error())
				
				return
			}

			result := m.Set(testCase.row, testCase.col, testCase.value)

			if result != testCase.result {
				t.Errorf("Couldn't set value")

				return
			}

			if testCase.result == true {
				rows := m.Rows()
				
				if rows[testCase.row][testCase.col] != testCase.value {
					t.Errorf("Expected %d at position [%d][%d], got %d", testCase.value, testCase.row, testCase.col, rows[testCase.row][testCase.col])
				}
			}
		})
	}
}