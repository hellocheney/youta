package say

import (
	"fmt"
	"testing"
)

func TestSay(t *testing.T) {
	name := "cheney"
	word := Say(name)
	if word != name {
		t.Error("不通过\n")
	}
}

type sayTestCase struct {
	name   string
	except string
}

var sayTestCases = []sayTestCase{
	{name: "小丽", except: "小丽"},
	{name: "小米", except: "小米"},
	{name: "xiaohua", except: "小花"},
	{name: "王丽", except: "王丽"},
	{name: "赵敏", except: "赵敏"},
	{name: "张无忌", except: "张无急"},
}

func TestSayA(t *testing.T) {
	for name, test := range sayTestCases {
		t.Run(fmt.Sprint(name), func(tt *testing.T) {
			out := Say(test.name)
			if out != test.except {
				tt.Errorf("except:%s,but got:%s\n", test.except, out)
			}
		})
	}

}
