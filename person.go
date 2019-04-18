package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main(){
	fmt.Println("Hello world")
	dude:= &Person{
		Name:  "Dude",
		Age: 22,
		Bmc: &BMC{
			Height:170,
			Weight:70,
		},
	}
	data, err:= proto.Marshal(dude)
	if err!=nil{
		fmt.Println("Marshalling error:", err)
	}
	fmt.Println("Marshalled data:", data)

	newDude:= &Person{}
	err= proto.Unmarshal(data, newDude)

	if err!=nil{
		fmt.Println("Unmarshalling error:", err)
	}
	fmt.Println(newDude.GetName())
	fmt.Println(newDude.GetAge())
	fmt.Println(newDude.Bmc.GetHeight())
	fmt.Println(newDude.Bmc.GetWeight())
	fmt.Println(newDude.ProtoMessage)

}
