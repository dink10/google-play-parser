package collection

import (
	"github.com/dink10/google-play-parser/pkg/scraper"
	"github.com/dink10/google-play-parser/pkg/store"
)

// Options type alias
type Options = scraper.Options

// New return collection list instance
func New(collection store.Collection, options Options) *scraper.Scraper {
	return scraper.New(scraper.BaseURL+"/collection/"+string(collection), &options)
}
