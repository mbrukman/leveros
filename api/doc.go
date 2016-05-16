/*

Package api provides Lever OS libraries for both implementing Lever services in
Go and invoking Lever methods, as a client.

Quick example of a Lever service implementation

    $ mkdir hello
    $ cd hello

server.go

    package main

    import (
        "fmt"
        "log"

        leverapi "github.com/leveros/leveros/api"
    )

    func main() {
    	server, err := leverapi.NewServer()
    	if err != nil {
    		log.Fatalf("Error: %v\n", err)
    	}
    	err = server.RegisterHandlerObject(new(Handler))
    	if err != nil {
    		log.Fatalf("Error: %v\n", err)
    	}
    	server.Serve()
    }

    type Handler struct {
    }

    func (*Handler) SayHello(name string) (result string, err error) {
    	return fmt.Sprintf("Hello, %s!", name), nil
    }

lever.json

    {
        "name": "helloService",
        "description": "A hello service.",
        "entry": ["./serve"]
    }

Compile and deploy

    $ GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./serve server.go
    $ lever deploy

Note the env vars for compiling for Lever OS. These need to be
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 even when running on Mac or Windows.

If you have problems building, you may need to reinstall go to include
cross-compilation support. On a Mac, you can achieve this with
brew install go --with-cc-common.

Quick example of a Lever client

    package main

    import (
    	"log"
    	"os"

    	leverapi "github.com/leveros/leveros/api"
    )

    func main() {
        client, err := leverapi.NewClient()
        if err != nil {
    		log.Fatalf("Error: %v\n", err)
    	}
    	client.ForceHost = os.Getenv("LEVEROS_IP_PORT")
        leverService := client.Service("dev.lever", "helloService")
        var reply string
        err = leverService.Invoke(&reply, "SayHello", "world")
        if err != nil {
            log.Printf("Error: %v\n", err)
        }
        log.Printf("%s\n", reply) // Hello, world!
    }

To run this

    # Without docker-machine
    $ LEVEROS_IP_PORT="127.0.0.1:8080" go run client.go

    # With docher-machine
    $ LEVEROS_IP_PORT="$(docker-machine ip default):8080" go run client.go

Setting LEVEROS_IP_PORT is necessary so that you can invoke the dev.lever
environment without adding an entry for it in /etc/hosts and setting the
listen port to 80.

*/
package api
