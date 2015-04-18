package string

import (
	"testing"
	"github.com/nbio/st"
	"github.com/vinayak/win"
	)

func Test(t *testing.T){
	var tests =[]struct{
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"", ""},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, c :=range tests{
		got :=Reverse(c.s)
		if(got != c.want){
			t.Errorf("Reverse(%q) == %q, want %q", c.s, got, c.want)
		}
	}
}

func TestSt(t *testing.T){
	var tests =[]struct{
		s, want string
	}{
		{"Backward", "drawkcaB"},
		{"", ""},
		{"Hello, 世界", "界世 ,olleH"},
	}
	for _, c :=range tests{
		st.Assert(t, Reverse(c.s), c.want)
		win.Expect(t,Reverse(c.s), c.want)
	}
	
}