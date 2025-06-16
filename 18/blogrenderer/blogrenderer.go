package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	blogposts "github.com/bgics/Learning-Go/18"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed templates/*
	postTemplates embed.FS
)

type PostRenderer struct {
	tmpl *template.Template
}

type PostViewModel struct {
	Title, SanitisedTitle, Description string
	Body                               template.HTML
	Tags                               []string
}

func NewPostView(p blogposts.Post) PostViewModel {
	sanitizedTitle := strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
	body := template.HTML(markdownToHTML(p.Body))

	return PostViewModel{
		Title:          p.Title,
		SanitisedTitle: sanitizedTitle,
		Description:    p.Description,
		Body:           body,
		Tags:           p.Tags,
	}
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	return r.tmpl.ExecuteTemplate(w, "blog.gohtml", NewPostView(p))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	var postViews []PostViewModel

	for _, p := range posts {
		postViews = append(postViews, NewPostView(p))
	}

	return r.tmpl.ExecuteTemplate(w, "index.gohtml", postViews)
}

func markdownToHTML(body string) string {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(body))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
