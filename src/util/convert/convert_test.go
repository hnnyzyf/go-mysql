package convert

import "testing"

func Test_String2Number(t *testing.T) {
	s2 := []string{
		"\"1234aaaa\"",
		"\"001234aaaa\"",
		"\"0aaaa\"",
		"\"0000aaaa\"",
		"\"0000\"",
		"\"0\"",
		"\"100\"",
		"\"1\"",
	}

	res2 := []string{
		"1234",
		"1234",
		"0",
		"0",
		"0",
		"0",
		"100",
		"1",
	}
	for i, s_test := range s2 {
		r_test := String2Number(s_test)
		if r_test == res2[i] {
			t.Log("string2number测试通过")
		} else {
			t.Error("string2number测试失败，返回", r_test, ",应该是", res2[i])
		}
	}
}
