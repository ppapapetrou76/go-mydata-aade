package models

import (
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestResponseDoc_HasErrors(t *testing.T) {
	t.Run("should return false if response has no errors", func(t *testing.T) {
		rDoc := ResponseDoc{Response: []*Response{}}
		assert.That(t, rDoc.HasErrors()).IsFalse()
	})
	t.Run("should return true if response has at least one error", func(t *testing.T) {
		rDoc := ResponseDoc{Response: []*Response{
			{Errors: Errors{Error: []Error{mockError()}}},
		}}
		assert.That(t, rDoc.HasErrors()).IsTrue()
	})
}

func TestResponseDoc_Errors(t *testing.T) {
	rDoc := ResponseDoc{Response: []*Response{
		{Errors: Errors{Error: []Error{mockError(), mockError()}}},
	}}
	assert.ThatError(t, rDoc.Errors("some prefix")).
		HasExactMessage("some prefix 2 errors:[ Code:code - Message:message, Code:code - Message:message ]")
}

func TestError_Error(t *testing.T) {
	assert.That(t, mockError().Error()).IsEqualTo("Code:code - Message:message")
}

func mockError() Error {
	return Error{Message: "message", Code: "code"}
}
