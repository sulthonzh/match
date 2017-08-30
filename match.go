package match

type AnyType interface{}

// GetCombination is a function for getcombination with multiple array
// by me => inspird from http://stackoverflow.com/users/3447216/priyesh
// http://stackoverflow.com/questions/22632826/combination-of-elements-of-multiple-arrays
func GetCombination(currentIndex int, containers [][]AnyType) (combinations [][]AnyType) {
	if currentIndex == len(containers) {
		// Skip the items for the last container
		combinations = append(combinations, []AnyType{})
		return combinations
	}
	// Get combination from next index
	suffixList := GetCombination(currentIndex+1, containers)
	for _, containerItem := range containers[currentIndex] {
		// Check suffixList
		if suffixList != nil {
			for _, suffix := range suffixList {
				combinations = append(combinations, append([]AnyType{containerItem}, suffix...))
			}
		}
	}
	return combinations
}

// ComboChain => Thanks for https://github.com/angch
func ComboChain(i [][]AnyType) chan []AnyType {
	var channel chan []AnyType
	for _, j := range i {
		channel = combo(channel, j)
	}
	return channel
}

func combo(input chan []AnyType, stuff []AnyType) chan []AnyType {
	c := make(chan []AnyType)
	go func() {
		defer close(c)
		if input == nil {
			for _, s := range stuff {
				c <- []AnyType{s}
			}
		} else {
			for i := range input {
				for _, s := range stuff {
					if false {
						// Race Condition
						c <- append(i, s)
					} else {
						// No Race Condition
						j := make([]AnyType, len(i))
						copy(j, i)
						c <- append(j, s)
					}
				}
			}
		}
	}()
	return c
}

// Combine => Thanks for this https://github.com/erikdubbelboer
func Combine(cb func([]AnyType), inputs ...[]AnyType) {
	_combine(cb, len(inputs), inputs...)
}

func _combine(cb func([]AnyType), inputsLen int, inputs ...[]AnyType) {
	if len(inputs) == 0 {
		cb(make([]AnyType, 0, inputsLen))
		return
	}

	end := len(inputs) - 1
	_combine(func(row []AnyType) {
		for _, input := range inputs[end] {
			cb(append(row, input))
		}
	}, inputsLen, inputs[:end]...)
}
