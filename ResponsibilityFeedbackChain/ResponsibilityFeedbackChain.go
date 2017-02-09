package ResponsibilityFeedbackChain

type Chain struct {
	members []*Member
}

type Member interface {
	Handle(param string) (interface{}, error)
	Feedback(param string, result interface{})
}

func NewRfChain() Chain {
	return Chain{members: make([]*Member, 0, 4)}
}

func (this *Chain) AddMember(member *Member) int {
	this.members = append(this.members, member)
	return len(this.members)
}

func (this *Chain) RunChain(param string, feedback bool) interface{} {
	index := 0
	var result interface{} = nil
	for i, v := range this.members {
		r, e := (*v).Handle(param)
		if e == nil {
			result = r
			index = i
			break
		}
	}

	if result == nil {
		return nil
	}

	for i := index - 1; feedback && i >= 0; i-- {
		(*this.members[i]).Feedback(param, result)
	}

	return result
}
