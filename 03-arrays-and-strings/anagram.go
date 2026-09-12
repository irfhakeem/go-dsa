package arraysandstrings

import "slices"

func IsAnagram(s1, s2 string) bool {
	runesS1 := []rune(s1)
	runesS2 := []rune(s2)
	n := len(runesS1)
	m := len(runesS2)

	if n != m {
		return false
	}

	// sorting
	// Time Complexity: O(n log n)
	// Karena kita melakukan sorting dimana algoritma sorting yang digunakan oleh Go adalah hybrid dari QuickSort dan HeapSort O(n log n).

	slices.Sort(runesS1)
	slices.Sort(runesS2)

	return slices.Equal(runesS1, runesS2)

	// frequency counting
	// Time Complexity: O(n)
	// Karena kita melakukan iterasi dan menyimpan frekuensi setiap karakter string (n) dalam map.

	mapS1 := make(map[rune]int)
	mapS2 := make(map[rune]int)

	for i := range n {
		mapS1[runesS1[i]]++
		mapS2[runesS2[i]]++
	}

	for r, count := range mapS1 {
		_, exists := mapS2[r]
		if !exists {
			return false
		}

		if count != mapS2[r] {
			return false
		}
	}

	return true

}
