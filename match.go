package match

type anyType interface{}

// GetCombination is a function for getcombination with multiple array
func getCombination(currentIndex int, containers [][]anyType) (combinations [][]anyType) {
	if currentIndex == len(containers) {
		// Skip the items for the last container
		combinations = append(combinations, []anyType{})
		return combinations
	}
	// Get combination from next index
	suffixList := getCombination(currentIndex+1, containers)
	for _, containerItem := range containers[currentIndex] {
		// Check suffixList
		if suffixList != nil {
			for _, suffix := range suffixList {
				combinations = append(combinations, append([]anyType{containerItem}, suffix...))
			}
		}
	}
	return combinations
}

func comboChain(i [][]anyType) chan []anyType {
	var channel chan []anyType
	for _, j := range i {
		channel = combo(channel, j)
	}
	return channel
}

func combo(input chan []anyType, stuff []anyType) chan []anyType {
	c := make(chan []anyType)
	go func() {
		defer close(c)
		if input == nil {
			for _, s := range stuff {
				c <- []anyType{s}
			}
		} else {
			for i := range input {
				for _, s := range stuff {
					if false {
						// Race Condition
						c <- append(i, s)
					} else {
						// No Race Condition
						j := make([]anyType, len(i))
						copy(j, i)
						c <- append(j, s)
					}
				}
			}
		}
	}()
	return c
}

func combine(cb func([]anyType), inputs ...[]anyType) {
	_combine(cb, len(inputs), inputs...)
}

func _combine(cb func([]anyType), inputsLen int, inputs ...[]anyType) {
	if len(inputs) == 0 {
		cb(make([]anyType, 0, inputsLen))
		return
	}

	end := len(inputs) - 1
	_combine(func(row []anyType) {
		for _, input := range inputs[end] {
			cb(append(row, input))
		}
	}, inputsLen, inputs[:end]...)
}
