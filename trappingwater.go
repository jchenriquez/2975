package trappingwater

type Pillar struct {
	Index     int
	LoweAfter int
	Val       int
}

func Trap(height []int) int {
	queue := make([]Pillar, len(height))

	for index, h := range height {

		curr := Pillar{Index: index, Val: h}

		if len(queue) == 0 {
			queue = append(queue, curr)
		} else {
			if queue[len(queue)-1].Val > curr.Val {
				queue[len(queue)-1].LoweAfter = curr.Val
				queue = append(queue, curr)
			}
		}

	}

}
