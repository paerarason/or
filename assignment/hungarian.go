package assignment 

import (
	"math"
	"encoding/json"
)

type JSONData struct {
    Data [][] int `json:"data"`
}

func Hungarian(matrix CostMatrix.Matrix) int,error {

	mat:=makeSquareMatrix(&matrix)
	
	//perform row minimum
	for i:=0;i<len(matrix);i++{
		min:=findMinRow(mat[i])  	
		for j:=0;j<len(mat[i]);j++{

			mat[i][j]-=min
		}
	}


	//perform coloum minimum
    for i:=0;i<len(mat[i]);i++{
		min:=findMinRow(mat[][i])
		for j:=0;j<len(matrix);j++{
			mat[j][i]-=min
		}
	}
    //Assignment
	row,col:=coverZeroes(mat)
	if row+col==len(mat){
		 se
	}

}




//find min in the array 
func findMinRow(arr []int) int {
   min:=math.MaxInt64
   for x,_:=range arr {
      if min>x {
		min=x
	  }
   }
   return min
}




func makeSquareMatrix(matrix [][] byte) [][]byte{
	var square bool 
	square=true
	max_length:=math.MinInt64

	if len(matrix)>max_length {
		max_length=len(matrix)
	}
	//iterate over the matrix to make the array  square matrix 
	for i:=0;i<len(matrix);i++{
		if len(matrix[0])!=len(matrix[i]){
			max_length=math.max(max_length,len(matrix[i]))
			square=false
		}
	}

    if square!=true{
        //coloum setting up 
		for i:=0;i<len(matrix);i++{
		 if len(matrix[i])<max_length{
			for j:=0;j<(max_length-matrix[i]);j++{
				append(matrix[i],0)
			}
		   }
	     }
		
		// row setup 
		for i:=0;i<max_length-matrix[i];i++{
			for j:=0;j<max_length;j++{
				append(matrix[len(matrix)+i],0)
			}
		}
	   return matrix
	       }	
	  return matrix
}


func coverZeroes(matrix [][]int) ([]int, []int) {
	rows, cols := len(matrix), len(matrix[0])
	coveredRows := make([]bool, rows)
	coveredCols := make([]bool, cols)
	// Step 1: Find the minimum value in each row
	minInRow := make([]int, rows)
	for i := 0; i < rows; i++ {
		minInRow[i] = matrix[i][0]
		for j := 1; j < cols; j++ {
			if matrix[i][j] < minInRow[i] {
				minInRow[i] = matrix[i][j]
			}
		}
	}

	// Step 2: Find the minimum value in each column
	minInCol := make([]int, cols)
	for j := 0; j < cols; j++ {
		minInCol[j] = matrix[0][j] - minInRow[0]
		for i := 1; i < rows; i++ {
			if matrix[i][j]-minInRow[i] < minInCol[j] {
				minInCol[j] = matrix[i][j] - minInRow[i]
			}
		}
	}

	// Step 3: Mark the rows and columns that should be covered
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j]-minInRow[i]-minInCol[j] == 0 {
				if !coveredRows[i] {
					coveredRows[i] = true
				}
				if !coveredCols[j] {
					coveredCols[j] = true
				}
			}
		}
	}

	// Step 4: Create lines to cover all zeroes
	rowLines := make([]int, 0)
	colLines := make([]int, 0)
	for i := 0; i < rows; i++ {
		if !coveredRows[i] {
			rowLines = append(rowLines, i)
		}
	}
	for j := 0; j < cols; j++ {
		if coveredCols[j] {
			colLines = append(colLines, j)
		}
	}
	return rowLines, colLines
}