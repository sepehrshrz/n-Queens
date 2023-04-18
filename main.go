package main

import (
	"fmt"
)

type BlockPoint struct {
	i int
	j int
	mainDiameter int
	subDiameter int
}

var(
	blockRows=make([]BlockPoint,0)
	successCount int
	n int
)

func main(){
    fmt.Print("Enter n for make matris n * n: ")
	fmt.Scan(&n)
	successCount=0
	matris:=make([]int, n)
	//go to MakeMatris function with i and j = 0 for search for true point
	SetQueen(0,0,matris)
	
}

func SetQueen(i,j int,matris []int){
	if i >= n {
		return
	}
	// etrate for find queen and call recursive function
		for j := 0; j < n; j++ {
			//Checking queen for non-collision of i,j, main diameter and minor diameter
			if findTruePoint(i,j) {
				matris[i]=j
				//set block rows for none-colision in next queen
				newBlockRows:=BlockPoint{
					i: i,
					j: j,
					mainDiameter: j-i,
					subDiameter: j+i,
				}
				blockRows=append(blockRows, newBlockRows)
				// if len of blockrows equal to n it means that n queens have been pawned
				if len(blockRows) == n {
					fmt.Println(matris)
					successCount++
				}
				//go to recursive function for find next queen
				SetQueen(i+1,0,matris)
				//If finding the next queen is unsuccessful, the replacement of the current queen is cleared and shifted to the right
				blockRows=blockRows[:len(blockRows) -1]
			}
		}
	
		//if i == 0 It means that the search to find the queen is over
		if i == 0{
			fmt.Println("count of true matris:",successCount)
		}

}

func findTruePoint(i,j int)bool{

	for _, v := range blockRows {
		//Check point for non-collision in j
		if v.j == j {
			return false
		}
		//Check point for non-collision in mainDiameter

		if j-i== v.mainDiameter {
			return false
		}
		//Check point for non-collision in subDiameter

		if j+i== v.subDiameter {
			return false
		}
	}

	return true
}

