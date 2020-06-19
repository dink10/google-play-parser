package similar

import (
	"github.com/dink10/google-play-parser/pkg/app"
	"github.com/dink10/google-play-parser/pkg/scraper"
)

// Options type alias
type Options = scraper.Options

// New return similar list instance
func New(appID string, options Options) *scraper.Scraper {
	a := app.New(appID, app.Options{options.Country, options.Language})
	err := a.LoadDetails()
	if err != nil || a.SimilarURL == "" {
		return nil
	}
	return scraper.New(a.SimilarURL, &options)
}
