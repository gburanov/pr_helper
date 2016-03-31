package pr_helper

type Authors map[Author]int

func (authors *Authors) GetLinesStat() (int, int) {
  total := 0
  left := 0
  for author, lines := range *authors {
    if !author.AtWimdu() {
      left += lines
    }
    total += lines
  }
  return left, total
}
