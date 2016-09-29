package godoku

import "math/rand"

//Tracker is a struct that will contain all attempted
type Tracker struct {
	attempts []map[byte]bool
}

//NewTracker creates a tracker based on a passed size
func NewTracker(size int) *Tracker {
	out := new(Tracker)
	out.attempts = make([]map[byte]bool, size*size)
	for i := 0; i < size*size; i++ {
		out.attempts[i] = make(map[byte]bool, size)
	}
	return out
}

//Clear replaces one of the maps with a new one
func (t *Tracker) Clear(index int) {
	//This is probably a naive way of clearing. I should probably just reset each key in the map to false
	//but it will work for now.
	size := len(t.attempts[index])
	for i := 0; i < size; i++ {
		t.attempts[index] = make(map[byte]bool, size)
	}
}

func availableNum(size int, used map[byte]bool) byte {
	//Another function that could be better managed.
	//TODO: Keep track of UNUSED numbers, as well as used, and then select from the unused.
	//This method *theoretically* could take a very long time. Until I have a working verion it's fine.
	for {
		out := byte(rand.Intn(size) + 1)
		if _, contained := used[out]; contained {
			continue
		}
		return out
	}
}
