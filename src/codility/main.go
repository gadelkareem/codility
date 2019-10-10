package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"path/filepath"
	"sort"
	"strings"
	"time"
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
	//fmt.Printf("%+v\n", NumberOfDiscIntersections([]int{1, 5, 2, 1, 4, 0}))
	//fmt.Printf("%+v\n", Brackets("([)()]"))
	//fmt.Printf("%+v\n", Nesting("())"))
	//fmt.Printf("%+v\n", Dominator([]int{3, 4, 3, 2, 3, -1, 3, 3}))
	//fmt.Printf("%+v\n", Solution([]int{9, 1, 4, 9, 0, 4, 8, 9, 0, 1}))
	fmt.Printf("%+v\n", f([]string{"a", "z", "c", "t", "v", "a"}, "cat"))
}

func Solution(T []int) []int {
	l := -1
	h := make(map[int][]int)
	for i, n := range T {
		if n == i {
			l = i
		} else {
			h[n] = append(h[n], i)
		}
	}
	var counts []int
	done := false
	for range T {
		if _, ok := h[l]; !ok || len(h[l]) == 0 {
			break
		}
		j := len(h[l])
		if done {
			j = 0
		}
		counts = append(counts, j)
		done = true
		for _, x := range h[l] {
			if _, ok := h[x]; ok && x != l {
				h[l] = nil
				l = x
				done = false
				break
			}
		}
	}

	return counts
}

var t = `photo.jpg, Warsaw, 2013-09-05 14:08:15
john.png, London, 2015-06-20 15:13:22
myFriends.png, Warsaw, 2013-09-05 14:07:13
Eiffel.jpg, Paris, 2015-07-23 08:03:02
pisatower.jpg, Paris, 2015-07-22 23:59:59
BOB.jpg, London, 2015-08-05 00:02:03
notredame.png, Paris, 2015-09-01 12:00:00
me.jpg, Warsaw, 2013-09-06 15:40:22
a.png, Warsaw, 2016-02-13 13:33:50
b.jpg, Warsaw, 2016-01-02 15:12:22
c.jpg, Warsaw, 2016-01-02 14:34:30
d.jpg, Warsaw, 2016-01-02 15:15:01
e.png, Warsaw, 2016-01-02 09:49:09
f.png, Warsaw, 2016-01-02 10:55:32
g.jpg, Warsaw, 2016-02-29 22:13:11`

func sol(S string) string {
	r := csv.NewReader(strings.NewReader(S))
	// Read all records.
	lines, _ := r.ReadAll()
	//count locations
	loc := make(map[string]map[float64]int)
	dates := make([]float64, len(lines))
	for i, l := range lines {
		//count locations
		c := strings.TrimSpace(l[1])
		if _, ok := loc[c]; !ok {
			loc[c] = make(map[float64]int)
		}
		t, err := time.Parse("2006-01-02 15:04:05", strings.TrimSpace(l[2]))
		if err != nil {
			log.Fatal(err)
			continue
		}
		tu := t.Unix()
		loc[c][float64(tu)] = i
		dates = append(dates, float64(tu))
	}
	//sort dates
	sort.Float64s(dates)
	//generate names
	names := make([]string, len(lines))
	for c, n := range loc {
		ln := len(n)
		format := "%s%0d"
		numOfZ := len(fmt.Sprintf("%d", ln))
		if numOfZ > 1 {
			format = "%s%0" + fmt.Sprintf("%d", numOfZ) + "d"
		}
		j := 1
		for _, tu := range dates {
			if _, ok := n[tu]; !ok {
				continue
			}
			names[n[tu]] = fmt.Sprintf(format, c, j)
			j++
		}
	}
	var list []string
	for i, l := range lines {
		list = append(list, fmt.Sprintf("%s%s", names[i], filepath.Ext(l[0])))
	}
	return strings.Join(list, "\n")
}

