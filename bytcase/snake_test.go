/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package bytcase_test

import (
	"bytes"
	"testing"

	"github.com/peyton-spencer/caseconv/bytcase"
)

var snakeCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("test_case")},
	{[]byte("TestCase"), []byte("test_case")},
	{[]byte("Test Case"), []byte("test_case")},
	{[]byte(" Test Case"), []byte("test_case")},
	{[]byte("Test Case "), []byte("test_case")},
	{[]byte(" Test Case "), []byte("test_case")},
	{[]byte("test"), []byte("test")},
	{[]byte("test_case"), []byte("test_case")},
	{[]byte("Test"), []byte("test")},
	{[]byte(""), []byte("")},
	{[]byte("ManyManyWords"), []byte("many_many_words")},
	{[]byte("manyManyWords"), []byte("many_many_words")},
	{[]byte("AnyKind of_string"), []byte("any_kind_of_string")},
	{[]byte("numbers2and55with000"), []byte("numbers_2_and_55_with_000")},
	{[]byte("JSONData"), []byte("json_data")},
	{[]byte("userID"), []byte("user_id")},
	{[]byte("AAAbbb"), []byte("aa_abbb")},
	{[]byte("1A2"), []byte("1_a_2")},
	{[]byte("A1B"), []byte("a_1_b")},
	{[]byte("A1A2A3"), []byte("a_1_a_2_a_3")},
	{[]byte("A1 A2 A3"), []byte("a_1_a_2_a_3")},
	{[]byte("AB1AB2AB3"), []byte("ab_1_ab_2_ab_3")},
	{[]byte("AB1 AB2 AB3"), []byte("ab_1_ab_2_ab_3")},
	{[]byte("some string"), []byte("some_string")},
	{[]byte(" some string"), []byte("some_string")},
}

func toSnake(tb testing.TB) {
	for _, c := range snakeCases {
		result := bytcase.ToSnake(c.input)
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToSnake(t *testing.T) { toSnake(t) }

func BenchmarkToSnake(b *testing.B) {
	benchmarkSnakeTest(b, toSnake)
}

var snakeWithIgnoreCases = []struct {
	input  []byte
	output []byte
	ignore []byte
}{
	{[]byte("testCase"), []byte("test_case"), nil},
	{[]byte("TestCase"), []byte("test_case"), nil},
	{[]byte("Test Case"), []byte("test_case"), nil},
	{[]byte(" Test Case"), []byte("test_case"), nil},
	{[]byte("Test Case "), []byte("test_case"), nil},
	{[]byte(" Test Case "), []byte("test_case"), nil},
	{[]byte("test"), []byte("test"), nil},
	{[]byte("test_case"), []byte("test_case"), nil},
	{[]byte("Test"), []byte("test"), nil},
	{[]byte(""), []byte(""), nil},
	{[]byte("ManyManyWords"), []byte("many_many_words"), nil},
	{[]byte("manyManyWords"), []byte("many_many_words"), nil},
	{[]byte("AnyKind of_string"), []byte("any_kind_of_string"), nil},
	{[]byte("numbers2and55with000"), []byte("numbers_2_and_55_with_000"), nil},
	{[]byte("JSONData"), []byte("json_data"), nil},
	{[]byte("AwesomeActivity.UserID"), []byte("awesome_activity.user_id"), []byte(".")},
	{[]byte("AwesomeActivity.User.Id"), []byte("awesome_activity.user.id"), []byte(".")},
	{[]byte("AwesomeUsername@Awesome.Com"), []byte("awesome_username@awesome.com"), []byte(".@")},
	{[]byte("lets-ignore all.of dots-and-dashes"), []byte("lets-ignore_all.of_dots-and-dashes"), []byte(".-")},
}

func toSnakeWithIgnore(tb testing.TB) {
	for _, c := range snakeWithIgnoreCases {
		result := bytcase.ToSnakeWithIgnore(c.input, c.ignore)
		if !bytes.Equal(result, c.output) {
			istr := ""
			if c.ignore != nil {
				istr = " ignoring '" + string(c.ignore) + "'"
			}
			tb.Errorf("%q (%q != %q%s)", c.input, result, c.output, istr)
		}
	}
}

func TestToSnakeWithIgnore(t *testing.T) { toSnakeWithIgnore(t) }

func BenchmarkToSnakeWithIgnore(b *testing.B) {
	benchmarkSnakeTest(b, toSnakeWithIgnore)
}

var delimitedCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("test@case")},
	{[]byte("TestCase"), []byte("test@case")},
	{[]byte("Test Case"), []byte("test@case")},
	{[]byte(" Test Case"), []byte("test@case")},
	{[]byte("Test Case "), []byte("test@case")},
	{[]byte(" Test Case "), []byte("test@case")},
	{[]byte("test"), []byte("test")},
	{[]byte("test_case"), []byte("test@case")},
	{[]byte("Test"), []byte("test")},
	{[]byte(""), []byte("")},
	{[]byte("ManyManyWords"), []byte("many@many@words")},
	{[]byte("manyManyWords"), []byte("many@many@words")},
	{[]byte("AnyKind of_string"), []byte("any@kind@of@string")},
	{[]byte("numbers2and55with000"), []byte("numbers@2@and@55@with@000")},
	{[]byte("JSONData"), []byte("json@data")},
	{[]byte("userID"), []byte("user@id")},
	{[]byte("AAAbbb"), []byte("aa@abbb")},
	{[]byte("test-case"), []byte("test@case")},
}

