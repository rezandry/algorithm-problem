package main

import (
	"fmt"
)

func walk_switch(m map[int]bool, num_of_lamp, i int) map[int]bool {
	j := i-1
	for j<num_of_lamp {
		m[j] = !m[j]
		j = j+i
	}
	return m
}

func firstApproach(){
	m := make(map[int]bool)
	num_of_lamp := 100
	num_of_walk := 100
	iteration_addion := 3
	for i := 0; i<num_of_lamp; i++ {
		m[i] = false
	}
	walk := 0
	for walk < num_of_walk {
		for i:= 1; i<=iteration_addion; i++{
			m = walk_switch(m, num_of_lamp, i)
			walk++
			
			lamp_on := 0
			for k:=0; k<num_of_lamp; k++ {
				if m[k]{
					lamp_on++
				}
			}
			fmt.Printf("walk:%d lamp:%d \n",walk,lamp_on)
			
			if walk == num_of_walk {
				break
			}
				
		}
	}
}

func secondApproach(){
	m := map[int]int {
		1: 100,
		2: 50,
		3: 49,
		4: 51,
		5: 33,
		6: 0,
	}
	// num_of_lamp := 100
	// iteration_addion := 3
	num_of_walk := 100
	fmt.Println(m[num_of_walk % 6])
}

func main(){
	firstApproach()
	/*
	disclaimer: 
		this function is only work with pattern addition 1-3 with num of lamp is 100

	From first approach, we could know the pattern is
	walk:1 lamp:100
	walk:2 lamp:50
	walk:3 lamp:49
	walk:4 lamp:51
	walk:5 lamp:33
	walk:6 lamp:0
	and back to first number of lamp on
	so, we can make new function that more efficient
	from complexity O(n^3) to O(1)

	Actually we could to find another pattern for addition and num_of_lamp to decide how many lamp is on, but for this case, is solved.
	*/
	secondApproach()
}