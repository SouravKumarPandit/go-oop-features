package packet

type SampleObject struct {
	Name string
	Number int
	Hidden string
}

func (obj SampleObject) NewSampleObject() *SampleObject {
	return &SampleObject{Name: "Sourav Pandit",Number:4055 ,Hidden: "hidden property for this Package"}
}


func  NewSampleObject(name string,number int) *SampleObject {
	return &SampleObject{Name: name,Number: number,Hidden: "hidden property for this Package"}
}