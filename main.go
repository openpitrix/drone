package main

import (
	_ "openpitrix.io/libconfd"
	_ "openpitrix.io/libconfd/etcdv3/backend"
	_ "openpitrix.io/libconfd/metad/backend"
	_ "openpitrix.io/logger"
)

func main() {
	println("hello drone")
}
