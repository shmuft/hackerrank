package main
import "fmt"

func sort(arr []int, n int){
  var i, k, last int
  last = arr[n]
  for i=n-1; (i >= 0); i--{
    if (arr[i] <= last){
        k = i+1
        break
    }
    arr[i+1] = arr[i]
  }
  arr[k] = last

}

func main() {
  var size, i, j int
  fmt.Scanf("%v", &size)
  arr := make([]int, size, size)
  for i=0; i<size; i++{
      fmt.Scanf("%v", &arr[i])
  } 
  for i=1; i<size; i++{
    sort(arr, i)
    for j=0; j<size; j++{
      fmt.Printf("%v ", arr[j])
    }
    fmt.Printf("\n")
  }
}
