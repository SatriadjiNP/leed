package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// soal no 1 = easy
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 1 + i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// soal no 2 = easy
func containsDuplicate(nums []int) bool {
	visited := make(map[int]bool)
	for _, num := range nums {
		if visited[num] {
			return true
		}
		visited[num] = true
	}
	return false
}

// soal no 3 = easy
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !startsWith(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

// func pemanggilan longestCommonPrefix
func startsWith(str, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}

	for i := 0; i < len(prefix); i++ {
		if str[i] != prefix[i] {
			return false
		}
	}

	return true
}

// soal no 4 = easy
func maxProfit(prices []int) int {
	minPrice := math.MaxInt32
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

// soal no 5 = easy
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	i, j := 0, len(str)-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}

// soal no 6 = easy
func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else if char == ')' || char == ']' || char == '}' {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if (char == ')' && last != '(') || (char == ']' && last != '[') || (char == '}' && last != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

// soal no 7 = easy
func romanToInt(s string) int {
	var romanValues = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	prevValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		value := romanValues[s[i]]

		if value >= prevValue {
			result += value
		} else {
			result -= value
		}

		prevValue = value
	}

	return result
}

// soal no 8 = easy
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChars := []rune(s)
	tChars := []rune(t)

	sort.Slice(sChars, func(i, j int) bool {
		return sChars[i] < sChars[j]
	})

	sort.Slice(tChars, func(i, j int) bool {
		return tChars[i] < tChars[j]
	})

	for i := 0; i < len(sChars); i++ {
		if sChars[i] != tChars[i] {
			return false
		}
	}

	return true
}

// soal no 9 = easy
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}

	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]
		if (y1-y0)*(x-x0) != (y-y0)*(x1-x0) {
			return false
		}
	}

	return true
}

// soal no 10 = easy
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Indeks berikutnya untuk elemen unik
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// soal no 11 = hard
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeArrays(nums1, nums2)
	length := len(merged)

	if length%2 == 0 {
		// Jika jumlah elemen genap, ambil rata-rata dari dua elemen di tengah
		mid := length / 2
		return float64(merged[mid-1]+merged[mid]) / 2.0
	} else {
		// Jika jumlah elemen ganjil, ambil elemen di tengah
		mid := length / 2
		return float64(merged[mid])
	}
}

func mergeArrays(nums1 []int, nums2 []int) []int {
	merged := append(nums1, nums2...)
	sort.Ints(merged)
	return merged
}

// soal no 12 = medium
// Definisikan struktur Node
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Definisikan struktur Node
	dummy := &ListNode{} // Node dummy sebagai awal dari hasil penjumlahan
	curr := dummy        // Node saat ini untuk traversal
	carry := 0           // Penyimpanan carry

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next // Mengembalikan node setelah dummy sebagai hasil penjumlahan

}

// soal no 13 = medium
func lengthOfLongestSubstring(s string) int {
	// Membuat map untuk melacak indeks terakhir dari setiap karakter
	charMap := make(map[byte]int)

	start := 0     // Indeks awal substring non-berulang saat ini
	maxLength := 0 // Panjang maksimum substring non-berulang

	for end := 0; end < len(s); end++ {
		// Periksa apakah karakter saat ini sudah ada dalam substring saat ini
		if idx, ok := charMap[s[end]]; ok && idx >= start {
			// Jika ditemukan karakter yang berulang, perbarui start ke indeks berikutnya
			start = idx + 1
		}

		// Perbarui atau tambahkan indeks karakter saat ini ke charMap
		charMap[s[end]] = end

		// Periksa apakah panjang substring saat ini lebih panjang dari maxLength
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

// soal no 14 = medium
func topKFrequent(nums []int, k int) []int {
	// Membuat map untuk menghitung frekuensi setiap elemen
	freqMap := make(map[int]int)

	// Menghitung frekuensi setiap elemen
	for _, num := range nums {
		freqMap[num]++
	}

	// Membuat slice untuk menyimpan k elemen unik dengan frekuensi terbanyak
	topK := make([]int, 0, k)

	// Mengambil k elemen unik dengan frekuensi terbanyak
	for num := range freqMap {
		topK = append(topK, num)
	}

	// Mengurutkan slice berdasarkan frekuensi secara menurun
	sort.Slice(topK, func(i, j int) bool {
		return freqMap[topK[i]] > freqMap[topK[j]]
	})

	// Mengembalikan k elemen unik dengan frekuensi terbanyak
	return topK[:k]
}

// soal no 15 = medium
func threeSum(nums []int) [][]int {
	result := [][]int{}
	n := len(nums)

	// Mengurutkan slice secara ascending
	sort.Ints(nums)

	// Iterasi pertama untuk memilih elemen pertama dari triplet
	for i := 0; i < n-2; i++ {
		// Menghindari duplikasi elemen pertama
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Menyisakan dua pointer di bagian kanan dari elemen pertama
		left := i + 1
		right := n - 1

		// Iterasi kedua untuk mencari pasangan angka yang jumlahnya sama dengan -nums[i]
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// Menambahkan triplet ke dalam hasil
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Menghindari duplikasi elemen kedua dan ketiga
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Memindahkan pointer ke elemen berikutnya
				left++
				right--
			} else if sum < 0 {
				// Jika jumlah kurang dari 0, kita perlu meningkatkan jumlah dengan memindahkan pointer left
				left++
			} else {
				// Jika jumlah lebih dari 0, kita perlu mengurangi jumlah dengan memindahkan pointer right
				right--
			}
		}
	}

	return result
}

