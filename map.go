package main
import(
  "fmt"
  "sort"
)
func main(){
  m := make(map [string]int )
  m["key"] =42
  fmt.Println(m)

  states := make(map [string]string)
  // fmt.Println(states)
  states["WA"]= "Washington"
  states["OR"]= "Oregon"
  states["CA"]= "California"
  fmt.Println(states)

  //set value
  cali := states["CA"]
  fmt.Println(cali)
  //delete
  delete(states, "CA")
  fmt.Println(states)

  states["NY"]= "New York"
  //loop
  for k, v := range states{
    fmt.Printf("%v : %v \n",k,v)
  }

  keys := make([]string, len(states))
  i:=0

  for k:= range states{
    keys[i] =k
    i++
  }
  sort.Strings(keys)
  fmt.Println("\n", keys)

  for i := range keys{
    fmt.Println(states[keys[i]])
  }

}
