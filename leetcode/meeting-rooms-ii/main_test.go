package meeting_rooms_ii

import "testing"

func Test1(t *testing.T) {
	t.Log(minMeetingRooms([][]int{
		[]int{15, 20},
		[]int{0, 30},
		[]int{5, 10},
	}))
}
