package models

import (
	"encoding/xml"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestExpensesClassificationType_MarshalXML(t *testing.T) {
	t.Run("should have `ecls` namespace", func(t *testing.T) {
		id := byte(1)
		classification := ExpensesClassificationType{
			ClassificationType:     "type",
			ClassificationCategory: "category",
			Amount:                 12,
			ID:                     &id,
		}

		ft := assert.NewFluentT(t)

		got, err := xml.Marshal(classification)
		ft.AssertThat(err).IsNil()

		want := []byte(`<ExpensesClassificationType><ecls:classificationType>type</ecls:classificationType><ecls:classificationCategory>category</ecls:classificationCategory><ecls:amount>12</ecls:amount><ecls:id>1</ecls:id></ExpensesClassificationType>`)
		ft.AssertThat(got).IsEqualTo(want)
	})
}
