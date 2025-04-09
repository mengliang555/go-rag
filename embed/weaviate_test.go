package embed

import (
	"context"
	"fmt"
	"testing"
)

func TestClientPing(t *testing.T) {
	CreateClient()
}

var content01 = "John是后端开发,擅长go语言"
var content02 = "Milk是前端开发，擅长react,John最好的朋友"
var className = "PersonDescript"

var ctx = context.Background()

func TestCreateObject(t *testing.T) {
	client := CreateClient()
	CreateObject(ctx, client, className, map[string]interface{}{
		"description": content01,
		// "answer": "Weaviate", // schema properties can be omitted
		"person": "John", // will be automatically added as a number property
	}, MustEmbeddingSingleString(ctx, content01))
	CreateObject(ctx, client, className, map[string]interface{}{
		"description": content02,
		// "answer": "Weaviate", // schema properties can be omitted
		"person": "Milk", // will be automatically added as a number property
	}, MustEmbeddingSingleString(ctx, content02))

}

func TestQueryWithVector(t *testing.T) {
	client := CreateClient()
	vectorValue := QueryWithText(ctx, client, className, MustEmbeddingSingleString(ctx, "John"))
	for k, v := range vectorValue {
		fmt.Printf("key:%s value:%v\n", k, v)
	}
}
func TestQueryObject(t *testing.T) {
	client := CreateClient()
	QueryObject(context.TODO(), client, className)
}

func TestDeleteObject(t *testing.T) {
	client := CreateClient()
	DeleteObject(context.TODO(), client, "PersonDescript", "429e58e6-7dce-44d6-ad00-464b0b287309")
}
