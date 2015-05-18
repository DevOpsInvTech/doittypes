package doittypes

//Ansile interface to marshal a struct into an ansible consumable JSON structure
type Ansible interface {
	//MarshalAnsilbe
	MarshalAnsible() string
}
