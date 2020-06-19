package category

import (
	"testing"

	"github.com/dink10/google-play-parser/pkg/store"
)

var resultsCount = 10

func TestCategory(t *testing.T) {
	sortList := []store.Sort{store.SortHelpfulness, store.SortNewest, store.SortRating}
	for _, sort := range sortList {
		l, err := New(store.Game, sort, store.AgeFiveUnder, Options{
			Country:  "ru",
			Language: "us",
			Number:   resultsCount,
		})
		if err != nil {
			t.Error(err)
		}

		for key, cluster := range l {
			err := cluster.Run()
			if err != nil {
				t.Error(err)
			}

			if len(cluster.Results) != resultsCount {
				t.Errorf("[%s] Expected Results length is %d, got %d", key, resultsCount, len(cluster.Results))
			}
		}
	}
}
