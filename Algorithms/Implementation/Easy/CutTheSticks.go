package main
import "fmt"

func sort(arr []int, n int){
  var i, j int
  for i=0; i<n-1; i++{
    for j=i+1; j<n; j++{
      if (arr[i]<arr[j]){
        arr[i], arr[j] = arr[j], arr[i]
      }
    }
  }
}

func main() {
  var n, max, i int
  fmt.Scanf("%v", &n)
  var arr = make([]int, n, n)
  for i=0; i<n; i++{
    fmt.Scanf("%v", &arr[i])
  }
  sort(arr, n)
  max = arr[n-1]
  fmt.Println(n)
  for i=n-1; i>=0; i--{
    if arr[i] > max {
      fmt.Println(i+1)
      max=arr[i]
    }
  }
}
