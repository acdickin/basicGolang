package main
import(
  "fmt"
)
type Animal interface{
  Speak() string
}
type Dog struct{
}
func (d Dog) Speak() string{
  return "Woof!"
}
type Cat struct{
}
func (c Cat) Speak() string{
  return "Meow!"
}
type Sheep struct{
}
func (s Sheep) Speak() string{
  return "HEY!"
}
func main(){
  poodle:=Animal(Dog{})
  fmt.Println(poodle)

  animals:=[]Animal{Dog{},Cat{},Sheep{}}
  for _, animal := range animals{
    fmt.Println(animal.Speak())
  }
}
