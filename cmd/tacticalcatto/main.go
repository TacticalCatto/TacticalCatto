package main

import (
	"context"
	"html/template"
	"log"
	"os"

	"github.com/google/go-github/v43/github"
)

func main() {
	t, err := template.New("").Parse(`
# Hey there!
I'm {{.Name}}, from {{.Location}}.
I spend most of my time working on [amogus.systems](https://amogus.systems), and [BerryByte](https://berrybyte.net).

## Programming languages, and libraries I use.
- Go
- Rust
- TypeScript/JavaScript
- React
- Vue
- C
- C++

## Want to contact me?
If you want to, your best bet is by Discord: TacticalCatto#1327.
`)
	if err != nil {
		log.Fatal(err)
	}

	c := github.NewClient(nil)
	u, _, err := c.Users.Get(context.Background(), "tacticalcatto")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("README.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := t.Execute(f, map[string]any{
		"Name":     u.GetName(),
		"Location": u.GetLocation(),
	}); err != nil {
		log.Fatal(err)
	}
}
