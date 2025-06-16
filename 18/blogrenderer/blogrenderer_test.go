package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	"github.com/approvals/go-approval-tests"
	blogposts "github.com/bgics/Learning-Go/18"
	"github.com/bgics/Learning-Go/18/blogrenderer"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogposts.Post{
			Title: "Welcome to my blog",
			Body: `# First recipe!
Welcome to my **amazing recipe blog**. I am going to write about my family recipes, and make sure I write a long, irrelevant and boring story about my family before you get to the actual instructions.`,
			Description: "Introduction to my blog",
			Tags:        []string{"cooking", "family"},
		}
	)
	postrenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postrenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postrenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogposts.Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postrenderer, err := blogrenderer.NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for range b.N {
		postrenderer.Render(io.Discard, aPost)
	}
}
