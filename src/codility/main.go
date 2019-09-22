package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Printf("%+v\n", MaxDoubleSliceSum([]int{3, 2, 6, -1, 4, 5, -1, 2}))
	//fmt.Printf("%+v\n", CyclicRotation([]int{3, 8, 9, 7, 6}, 3))
	//fmt.Printf("%+v\n", PermMissingElem([]int{2, 3, 1, 5}))
	//fmt.Printf("%+v\n", TapeEquilibrium([]int{-1000, 1000}))
	//fmt.Printf("%+v\n", BinaryGap(1041))
	//fmt.Printf("%+v\n", OddOccurrencesInArray([]int{2, 2, 3, 3, 4}))
	//fmt.Printf("%+v\n", PermCheck([]int{4, 1, 2, 3}))
	//fmt.Printf("%+v\n", FrogRiverOne(5, []int{1, 3, 1, 4, 2, 3, 5, 4}))
	//fmt.Printf("%+v\n", MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4}))
	//fmt.Printf("%+v\n", MissingInteger([]int{1, 3, 6, 4, 1, 2}))
	//fmt.Printf("%+v\n", PassingCars([]int{0, 1, 0, 1, 1}))
	//fmt.Printf("%+v\n", GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
	//fmt.Printf("%+v\n", MinAvgTwoSlice([]int{4, 2, 2, 5, 1, 5, 8}))
	//fmt.Printf("%+v\n", MaxProductOfThree([]int{-3, 1, 2, -2, 5, 6}))
	//fmt.Printf("%+v\n", Distinct([]int{2, 1, 1, 2, 3, 1}))
	//fmt.Printf("%+v\n", Triangle([]int{50, 40, 20, 10, 1}))
	fmt.Printf("%+v\n", NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}))
}

// https://app.codility.com/demo/results/trainingJZ8UZK-4ET/
func NumberOfDiscIntersections(A []int) int {
	res := 0
	n := len(A)
	r, l := make([]int, n), make([]int, n)
	for i := range A {
		l[i] = i - A[i]
		r[i] = i + A[i]
	}
	sort.Ints(l)
	sort.Ints(r)
	for i := range A {
		end := r[i]
		count := sort.Search(len(l), func(i int) bool {
			return l[i] > end
		})
		count -= i + 1
		res += count
		if res > 10000000 {
			return -1
		}
	}
	return res
}

// https://app.codility.com/demo/results/training6HR5FG-G2F/
func Distinct(A []int) int {
	h := make(map[int]bool)

	for _, n := range A {
		h[n] = true
	}

	return len(h)
}

// https://app.codility.com/demo/results/trainingFW7JRE-ZEX/
func MaxProductOfThree(A []int) int {
	N := len(A)
	a, b, c := -1<<31, -1<<31, -1<<31
	d, e := 1<<31-1, 1<<31-1

	for i := 0; i < N; i++ {
		if A[i] > a {
			c = b
			b = a
			a = A[i]
		} else if A[i] > b {
			c = b
			b = A[i]
		} else if A[i] > c {
			c = A[i]
		}

		if A[i] < d {
			e = d
			d = A[i]
		} else if A[i] < e {
			e = A[i]
		}

	}
	return int(math.Max(float64(a)*float64(b)*float64(c), float64(a)*float64(d)*float64(e)))
}

// https://app.codility.com/demo/results/trainingBHJWDB-3YE/
func MinAvgTwoSlice(A []int) int {
	N := len(A)
	minAvg := float64(1<<32 - 1)
	r := 0
	var avg float64
	for i := 0; i < N-1; i++ {
		if i+1 < N {
			avg = (float64(A[i]) + float64(A[i+1])) / 2
			if avg < minAvg {
				minAvg = avg
				r = i
			}
		}
		if i+2 < N {
			avg = (float64(A[i]) + float64(A[i+1]) + float64(A[i+2])) / 3
			if avg < minAvg {
				minAvg = avg
				r = i
			}
		}
	}
	return r
}

// https://app.codility.com/demo/results/training6R6GDV-9CS/
func GenomicRangeQuery(S string, P []int, Q []int) []int {
	sLn := len(S) + 1
	pSum := make(map[rune][]int, 3)
	pSum['A'] = make([]int, sLn)
	pSum['C'] = make([]int, sLn)
	pSum['G'] = make([]int, sLn)
	pSum['T'] = make([]int, sLn)
	for i := 1; i < sLn; i++ {
		for x := range pSum {
			pSum[x][i] = pSum[x][i-1]
		}
		pSum[rune(S[i-1])][i]++
	}
	pLn := len(P)
	r := make([]int, pLn)
	for i := 0; i < pLn; i++ {
		if pSum['A'][Q[i]+1]-pSum['A'][P[i]] > 0 {
			r[i] = 1
		} else if pSum['C'][Q[i]+1]-pSum['C'][P[i]] > 0 {
			r[i] = 2
		} else if pSum['G'][Q[i]+1]-pSum['G'][P[i]] > 0 {
			r[i] = 3
		} else {
			r[i] = 4
		}
	}
	return r
}

