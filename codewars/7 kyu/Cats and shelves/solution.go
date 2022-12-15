package Cats_and_shelves

func Cats(start, finish int) int {
	numberOfJumps := 0
	bigJump := 3
	smallJump := 1

	for start < finish {
		if finish-bigJump >= start {
			numberOfJumps++
			start = start + bigJump
		} else if finish-smallJump >= start {
			numberOfJumps++
			start = start + smallJump
		}
	}

	return numberOfJumps
}
