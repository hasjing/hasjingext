// RandPWD_test.go
package RandPWD

import (
	"testing"
)

func TestGenerateRNDString(t *testing.T) {
	t.Log("=============  测试 GenerateRNDString =================")
	tests := []struct {
		char   string
		length int
		exp    int
	}{
		{"num", 20, 20},
		{"char", 15, 15},
		{"mix", 20, 20},
		{"advance", 30, 30},
		{"advance", 50, 50},
	}
	for _, tt := range tests {
		actual := GenerateRNDString(tt.char, tt.length)
		if len(actual) != tt.exp {
			t.Fail()
		}
		t.Log("返回的密码串", actual)
	}

}

func BenchmarkGenerateRNDString(b *testing.B) {
	//	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	//	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		GenerateRNDString("advance", 8)
	}
}

func TestCheckPasswd(t *testing.T) {
	t.Log("=============  测试 CheckPasswd =================")
	tests := []struct {
		passwd string //要检查的密码串
		length int    //要满足的长度
		exp    int    //期望的密码强度等级
	}{
		{"12939", 8, -1},
		{"12345678", 8, 1},
		{"mix12312", 8, 1},
		{"miX12312", 8, 2},
		{"miX1231{", 8, 3},
	}
	for _, tt := range tests {
		actual := CheckPasswd(tt.passwd, tt.length)
		if actual != tt.exp {
			t.Fail()
		}
		t.Log("返回密码等级", actual)
	}

}

func BenchmarkCheckPasswd(b *testing.B) {
	//	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	//	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		CheckPasswd("miX1231}", 8)
	}
}

func TestGeneratePasswd(t *testing.T) {
	t.Log("=============  测试 GeneratePasswd =================")
	tests := []struct {
		char   string
		length int
	}{
		{"num", 8},
		{"char", 12},
		{"mix", 12},
		{"advance", 6},
		{"advance", 15},
	}
	for _, tt := range tests {
		actual := GeneratePasswd(tt.char, tt.length)
		if actual == "" {
			t.Fail()
		}
		t.Log("返回的密码", tt.char, actual)
	}

}
func BenchmarkGeneratePasswd(b *testing.B) {
	//	b.StopTimer() //调用该函数停止压力测试的时间计数

	//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	//这样这些时间不影响我们测试函数本身的性能

	//	b.StartTimer() //重新开始时间
	for i := 0; i < b.N; i++ {
		if GeneratePasswd("advance", 8) == "" {
			b.Fail()
		}
	}
}
