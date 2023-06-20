package main

import (
	"github.com/Msaorc/MaSystemUtils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	rabbitmq.Producer(ch, "{'teste':'teste'}", "amq.direct")
}
