package embed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

// Create the client
func CreateClient() {
	cfg := weaviate.Config{
		Host:    "localhost:8080",
		Scheme:  "http",
		Headers: nil,
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	live, err := client.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", live)

}

func CreateObject(ctx context.Context, client *weaviate.Client) {
	w, err := client.Data().Creator().
		WithClassName("JeopardyQuestion").
		WithProperties(map[string]interface{}{
			"question": "This vector DB is OSS and supports automatic property type inference on import",
			// "answer": "Weaviate", // schema properties can be omitted
			"newProperty": 123, // will be automatically added as a number property
		}).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	// the returned value is a wrapped object
	b, err := json.MarshalIndent(w.Object, "", "  ")
	fmt.Println(string(b))
}
