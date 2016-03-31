package pr_helper

import "sort"

func arrayToMap(authors []Author) *Authors {
  ret := make(Authors)
  for _, author := range authors {
    ret[author] += 1
  }
  return &ret
}

func FilterTop(num int, authors *Authors) *Authors {
  reverse := map[int][]Author{}
  for k, v := range *authors {
    // also skip blacklisted
    if k.filtered() {
      continue
    }

    reverse[v] = append(reverse[v], k)
  }

  var a []int
  for k := range reverse {
    a = append(a, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(a)))

  ret := Authors{}
  for _, key := range a {
    authors := reverse[key]
    for _, author := range authors {
      ret[author] = key
      num--
      if num == 0 {
        return &ret
      }
    }
  }
  return &ret
}
