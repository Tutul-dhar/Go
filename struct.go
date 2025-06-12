package main

import "fmt"

type student struct {
	name string
	grades []int
	age int
}

func (s *student) getAvgGrade() float32 {
	sum := 0
	for _,v := range s.grades {
		sum += v
	}

	return float32(sum)/float32(len(s.grades))
}


func (s student) getMax() int {
	mx := 0

	for _, v := range s.grades {
		if v > mx {
			mx = v
		}
	}


	return mx
}

func main() {
	s1 := student{"tutul", []int{90, 34, 45,90 ,60}, 24}
	fmt.Println(s1.getAvgGrade())
	s2 := student{"Urmi", []int{90,80,70,93,34},23}

	p := s2.getMax()
	fmt.Println(p)
}