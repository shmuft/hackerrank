package main
import ("fmt")

func main() {
    var  minVal, n, val int
    fmt.Scanf("%v", &n)
    arr := make(map[int]int)
    for i:=0; i<n; i++{
        fmt.Scanf("%v", &val)
        _, ok := arr[val]
        if ok {
            arr[val]++
        } else {
            arr[val] = 1
        }
    }
    x := 0
    minVal = 0
    for key, val := range arr {
        if val >= minVal {
            minVal = val
            x = key
        }
    }
    fmt.Println(n - arr[x])
 //Enter your code here. Read input from STDIN. Print output to STDOUT
}
