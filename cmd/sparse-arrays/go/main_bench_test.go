package main

import (
    "math/rand"
    "strconv"
    "testing"
)

func makeTestData(n, m int) ([]string, []string) {
    stringList := make([]string, n)
    queries := make([]string, m)

    for i := 0; i < n; i++ {
        stringList[i] = strconv.Itoa(rand.Intn(1000))
    }
    for i := 0; i < m; i++ {
        queries[i] = strconv.Itoa(rand.Intn(1000))
    }
    return stringList, queries
}

func BenchmarkMatchingStringsNaive(b *testing.B) {
    stringList, queries := makeTestData(100000, 1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = matchingStrings(stringList, queries)
    }
}

func BenchmarkMatchingStringsOfficial(b *testing.B) {
    stringList, queries := makeTestData(100000, 1000)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = matchingStringsOffical(stringList, queries)
    }
}
