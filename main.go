package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	mapValidation()
}

func mapValidation() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		in.ReadLine()
		rows := make([][]rune, n)
		for j := 0; j < n; j++ {
			rows[j] = make([]rune, m)
			for l := 0; l < m; l++ {
				rows[j][l], _, _ = in.ReadRune()
			}
			in.ReadLine()
		}
		processMap(rows, out)
	}
}

func processMap(arr [][]rune, out *bufio.Writer) {

}

func mountainsFromASCII() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var num, n, m int
		fmt.Fscan(in, &num, &n, &m)
		in.ReadLine()
		rows := make([][][]rune, num)
		for k := 0; k < num; k++ {
			rows[k] = make([][]rune, n)
			for j := 0; j < n; j++ {
				rows[k][j] = make([]rune, m)
				for l := 0; l < m; l++ {
					rows[k][j][l], _, _ = in.ReadRune()
				}
				in.ReadLine()
			}
			if k < num-1 {
				in.ReadLine()
			}
		}
		processMountains(rows, out)
		if i < count-1 {
			fmt.Fprintln(out, "")
		}
	}
}

func processMountains(rows [][][]rune, out *bufio.Writer) {
	countOfPictures := len(rows)
	heightSize := len(rows[0])
	widthSize := len(rows[0][0])

	for i := 0; i < widthSize; i++ {
		for j := heightSize - 1; j >= 0; j-- {
			if rows[0][j][i] == '.' {
				diffFound := false
				for k := 1; k < countOfPictures; k++ {
					if rows[k][j][i] != '.' {
						rows[0][j][i] = rows[k][j][i]
						diffFound = true
						break
					}
				}
				if !diffFound {
					break
				}
			}
		}
	}

	for i := 0; i < heightSize; i++ {
		for j := 0; j < widthSize; j++ {
			fmt.Fprint(out, string(rows[0][i][j]))
		}
		fmt.Fprintln(out, "")
	}
}

func robotRace() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var n, m int
		fmt.Fscan(in, &n, &m)
		in.ReadLine()
		rows := make([][]rune, n)
		for j := 0; j < n; j++ {
			rows[j] = make([]rune, m)
			for k := 0; k < m; k++ {
				rows[j][k], _, _ = in.ReadRune()
			}
			in.ReadLine()
		}
		//fmt.Fprintln(out, rows)
		processRobots(rows, out)
	}
}

func processRobots(rows [][]rune, out *bufio.Writer) {

	if rows[0][0] != '.' && rows[len(rows)-1][len(rows[0])-1] != '.' {
		for _, row := range rows {
			for _, element := range row {
				fmt.Fprint(out, string(element))
			}
			fmt.Fprintln(out, "")
		}
		return
	}
	robotsCount := 0

mainLoop:
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[0]); j++ {
			if rows[i][j] != '.' && rows[i][j] != '#' {
				robotsCount++
				drawRobotWay(rows, i, j, robotsCount)
				if robotsCount == 2 {
					break mainLoop
				}
			}
		}
	}

	for _, row := range rows {
		for _, element := range row {
			fmt.Fprint(out, string(element))
		}
		fmt.Fprintln(out, "")
	}
}

func drawRobotWay(arr [][]rune, i int, j int, count int) {
	runeToDraw := arr[i][j] + 32

	if (i == 0 && j == 0) || (i == len(arr)-1 && j == len(arr[0])-1) {
		return
	}

	if count == 1 {
		if i > 0 && arr[i-1][j] != '#' {
			for k := 0; k < i; k++ {
				arr[k][j] = runeToDraw
			}
		} else {
			j--
			for k := 0; k <= i; k++ {
				arr[k][j] = runeToDraw
			}
		}
		for k := 0; k < j; k++ {
			arr[0][k] = runeToDraw
		}
	} else {
		if i < len(arr)-1 && arr[i+1][j] != '#' {
			for k := i + 1; k < len(arr); k++ {
				arr[k][j] = runeToDraw
			}
		} else {
			j++
			for k := i; k < len(arr); k++ {
				arr[k][j] = runeToDraw
			}
		}
		for k := j + 1; k < len(arr[0]); k++ {
			arr[len(arr)-1][k] = runeToDraw
		}
	}
}

