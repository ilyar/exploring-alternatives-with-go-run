package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Item struct {
	Name     string
	Quantity int
}

func unmarshal(b []byte) ([]Item, error) {
	var items []Item
	err := json.Unmarshal(b, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func main() {
	log.SetFlags(log.Lshortfile)

	wanted := []Item{
		{Name: "Gopher Plush", Quantity: 5},
		{Name: "Gopher Sticker", Quantity: 77},
	}

	example, err := ioutil.ReadFile("example.json")
	if err != nil {
		log.Fatalln(err)
	}

	// test unmarshal
	object, err := unmarshal(example)
	if err != nil {
		log.Fatalln(err)
	}
	if !cmp.Equal(object, wanted) {
		log.Fatalln(cmp.Diff(object, wanted))
	}

	// benchmark unmarshal
	// need more than one result for benchstat
	for i := 0; i < 5; i++ {
		bResult := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := unmarshal(example)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
		fmt.Printf("BenchmarkUnmarshal\t%s\t%s\n",
			bResult, bResult.MemString())
	}
}
