package main

import "sort"

func arrayToMap(authors []Author) map[Author]int {
  ret := make(map[Author] int )
  for _, author := range authors {
    ret[author] += 1
  }
  return ret
}

func filterTop(num int, authors map[Author]int) map[Author]int {
  reverse := map[int][]Author{}
  for k, v := range authors {
    // also skip blacklisted
    if !k.available() {
      continue
    }

    reverse[v] = append(reverse[v], k)
  }

  var a []int
  for k := range reverse {
    a = append(a, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  ret := map[Author]int{}
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
