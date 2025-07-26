//两数之和
//这里可以巧妙用map的方式实现，一下做法可优化
func twoSum(nums []int, target int) (res []int) {
    res = make([]int,2)
    for i:=0;i<len(nums)-1;i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i] + nums[j] == target{
                res[0], res[1] = i, j
                return 
            }
        }
    }
    return 
}

//合并区间
//此题关键在于条件排序
func merge(intervals [][]int) [][]int {
    res := make([][]int, 1)
    slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] })
    count := 0
    res[count] = intervals[0]
    for i:=1;i<len(intervals);i++{
        if intervals[i][0] <= res[count][1]{
            res[count][1] = max(intervals[i][1], res[count][1])
        }else{
            res = append(res, intervals[i])
            count++
        }
    }
    return res
}

//删除有序数组中的重复项
func removeDuplicates(nums []int) int {
    i := 0
    j := 1
    for {
        if j>=len(nums){
            break
        }
        if nums[i] == nums[j]{
            nums = append(nums[:j], nums[j+1:]...)
            
        }else{
            i = j
            j++
        }
    }
    return len(nums)
}

//加一
func plusOne(digits []int) []int {
    if digits[len(digits)-1] != 9{
        digits[len(digits) - 1] += 1
        return digits
    }
    digits[len(digits) - 1] = 0
    inc := true
    for i:=len(digits)-2; i>=0; i--{
        if digits[i] == 9 && inc {
            digits[i] = 0
        }else{
            digits[i] += 1
            inc = false
            break
        }
    }
    if inc{
        tlist := []int{1}
        digits = append(tlist, digits...)
    }
    return digits

}

//最长公共前缀
//此题关键在于第一个字符开始匹配，并不关心中间字符串是否匹配
func longestCommonPrefix(strs []string) string {
    for i, c := range strs[0]{
        for _, s := range strs[1:]{
            if i == len(s) || s[i] != byte(c){
                return strs[0][:i]
            }
        }
    }
    return strs[0]
}

//有效的括号
//此题关键在于栈的理解，字符压入栈之后得从最后一个字符匹配
func isValid(s string) bool {

    rs := []rune(s)
    ts := []rune{}
    if len(rs) % 2 != 0{
        return false
    }
    for _, v := range rs{
        if v == '{' || v == '(' || v == '['{
            ts = append(ts, v)
        }else if v == '}'{
            if len(ts) == 0 || ts[len(ts) -1] != '{' {
                return false
            }
            ts = ts[:len(ts) -1]
        }else if v == ')'{
            if len(ts) == 0 || ts[len(ts) -1] != '('{
                return false
            }
           ts = ts[:len(ts) -1]
        }else if v == ']'{
            if len(ts) == 0 || ts[len(ts) -1] != '['{
                return false
            }
            ts = ts[:len(ts) -1]
        }
    }
    if len(ts) != 0 {
        return false
    }
    return true
}

//回文数
func isPalindrome(x int) bool {
    if x < 0{
        return false
    }
    nlist := make([]int,0)
    for {
        y := x % 10
        nlist = append(nlist, y)
        x /= 10
        if x == 0{
            break
        }
    }
    for i,v := range nlist{
        if v != nlist[len(nlist) - (i+1)] {
            return false
        }
    }
    return true
}

//只出现一次的数字
func singleNumber(nums []int) int {
    d := make(map[int]int,0)
    for _,n := range nums{
        d[n]++
    }
    for k,v := range d{
        if v == 1{
            return k
        }
    }
    return 0
}