// soal no 16 = medium
func reverse(x int) int {
	result := 0

	for x != 0 {
		// Mendapatkan digit terakhir dari x
		digit := x % 10

		// Mengecek batasan overflow dan underflow
		if result > (1<<31-1)/10 || (result == (1<<31-1)/10 && digit > 7) {
			return 0
		}
		if result < -(1<<31)/10 || (result == -(1<<31)/10 && digit < -8) {
			return 0
		}

		// Memperbarui hasil dengan menambahkan digit terakhir
		result = result*10 + digit

		// Mengurangi digit terakhir dari x
		x /= 10
	}

	return result
}

// soal no 17 = easy
func isPalindromeString(s string) bool {
	// Mengubah string menjadi lowercase dan menghapus karakter non-alphanumeric
	s = normalizeString(s)

	// Memeriksa apakah string adalah palindrom
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}

	return true
}

func normalizeString(s string) string {
	// Mengubah string menjadi lowercase
	s = strings.ToLower(s)

	// Menghapus karakter non-alphanumeric
	var normalized strings.Builder
	for _, ch := range s {
		if isAlphanumeric(ch) {
			normalized.WriteRune(ch)
		}
	}

	return normalized.String()
}

func isAlphanumeric(ch rune) bool {
	return ('a' <= ch && ch <= 'z') || ('0' <= ch && ch <= '9')
}

// soal no 18 = easy
func merge(nums1 []int, m int, nums2 []int, n int) {
	// Menggunakan dua pointer untuk melacak posisi saat penggabungan
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// Melakukan penggabungan dari belakang
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// Jika masih ada elemen yang tersisa di nums2, menyalinnya ke nums1
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}

// soal no 19 = easy
func removeElement(nums []int, val int) int {
	k := 0 // variabel untuk menghitung elemen yang tidak sama dengan val

	// Mengiterasi semua elemen dalam array nums
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}

	return k
}

// soal no 20 = easy
func mergeTwoLists(list1 []int, list2 []int) []int {
	merged := make([]int, 0)
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			merged = append(merged, list1[i])
			i++
		} else {
			merged = append(merged, list2[j])
			j++
		}
	}

	for i < len(list1) {
		merged = append(merged, list1[i])
		i++
	}

	for j < len(list2) {
		merged = append(merged, list2[j])
		j++
	}

	return merged
}

func main() {
	fmt.Print("No. 1 : ")
	fmt.Println(twoSum([]int{2, 4, 5, 3, 7}, 11))

	fmt.Print("No. 2 : ")
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

	fmt.Print("No. 3 : ")
	// strs := []string{"flower", "flow", "flight"}
	// commonPrefix := longestCommonPrefix(strs)
	// fmt.Println(commonPrefix)
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))

	fmt.Print("No. 4 : ")
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))

	fmt.Print("No. 5 : ")
	fmt.Println(isPalindrome(121))

	fmt.Print("No. 6 : ")
	fmt.Println(isValid("()[]{}"))

	fmt.Print("No. 7 : ")
	fmt.Println(romanToInt("MCMXCIV"))

	fmt.Print("No. 8 : ")
	s := "anagram"
	t := "nagaram"
	isAnagram := isAnagram(s, t)
	fmt.Println(isAnagram)

	fmt.Print("No. 9 : ")
	coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	isStraightLine := checkStraightLine(coordinates)
	fmt.Println(isStraightLine)

	fmt.Println("No. 10 : ")
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println("k =", k)
	fmt.Println("nums =", nums[:k])

	fmt.Print("No. 11 : ")
	nums1 := []int{1, 3}
	nums2 := []int{2}
	median := findMedianSortedArrays(nums1, nums2)
	fmt.Printf("%.5f\n", median)

	fmt.Print("No. 12 : ")
	// Membentuk linked list l1: 2 -> 4 -> 3
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	// Membentuk linked list l2: 5 -> 6 -> 4
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	// Memanggil fungsi addTwoNumbers untuk menjumlahkan kedua linked list
	result := addTwoNumbers(l1, l2)

	// Mencetak hasil penjumlahan sebagai linked list
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}
	fmt.Println("nil")

	fmt.Print("No. 13 : ")
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))

	fmt.Print("No. 14 : ")
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))

	fmt.Print("No. 15 : ")
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

	fmt.Print("No. 16 : ")
	fmt.Println(reverse(-3210))

	fmt.Print("No. 17 : ")
	fmt.Println(isPalindromeString("A man, a plan, a canal: Panama"))

	fmt.Print("No. 18 : ")
	nums1 = []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 = []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	fmt.Println("No. 19 : ")
	nums = []int{3, 2, 2, 3}
	val := 3

	k = removeElement(nums, val)
	fmt.Println("k =", k)
	fmt.Println("nums =", nums[:k])

	fmt.Print("No. 20 : ")
	list1 := []int{1, 2, 4}
	list2 := []int{1, 3, 4}

	merged := mergeTwoLists(list1, list2)
	fmt.Println(merged)

}
