package main

import "testing"

func TestMatchingStrings(t *testing.T) {
    cases := []struct {
        name    string
        fn      func([]string, []string) []int32
        list    []string
        queries []string
        want    []int32
    }{
        {
            name:    "naive version simple case",
            fn:      matchingStrings,
            list:    []string{"ab", "ab", "abc"},
            queries: []string{"ab", "abc", "bc"},
            want:    []int32{2, 1, 0},
        },
        {
            name:    "official version simple case",
            fn:      matchingStringsOffical,
            list:    []string{"ab", "ab", "abc"},
            queries: []string{"ab", "abc", "bc"},
            want:    []int32{2, 1, 0},
        },
        {
            name:    "mixed characters",
            fn:      matchingStrings,
            list:    []string{"a", "b", "a"},
            queries: []string{"a", "b", "c"},
            want:    []int32{2, 1, 0},
        },
        {
            name:    "mixed characters official",
            fn:      matchingStringsOffical,
            list:    []string{"a", "b", "a"},
            queries: []string{"a", "b", "c"},
            want:    []int32{2, 1, 0},
        },
    }

    for _, c := range cases {
        got := c.fn(c.list, c.queries)
        for i := range got {
            if got[i] != c.want[i] {
                t.Errorf("%s: got %v, want %v", c.name, got, c.want)
                break
            }
        }
    }
}
