package bubble_sort

import "testing"

func TestBubbleSort(t *testing.T) {
	result := BubbleSort([]int {1,2,3,4,7})
	expected := 3
	if result != expected { 
		t.Errorf("result: %v\n expected: %v", result, expected) 
	}
}