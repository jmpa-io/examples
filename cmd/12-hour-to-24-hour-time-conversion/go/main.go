package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "time"
    "strings"
)

// timeConversion converts the given 12-hour time (hh:mm:ssAM/PM) to 24-hour time (HH:mm:ss)
// and returns the result as a formatted string.
func timeConversion(s string) string {
    t, err := time.Parse("03:04:05PM", s)
    checkError(err)
    return t.Format("15:04:05")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    outputPath := os.Getenv("OUTPUT_PATH")
    if outputPath == "" {
        fmt.Fprintln(os.Stderr, "error: OUTPUT_PATH environment variable is not set")
        os.Exit(1)
    }
    stdout, err := os.Create(outputPath)
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
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
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

