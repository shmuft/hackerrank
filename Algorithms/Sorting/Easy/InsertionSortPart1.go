package main
import "fmt"

func main() {
  var size, i, j, k, last int
  fmt.Scanf("%v", &size)
  arr := make([]int, size, size)
  for i=0; i<size; i++{
      fmt.Scanf("%v", &arr[i])
  } 
  last = arr[size-1]
  for i=size-2; (i >= 0); i--{
    if (arr[i] <= last){
        k = i+1
        break
    }
    arr[i+1] = arr[i]
    for j=0; j<size; j++{
      fmt.Printf("%v ", arr[j])
    }
    fmt.Printf("\n")
  }
  arr[k] = last
  for j=0; j<size; j++{
    fmt.Printf("%v ", arr[j])
  }
  
}
