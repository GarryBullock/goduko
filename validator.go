package godoku

//Check to see if a given board is valid solution.
func Validate(arr *[]byte, length int) bool {
	if validateBlocks(arr, length) {
		if validateColums(arr, length) {
			if validateColums(arr, length) {
				return true
			}
		}
	}
	return false
}

func validateBlocks(arr *[]byte, length int) bool {
	//Validate block by block... Three loops is obviously terribly inefficient. But for the time
	//being, I can't figure out a better way. Its ugly and slow but it works.
	block := make([]byte, length)
	for i := 0; i < length; i += length / 3 {
		for j := 0; j < length; j += length / 3 {
			for k := 0; k < length/3; k++ {
				val := i + length*j + k*length
				block[k*length/3] = (*arr)[val]
				block[k*length/3+1] = (*arr)[val+1]
				block[k*length/3+2] = (*arr)[val+2]
			}
			ok := validate(&block, 9)
			if !ok {
				return false
			}
		}
	}
	return true
}

func validateRows(arr *[]byte, length int) bool {
	for i := 0; i < length; i++ {
		row := (*arr)[i*length : i*length+length]
		ok := validate(&row, length)
		if !ok {
			return false
		}
	}
	return true
}

func validateColums(arr *[]byte, length int) bool {
	for i := 0; i < length; i++ {
		column := make([]byte, length)
		for j := 0; j < length; j++ {
			column[j] = (*arr)[i+j*length]
		}
		ok := validate(&column, length)
		if !ok {
			return false
		}
	}
	return true
}

func validate(arr *[]byte, length int) bool {
	entries := make(map[int]byte, length)
	count := 0
	for i := 0; i < length; i++ {
		//TODO: Casting int to byte can lose info... not expecting size to be bigger than 255.
		if (*arr)[i] > 0 && (*arr)[i] <= byte(length) {
			entries[int((*arr)[i])] = (*arr)[i]
			count++
		}
	}
	//check that all the entries added to the map were unique
	result := len(entries) == count
	return result
}
