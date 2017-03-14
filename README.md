ResponsibilityFeedbackChain
============================
### Introduction
This is a tiny framework for acquiring a result from a Responsibility Chain.  
After acquiring the result without errors, it can feed the result back to previous Chain Members.

### Internal Logic 
    // todo
    // To difficult for me to draw a graph

### Installation
    go get github.com/nuo9/ResponsibilityFeedbackChain.go/ResponsibilityFeedbackChain
    
### Usage
##### For more details, see [testChain.go](https://github.com/nuo9/ResponsibilityFeedbackChain.go/blob/master/testChain.go)
    // create a chain
    chain := ResponsibilityFeedbackChain.NewRfChain()
    // create members
    type obj1 struct {
    	ResponsibilityFeedbackChain.Member
    }
    
    func (this *obj1) Handle(param string) (interface{}, error) {
    	v, e := tryToHandle(param)
    	if e != nil {
    		return 0, e
    	} else {
    		return v, nil
    	}
    }
    
    func (this *obj1) Feedback(param string, result interface{}) {
    	// do something or just empty
    }
    // add members to chain
    var m1 ResponsibilityFeedbackChain.Member = &obj1{}
    chain.AddMember(&m1)
    // acquiring a result
    result := chain.RunChain("3", true)

### Java Version
[ResponsibilityFeedbackChain.java](https://github.com/nuo9/ResponsibilityFeedbackChain.java, "daye come to play!")

### MD Reference
    https://github.com/docker/docker/blob/master/README.md
