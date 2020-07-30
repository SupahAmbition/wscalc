// A state manager for keeping track of all the
//calculations made by users.
//
// There are a few patterns being used for this datastructure.
// 1. The data strucutre  resembles a stack, but is just an array
//		where add to the end of the array.
//		This allows for an easy way to get the last ten calculations.
// 2. The data structure utilizes the Singleton pattern so that there is
//		only one version of the stack.
// 3. The data structure is thread safe. A mutex is aquired on writing,
//		and copies are made for reads.
//		A soft lock is used for reading, to allow for simultaneous reading.

package calculations

import (
	"sync"
	"time"
)

var instance *Calculations

type Calculations struct {
	calculations []Calculation
	lock         sync.RWMutex
}

type Calculation struct {
	equation  string
	timestamp time.Time // not neccessary, but could be useful in the future.
	user      string
}

func NewCalculation(equation string) *Calculation {
	return &Calculation{
		equation: equation,
		user:     "", // for now we dont have usernames.
	}
}

//only one calculations will be used.
func GetInstance() *Calculations {
	if instance == nil {
		instance = newCalculations()
	}
	return instance
}

//private constructor.
func newCalculations() *Calculations {
	return &Calculations{
		calculations: []Calculation{},
	}
}

func (s *Calculations) Push(c Calculation) {
	s.lock.Lock()
	s.calculations = append(s.calculations, c)
	s.lock.Unlock()
}

//look at the last Calculation made.
//make a copy so the underlying array isnt changed.
func (s *Calculations) Peek() Calculation {

	s.lock.RLock()
	var last Calculation = s.calculations[len(s.calculations)-1]
	var result Calculation = Calculation{
		equation:  last.equation,
		timestamp: last.timestamp,
		user:      last.user,
	}

	s.lock.RUnlock()
	return result
}

//gets the last 10 calculations as a copy.
func (s *Calculations) Peek10() [10]Calculation {

	s.lock.RLock()
	var result [10]Calculation = [10]Calculation{}
	copy(result[:], s.calculations[len(s.calculations)-11:])

	s.lock.RUnlock()
	return result
}
