// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package patterns

// Matches Searches for a pattern in a given text using the Knuth-Morris-Pratt algorithm.
// It is used to avoid a brute-force search for a pattern in a text.
//
// Time complexity:  O(m + n), where m is the length of the text and n is the length of the pattern.
// This is efficient because the algorithm preprocesses the pattern in O(n) and then searches the text in O(m).
// Space Complexity: O(n) for the LPS array.
//
// Parameters:
// - text: the text to search in
// - pattern: the pattern to search for
//
// Returns:
// - []int: the positions of the pattern in the text
func Matches(text string, pattern string) []int {
	lps := clps(pattern)
	var result []int
	var i, j int

	for i < len(text) {
		// While the algorithm advances in equalities, it increases the LPS pointer by 1.
		if pattern[j] == text[i] {
			i, j = i+1, j+1
		}

		// if j reaches the end of the pattern, then we have found a match
		if j == len(pattern) {
			result = append(result, i-j) // Match found, store index
			j = lps[j-1]                 // Reset j using lps
		} else if i < len(text) && pattern[j] != text[i] { // If the character in the text doesn't match the pattern and the
			// text has not reached the end, j is moved to the value on the previous lps array if it isn't a start of the pattern.
			// this is a smart move because it will check it again against the next character from the first pattern ocurrency
			// which can be different from the current one. This way that ambiguity is checked and the algorithm can
			// move along the text linearly.
			if j != 0 {
				j = lps[j-1] // Reset j using lps
			} else {
				i++ // Move to the next character in text
			}
		}
	}

	return result
}

// clps Computes the Longest-Prefix-Suffixes and returns its P(i) table
// A longest prefix suffix (LPS) array is an array that contains the length of the longest prefix that's also a suffix
// on each start -> (current loop) index of the pattern.
// That means it returns an array pointing on each one of its indexes the length of the longest prefix that's also a suffix
// from the beginning of the pattern until the current index it's traversing.
//
// Ex: a | b | c | d | a | b | e | a | b | f |
// Pi: 0 | 0 | 0 | 0 | 1 | 2 | 0 | 1 | 2 | 0 |
func clps(pattern string) []int {
	// Example pattern: "ababababba"
	lps := make([]int, len(pattern)) // lps[i] stores the longest-prefix-suffix length of pattern[0:i]
	var length int                   // length of the previous longest prefix suffix
	i := 1

	// the for isn't a range because there's no need to compare the first character with itself because
	// there’s no possible prefix-suffix pair in a single character.
	for i < len(pattern) {
		// If the characters at pattern[i] and pattern[length] match,
		// increase the length and add its value to lps
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else { // if the characters are different, reset the length and lps
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return lps
}
