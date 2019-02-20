package ResponsibilityFeedbackChain

import (
	"errors"
	"fmt"
	"testing"
)

type obj1 struct {
	Member
}

func (this *obj1) Handle(params interface{}) (interface{}, error) {
	v, e := cache1[params]
	if !e {
		return 0, errors.New("not in cache1")
	} else {
		return v, nil
	}
}

func (this *obj1) Feedback(params interface{}, result interface{}) {
	cache1[params] = result.(int)
}

type obj2 struct {
	Member
}

func (this *obj2) Handle(params interface{}) (interface{}, error) {
	v, e := cache2[params]
	if !e {
		return 0, errors.New("not in cache2")
	} else {
		return v, nil
	}
}

func (this *obj2) Feedback(params interface{}, result interface{}) {
	cache2[params] = result.(int)
}

type obj3 struct {
	Member
}

func (this *obj3) Handle(params interface{}) (interface{}, error) {
	v, e := cache3[params]
	if !e {
		return 0, errors.New("not in cache3")
	} else {
		return v, nil
	}
}

func (this *obj3) Feedback(params interface{}, result interface{}) {
	cache3[params] = result.(int)
}

var cache1 = make(map[interface{}]int, 8)
var cache2 = make(map[interface{}]int, 8)
var cache3 = make(map[interface{}]int, 8)

func TestRunChain(t *testing.T) {
	cache3["3"] = 300

	chain := NewRfChain()
	var m1 Member = &obj1{}
	chain.AddMember(&m1)
	var m2 Member = &obj2{}
	chain.AddMember(&m2)
	var m3 Member = &obj3{}
	chain.AddMember(&m3)

	v := chain.RunChain("3", true)

	fmt.Println(v)
	fmt.Println(cache1["3"])
	fmt.Println(cache2["3"])
}
