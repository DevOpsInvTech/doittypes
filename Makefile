
test_ansible:
	godep go test -v -run=TestDomainAnsible*

build_deps:
	godep go install -a .
