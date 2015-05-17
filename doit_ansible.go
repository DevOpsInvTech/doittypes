package doittypes

const (
	//Struct tag for ansible style output
	AnsileTag = "ansible"
)

type Ansible interface {
	MarshalAnsilbe()
}
