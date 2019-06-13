module github.com/owais/demo-service/cmd/service

go 1.12

require (
	github.com/owais/demo-contrib v0.0.0-20190613193153-ac4f35757c44
	github.com/owais/demo-service v0.0.0-20190613192838-7b69d8aac58b
)

// way to use code from same revision
replace github.com/owais/demo-service => ../../
