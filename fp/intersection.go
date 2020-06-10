package fp

// IntersectionInt return a set that is the intersection of the input sets
func IntersectionInt(arrList ...[]int) []int {
	if arrList == nil {
		return []int{}
	}

	resultMap := make(map[int]bool)
	if len(arrList) == 1 {
		var newList []int
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []int
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionIntPtr(arrList ...[]*int) []*int {
	if arrList == nil {
		return []*int{}
	}

	resultMap := make(map[int]bool)
	if len(arrList) == 1 {
		var newList []*int
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionInt64 return a set that is the intersection of the input sets
func IntersectionInt64(arrList ...[]int64) []int64 {
	if arrList == nil {
		return []int64{}
	}

	resultMap := make(map[int64]bool)
	if len(arrList) == 1 {
		var newList []int64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []int64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionInt64Ptr(arrList ...[]*int64) []*int64 {
	if arrList == nil {
		return []*int64{}
	}

	resultMap := make(map[int64]bool)
	if len(arrList) == 1 {
		var newList []*int64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionInt32 return a set that is the intersection of the input sets
func IntersectionInt32(arrList ...[]int32) []int32 {
	if arrList == nil {
		return []int32{}
	}

	resultMap := make(map[int32]bool)
	if len(arrList) == 1 {
		var newList []int32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []int32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionInt32Ptr(arrList ...[]*int32) []*int32 {
	if arrList == nil {
		return []*int32{}
	}

	resultMap := make(map[int32]bool)
	if len(arrList) == 1 {
		var newList []*int32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionInt16 return a set that is the intersection of the input sets
func IntersectionInt16(arrList ...[]int16) []int16 {
	if arrList == nil {
		return []int16{}
	}

	resultMap := make(map[int16]bool)
	if len(arrList) == 1 {
		var newList []int16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []int16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionInt16Ptr(arrList ...[]*int16) []*int16 {
	if arrList == nil {
		return []*int16{}
	}

	resultMap := make(map[int16]bool)
	if len(arrList) == 1 {
		var newList []*int16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionInt8 return a set that is the intersection of the input sets
func IntersectionInt8(arrList ...[]int8) []int8 {
	if arrList == nil {
		return []int8{}
	}

	resultMap := make(map[int8]bool)
	if len(arrList) == 1 {
		var newList []int8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []int8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionInt8Ptr(arrList ...[]*int8) []*int8 {
	if arrList == nil {
		return []*int8{}
	}

	resultMap := make(map[int8]bool)
	if len(arrList) == 1 {
		var newList []*int8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionUint return a set that is the intersection of the input sets
func IntersectionUint(arrList ...[]uint) []uint {
	if arrList == nil {
		return []uint{}
	}

	resultMap := make(map[uint]bool)
	if len(arrList) == 1 {
		var newList []uint
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []uint
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionUintPtr(arrList ...[]*uint) []*uint {
	if arrList == nil {
		return []*uint{}
	}

	resultMap := make(map[uint]bool)
	if len(arrList) == 1 {
		var newList []*uint
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionUint64 return a set that is the intersection of the input sets
func IntersectionUint64(arrList ...[]uint64) []uint64 {
	if arrList == nil {
		return []uint64{}
	}

	resultMap := make(map[uint64]bool)
	if len(arrList) == 1 {
		var newList []uint64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []uint64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionUint64Ptr(arrList ...[]*uint64) []*uint64 {
	if arrList == nil {
		return []*uint64{}
	}

	resultMap := make(map[uint64]bool)
	if len(arrList) == 1 {
		var newList []*uint64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionUint32 return a set that is the intersection of the input sets
func IntersectionUint32(arrList ...[]uint32) []uint32 {
	if arrList == nil {
		return []uint32{}
	}

	resultMap := make(map[uint32]bool)
	if len(arrList) == 1 {
		var newList []uint32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []uint32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionUint32Ptr(arrList ...[]*uint32) []*uint32 {
	if arrList == nil {
		return []*uint32{}
	}

	resultMap := make(map[uint32]bool)
	if len(arrList) == 1 {
		var newList []*uint32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionUint16 return a set that is the intersection of the input sets
func IntersectionUint16(arrList ...[]uint16) []uint16 {
	if arrList == nil {
		return []uint16{}
	}

	resultMap := make(map[uint16]bool)
	if len(arrList) == 1 {
		var newList []uint16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []uint16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionUint16Ptr(arrList ...[]*uint16) []*uint16 {
	if arrList == nil {
		return []*uint16{}
	}

	resultMap := make(map[uint16]bool)
	if len(arrList) == 1 {
		var newList []*uint16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionUint8 return a set that is the intersection of the input sets
func IntersectionUint8(arrList ...[]uint8) []uint8 {
	if arrList == nil {
		return []uint8{}
	}

	resultMap := make(map[uint8]bool)
	if len(arrList) == 1 {
		var newList []uint8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []uint8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionUint8Ptr(arrList ...[]*uint8) []*uint8 {
	if arrList == nil {
		return []*uint8{}
	}

	resultMap := make(map[uint8]bool)
	if len(arrList) == 1 {
		var newList []*uint8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionStr return a set that is the intersection of the input sets
func IntersectionStr(arrList ...[]string) []string {
	if arrList == nil {
		return []string{}
	}

	resultMap := make(map[string]bool)
	if len(arrList) == 1 {
		var newList []string
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []string
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionStrPtr(arrList ...[]*string) []*string {
	if arrList == nil {
		return []*string{}
	}

	resultMap := make(map[string]bool)
	if len(arrList) == 1 {
		var newList []*string
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*string
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionBool return a set that is the intersection of the input sets
func IntersectionBool(arrList ...[]bool) []bool {
	if arrList == nil {
		return []bool{}
	}

	resultMap := make(map[bool]bool)
	if len(arrList) == 1 {
		var newList []bool
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []bool
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionBoolPtr(arrList ...[]*bool) []*bool {
	if arrList == nil {
		return []*bool{}
	}

	resultMap := make(map[bool]bool)
	if len(arrList) == 1 {
		var newList []*bool
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*bool
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionFloat32 return a set that is the intersection of the input sets
func IntersectionFloat32(arrList ...[]float32) []float32 {
	if arrList == nil {
		return []float32{}
	}

	resultMap := make(map[float32]bool)
	if len(arrList) == 1 {
		var newList []float32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []float32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionFloat32Ptr(arrList ...[]*float32) []*float32 {
	if arrList == nil {
		return []*float32{}
	}

	resultMap := make(map[float32]bool)
	if len(arrList) == 1 {
		var newList []*float32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*float32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionFloat64 return a set that is the intersection of the input sets
func IntersectionFloat64(arrList ...[]float64) []float64 {
	if arrList == nil {
		return []float64{}
	}

	resultMap := make(map[float64]bool)
	if len(arrList) == 1 {
		var newList []float64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
		return newList
	}

	var newList []float64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = true
			}
		}
	}
	return newList
}

// IntersectionIntPtr return a set that is the intersection of the input sets
func IntersectionFloat64Ptr(arrList ...[]*float64) []*float64 {
	if arrList == nil {
		return []*float64{}
	}

	resultMap := make(map[float64]bool)
	if len(arrList) == 1 {
		var newList []*float64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = true
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*float64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == len(arrList)-1 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = true
			}
		}
	}
	return newList
}