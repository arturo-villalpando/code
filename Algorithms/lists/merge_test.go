package lists

import "testing"

func TestMergeTwoLists(t *testing.T) {
	list1 := &LinkedList{}
	list1.Append(1)
	list1.Append(2)
	list1.Append(4)

	list2 := &LinkedList{}
	list2.Append(1)
	list2.Append(3)
	list2.Append(4)

	response := MergeTwoLists(list1.Head, list2.Head)
	t.Errorf("Response is %v", response)
}
