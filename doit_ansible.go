package doittypes

//http://docs.ansible.com/developing_inventory.html

//Ansile interface to marshal a struct into an ansible consumable JSON structure
type Ansible interface {
	//MarshalAnsilbe
	MarshalAnsible() map[string]interface{}
}
