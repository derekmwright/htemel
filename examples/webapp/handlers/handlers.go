package handlers

import (
	"bytes"
	"net/http"

	"github.com/derekmwright/htemel/examples/webapp/views"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	view := views.Layout(
		views.PageData{
			Title:   "htemel Example WebApp",
			Request: r,
		},
		views.Index(),
	)

	buf := &bytes.Buffer{}
	if err := view.Render(buf); err != nil {
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf.Bytes())
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	view := views.Layout(
		views.PageData{
			Title:   "htemel Example WebApp - About",
			Request: r,
		},
		views.About(),
	)

	buf := &bytes.Buffer{}
	if err := view.Render(buf); err != nil {
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf.Bytes())
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	view := views.Layout(
		views.PageData{
			Title:   "htemel Example WebApp - Contact",
			Request: r,
		},
		views.Contact(),
	)

	buf := &bytes.Buffer{}
	if err := view.Render(buf); err != nil {
		http.Error(w, "Failed to render view", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buf.Bytes())
}
