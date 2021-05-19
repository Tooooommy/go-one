package meili

import (
	"fmt"
	"github.com/meilisearch/meilisearch-go"
	"testing"
)

type Movies struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Poster      string   `json:"poster"`
	Overview    string   `json:"overview"`
	ReleaseDate int      `json:"release_date"`
	Genres      []string `json:"genres"`
}

var movies = []Movies{
	{
		ID:          "42312",
		Title:       "你好",
		Poster:      "https://image.tmdb.org/t/p/w500/xnopI5Xtky18MPhK40cZAGAOVeV.jpg",
		Overview:    "我叫子健， 天下第一",
		ReleaseDate: 1553299200,
		Genres:      []string{"小孩", "笨小孩", "三重刘德华"},
	},
}

func TestInitPool(t *testing.T) {
	p := NewPool("http://127.0.0.1:7700", "")
	client := p.Get().(meilisearch.ClientInterface)
	fmt.Println(client.Version())
	update, err := client.Documents("movies").AddOrUpdate(movies)
	if err != nil {
		panic(err)
	}
	fmt.Print(update)
}
