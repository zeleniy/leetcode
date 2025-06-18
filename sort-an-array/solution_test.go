package sort_an_array

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseDataSet struct {
	answer []int
	nums   []int
}

func BenchmarkMergeSortRecursiveWithSlices(b *testing.B) {

	dataSet := getTestDataSet()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			MergeSortRecursiveWithSlices(data.nums)
		}
	}
}

func TestMergeSortRecursiveWithSlices(t *testing.T) {

	for _, data := range getTestDataSet() {
		assert.Equal(t, data.answer, MergeSortRecursiveWithSlices(data.nums))
	}
}

func BenchmarkMergeSortRecursiveWithIndexes(b *testing.B) {

	dataSet := getTestDataSet()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, data := range dataSet {
			MergeSortRecursiveWithIndexes(data.nums)
		}
	}
}

func TestMergeSortRecursiveWithIndexes(t *testing.T) {

	for _, data := range getTestDataSet() {
		MergeSortRecursiveWithIndexes(data.nums)
		assert.Equal(t, data.answer, data.nums)
	}
}

var testDataSets []TestCaseDataSet

func init() {

	baseArrays := map[string][]int{
		"simple":        {2, 0, -1, 1, -2},
		"sorted":        generateSortedArray(1000),
		"reversed":      reverseArray(generateSortedArray(1000)),
		"all_equal":     make([]int, 1000),
		"random":        generateRandomArray(1000),
		"almost_sorted": append(generateSortedArray(990), generateRandomArray(10)...),
		"few_unique":    generateSortedArrayWithDuplicates(1000, 10),
		"large_nums":    generateRandomArrayWithLargeNumbers(1000, 1e9),
		"with_negs":     generateRandomArrayWithNegatives(1000),
		"alternating":   generateAlternatingArray(1000),
		"big_first":     append([]int{1e9}, generateSortedArray(999)...),
	}

	testDataSets = make([]TestCaseDataSet, 0, len(baseArrays))
	for _, arr := range baseArrays {
		testDataSets = append(testDataSets, generateTestCase(arr))
	}
}

func getTestDataSet() []TestCaseDataSet {
	// Возвращаем копию, чтобы тесты не влияли друг на друга
	result := make([]TestCaseDataSet, len(testDataSets))
	for i, ds := range testDataSets {
		numsCopy := make([]int, len(ds.nums))
		copy(numsCopy, ds.nums)
		result[i] = TestCaseDataSet{
			answer: ds.answer,
			nums:   numsCopy,
		}
	}
	return result
}

// Вспомогательная функция для создания тестового набора данных
func generateTestCase(input []int) TestCaseDataSet {
	answer := make([]int, len(input))
	copy(answer, input)
	sort.Ints(answer) // Создаем отсортированную версию массива
	return TestCaseDataSet{
		answer: answer,
		nums:   input,
	}
}

// Вспомогательные функции для генерации данных

func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i + 1
	}
	return arr
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size * 10)
	}
	return arr
}

func generateRandomArrayWithLargeNumbers(size, max int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(max)
	}
	return arr
}

func generateRandomArrayWithNegatives(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size*10) - size*5 // Диапазон от -size*5 до size*5
	}
	return arr
}

func generateSortedArrayWithDuplicates(size, uniqueCount int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i%uniqueCount + 1
	}
	return arr
}

func generateAlternatingArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		if i%2 == 0 {
			arr[i] = rand.Intn(size / 10)
		} else {
			arr[i] = rand.Intn(size * 10)
		}
	}
	return arr
}
