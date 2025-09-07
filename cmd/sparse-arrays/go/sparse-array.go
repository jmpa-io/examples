package main

// my solution.
func matchingStrings(stringList []string, queries []string) []int32 {

    // setup counter.
    counter := make(map[string]int)
    for _, q := range queries {
        counter[q] = 0
    }


    // populate counter.
    for _, q := range queries {
        for _, s := range stringList {
            if q != s {
                continue
            }
            counter[q]++
        }
    }

    // build output.
    out := make([]int32, len(queries))
    for i, q := range queries {
        out[i] = int32(counter[q])
    }

    // ~return result.
   return out

}

// official solution.
func matchingStringsOffical(stringList []string, queries []string) []int32 {

		// count all strings in stringList
    counter := make(map[string]int)
    for _, s := range stringList {
        counter[s]++
    }

    // answer queries in order
    out := make([]int32, len(queries))
    for i, q := range queries {
        out[i] = int32(counter[q])
    }
    return out
}


