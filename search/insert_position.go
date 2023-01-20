package main

import "fmt"


func searchInsert(nums []int, target int) int {
    for idx, num := range(nums) {
        if num >= target {
            return idx
        }
    }
    return len(nums)
}


func main() {

    var arr[] int
    arr = append(arr, 1,3,5,6)
    result := searchInsert(arr, 7)
    fmt.Println(result)
}
