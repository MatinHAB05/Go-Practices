package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/andyfusniak/base58"
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/joho/godotenv"
)

var client *elasticsearch.TypedClient
var ctx = context.Background()
var Marshal = func(v any) string {
	if res, err := json.MarshalIndent(v, "", " "); err != nil {
		log.Println("error:", err)
		return ""
	} else {
		return string(res)
	}
}
var rand_str <-chan *string
var NullStr = func(x string) *string {
	xx := x
	return &xx
}
var rand_int <-chan *int
var falsy = false
var truthly = true

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	inport := os.Getenv("IN_ES")

	logfile, err := os.OpenFile("log/es.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	client, err = elasticsearch.NewTyped(
		elasticsearch.WithAddresses("http://localhost:"+inport),
		elasticsearch.WithLogger(&elastictransport.ColorLogger{
			Output:             logfile,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	if status, err := client.Ping().Do(ctx); err != nil || !status {
		log.Fatal(status, err)
	} else {
		log.Println("Ping-status:", status)
	}

	c := 30
	buf_str := make(chan *string, c)
	buf_int := make(chan *int, c)
	for i := 0; i < c; i++ {
		x, _ := base58.RandString(8)
		x = strings.ToLower(x)
		buf_str <- &x
		xx := rand.Int() % 10000
		buf_int <- &xx
	}
	rand_str = buf_str
	rand_int = buf_int
}

func main() {
	// if res, err := client.Cat.Health().V(true).Do(ctx); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	for _, item := range res {
	// 		fmt.Println(Marshal(item))
	// 	}
	// }

	/////////////////////////////////////////////////////////////////////////////

	// iiiiinnnnnn = <-rand_str
	// if res, err :=
	// 	client.Indices.Create(iiiiinnnnnn).
	// 		Mappings(
	// 			dsl.NewTypeMapping().
	// 				Dynamic(dynamicmapping.Strict).
	// 				AddProperty("name", dsl.NewTextProperty()).
	// 				AddProperty("sex", dsl.NewKeywordProperty()).
	// 				AddProperty("age", dsl.NewIntegerNumberProperty()),
	// 		).
	// 		Settings(
	// 			&types.IndexSettings{
	// 				Index: &types.IndexSettings{
	// 					Mapping: &types.MappingLimitSettings{
	// 						Coerce: &falsy,
	// 					},
	// 				},
	// 			},
	// 		).
	// 		Do(ctx); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(Marshal(res))
	// }

	/////////////////////////////////////////////////////////////////////////////

	// x := struct {
	// 	Name    string `json:"name"` // مهم!!!!!!!!!!!!!!!!!!!!!
	// 	Price   int    `json:"price"`
	// 	InStock int    `json:"in_stock"`
	// }{
	// 	Name:    "Toaster",
	// 	Price:   97,
	// 	InStock: 142,
	// }
	// if res, err := client.Index("products").Document(x).Do(ctx); err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println(Marshal(res))
	// }

	/////////////////////////////////////////////////////////////////////////////

	// { "index": { "_index": "products", "_id": 200 } }
	// { "name": "Espresso Machine", "price": 199, "in_stock": 5 }
	// { "create": { "_index": "products", "_id": 201 } }
	// { "name": "Milk Frother", "price": 149, "in_stock": 14 }
	batch := client.Bulk()

	batch.IndexOp(types.IndexOperation{
		Index_: NullStr("products"),
		Id_:    <-rand_str,
	}, map[string]any{
		"name": *<-rand_str, "price": 97, "in_stock": *<-rand_int,
	})
	batch.CreateOp(types.CreateOperation{
		Index_: NullStr("products"),
		Id_:    <-rand_str,
	}, map[string]any{
		"name": *<-rand_str, "price": 149, "in_stock": *<-rand_int,
	})
	if res, err := batch.Do(ctx); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(Marshal(res))
	}
}
