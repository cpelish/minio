// Copyright (c) 2015-2021 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package condition

import (
	"reflect"
	"testing"
)

func TestStringEqualsFuncEvaluate(t *testing.T) {
	case1Function, err := newStringEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		function       Function
		values         map[string][]string
		expectedResult bool
	}{
		{case1Function, map[string][]string{"x-amz-copy-source": {"mybucket/myobject"}}, true},
		{case1Function, map[string][]string{"x-amz-copy-source": {"yourbucket/myobject"}}, false},
		{case1Function, map[string][]string{}, false},
		{case1Function, map[string][]string{"delimiter": {"/"}}, false},

		{case2Function, map[string][]string{"x-amz-server-side-encryption": {"AES256"}}, true},
		{case2Function, map[string][]string{}, false},
		{case2Function, map[string][]string{"delimiter": {"/"}}, false},

		{case3Function, map[string][]string{"x-amz-metadata-directive": {"REPLACE"}}, true},
		{case3Function, map[string][]string{"x-amz-metadata-directive": {"COPY"}}, false},
		{case3Function, map[string][]string{}, false},
		{case3Function, map[string][]string{"delimiter": {"/"}}, false},

		{case4Function, map[string][]string{"LocationConstraint": {"eu-west-1"}}, true},
		{case4Function, map[string][]string{"LocationConstraint": {"us-east-1"}}, false},
		{case4Function, map[string][]string{}, false},
		{case4Function, map[string][]string{"delimiter": {"/"}}, false},
	}

	for i, testCase := range testCases {
		result := testCase.function.evaluate(testCase.values)

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestStringEqualsFuncKey(t *testing.T) {
	case1Function, err := newStringEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		function       Function
		expectedResult Key
	}{
		{case1Function, S3XAmzCopySource},
		{case2Function, S3XAmzServerSideEncryption},
		{case3Function, S3XAmzMetadataDirective},
		{case4Function, S3LocationConstraint},
	}

	for i, testCase := range testCases {
		result := testCase.function.key()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestStringEqualsFuncToMap(t *testing.T) {
	case1Function, err := newStringEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case1Result := map[Key]ValueSet{
		S3XAmzCopySource: NewValueSet(NewStringValue("mybucket/myobject")),
	}

	case2Function, err := newStringEqualsFunc(S3XAmzCopySource,
		NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Result := map[Key]ValueSet{
		S3XAmzCopySource: NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	}

	case3Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Result := map[Key]ValueSet{
		S3XAmzServerSideEncryption: NewValueSet(NewStringValue("AES256")),
	}

	case4Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption,
		NewValueSet(
			NewStringValue("AES256"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Result := map[Key]ValueSet{
		S3XAmzServerSideEncryption: NewValueSet(
			NewStringValue("AES256"),
		),
	}

	case5Function, err := newStringEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case5Result := map[Key]ValueSet{
		S3XAmzMetadataDirective: NewValueSet(NewStringValue("REPLACE")),
	}

	case6Function, err := newStringEqualsFunc(S3XAmzMetadataDirective,
		NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case6Result := map[Key]ValueSet{
		S3XAmzMetadataDirective: NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	}

	case7Function, err := newStringEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case7Result := map[Key]ValueSet{
		S3LocationConstraint: NewValueSet(NewStringValue("eu-west-1")),
	}

	case8Function, err := newStringEqualsFunc(S3LocationConstraint,
		NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case8Result := map[Key]ValueSet{
		S3LocationConstraint: NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	}

	testCases := []struct {
		f              Function
		expectedResult map[Key]ValueSet
	}{
		{case1Function, case1Result},
		{case2Function, case2Result},
		{case3Function, case3Result},
		{case4Function, case4Result},
		{case5Function, case5Result},
		{case6Function, case6Result},
		{case7Function, case7Result},
		{case8Function, case8Result},
		{&stringEqualsFunc{}, nil},
	}

	for i, testCase := range testCases {
		result := testCase.f.toMap()

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestStringNotEqualsFuncEvaluate(t *testing.T) {
	case1Function, err := newStringNotEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringNotEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		function       Function
		values         map[string][]string
		expectedResult bool
	}{
		{case1Function, map[string][]string{"x-amz-copy-source": {"mybucket/myobject"}}, false},
		{case1Function, map[string][]string{"x-amz-copy-source": {"yourbucket/myobject"}}, true},
		{case1Function, map[string][]string{}, true},
		{case1Function, map[string][]string{"delimiter": {"/"}}, true},

		{case2Function, map[string][]string{"x-amz-server-side-encryption": {"AES256"}}, false},
		{case2Function, map[string][]string{}, true},
		{case2Function, map[string][]string{"delimiter": {"/"}}, true},

		{case3Function, map[string][]string{"x-amz-metadata-directive": {"REPLACE"}}, false},
		{case3Function, map[string][]string{"x-amz-metadata-directive": {"COPY"}}, true},
		{case3Function, map[string][]string{}, true},
		{case3Function, map[string][]string{"delimiter": {"/"}}, true},

		{case4Function, map[string][]string{"LocationConstraint": {"eu-west-1"}}, false},
		{case4Function, map[string][]string{"LocationConstraint": {"us-east-1"}}, true},
		{case4Function, map[string][]string{}, true},
		{case4Function, map[string][]string{"delimiter": {"/"}}, true},
	}

	for i, testCase := range testCases {
		result := testCase.function.evaluate(testCase.values)

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestStringNotEqualsFuncKey(t *testing.T) {
	case1Function, err := newStringNotEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringNotEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		function       Function
		expectedResult Key
	}{
		{case1Function, S3XAmzCopySource},
		{case2Function, S3XAmzServerSideEncryption},
		{case3Function, S3XAmzMetadataDirective},
		{case4Function, S3LocationConstraint},
	}

	for i, testCase := range testCases {
		result := testCase.function.key()

		if result != testCase.expectedResult {
			t.Fatalf("case %v: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestStringNotEqualsFuncToMap(t *testing.T) {
	case1Function, err := newStringNotEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case1Result := map[Key]ValueSet{
		S3XAmzCopySource: NewValueSet(NewStringValue("mybucket/myobject")),
	}

	case2Function, err := newStringNotEqualsFunc(S3XAmzCopySource,
		NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Result := map[Key]ValueSet{
		S3XAmzCopySource: NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	}

	case3Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Result := map[Key]ValueSet{
		S3XAmzServerSideEncryption: NewValueSet(NewStringValue("AES256")),
	}

	case4Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption,
		NewValueSet(
			NewStringValue("AES256"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Result := map[Key]ValueSet{
		S3XAmzServerSideEncryption: NewValueSet(
			NewStringValue("AES256"),
		),
	}

	case5Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case5Result := map[Key]ValueSet{
		S3XAmzMetadataDirective: NewValueSet(NewStringValue("REPLACE")),
	}

	case6Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective,
		NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case6Result := map[Key]ValueSet{
		S3XAmzMetadataDirective: NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	}

	case7Function, err := newStringNotEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case7Result := map[Key]ValueSet{
		S3LocationConstraint: NewValueSet(NewStringValue("eu-west-1")),
	}

	case8Function, err := newStringNotEqualsFunc(S3LocationConstraint,
		NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case8Result := map[Key]ValueSet{
		S3LocationConstraint: NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	}

	testCases := []struct {
		f              Function
		expectedResult map[Key]ValueSet
	}{
		{case1Function, case1Result},
		{case2Function, case2Result},
		{case3Function, case3Result},
		{case4Function, case4Result},
		{case5Function, case5Result},
		{case6Function, case6Result},
		{case7Function, case7Result},
		{case8Function, case8Result},
		{&stringNotEqualsFunc{}, nil},
	}

	for i, testCase := range testCases {
		result := testCase.f.toMap()

		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
		}
	}
}

func TestNewStringEqualsFunc(t *testing.T) {
	case1Function, err := newStringEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringEqualsFunc(S3XAmzCopySource,
		NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringEqualsFunc(S3XAmzServerSideEncryption,
		NewValueSet(
			NewStringValue("AES256"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case5Function, err := newStringEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case6Function, err := newStringEqualsFunc(S3XAmzMetadataDirective,
		NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case7Function, err := newStringEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case8Function, err := newStringEqualsFunc(S3LocationConstraint,
		NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		key            Key
		values         ValueSet
		expectedResult Function
		expectErr      bool
	}{
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")), case1Function, false},
		{S3XAmzCopySource,
			NewValueSet(
				NewStringValue("mybucket/myobject"),
				NewStringValue("yourbucket/myobject"),
			), case2Function, false},

		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")), case3Function, false},
		{S3XAmzServerSideEncryption,
			NewValueSet(
				NewStringValue("AES256"),
			), case4Function, false},

		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")), case5Function, false},
		{S3XAmzMetadataDirective,
			NewValueSet(
				NewStringValue("REPLACE"),
				NewStringValue("COPY"),
			), case6Function, false},

		{S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")), case7Function, false},
		{S3LocationConstraint,
			NewValueSet(
				NewStringValue("eu-west-1"),
				NewStringValue("us-west-1"),
			), case8Function, false},

		// Unsupported value error.
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject"), NewIntValue(7)), nil, true},
		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256"), NewIntValue(7)), nil, true},
		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE"), NewIntValue(7)), nil, true},
		{S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1"), NewIntValue(7)), nil, true},

		// Invalid value error.
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket")), nil, true},
		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("SSE-C")), nil, true},
		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("DUPLICATE")), nil, true},
	}

	for i, testCase := range testCases {
		result, err := newStringEqualsFunc(testCase.key, testCase.values)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}

func TestNewStringNotEqualsFunc(t *testing.T) {
	case1Function, err := newStringNotEqualsFunc(S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case2Function, err := newStringNotEqualsFunc(S3XAmzCopySource,
		NewValueSet(
			NewStringValue("mybucket/myobject"),
			NewStringValue("yourbucket/myobject"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case3Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case4Function, err := newStringNotEqualsFunc(S3XAmzServerSideEncryption,
		NewValueSet(
			NewStringValue("AES256"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case5Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case6Function, err := newStringNotEqualsFunc(S3XAmzMetadataDirective,
		NewValueSet(
			NewStringValue("REPLACE"),
			NewStringValue("COPY"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case7Function, err := newStringNotEqualsFunc(S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")))
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	case8Function, err := newStringNotEqualsFunc(S3LocationConstraint,
		NewValueSet(
			NewStringValue("eu-west-1"),
			NewStringValue("us-west-1"),
		),
	)
	if err != nil {
		t.Fatalf("unexpected error. %v\n", err)
	}

	testCases := []struct {
		key            Key
		values         ValueSet
		expectedResult Function
		expectErr      bool
	}{
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject")), case1Function, false},
		{S3XAmzCopySource,
			NewValueSet(
				NewStringValue("mybucket/myobject"),
				NewStringValue("yourbucket/myobject"),
			), case2Function, false},

		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256")), case3Function, false},
		{S3XAmzServerSideEncryption,
			NewValueSet(
				NewStringValue("AES256"),
			), case4Function, false},

		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE")), case5Function, false},
		{S3XAmzMetadataDirective,
			NewValueSet(
				NewStringValue("REPLACE"),
				NewStringValue("COPY"),
			), case6Function, false},

		{S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1")), case7Function, false},
		{S3LocationConstraint,
			NewValueSet(
				NewStringValue("eu-west-1"),
				NewStringValue("us-west-1"),
			), case8Function, false},

		// Unsupported value error.
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket/myobject"), NewIntValue(7)), nil, true},
		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("AES256"), NewIntValue(7)), nil, true},
		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("REPLACE"), NewIntValue(7)), nil, true},
		{S3LocationConstraint, NewValueSet(NewStringValue("eu-west-1"), NewIntValue(7)), nil, true},

		// Invalid value error.
		{S3XAmzCopySource, NewValueSet(NewStringValue("mybucket")), nil, true},
		{S3XAmzServerSideEncryption, NewValueSet(NewStringValue("SSE-C")), nil, true},
		{S3XAmzMetadataDirective, NewValueSet(NewStringValue("DUPLICATE")), nil, true},
	}

	for i, testCase := range testCases {
		result, err := newStringNotEqualsFunc(testCase.key, testCase.values)
		expectErr := (err != nil)

		if expectErr != testCase.expectErr {
			t.Fatalf("case %v: error: expected: %v, got: %v\n", i+1, testCase.expectErr, expectErr)
		}

		if !testCase.expectErr {
			if !reflect.DeepEqual(result, testCase.expectedResult) {
				t.Fatalf("case %v: result: expected: %v, got: %v\n", i+1, testCase.expectedResult, result)
			}
		}
	}
}
