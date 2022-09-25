package designcircularqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueOperations(t *testing.T) {
	


//   Your MyCircularQueue object will be instantiated and called as such:
// ["MyCircularQueue","enQueue","enQueue","enQueue","enQueue","Rear","isFull","deQueue","enQueue","Rear"]
// [[3],[1],[2],[3],[4],[],[],[],[4],[]]
  obj := Constructor(3);
  assert.True(t,obj.EnQueue(1) )
  assert.True(t,obj.EnQueue(2) )
  assert.True(t,obj.EnQueue(3) )
  assert.False(t,obj.EnQueue(4) )
  assert.Equal(t,3, obj.Rear() )
  assert.True(t, obj.IsFull() )
  assert.True(t,obj.DeQueue() )
  assert.True(t,obj.EnQueue(4) )
  assert.Equal(t,4, obj.Rear() )
  
  
}