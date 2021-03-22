package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"oopfeatures/souravkumarpandit/objectaddress/packet"
	"reflect"
)

func main() {
	obj1 := packet.SampleObject{
		Name:   "Sourav Pandit",
		Number: 4055,
		Hidden: "hidden property for this Package",
	}
	//obj1 = *(obj1.NewSampleObject())
	obj2 := packet.NewSampleObject("Sourav Pandit", 4055)
	//fmt.Println(obj1.Number,"  : ",obj1.Name,obj1.hidden)
	fmt.Println(obj1.Number, "  : ", obj1.Name)
	fmt.Println(obj2.Number, "  : ", obj2.Name)

	fmt.Println("---------------------1--------------------------")
	fmt.Println(&obj1)
	fmt.Println(&obj2)


	fmt.Println("----------------------2-------------------------")
	fmt.Printf("%p\n", &obj1)
	fmt.Println(&obj2)


	fmt.Println("-----------------------3------------------------")
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, *obj2)
	fmt.Println((obj1 == *obj2 ))


	fmt.Println("------------------------4-----------------------")
	fmt.Printf("obj1: %v \nobj2: %v \n", &obj1, obj2)
	fmt.Println((&obj1 == obj2))


	fmt.Println("------------------------5-----------------------")
	/* address comparison with struct will be false*/
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, obj2)
	fmt.Println(reflect.DeepEqual(obj1, obj2))


	fmt.Println("------------------------6-----------------------")
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, *obj2)
	fmt.Println(reflect.DeepEqual(&obj1, *obj2))


	fmt.Println("------------------------7-----------------------")
	/* address comparison with struct will be false*/
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, obj2)
	fmt.Println(cmp.Equal(obj1, obj2))


	fmt.Println("------------------------8----------------------")
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, *obj2)
	fmt.Println(cmp.Equal(&obj1, *obj2))


	obj3 := packet.SampleObject{
		Name:   "Sourav Pandit",
		Number: 4055,
		Hidden: "hidden property for this Package",
	}

	fmt.Println("---------------------9--------------------------")
	fmt.Println(&obj1)
	fmt.Println(&obj3)


	fmt.Println("----------------------10-------------------------")
	fmt.Printf("%p\n", &obj1)
	fmt.Println(&obj3)


	fmt.Println("-----------------------12------------------------")
	fmt.Printf("obj1: %v \nobj3: %v \n", obj1, obj3)
	fmt.Println((obj1 == obj3 ))


	fmt.Println("------------------------13-----------------------")
	fmt.Printf("obj1: %v \nobj3: %v \n", &obj1, &obj3)
	fmt.Println((&obj1 == &obj3))


	fmt.Println("------------------------14-----------------------")
	/* address comparison with struct will be false*/
	fmt.Printf("obj1: %v \nobj3: %v \n", obj1, obj3)
	fmt.Println(reflect.DeepEqual(obj1, obj3))


	fmt.Println("-----------------------15-----------------------")
	fmt.Printf("obj1: %v \nobj2: %v \n", obj1, obj3)
	fmt.Println(reflect.DeepEqual(obj1, obj3))


	fmt.Println("-----------------------16-----------------------")
	/* address comparison with struct will be false*/
	fmt.Printf("obj1: %v \nobj3: %v \n", obj1, obj3)
	fmt.Println(cmp.Equal(obj1, obj3))


	fmt.Println("-----------------------17----------------------")
	fmt.Printf("obj1: %v \nobj3: %v \n", &obj1, &obj3)
	fmt.Println(cmp.Equal(&obj1, &obj3))


}
