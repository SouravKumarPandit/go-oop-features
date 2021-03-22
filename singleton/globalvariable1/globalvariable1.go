package globalvariable1
// type global
type singleton map[string]string

var (
	instance singleton
)

func NewClass() singleton {

	if instance == nil {

		instance = make(singleton) // <-- not thread safe
	}

	return instance
}