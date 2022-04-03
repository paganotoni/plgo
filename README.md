## PLGO

Plgo is a sample repo for how to use [Plush](https://github.com/gobuffalo/plush) without Buffalo. While this is not a complete application it is a good starting point for your own projects.

### Running this app

To run this app you need Go 1.16+. The app can be run within the root folder invoking the following command:

```sh
go run cmd/plgo
```

Which should start the app on port 8080.

### Plush usage

Main plush usage is in the [landing.go](handlers/landing.go) file. In there we set some variables on the http handlerFunc that then we execute with plush.

```go
func Landing(w http.ResponseWriter, r *http.Request) {
	c := plush.NewContextWithContext(r.Context())   // Setup a plush context from the request context
	c.Set("title", "This is a plush template")      // Setting some variables to be used by plush
	c.Set("items", []string{"one", "two", "three"})

	s, err := plush.Render(landingHTML, c) // rendering the template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Write([]byte(s)) // Writing the template to the response writer
}
```

And the html file is:

```html
...
<body>
    <h2><%= title %></h2>
    <p>This is a sample of using plush in Go</p>
    <ul>
        <%= for (k) in items { %>
            <li><%= k %></li>
        <% } %>
    </ul>
</body>
...
```
