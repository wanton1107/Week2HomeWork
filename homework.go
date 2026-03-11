package homework01

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	dict := make(map[int]int)

	for _, v := range nums {
		if num, ok := dict[v]; ok {
			num++
			dict[v] = num
		} else {
			dict[v] = 1
		}
	}

	for k, v := range dict {
		if v == 1 {
			return k
		}
	}

	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	l := make([]int, 0, 10)
	for x != 0 {
		l = append(l, x%10)
		x /= 10
	}
	for i := 0; i < len(l)/2; i++ {
		if l[i] != l[len(l)-1-i] {
			return false
		}
	}
	return true
}

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	stack := make([]rune, 0, 10)
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			if v == ')' && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if v == ']' && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else if v == '}' && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	l := len(strs[0])
	mi := 0
	for i, v := range strs {
		if len(v) < l {
			l = len(v)
			mi = i
		}
	}
	maxPrefix := ""
	for i := 0; i < len(strs[mi]); i++ {
		c := strs[mi][i]
		for _, v := range strs {
			if v[i] != c {
				return maxPrefix
			}
		}
		maxPrefix += string(c)
	}
	return maxPrefix
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	flag := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + flag
		flag = digits[i] / 10
		digits[i] %= 10
	}
	if flag == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	for i := 1; i < len(intervals); {
		if intervals[i][0] <= intervals[i-1][1] {
			if intervals[i][1] > intervals[i-1][1] {
				intervals[i-1][1] = intervals[i][1]
			}
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			i++
		}
	}
	return intervals
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		temp := target - num
		if _, ok := m[temp]; ok {
			return []int{m[temp], i}
		} else {
			m[num] = i
		}
	}
	return nil
}
