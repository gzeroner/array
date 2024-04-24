package array

import "math/rand"

type number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | complex64 | complex128
}

// Pad pads the array with the given value to the given size.
// If the array is already larger than or equal to the size, it is returned unchanged.
func Pad[T any](arr []T, size int, value T) []T {
	if len(arr) >= size {
		return arr
	}
	newArr := make([]T, size)
	copy(newArr, arr)
	for i := len(arr); i < size; i++ {
		newArr[i] = value
	}
	return newArr
}

// Sum returns the sum of all elements in the array.
func Sum[T number](arr []T) T {
	var sum T
	for _, v := range arr {
		sum += v
	}
	return sum
}

// Fill returns an array of the given size filled with the given value.
func Fill[T any](num int, value T) []T {
	if num < 0 {
		panic("array size cannot be negative")
	}
	arr := make([]T, num)
	for i := 0; i < num; i++ {
		arr[i] = value
	}
	return arr
}

// Push returns a new array with the given value appended to the end of the original array.
func Push[T any](arr []T, value T) []T {
	newArr := make([]T, len(arr)+1)
	copy(newArr, arr)
	newArr[len(arr)] = value
	return newArr
}

// Rand returns a random element from the array.
func Rand[T any](arr []T) T {
	return arr[rand.Intn(len(arr))]
}

// Merge returns a new array that contains all elements of the given arrays.
func Merge[T any](list ...[]T) []T {
	var result []T
	for _, arr := range list {
		result = append(result, arr...)
	}
	return result
}

// Filter returns a new array that contains only the elements that satisfy the given predicate.
func Filter[T any](arr []T, f func(T) bool) []T {
	var result []T
	for _, v := range arr {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// Search returns the index of the first occurrence of the given value in the array, or -1 if not found.
func Search[T comparable](needle T, haystack []T) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}

// Unique returns a new array that contains only the unique elements of the given array.
func Unique[T comparable](arr []T) []T {
	m := make(map[T]struct{})
	var result []T
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// Product returns the product of all elements in the array.
func Product[T number](arr []T) T {
	var result T = 1
	for _, v := range arr {
		result *= v
	}
	return result
}

// Reverse returns a new array that contains the elements of the original array in reverse order.
func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// Unshift returns a new array with the given value prepended to the beginning of the original array.
func Unshift[T any](arr []T, value T) []T {
	newArr := make([]T, len(arr)+1)
	newArr[0] = value
	copy(newArr[1:], arr)
	return newArr
}
