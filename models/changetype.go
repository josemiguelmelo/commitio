package models

import (
	"errors"
	"fmt"
	"reflect"
)

type ChangeType string

const (
	FeatureType       ChangeType = "feature"
	FixType           ChangeType = "fix"
	DocumentationType ChangeType = "docs"
	TestType          ChangeType = "test"
)

const DefaultType = FixType

var ChangeTypeEnum map[string]ChangeType = map[string]ChangeType{
	string(FeatureType):       FeatureType,
	string(FixType):           FixType,
	string(DocumentationType): DocumentationType,
	string(TestType):          TestType,
}

func ChangeTypeFromString(v string) (*ChangeType, error) {
	if val, exists := ChangeTypeEnum[v]; exists {
		return &val, nil
	}

	errorMsg := fmt.Sprintf("Change type %s not supported", v)
	return nil, errors.New(errorMsg)
}

func ChangeTypeValuesAsString() []string {
	values := []string{}
	for _, v := range reflect.ValueOf(ChangeTypeEnum).MapKeys() {
		values = append(values, v.String())
	}
	return values
}
