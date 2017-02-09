package main

import (
	"errors"
	"fmt"
	"./ResponsibilityFeedbackChain"
)

type obj1 struct {
	ResponsibilityFeedbackChain.Member
}

func (this *obj1) Handle(param string) (interface{}, error) {
	v, e := cache1[param]
	if !e {
		return 0, errors.New("not in cache1")
	} else {
		return v, nil
	}
}

func (this *obj1) Feedback(param string, result interface{}) {
	cache1[param] = result.(int)
}

type obj2 struct {
	ResponsibilityFeedbackChain.Member
}

func (this *obj2) Handle(param string) (interface{}, error) {
	v, e := cache2[param]
	if !e {
		return 0, errors.New("not in cache2")
	} else {
		return v, nil
	}
}

func (this *obj2) Feedback(param string, result interface{}) {
	cache2[param] = result.(int)
}

type obj3 struct {
	ResponsibilityFeedbackChain.Member
}

func (this *obj3) Handle(param string) (interface{}, error) {
	v, e := cache3[param]
	if !e {
		return 0, errors.New("not in cache3")
	} else {
		return v, nil
	}
}

func (this *obj3) Feedback(param string, result interface{}) {
	cache3[param] = result.(int)
}

var cache1 map[string]int = make(map[string]int, 8)
var cache2 map[string]int = make(map[string]int, 8)
var cache3 map[string]int = make(map[string]int, 8)

func main() {
	cache3["3"] = 300

	chain := ResponsibilityFeedbackChain.NewRfChain()
	var m1 ResponsibilityFeedbackChain.Member = &obj1{}
	chain.AddMember(&m1)
	var m2 ResponsibilityFeedbackChain.Member = &obj2{}
	chain.AddMember(&m2)
	var m3 ResponsibilityFeedbackChain.Member = &obj3{}
	chain.AddMember(&m3)

	v := chain.RunChain("3", true)

	fmt.Println(v)
	fmt.Println(cache1["3"])
	fmt.Println(cache2["3"])
}
