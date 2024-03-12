package math

func Average(arrs []int) {
  len := len(arrs)
  sum int

  for _, v := range(arrs) {
    sum = sum + v
  }

  return sum / float64(len)
}
