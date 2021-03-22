package globalvariable2

import "sync"

var lock = &sync.Mutex{}

// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	lock.Lock()
	defer lock.Unlock()

	if instance == nil {

		instance = make(singleton) // <-- thread safe
	}

	return instance
}