package doittypes

import (
	"log"
	"reflect"
)

const (
	//AnsileTag Struct tag for ansible style output
	AnsibleTag = "ansible"
	//AnsilbeTagSkip Struct tag for ansible to skip field
	AnsilbeTagSkip = "-"
	//AnsibleTagRoot root means take that as the root value
	AnsibleTagRoot = "root"
	//AnsibleTagParent parent means to link the item to its parent
	AnsibleTagParent = "parent"
	//AnsibleTagValue value means value for item
	AnsibleTagValue = "value"
	//AnsibleTagMain main means this is a top tier item to embed into the json
	AnsibleTagMain = "main"
)

//Ansile interface to marshal a struct into an ansible consumable JSON structure
type Ansible interface {
	//MarshalAnsilbe
	MarshalAnsible(n *AnsibleNode)
}

type AnsibleNode struct {
	Main  []*AnsibleNode
	Prev  *AnsibleNode
	Next  *AnsibleNode
	Key   string
	Value string
}

/*
	n ->
		has child:
			recurse
		else:
			extract value
*/

func foo(n *AnsibleNode, e interface{}) {
	switch e.(type) {
	case string:
		n.Value = e.(string)
	case *Host:
		n.Value = e.(*Host).Name
	case *Group:
		n.Value = e.(*Group).Name
	case *Var:
		n.Value = e.(*Var).Value
	}
}

//AnsibleCheckTag Check struct tags
func AnsibleCheckTag(field reflect.StructField, n *AnsibleNode, e interface{}) *AnsibleNode {
	tag := field.Tag.Get(AnsibleTag)
	switch tag {
	case AnsibleTagMain:
		log.Println(field.Name, field.Type.Name(), field.Tag.Get(AnsibleTag))
	case AnsibleTagRoot:
		log.Println(field.Name, field.Type.Name(), field.Tag.Get(AnsibleTag))
		nn := &AnsibleNode{}
		InsertValue(n, e)
		n.Main = append(n.Main, nn)
	case AnsibleTagParent:
		log.Println(field.Name, field.Type.Name(), field.Tag.Get(AnsibleTag))
		InsertNode(n, e)
	case AnsibleTagValue:
		InsertValue(n, e)
		log.Println(field.Name, field.Type.Name(), field.Tag.Get(AnsibleTag))
	case AnsilbeTagSkip:
		//Ignore tag
	}
	return n
}

func InsertValue(n *AnsibleNode, e interface{}) {
	switch e.(type) {
	case string:
		n.Value = e.(string)
	case *Host:
		n.Value = e.(*Host).Name
	case *Group:
		n.Value = e.(*Group).Name
	case *Var:
		n.Value = e.(*Var).Value
	}
}

func InsertNode(n *AnsibleNode, e interface{}) {
	switch e.(type) {
	case string:
		n.Value = e.(string)
	case *Host:
		n.Value = e.(*Host).Name
	case *Group:
		n.Key = e.(*Group).Name
	case *Var:
		n.Key = e.(*Var).Name
		n.Value = e.(*Var).Value
	}
}
