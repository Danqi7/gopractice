package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	Name string
	Phone string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	session.SetMode(mgo.Monotonic, true)
	
	c := session.DB("gotest").C("people")
	err = c.Insert(&Person{"Danqi", "2243087653"},
		&Person{"Tobin", "2243973647"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name":"Danqi"}).One(&result)
	if err != nil {
		panic(err)
	}

	fmt.Println("Phone: ", result.Phone)

}
