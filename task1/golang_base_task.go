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
