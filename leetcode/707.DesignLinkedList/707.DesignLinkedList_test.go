package designlinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListOperations(t *testing.T) {


//   Your MyLinkedList object will be instantiated and called as such:
  obj := Constructor();
  assert.Equal(t,-1, obj.Get(0))
  
  obj.AddAtHead(10);
  obj.AddAtTail(20);

  obj.AddAtIndex(1,15);
  assert.Equal(t, 10, obj.head.Val)
  assert.Equal(t, 15, obj.head.Next.Val)
  assert.Equal(t, 20, obj.head.Next.Next.Val)
  
  obj.DeleteAtIndex(1);
  
  assert.Equal(t, 10, obj.head.Val)
  assert.Equal(t, 20, obj.head.Next.Val) 
  obj.AddAtIndex(2,30)
  assert.Equal(t, 30, obj.head.Next.Next.Val) 
  

//   ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get"]
//   [[],[1],[3],[1,2],[1],[0],[0]]

	obj1 := Constructor()
	obj1.AddAtHead(1);
	obj1.AddAtTail(3);
	obj1.AddAtIndex(1,2);
	assert.Equal(t, 2, obj1.Get(1))
	obj1.DeleteAtIndex(0);
	assert.Equal(t, 2, obj1.Get(0))

	// ["MyLinkedList","addAtTail","get"]
	// [[],[1],[0]]
	obj3 := Constructor()
	obj3.AddAtTail(1)
	assert.Equal(t, 1, obj3.Get(0))


	// ["MyLinkedList","addAtHead","deleteAtIndex","addAtHead","addAtHead","addAtHead","addAtHead","addAtHead","addAtTail","get","deleteAtIndex","deleteAtIndex"]
	// [[],[2],[1],[2],[7],[3],[2],[5],[5],[5],[6],[4]]

}