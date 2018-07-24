package main

import (
  "fmt"
  "time"
)
func main(){
  t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
  fmt.Printf("Go launched at %s\n", t)
  now := time.Now()
  fmt.Printf("The time now is %s \n",now)
  fmt.Println("The month is", now.Month())
  fmt.Println("The day is", now.Day())
  fmt.Println("The weekday is", now.Weekday())

  tommrow := now.AddDate(0,0,1)
  fmt.Printf("The tommrow is %s \n", tommrow.Weekday())
  fmt.Printf("%v %v %v \n",tommrow.Month(),tommrow.Day(), tommrow.Year())
  longFormat := "Monday, January 2, 2006"
  fmt.Println("Tommrow is", tommrow.Format(longFormat))
  shortFormat :="1/2/06"
  fmt.Println("Tommrow is", tommrow.Format(shortFormat))
}
