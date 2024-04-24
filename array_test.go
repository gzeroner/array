package array

import (
	"testing"
)

func TestPad(t *testing.T) {
	arr := Pad([]string{"a", "b", "c"}, 5, "x")
	if len(arr) != 5 {
		t.Errorf("expected length of 5, got %d", len(arr))
	}
	if arr[3] != "x" && arr[4] != "x" {
		t.Errorf("expected 'x' at index 3 and 4, got %s and %s", arr[3], arr[4])
	}
}

func BenchmarkPad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pad([]string{"a", "b", "c"}, 5, "x")
	}
}

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	sum := Sum(arr)
	if sum != 15 {
		t.Errorf("expected sum of 15, got %d", sum)
	}
}

func BenchmarkSum(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(arr)
	}
}

func TestFill(t *testing.T) {
	arr := Fill(5, 1)
	if arr[0] != 1 || arr[1] != 1 || arr[2] != 1 || arr[3] != 1 || arr[4] != 1 {
		t.Errorf("expected all elements to be 1, got %d, %d, %d, %d, %d", arr[0], arr[1], arr[2], arr[3], arr[4])
	}
}

func BenchmarkFill(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fill(5, 1)
	}
}

func TestPush(t *testing.T) {
	arr := Push([]int{1, 2, 3}, 4)
	if len(arr) != 4 || arr[3] != 4 {
		t.Errorf("expected length of 4 and last element to be 4, got %d and %d", len(arr), arr[3])
	}
}

func BenchmarkPush(b *testing.B) {
	arr := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Push(arr, 4)
	}
}

func TestRand(t *testing.T) {
	element := Rand([]string{"a", "b", "c", "d", "e"})
	if element != "a" && element != "b" && element != "c" && element != "d" && element != "e" {
		t.Errorf("expected random element from the array, got %s", element)
	}
}

func BenchmarkRand(b *testing.B) {
	arr := []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		Rand(arr)
	}
}

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	merged := Merge(arr1, arr2)
	if len(merged) != 6 || merged[0] != 1 || merged[1] != 2 || merged[2] != 3 || merged[3] != 4 || merged[4] != 5 || merged[5] != 6 {
		t.Errorf("expected merged array to be [1, 2, 3, 4, 5, 6], got %v", merged)
	}
}

func BenchmarkMerge(b *testing.B) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	for i := 0; i < b.N; i++ {
		Merge(arr1, arr2)
	}
}

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filterArr := Filter(arr, func(v int) bool {
		if v%3 == 0 {
			return true
		}
		return false
	})
	if len(filterArr) != 3 {
		t.Errorf("expected length of 3, got %d", len(filterArr))
	}
}

func BenchmarkFilter(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Filter(arr, func(v int) bool {
			if v%3 == 0 {
				return true
			}
			return false
		})
	}
}

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	index := Search(5, arr)
	if index != 4 {
		t.Errorf("expected index of 4, got %d", index)
	}
}

func BenchmarkSearch(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		Search(5, arr)
	}
}

func TestUnique(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	uniqueArr := Unique(arr)
	if len(uniqueArr) != 5 || uniqueArr[0] != 1 || uniqueArr[1] != 2 || uniqueArr[2] != 3 || uniqueArr[3] != 4 || uniqueArr[4] != 5 {
		t.Errorf("expected unique array to be [1, 2, 3, 4, 5], got %v", uniqueArr)
	}
}

func BenchmarkUnique(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		Unique(arr)
	}
}

func TestProduct(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	product := Product(arr)
	if product != 120 {
		t.Errorf("expected product of 120, got %d", product)
	}
}

func BenchmarkProduct(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Product(arr)
	}
}

func TestReverse(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	reversedArr := Reverse(arr)
	if len(reversedArr) != 5 || reversedArr[0] != 5 || reversedArr[1] != 4 || reversedArr[2] != 3 || reversedArr[3] != 2 || reversedArr[4] != 1 {
		t.Errorf("expected reversed array to be [5, 4, 3, 2, 1], got %v", reversedArr)
	}
}

func BenchmarkReverse(b *testing.B) {
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Reverse(arr)
	}
}

func TestUnshift(t *testing.T) {
	arr := []int{1, 2, 3}
	unshiftArr := Unshift(arr, 0)
	if len(unshiftArr) != 4 || unshiftArr[0] != 0 || unshiftArr[1] != 1 || unshiftArr[2] != 2 || unshiftArr[3] != 3 {
		t.Errorf("expected unshifted array to be [0, 1, 2, 3], got %v", unshiftArr)
	}
}

func BenchmarkUnshift(b *testing.B) {
	arr := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		Unshift(arr, 0)
	}
}
