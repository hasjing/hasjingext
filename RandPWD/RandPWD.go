// package RandPWD
// 产生随机的字符串或密码串
package RandPWD

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

const (
	NUmStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[](){}!%^*$"
)

// init  利用当前时间初始化随机种子
func init() {
	// 初始化随机种子
	rand.Seed(time.Now().UnixNano())
}

// GenerateRNDString 产生制定类型长度的随机字符串
//	string := GenerateRNDString("advance",256)  产生一个由数字字母符号组成的256个字符的串
//		num     数字
//		char    字母（含大小写)
//		mix     数字+字母
//		advance 数字+字母+符号
func GenerateRNDString(charset string, length int) string {
	//初始化密码切片
	var passwd []byte = make([]byte, length, length)
	//源字符串
	var sourceStr string

	//产生的字符串类型
	switch charset {
	case "num":
		sourceStr = NUmStr
	case "char":
		sourceStr = CharStr
	case "mix":
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
	case "advance":
		sourceStr = fmt.Sprintf("%s%s%s", NUmStr, CharStr, SpecStr)
	default:
		sourceStr = fmt.Sprintf("%s%s", NUmStr, CharStr)
	}

	// fmt.Println("source:", sourceStr)
	slen := len(sourceStr)
	//遍历，生成一个随机index索引,
	for i := 0; i < length; i++ {
		passwd[i] = sourceStr[rand.Intn(slen)]
	}
	return string(passwd)
}

//  GeneratePasswd  返回指定强度的密码串
func GeneratePasswd(charset string, length int) string {
	//初始化密码切片
	var passwd string
	var level, i int

	switch charset {
	case "num":
		level = 1
	case "char":
		level = 1
	case "mix":
		level = 2
	case "advance":
		level = 3
	default:
		level = 3
	}
	for i = 0; i < 1000; i++ {
		passwd = GenerateRNDString(charset, length)
		if CheckPasswd(passwd, length) >= level {
			break
		}
	}
	if i < 100 {
		return passwd
	} else {
		return ""
	}
}

//	CheckPasswd  检查密码的强度
//		ret := CheckPasswd(passwd,8)  返回一个满足长度大于等于8位的密码强度
//		密码长度必须满足；否则返回-1；
//		强度规则："含数字"、"含大小写"、"含特殊字符"  每满足一条 强度+1
func CheckPasswd(passwd string, expLength int) int {
	level := -1
	//	定义检查正则
	re1 := regexp.MustCompile(`[0-9]`)
	re2 := regexp.MustCompile(`[A-Z]`)
	re3 := regexp.MustCompile(`[a-z]`)
	re4 := regexp.MustCompile(`[+=-@#~,.\[\]\(\)!%^*$\{\}]`)
	if len(passwd) >= expLength { // 长度满足
		level++
		if re1.MatchString(passwd) { // 包含数字
			level++
		}
		if re2.MatchString(passwd) && re3.MatchString(passwd) { // 包含英文大小写字母
			level++
		}
		if re4.MatchString(passwd) {
			level++
		}
	}
	return level
}
