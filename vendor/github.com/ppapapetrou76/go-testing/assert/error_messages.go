package assert

import (
	"fmt"
	"reflect"
	"strings"

	utils2 "github.com/ppapapetrou76/go-testing/internal/pkg/utils"
	"github.com/ppapapetrou76/go-testing/types"
	"github.com/r3labs/diff/v2"
)

func shouldBeEqual(actual types.Assertable, expected interface{}) string {
	diffMessage := strings.Builder{}
	skipDetailedDiff := utils2.HasUnexportedFields(reflect.ValueOf(expected)) || utils2.HasUnexportedFields(reflect.ValueOf(actual.Value()))

	if !skipDetailedDiff {
		diffs, _ := diff.Diff(expected, actual.Value())
		for _, d := range diffs {
			if len(d.Path) == 0 {
				continue
			}
			switch d.Type {
			case "delete":
				path := strings.Join(d.Path, ":")
				diffMessage.WriteString(fmt.Sprintf("actual value of %+v is expected but missing from %s\n", d.To, path))
			case "create":
				path := strings.Join(d.Path, ":")
				diffMessage.WriteString(fmt.Sprintf("actual value of %+v is not expected in %s\n", d.To, path))
			case "update":
				path := strings.Join(d.Path, ":")
				diffMessage.WriteString(fmt.Sprintf("actual value of %+v is different in %s from %+v\n", d.To, path, d.From))
			}
		}
	}

	return fmt.Sprintf("assertion failed:\nexpected value\t:%+v\nactual value\t:%+v\n%s", expected, actual.Value(), diffMessage.String())
}

func shouldNotBeEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be other than %+v", actual.Value(), expected)
}

func shouldBeGreater(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than %+v", actual.Value(), expected)
}

func shouldBeGreaterOrEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than or equal to %+v", actual.Value(), expected)
}

func shouldBeLessThan(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than %+v", actual.Value(), expected)
}

func shouldBeLessOrEqual(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be less than or equal to %+v", actual.Value(), expected)
}

func shouldBeEmpty(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be empty, but it's not", actual.Value())
}

func shouldNotBeEmpty(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v not to be empty, but it is", actual.Value())
}

func shouldBeNil(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be nil but it wasn't", actual.Value())
}

func shouldNotBeNil(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be non-nil but it was", actual.Value())
}

func shouldHaveSize(actual types.Sizeable, expected int) string {
	return fmt.Sprintf("assertion failed: expected size of = [%d], to be but it has size of [%d] ", actual.Size(), expected)
}

func shouldContain(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain [%+v], but it doesn't", actual.Value(), elements)
}

func shouldContainIgnoringCase(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain [%+v] ignoring case, but it doesn't", actual.Value(), elements)
}

func shouldContainOnly(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain only [%+v], but it doesn't", actual.Value(), elements)
}

func shouldContainOnlyOnce(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain [%+v] only once, but it doesn't", actual.Value(), elements)
}

func shouldContainWhiteSpace(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain whitespace(s), but it doesn't", actual.Value())
}

func shouldNotContainAnyWhiteSpace(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: containable [%v] should not contain any whitespace, but it does", actual.Value())
}

func shouldContainOnlyWhiteSpaces(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: containable [%v] should contain only whitespaces, but it does not", actual.Value())
}

func shouldNotContainOnlyWhiteSpaces(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: containable [%v] should not contain only whitespaces, but it does", actual.Value())
}

func shouldNotContain(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: containable [%v] should not contain [%+v], but it does", actual.Value(), elements)
}

func shouldBeMap(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: assertable should be a map but it is %T", reflect.ValueOf(actual.Value()).Kind())
}

func shouldHaveKey(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: map [%v] should have the key [%+v], but it doesn't", actual.Value(), elements)
}

func shouldHaveValue(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: map [%v] should have the value [%+v], but it doesn't", actual.Value(), elements)
}

func shouldHaveEntry(actual types.Assertable, entry types.MapEntry) string {
	return fmt.Sprintf("assertion failed: map [%v] should have the entry [%+v], but it doesn't", actual.Value(), entry)
}

func shouldNotHaveKey(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: map [%v] should not have the key [%+v], but it does", actual.Value(), elements)
}

func shouldNotHaveValue(actual types.Assertable, elements interface{}) string {
	return fmt.Sprintf("assertion failed: map [%v] should not have the value [%+v], but it does", actual.Value(), elements)
}

func shouldNotHaveEntry(actual types.Assertable, entry types.MapEntry) string {
	return fmt.Sprintf("assertion failed: map [%v] should not have the entry [%+v], but it does", actual.Value(), entry)
}

func shouldStartWith(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected value of [%v] to start with [%+v], but it doesn't", actual.Value(), substr)
}

func shouldNotStartWith(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected value of [%v] to not start with [%+v], but it does", actual.Value(), substr)
}

func shouldEndWith(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected value of [%v] to end with [%+v], but it doesn't", actual.Value(), substr)
}

func shouldNotEndWith(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected value of [%v] to not end with [%+v], but it does", actual.Value(), substr)
}

func shouldHaveSameSizeAs(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected size of [%v] should be same as the size of [%+v], but it isn't", actual.Value(), substr)
}

func shouldHaveLessSizeThan(actual types.Assertable, substr string) string {
	return fmt.Sprintf("assertion failed: expected size of [%v] should be less than the size of [%+v], but it isn't", actual.Value(), substr)
}

func shouldHaveType(actual types.Assertable, value interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to have type of %T but it hasn't", actual.Value(), value)
}

func shouldBeShorter(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be greater than %+v", actual.Value(), expected)
}

func shouldBeLonger(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be longer than %+v", actual.Value(), expected)
}

func shouldContainOnlyDigits(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to have only digits, but it's not", actual.Value())
}

func shouldBeLowerCase(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be lower case, but it's not", actual.Value())
}

func shouldBeUpperCase(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected %+v to be upper case, but it's not", actual.Value())
}

func shouldBeAlmostSame(actual types.Assertable, expected interface{}) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be almost the same as %+v", actual.Value(), expected)
}

func shouldBeDefined(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be defined but it was not", actual.Value())
}

func shouldNotBeDefined(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be un-defined but it was", actual.Value())
}

func shouldContainUniqueElements(actual types.Assertable) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to contain unique elements but it doesn't", actual.Value())
}

func shouldBeSorted(actual types.Assertable, order string) string {
	return fmt.Sprintf("assertion failed: expected value of = %+v, to be sorted in %s order but it isn't", order, actual.Value())
}
