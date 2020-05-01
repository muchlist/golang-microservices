package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	elm := []int{9, 8, 7, 6, 5}

	Sort(elm)

	assert.NotNil(t, elm)
	assert.EqualValues(t, 5, len(elm))
	assert.EqualValues(t, 5, elm[0])
	assert.EqualValues(t, 6, elm[1])
	assert.EqualValues(t, 7, elm[2])
	assert.EqualValues(t, 8, elm[3])
	assert.EqualValues(t, 9, elm[4])
}

func TestBubbleSortBestCase(t *testing.T) {
	elm := []int{5, 6, 7, 8, 9}

	Sort(elm)

	assert.NotNil(t, elm)
	assert.EqualValues(t, 5, len(elm))
	assert.EqualValues(t, 5, elm[0])
	assert.EqualValues(t, 6, elm[1])
	assert.EqualValues(t, 7, elm[2])
	assert.EqualValues(t, 8, elm[3])
	assert.EqualValues(t, 9, elm[4])
}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	elm := getElements(5)

	assert.NotNil(t, elm)
	assert.EqualValues(t, 5, len(elm))
	assert.EqualValues(t, 4, elm[0])
	assert.EqualValues(t, 3, elm[1])
	assert.EqualValues(t, 2, elm[2])
	assert.EqualValues(t, 1, elm[3])
	assert.EqualValues(t, 0, elm[4])
}

func BenchmarkBubbleSort10(b *testing.B) {
	elm := getElements(10)

	for i := 0; i < b.N; i++ {
		Sort(elm)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elm := getElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elm)
	}
}

//Menggunakan sort bawaan golang, lebih lambat jika jumlah
//elementnya sedikit, namun jika sudah seratus ribuan jauh lebih baik
func BenchmarkSort10(b *testing.B) {
	elm := getElements(10)

	for i := 0; i < b.N; i++ {
		Sort(elm)
	}
}

func BenchmarkSort100000(b *testing.B) {
	elm := getElements(100000)

	for i := 0; i < b.N; i++ {
		Sort(elm)
	}
}
