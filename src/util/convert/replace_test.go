package convert

import "testing"

func Test_convert(t *testing.T) {
	s := []string{
		"000000001234",
		"00000",
		"00001",
		"123455",
	}

	res := []string{
		"1234",
		"0",
		"1",
		"123455",
	}
	for i, s_test := range s {
		r_test := Replace0(s_test, Raplace_prefix)
		if r_test == res[i] {
			t.Log("prefix测试通过")
		} else {
			t.Error("prefix测试失败，返回", r_test, ",应该是", res[i])
		}
	}

	s1 := []string{
		"00000000",
		"100000000",
		"12340000",
		"1234556",
	}

	res1 := []string{
		"0",
		"1",
		"1234",
		"1234556",
	}
	for i, s_test := range s1 {
		r_test := Replace0(s_test, Raplace_suffix)
		if r_test == res1[i] {
			t.Log("suffix测试通过")
		} else {
			t.Error("suffix测试失败，返回", r_test, ",应该是", res1[i])
		}
	}

	s2 := []string{
		"0x1234556",
		"0x000006",
		"0x000023456",
		"0x0000",
		"X'123456'",
		"X'0000000123456'",
		"X'0000001'",
		"X'000000'",
	}

	res2 := []string{
		"0x1234556",
		"0x6",
		"0x23456",
		"0x0",
		"X'123456'",
		"X'123456'",
		"X'1'",
		"X'0'",
	}
	for i, s_test := range s2 {
		r_test := Replace0(s_test, Raplace_bit_prefix)
		if r_test == res2[i] {
			t.Log("Hex和Bit suffix测试通过")
		} else {
			t.Error("Hex和Bit suffix测试失败，返回", r_test, ",应该是", res2[i])
		}
	}

}
