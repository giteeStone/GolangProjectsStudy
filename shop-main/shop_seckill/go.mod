module shop_seckill

go 1.14

require (
	github.com/garyburd/redigo v1.6.4
	github.com/golang/protobuf v1.5.2
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/streadway/amqp v1.0.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
