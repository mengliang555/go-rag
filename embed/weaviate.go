package embed

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/weaviate/weaviate-go-client/v4/weaviate/graphql"

	"github.com/weaviate/weaviate-go-client/v4/weaviate"
)

//var weaviateClient *weaviate.Client

// Create the client
func CreateClient() *weaviate.Client {
	cfg := weaviate.Config{
		Host:    "localhost:8080",
		Scheme:  "http",
		Headers: nil,
	}

	client, err := weaviate.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// Check the connection
	live, err := client.Misc().LiveChecker().Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("isSuccessConnect:%v\n", live)

	return client
}

func CreateObject(ctx context.Context, client *weaviate.Client, className string, properties map[string]interface{}, vector []float32) {
	w, err := client.Data().Creator().
		WithClassName(className).
		WithProperties(properties).
		WithVector(vector).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	_, err = json.MarshalIndent(w.Object, "", "  ")
	if err != nil {
		panic(err)

	}
}

func QueryWithText(ctx context.Context, client *weaviate.Client, className string, vector []float32) map[string]any {
	response, err := client.GraphQL().Get().
		WithClassName(className).
		WithFields(
			graphql.Field{Name: "description"},
			graphql.Field{Name: "person"},
			graphql.Field{
				Name: "_additional",
				Fields: []graphql.Field{
					{Name: "distance"},
				},
			},
		).WithNearVector(client.GraphQL().NearVectorArgBuilder().
		WithVector(vector)).
		WithLimit(3).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	if len(response.Errors) > 0 {
		for _, err := range response.Errors {
			fmt.Println("Errors: ", err)
		}
	}
	fmt.Println(response)

	ans := make(map[string]any)
	for k, v := range response.Data {
		ans[k] = v
	}

	return ans
}

func QueryObject(ctx context.Context, client *weaviate.Client, className string, id ...string) {
	objects := client.Data().ObjectsGetter().
		WithClassName(className)

	if len(id) > 0 {
		objects = objects.WithID(id[0])
	}
	allObject, err := objects.Do(ctx)
	if err != nil {
		// handle error
		panic(err)
	}

	for i, obj := range allObject {
		marshal, err := json.Marshal(obj)
		if err != nil {
			return
		}
		fmt.Printf("object[%v]: %s\n", i, string(marshal))
	}
}

func DeleteObject(ctx context.Context, client *weaviate.Client, className string, id ...string) {
	objects := client.Data().Deleter().
		WithClassName(className)

	if len(id) > 0 {
		objects = objects.WithID(id[0])
	}
	err := objects.Do(ctx)
	if err != nil {
		// handle error
		panic(err)
	}

}