func toDelimited(tb testing.TB) {
	for _, c := range delimitedCases {
		result := bytcase.ToDelimited(c.input, '@')
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToDelimited(t *testing.T) { toDelimited(t) }

func BenchmarkToDelimited(b *testing.B) {
	benchmarkSnakeTest(b, toDelimited)
}

var screamingSnakeCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("TEST_CASE")},
}

func toScreamingSnake(tb testing.TB) {
	for _, c := range screamingSnakeCases {
		result := bytcase.ToScreamingSnake(c.input)
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToScreamingSnake(t *testing.T) { toScreamingSnake(t) }

func BenchmarkToScreamingSnake(b *testing.B) {
	benchmarkSnakeTest(b, toScreamingSnake)
}

var kebabCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("test-case")},
}

func toKebab(tb testing.TB) {
	for _, c := range kebabCases {
		result := bytcase.ToKebab(c.input)
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToKebab(t *testing.T) { toKebab(t) }

func BenchmarkToKebab(b *testing.B) {
	benchmarkSnakeTest(b, toKebab)
}

var screamingKebabCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("TEST-CASE")},
}

func toScreamingKebab(tb testing.TB) {
	for _, c := range screamingKebabCases {
		result := bytcase.ToScreamingKebab(c.input)
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToScreamingKebab(t *testing.T) { toScreamingKebab(t) }

func BenchmarkToScreamingKebab(b *testing.B) {
	benchmarkSnakeTest(b, toScreamingKebab)
}

var screamingDelimitedCases = []struct {
	input  []byte
	output []byte
}{
	{[]byte("testCase"), []byte("TEST.CASE")},
}

func toScreamingDelimited(tb testing.TB) {
	for _, c := range screamingDelimitedCases {
		result := bytcase.ToScreamingDelimited(c.input, '.', nil, true)
		if !bytes.Equal(result, c.output) {
			tb.Errorf("%q (%q != %q)", c.input, result, c.output)
		}
	}
}

func TestToScreamingDelimited(t *testing.T) { toScreamingDelimited(t) }

func BenchmarkToScreamingDelimited(b *testing.B) {
	benchmarkSnakeTest(b, toScreamingDelimited)
}

var screamingDelimitedWithIgnoreCases = []struct {
	input     []byte
	output    []byte
	delimiter byte
	ignore    []byte
}{
	{[]byte("AnyKind of_string"), []byte("ANY.KIND OF.STRING"), '.', []byte(" ")},
}

func toScreamingDelimitedWithIgnore(tb testing.TB) {
	for _, c := range screamingDelimitedWithIgnoreCases {
		result := bytcase.ToScreamingDelimited(c.input, c.delimiter, c.ignore, true)
		if !bytes.Equal(result, c.output) {
			istr := ""
			if c.ignore != nil {
				istr = " ignoring '" + string(c.ignore) + "'"
			}
			tb.Errorf("%q (%q != %q%s)", c.input, result, c.output, istr)
		}
	}
}

func TestToScreamingDelimitedWithIgnore(t *testing.T) { toScreamingDelimitedWithIgnore(t) }

func BenchmarkToScreamingDelimitedWithIgnore(b *testing.B) {
	benchmarkSnakeTest(b, toScreamingDelimitedWithIgnore)
}

func benchmarkSnakeTest(b *testing.B, fn func(testing.TB)) {
	for n := 0; n < b.N; n++ {
		fn(b)
	}
}
