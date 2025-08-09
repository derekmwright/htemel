package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/derekmwright/htemel/internal/generators/spec"
)

type Config struct {
	outputDir    string
	all          bool
	htmlOnly     bool
	htmlSpecSite string
}

func main() {
	cfg := Config{}

	flag.StringVar(&cfg.outputDir, "output", "spec", "Directory to write spec files to")
	flag.BoolVar(&cfg.all, "all", true, "Generate all spec files")
	flag.BoolVar(&cfg.htmlOnly, "html", false, "Only generate HTML spec files")
	flag.StringVar(&cfg.htmlSpecSite, "html-spec-site", "https://html.spec.whatwg.org/", "HTML spec site name")
	flag.Parse()

	if _, err := os.Stat(cfg.outputDir); err != nil {
		if err = os.MkdirAll(cfg.outputDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	if cfg.htmlOnly || cfg.all {
		path := filepath.Join(cfg.outputDir, "html.json")

		f, err := os.Create(path)
		if err != nil {
			log.Fatal(err)
		}

		req, err := http.Get(cfg.htmlSpecSite)
		if err != nil {
			panic(err)
		}

		var out *spec.Spec
		if out, err = spec.GenerateHTMLSpec(req.Body); err != nil {
			log.Fatal(err)
		}

		jsonOut, err := json.MarshalIndent(out, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		if _, err = f.Write(jsonOut); err != nil {
			log.Fatal(err)
		}
	}
}
