package gotest

import "testing"

func Test_Division(t *testing.T) {
	if i, e := Division(10, 2); i != 5 || e != nil {
		t.Error("error")
	} else {
		t.Log("ok")
	}
}

func Test_Division_2(t *testing.T) {
	t.Error("就是不通过")
}
