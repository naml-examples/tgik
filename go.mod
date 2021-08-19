module github.com/naml-examples/simple

go 1.16

require (
	github.com/hexops/valast v1.4.0
	github.com/kris-nova/logger v0.2.2
	github.com/kris-nova/naml v0.2.9
	k8s.io/api v0.22.0
	k8s.io/apimachinery v0.22.0
	k8s.io/client-go v0.22.0
)

replace github.com/hexops/valast => github.com/fkautz/valast v1.4.1-0.20210806063143-f33a97256bcb