// https://app.codility.com/demo/results/trainingNDRTN3-Q5Q/
func PassingCars(A []int) int {
	N := len(A)
	P := make([]int, N)

	for i := N - 2; i > -1; i-- {
		P[i] = P[i+1] + A[i+1]
	}
	pairs := 0
	for i, v := range A {
		if v == 0 {
			pairs += P[i]
			if pairs > 1000000000 {
				return -1
			}
		}
	}

	return pairs
}

// https://app.codility.com/demo/results/trainingEARXAJ-4P7/
func MissingInteger(A []int) int {
	r := make(map[int]bool)
	for _, i := range A {
		if i > 0 {
			r[i] = true
		}
	}
	x := 1
	for ; x < len(r)+1; x++ {
		if _, ok := r[x]; !ok {
			return x
		}
	}
	return x
}

// https://app.codility.com/demo/results/training74S4ZC-NWJ/
func MaxCounters(N int, A []int) []int {
	var X, maxCounter, currentMax int
	r := make([]int, N)
	for i := 0; i < len(A); i++ {
		X = A[i]
		if X >= 1 && X <= N {
			if r[X-1] < maxCounter {
				r[X-1] = maxCounter
			}
			r[X-1]++
			if r[X-1] > currentMax {
				currentMax = r[X-1]
			}
		} else if X == N+1 {
			maxCounter = currentMax
		}
	}
	for j := 0; j < N; j++ {
		if r[j] < maxCounter {
			r[j] = maxCounter
		}
	}
	return r
}

// https://app.codility.com/demo/results/training8TSYUJ-JVJ/
func FrogRiverOne(X int, A []int) int {
	cache := make(map[int]bool, X)
	sum := 0
	expectedSum := X * (X + 1) / 2

	for i := 0; i < len(A); i++ {
		if !cache[A[i]] {
			sum += A[i]
			cache[A[i]] = true
		}
		if sum == expectedSum {
			return i
		}
	}

	return -1
}

// https://app.codility.com/demo/results/trainingZEHG8M-R7J/
func PermCheck(A []int) int {
	sort.Ints(A)
	if len(A) == 0 || A[0] != 1 {
		return 0
	}
	for i := 1; i < len(A); i++ {
		if A[i-1]+1 != A[i] {
			return 0
		}
	}
	return 1
}

// https://app.codility.com/demo/results/trainingUDJUH9-EB8/
func OddOccurrencesInArray(A []int) int {
	sort.Ints(A)
	nums := []int{0}
	for _, n := range A {
		if n == nums[len(nums)-1] {
			nums = nums[:len(nums)-1]
			continue
		}
		nums = append(nums, n)
	}
	return nums[len(nums)-1]
}

// binary gap https://app.codility.com/demo/results/trainingFR8VWJ-9NP/
func BinaryGap(N int) int {
	r := fmt.Sprintf("%b", N)
	gabs := []int{0}
	gab := 0
	for i := 0; i < len(r); i++ {
		if r[i] == '0' {
			gab++
		} else {
			gabs = append(gabs, gab)
			gab = 0
		}
	}
	sort.Ints(gabs)
	return gabs[len(gabs)-1]
}

// https://app.codility.com/demo/results/trainingPKC6RZ-DYW/
func TapeEquilibrium(A []int) int {
	if len(A) == 0 {
		return 0
	}
	min := 1 << 32
	totalSum := 0
	for _, n := range A {
		totalSum += n
	}
	l := 0
	r := totalSum
	for i := 0; i < len(A)-1; i++ {
		l += A[i]
		r -= A[i]
		m := int(math.Abs(float64(l) - float64(r)))
		if min > m {
			min = m
		}

	}
	return min
}

// https://app.codility.com/demo/results/trainingK5BG56-8BU/
func PermMissingElem(A []int) int {
	sort.Ints(A)
	if len(A) == 0 || A[0] != 1 {
		return 1
	}
	var missingNum int
	for i := 1; i < len(A); i++ {
		missingNum = A[i-1] + 1
		if missingNum != A[i] {
			return missingNum
		}
	}
	return A[len(A)-1] + 1
}

// https://app.codility.com/demo/results/trainingHVZF8P-YMQ/
func FrogJmp(X int, Y int, D int) int {
	return int(math.Ceil(float64(float64(Y-X) / float64(D))))
}

// https://app.codility.com/demo/results/training7CXZFN-GM2/
func CyclicRotation(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	var x []int
	n := len(A) - 1
	for i := 0; i < K; i++ {
		x = []int{A[n]}
		A = A[:n]
		A = append(x, A...)
	}
	return A
}

//
func MaxDoubleSliceSum(A []int) int {
	maxSliceSum := 0
	N := len(A) - 1
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				m := sum(A, i, j, k)
				if maxSliceSum < m {
					maxSliceSum = m
				}
			}
		}
	}

	return maxSliceSum
}

func sum(A []int, x, y, z int) int {
	v := 0
	for i := x + 1; i <= y-1; i++ {
		v += A[i]
	}
	for i := y + 1; i <= z-1; i++ {
		v += A[i]
	}
	return v
}

func minInSlice(v []int) int {
	sort.Ints(v)
	return v[0]
}
