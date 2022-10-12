package models

import (
	"encoding/xml"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestIncomeClassificationType_MarshalXML(t *testing.T) {
	t.Run("should have `icls` namespace", func(t *testing.T) {
		id := byte(1)
		classification := IncomeClassificationType{
			ClassificationType:     "type",
			ClassificationCategory: "category",
			Amount:                 12,
			ID:                     &id,
		}

		ft := assert.NewFluentT(t)

		got, err := xml.Marshal(classification)
		ft.AssertThat(err).IsNil()

		want := []byte(`<IncomeClassificationType><icls:classificationType>type</icls:classificationType><icls:classificationCategory>category</icls:classificationCategory><icls:amount>12</icls:amount><icls:id>1</icls:id></IncomeClassificationType>`)
		ft.AssertThat(got).IsEqualTo(want)
	})
}
