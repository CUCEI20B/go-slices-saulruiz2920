package main

import "fmt"

func main()  {
  var n int
  fmt.Scan(&n)
  var slice []int
  for i := 0; i < n ; i++ {
    var input int
    fmt.Scan(&input)
    slice = append(slice, input)
  }
  var res int
  for i := 0; i < len(slice); i++ {
    res += slice[i]
  }
  fmt.Print(res)
}