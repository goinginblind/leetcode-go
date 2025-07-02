package main

func countStudentsSimple(students []int, sandwiches []int) int {
	didntEatCount := 0
	for didntEatCount < len(students) {
		if students[0] == sandwiches[0] {
			didntEatCount = 0
			students, sandwiches = students[1:], sandwiches[1:]
		} else {
			didntEatCount++
			students = append(students, students[0])
			students = students[1:]
		}
	}
	return len(students)
}
