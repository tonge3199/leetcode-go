package recursive

// var letterMap = [10]string{
// 	"",     // 0
// 	"",     // 1
// 	"abc",  // 2
// 	"def",  // 3
// 	"ghi",  // 4
// 	"jkl",  // 5
// 	"mno",  // 6
// 	"pqrs", // 7
// 	"tuv",  // 8
// 	"wxyz", // 9
// }

var letterMap = map[byte]string{
	'0': "",
	'1': "",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var result []string
	backtrack(digits, 0, "", &result)
	return result
}

func backtrack(digits string, index int, current string, result *[]string) {
	// need *[]string, append([]string) assign a local array and will be discard in top level.
	// last level return
	if index == len(digits) {
		*result = append(*result, current)
		return
	}

	digit := digits[index]
	letters := letterMap[digit]

	for _, r := range letters {
		letter := string(r)
		current += letter // added letter: "" + "a", "a"+"b".
		backtrack(digits, index+1, current, result)
		current = current[:len(current)-1] // backtrack to prevState: "abc" -> "ab",
	}

}
