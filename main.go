package main

/*import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)*/

import (
	"fmt"
	"os"

	"github.com/armon/go-socks5"
)

func getEnv(key string, defval string) string {
	val := os.Getenv(key)
	if val == "" {
		return defval
	}
	return val
}

func main() {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		fmt.Printf("Error create sock server: %s\n", err)
		panic(err)
	}

	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), getEnv("OPENSHIFT_GO_PORT", "8080"))
	fmt.Printf("listening on %s...", bind)

	if err := server.ListenAndServe("tcp", bind); err != nil {
		fmt.Printf("Error listen and serve %s\n", err)
		panic(err)
	}
}

/*func main() {
	http.HandleFunc("/", hello)
	bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
	fmt.Printf("listening on %s...", bind)
	err := http.ListenAndServe(bind, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", runtime.Version())
}*/
