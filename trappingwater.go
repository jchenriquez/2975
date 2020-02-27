package trappingwater

import (
  "math"
  "fmt"
)

func Trap(height []int) (trapped int) {
  if len(height) < 2 {
    return 0
  }

  pillars := make([]int, 0, len(height))

  updateTrapped := func (i int) {
    if i - pillars[len(pillars)-1] >= 2 {
        top := pillars[len(pillars)-1]
        topNormalizedH := height[top] - height[top+1]
        currNormedH := height[i] - height[i-1]
        trap := int(math.Min(float64(topNormalizedH), float64(currNormedH))) * (i - (pillars[len(pillars)-1]+1))
        fmt.Printf("trap %d\n", trap)
        trapped += trap
    }
  }


  for i := 1; i < len(height); i++ {
    if height[i] == 0 {
      continue
    }
    
    enteredLoop := false

    for len(pillars) > 0 && height[i] >= height[pillars[len(pillars)-1]] {
      enteredLoop = true
      updateTrapped(i)
      pillars = pillars[:len(pillars)-1]
    }

    if len(pillars) > 0 && !enteredLoop {
      updateTrapped(i)
    }
    pillars = append(pillars, i)
  }

  return
}