// https://app.codility.com/demo/results/training3666ES-AR2/
func Dominator(A []int) int {
	halfLn := int(math.Ceil(float64(len(A) / 2)))
	h := make(map[int][]int)
	for i, n := range A {
		h[n] = append(h[n], i)
	}
	for _, r := range h {
		if len(r) > halfLn {
			return r[0]
		}
	}
	return -1
}

// https://app.codility.com/demo/results/training2XGHZ9-CHS/
func Nesting(S string) int {
	if S == "" {
		return 1
	}
	o := "("
	c := ")"
	var h []string
	m := map[string]string{"(": ")"}
	for i := 0; i < len(S); i++ {
		N := len(h)
		chr := string(S[i])
		if chr == o {
			h = append(h, chr)
		} else if chr == c {
			if N == 0 || chr != m[h[N-1]] {
				return 0
			}
			h = h[:N-1]
		}
	}
	if len(h) == 0 {
		return 1
	}
	return 0
}

// https://app.codility.com/demo/results/trainingUUMP9F-8JU/
func Brackets(S string) int {
	if S == "" {
		return 1
	}
	o := "{[("
	c := "}])"
	var h []string
	m := map[string]string{"[": "]", "{": "}", "(": ")"}
	for i := 0; i < len(S); i++ {
		N := len(h)
		chr := string(S[i])
		if strings.Contains(o, chr) {
			h = append(h, chr)
		} else if strings.Contains(c, chr) {
			if N == 0 || chr != m[h[N-1]] {
				return 0
			}
			h = h[:N-1]
		}
	}
	if len(h) == 0 {
		return 1
	}
	return 0
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

/*

example word:  "cat"
board:  [ 'a', 'z', 'c', 't', 'v', 'a' ]

best order (move and the state of the board after the move):

starting board: [ 'a', 'z', 'c', 't', 'v', 'a' ]

MOVE:         RESULTING BOARD (after the move):

seeking "c" (closer to the left, so take from the left) [ 'a', 'z', 'c', 't', 'v', 'a' ]

LEFT       => [ 'z', 'c', 't', 'v', 'a', 'a' ]
LEFT       => [ 'c', 't', 'v', 'a', 'a', 'z' ]
LEFT  "c"  => [ 't', 'v', 'a', 'a', 'z' ]       "c" was removed from the board

seeking "a" (closer to the right, so take from right)  [ 't', 'v', 'a', 'a', 'z' ]

RIGHT      => [ 'z', 't', 'v', 'a', 'a' ]
RIGHT "a"  => [ 'z', 't', 'v', 'a' ]            "a" was removed from the board

seeking "t" (closer to the left, so take from the left)

LEFT       => [ 't', 'v', 'a', 'z' ]
LEFT  "t"  => [ 'v', 'a', 'z' ]                 "t" was removed from the board, all done!

f(board, word) => moves

Result: data structure

LEFT:nil, LEFT:nil, LEFT:c, RIGHT:nil, RIGHT:a, LEFT:nil, LEFT:t

2nd example word: "tv"
2nd board:  [ 'a', 'z', 'c', 't', 'v', 'a' ]

Result:

RIGHT:nil, RIGHT:nil, RIGHT:t, LEFT:v


1.  STRATEGY: start from the side closest to the letter, and favor left on a tie.
2.  you only take from the left or the right
3.  unused letters go back on the other side (rotate the board)
4.  used letter is removed from the board

*/
func f(b []string, word string) (mv [][]string) {

	var dir string
	var x int

start:
	for _, s := range word {
		i := indexOf(string(s), b)
		if i > len(b)/2-1 {
			dir = "RIGHT"
			x = len(b) - 1
		} else {
			dir = "LEFT"
			x = 0
		}
		for {
			l := len(b)
			if x < 0 || x > l {
				break
			}
			lt := b[x]
			if dir == "LEFT" {
				b = b[x+1:]
			} else {
				b = b[:x]
			}
			if lt == string(s) {
				mv = append(mv, []string{dir, lt})
				continue start
			}

			mv = append(mv, []string{dir, ""})
			if dir == "LEFT" {
				b = append(b, lt)
			} else {
				b = append([]string{lt}, b...)
			}

		}

	}

	return mv
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
