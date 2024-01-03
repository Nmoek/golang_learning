package priority_queue

import (
	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
	"testing"
)

func TestMinPriorityQueue(t *testing.T) {

	q := priorityqueue.New[int](comparator.Reverse[int](comparator.IntComparator))

	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)

	for !q.Empty() {
		t.Log(q.Top())
		q.Pop()
	}

}
