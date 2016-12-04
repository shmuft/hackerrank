package main
import "fmt"

func main() {
  var n, k, q, i, val int
  fmt.Scanf("%v %v %v", &n, &k, &q)
  k = k % n
  var arr = make([]int, n, n)
  for i=k; i<n+k; i++{
    fmt.Scanf("%v", &val)
    if (i>=n) {
      arr[i-n] = val
    } else {
      arr[i] = val
    }
  }
  for i=0; i<q; i++{
    fmt.Scanf("%v", &val)
    fmt.Printf("%v\n", arr[val])
  }
}
