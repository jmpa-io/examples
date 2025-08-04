package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// miniMaxSum receives 5 integers.
// Prints two space-seperated integers.
// Returns a space separated string of minValue & maxValue.
// eg. arr = [1, 2, 3, 4, 5]; output = "10 14".
func miniMaxSum(arr []int32) {
    var total int64
    var min, max int32 = arr[0], arr[0] // NOTE: this must not be initalized to be zero.
    for _, n := range arr {
      total += int64(n)
      if n < min {
        min = n
      }
      if n > max {
        max = n
      }
    }
    fmt.Printf("%v %v\n", total-int64(max), total-int64(min))
}

// OVERVIEW:
// Given 5 positive integers.
// Find minimum and maximum values that can be calculated by summing exactly four of the five integers.
// Print the respective minimum and maximum values as a single line of two spaced-separated long integers.
// eg. arr = [1, 3, 5, 7, 9]; smallest value = 16 (1 + 3 + 5 + 7); largest value = 24 (3 + 5 + 7 + 9);
//
// CONSTRAINTS:
// 1 < arr[i] < 10^9 (eg. only positive numbers).
//
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)
    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32
    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