func whoIsDoSmth() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var rowsCount int
		fmt.Fscan(in, &rowsCount)
		in.ReadString('\n')
		rows := make([]string, rowsCount)
		for j := 0; j < rowsCount; j++ {
			tempStr, _ := in.ReadString('\n')
			rows[j] = strings.Replace(tempStr, "\n", "", -1)
		}
		processData(rows, out)
	}
}

func processData(rows []string, out *bufio.Writer) {
	namesMap := make(map[string]int, len(rows))
	results := make([]string, 0, len(rows))
	var activity string
	maxScore := math.MinInt

	for i := 0; i < len(rows); i++ {
		strData := strings.Split(rows[i], ": ")

		if _, ok := namesMap[strData[0]]; !ok {
			namesMap[strData[0]] = 0
		}

		words := strings.Split(strData[1], " ")
		nameFromWords := words[0]

		if len(words) > 3 {
			if words[1] == "am" {
				namesMap[strData[0]]--
			} else {
				namesMap[nameFromWords]--
			}
		} else {
			if words[1] == "am" {
				namesMap[strData[0]] += 2
			} else {
				namesMap[nameFromWords]++
			}
		}

		if activity == "" {
			activity = words[len(words)-1]
		}
	}

	for _, score := range namesMap {
		if score > maxScore {
			maxScore = score
		}
	}

	for name, score := range namesMap {
		if score == maxScore {
			results = append(results, name)
		}
	}

	if len(results) > 1 {
		sort.Strings(results)
	}

	for _, name := range results {
		fmt.Fprintln(out, name+" is "+activity[:len(activity)-1]+".")
	}
}

func longestValidParentheses(s string) int {
	stack := new(list.List)
	maxCount := 0
	counter := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.PushBack(i)
		} else if s[i] == ')' && stack.Len() > 0 {
			stack.Remove(stack.Back())
			counter += 2
			if stack.Len() == 0 && maxCount < counter {
				maxCount = counter
			}
		} else {
			if maxCount < counter {
				maxCount = counter
			}
			counter = 0
		}
	}

	if maxCount < counter-maxCount {
		maxCount = counter
	}

	return maxCount
}

func getPermutation(n int, k int) string {
	var divider, x int
	result := new(strings.Builder)
	nums := make(map[int]bool)
	countOfPerms := getCountOfPermutation(n)

	for i := 0; i < n; i++ {

		divider = countOfPerms / (n - i)
		x = (k + divider - 1) / divider

		if k == 0 {
			for j := n; j >= 1; j-- {
				if !nums[j] {
					result.WriteString(strconv.Itoa(j))
					nums[j] = true
				}
			}
			break
		} else {
			numsCounter := 0
			for j := 1; j <= n; j++ {
				if !nums[j] {
					numsCounter++
					if numsCounter == x {
						result.WriteString(strconv.Itoa(j))
						nums[j] = true
						break
					}
				}
			}
		}
		k = k % divider
		countOfPerms = getCountOfPermutation(n - i - 1)
	}

	return result.String()
}

