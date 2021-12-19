package tests

import "testing"


func TestSum(t *testing.T){

	numbers := [5]int{1,2,3,4,5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d, wanted %d, array is %v", got, want, numbers)
	}
}

