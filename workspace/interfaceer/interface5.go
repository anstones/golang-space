package main

import ("fmt")

type Animal interface{
	ScientificName() string
	Category() string
}

type Named interface{
	Name() string
}

// 包含以上两个接口
type Pet interface {
	Animal
	Named
}


type PetTag struct{
	name string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct{
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main()  {
	pet := PetTag{name: "little pig"}
	_, ok := interface{}(pet).(Named)

	fmt.Printf("PetTag implements interface Named: %v\n", ok)

	dog := Dog {
		PetTag: pet,
		scientificName: "Labrador Retriver",
	}
	_, ok = interface{}(dog).(Animal)  // 可以转换
	fmt.Printf("Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Named) // 可以转换
	fmt.Printf("Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Pet)  // 可以转换
	fmt.Printf("Dog implements interface Pet: %v\n", ok)

}

