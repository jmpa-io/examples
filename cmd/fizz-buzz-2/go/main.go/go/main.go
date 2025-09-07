package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func fizzBuzz(n int32) {
    var sb strings.Builder
    for i := int32(1); i <= n; i++ {
        var line string
        if i % 3 == 0 {
            line += "Fizz"
        }
        if i % 5 == 0 {
            line += "Buzz"
        }
        if line == "" {
            line = fmt.Sprint(i)
        }
        sb.WriteString(line + "\n")
    }
    fmt.Print(sb.String())
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    fizzBuzz(n)
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

