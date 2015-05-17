
test_ansible:
	godep go test -v -run=TestDomainAnsibleBasic

build_deps:
	godep go install -a .
