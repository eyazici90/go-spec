package main

import (
	"log"

	"github.com/eyazici90/go-spec/spec"
)

type User struct {
	name, surname, email string
}

type findByNameSpec struct {
	name string
}

func (f findByNameSpec) Eval() func(interface{}) bool {
	return func(i interface{}) bool {
		return i.(User).name == f.name
	}
}

type findByEmailSpec struct {
	email string
}

func (f findByEmailSpec) Eval() func(interface{}) bool {
	return func(i interface{}) bool {
		return i.(User).email == f.email
	}
}

func main() {
	user := User{
		email: "test.123@gmail.com",
		name:  "Emre",
	}

	spec1 := findByEmailSpec{
		email: "test.123@gmail.com",
	}
	spec2 := findByNameSpec{
		name: "Emre",
	}

	var chains []spec.Specifier
	chains = append(chains, spec.And(spec1, spec2))
	chains = append(chains, spec.And(spec1).AndNot(spec2))
	chains = append(chains, spec.Or(spec1, spec2))

	for _, v := range chains {
		ok := v.Eval()(user)
		log.Println(ok)
	}
}
