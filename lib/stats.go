package pr_helper

import (
  "time"
)

type Stats []Stat

type Stat struct {
  Author Author
  Time time.Time
}

func (stats *Stats) Authors() []Author {
  ret := []Author{}
  for _, stat := range *stats {
    ret = append(ret, stat.Author)
  }
  return ret
}

func (stats *Stats) EarliestTime() time.Time {
  var lowest time.Time = time.Now()
  for _, stat := range *stats {
    if stat.Time.Unix() < lowest.Unix() {
      lowest = stat.Time
    }
  }
  return lowest
}

func (stats *Stats) AverageTime() time.Time {
  var sum int64
  count := 0
  for _, stat := range *stats {
    count += 1
    sum += stat.Time.Unix()
  }
  return time.Unix(sum / int64(count), 0)
}
