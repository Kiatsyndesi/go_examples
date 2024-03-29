package algorithms

import "sort"

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.Slice(userIDs, func(i, j int) bool {
		return userIDs[i] < userIDs[j]
	})

	uniqPointer := 0

	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}
