package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// determineRations calculates and prints the ratio of positive, negative, and
// neutral numbers in the given slice, formated to 6 decimal places.
func determineRatios(numbers []int32) {
    var total = float64(len(numbers))
    var positive, negative, neutral float64
    for _, n := range numbers {
        if n > 0 {
            positive++
            continue
        }
        if n < 0 {
            negative++
            continue
        }
        neutral++
    }
    fmt.Printf("%.6f\n", positive/total)
    fmt.Printf("%.6f\n", negative/total)
    fmt.Printf("%.6f\n", neutral/total)
}

func main() {

		// Setup a buffered reader for large input.
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

		// Read and parse the number of elements in the slice.
    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int(nTemp)

		// Read and parse the array of integers from input.
    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
		var numbers = make([]int32, n)
    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        numbers[i] = int32(arrItemTemp)
    }

    // Calculate and print the ratio of positive, negative, and neutral values.
    determineRatios(numbers)
}

// readLine reads a line of input from the given buffered reader.
func readLine(reader *bufio.Reader) string {
    str, _, err := reader.Readslice
    if err == io.EOF {
        return ""
    }
    return strings.TrimRight(string(str), "\r\n")
}

// checkError panics if an error is encountered.
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
