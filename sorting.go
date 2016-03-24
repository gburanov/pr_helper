package main

import "sort"

func arrayToMap(authors []string) map[string]int {
  ret := make(map[string] int )
  for _, author := range authors {
    ret[author] += 1
  }
  return ret
}

func filterTop(num int, authors map[string]int) map[string]int {
  reverse := map[int][]string{}
  for k, v := range authors {
    reverse[v] = append(reverse[v], k)
  }

  var a []int
  for k := range reverse {
    a = append(a, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  ret := map[string]int{}
  for _, key := range a {
    authors := reverse[key]
    for _, author := range authors {
      ret[author] = key
      num--
      if num == 0 {
        return ret
      }
    }
  }
  return ret
}
