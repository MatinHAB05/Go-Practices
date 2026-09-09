package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ctx := context.Background()

	// اتصال به SQLite (فایل لوکال)
	db, err := sql.Open(
		dialect.SQLite,
		"file:ent.db?_fk=1",
	)
	if err != nil {
		log.Fatal(err)
	}

	client := ent.NewClient(ent.Driver(db))
	defer client.Close()

	// migrate (ساخت جدول‌ها)
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	// =====================
	// Create Movie
	// =====================
	movie, err := client.Movie.
		Create().
		SetTitle("Interstellar").
		SetReleaseYear(2014).
		SetQuality("1080p").
		Save(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// =====================
	// Create Casts
	// =====================
	actor1, _ := client.Cast.
		Create().
		SetName("Matthew McConaughey").
		Save(ctx)

	actor2, _ := client.Cast.
		Create().
		SetName("Anne Hathaway").
		Save(ctx)

	// =====================
	// Connect Movie ↔ Casts
	// =====================
	err = client.Movie.
		UpdateOne(movie).
		AddCasts(actor1, actor2).
		Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// =====================
	// Query with relation
	// =====================
	movies, err := client.Movie.
		Query().
		WithCasts().
		All(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// =====================
	// Print result
	// =====================
	for _, m := range movies {
		fmt.Printf("Movie: %s (%d)\n", m.Title, m.ReleaseYear)
		for _, c := range m.Edges.Casts {
			fmt.Printf("  - Cast: %s\n", c.Name)
		}
	}
}