func getCountOfPermutation(n int) int {
	countOfPerms := 1
	for i := 2; i <= n; i++ {
		countOfPerms *= i
	}

	return countOfPerms
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 || m == 1 || n == 1 {
		return n
	}

	result := make([][]int, n)

	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				result[i][j] = 1
			}
			if j > 0 {
				result[i][j] += result[i][j-1]
			}
			if i > 0 {
				result[i][j] += result[i-1][j]
			}
		}
	}

	return result[n-1][m-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 && (i != 0 || j != 0) {
				obstacleGrid[i][j] = 0
			} else {
				if i > 0 {
					obstacleGrid[i][j] += obstacleGrid[i-1][j]
				}
				if j > 0 {
					obstacleGrid[i][j] += obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func minPathSum(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i > 0 && j > 0 {
				grid[i][j] += min(grid[i-1][j], grid[i][j-1])
			} else {
				if i > 0 {
					grid[i][j] += grid[i-1][j]
				}
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	result := make([][]int, len(word1)+1)

	compareChars := func(i, j int) int {
		if word1[i-1] == word2[j-1] {
			return 0
		}
		return 1
	}

	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(word2)+1)
		for j := 0; j < len(result[i]); j++ {
			if j == 0 && i == 0 {
				result[i][j] = 0
			}
			if j > 0 && i == 0 {
				result[i][j] = j
			}
			if i > 0 && j == 0 {
				result[i][j] = i
			}
			if j > 0 && i > 0 {
				result[i][j] += min(result[i-1][j]+1, result[i][j-1]+1, result[i-1][j-1]+compareChars(i, j))
			}
		}
	}

	return result[len(word1)][len(word2)]

}

func jump(nums []int) int {
	lenNums := len(nums)
	if lenNums == 1 {
		return 0
	}
	dp := make([]int, lenNums)
	dp[0] = 0
	minJumps := math.MaxInt

	for i := 0; i < lenNums; i++ {
		for j := nums[i]; j >= 1; j-- {
			if i+j < lenNums-1 {
				if dp[i+j] == 0 {
					dp[i+j] += dp[i] + 1
				} else {
					dp[i+j] = min(dp[i+j], dp[i]+1)
				}
			}
			if i+j == lenNums-1 {
				minJumps = min(minJumps, dp[i]+1)
			}
		}

	}
	return minJumps
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	countOfDecs := 0
	countOfDuobleCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '0' && (i == 0 || s[i-1] == '0') {
			return 0
		}
		if (s[i] == '1' || s[i] == '2') && i < len(s)-1 && s[i+1] <= 54 {
			countOfDecs += 2
			countOfDuobleCount++
		} else if s[i] != '0' {
			countOfDecs++
		}

	}
	return countOfDecs - countOfDuobleCount
}

func firstTask() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c, d, result int

	fmt.Fscanln(in, &a, &b, &c, &d)

	if d > b {
		result = (d - b) * c
	}

	fmt.Fprintln(out, result+a)

}

func secondTask() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a int
	fmt.Fscanln(in, &a)
	if a == 1 {
		fmt.Fprintln(out, 0)
		return
	}
	res := 1
	count := 0

	for res < a {
		count++
		res *= 2
	}

	fmt.Fprintln(out, count)

}

func partitionString(s string) []string {
	seenSegments := make(map[string]bool, len(s))
	newSegment := new(strings.Builder)
	result := make([]string, 0, len(s))

	for _, item := range s {
		newSegment.WriteString(string(item))
		if _, ok := seenSegments[newSegment.String()]; !ok {
			strToAppend := newSegment.String()
			result = append(result, strToAppend)
			seenSegments[strToAppend] = true
			newSegment.Reset()
		}
	}

	if len(seenSegments) > 0 {
		result = append(result, newSegment.String())
	}

	return result[:len(result)-1]
}

func longestCommonPrefix(words []string) []int {
	result := make([]int, len(words))
	prefixes := make(map[int]string, len(words))
	maxLength := 0
	prevLength := 0

	for i := 0; i < len(words)-1; i++ {
		if result[i] != 0 && words[i] == words[i+1] {
			result[i+1] = result[i]
		} else {
			prefixLen, prefix := calcPrefix(words[i], words[i+1])
			if prefixLen > maxLength {
				prevLength = maxLength
				maxLength = prefixLen
			} else if prefixLen > prevLength {
				prevLength = prefixLen
			}
			if result[i] < prefixLen {
				result[i] = prefixLen
				prefixes[i] = prefix
			}
			result[i+1] = max(result[i+1], prefixLen)
			prefixes[i+1] = prefix
		}
	}

	for i := 0; i < len(result); i++ {
		if result[i] == maxLength {
			counter := 0
			for j := i; j < len(result) && result[j] == maxLength && prefixes[i] == prefixes[j]; j++ {
				counter++
			}
			switch counter {
			case 2:
				result[i] = prevLength
				result[i+1] = prevLength
			case 3:
				result[i+1] = prevLength
			}
			i += counter - 1
		} else if i > 0 && i < len(result)-1 && prefixes[i-1] == prefixes[i+1] {
			result[i] = result[i-1]
		} else {
			result[i] = maxLength
		}
	}

	return result
}

func calcPrefix(first string, second string) (int, string) {
	result := 0
	prefix := new(strings.Builder)

	for result < len(first) && result < len(second) && first[result] == second[result] {
		prefix.WriteByte(first[result])
		result++
	}

	return result, prefix.String()
}
