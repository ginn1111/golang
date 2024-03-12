package math

func Average(arrs []int) float64 {
  sum := float64(0)

  for _, v := range arrs {
    sum += float64(v)
  }

  result := sum / float64(len(arrs))

  return result
}
