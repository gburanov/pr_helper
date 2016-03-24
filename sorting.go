package main

import "sort"
import "fmt"

func arrayToMap(authors []string) {
  ret := make(map[string] int )
  for _, author := range authors {
    ret[author] += 1
  }

  reverse := map[int][]string{}
  for k, v := range ret {
    reverse[v] = append(reverse[v], k)
  }

  var a []int
  for k := range reverse {
    a = append(a, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  for _, key := range a {
    author := reverse[key]
    fmt.Println(author, " Values ", key)
  }
}
