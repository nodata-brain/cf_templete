package helloworld

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorld", HelloWorld)
}

// helloWorld writes "Hello, World!" to the HTTP response.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")

}
