// Code generated by pegomock. DO NOT EDIT.
package matchers

import (
	"reflect"
	"github.com/petergtz/pegomock"
	gits "github.com/jenkins-x/jx/pkg/gits"
)

func AnySliceOfPtrToGitsGitWebHookArguments() []*gits.GitWebHookArguments {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*([]*gits.GitWebHookArguments))(nil)).Elem()))
	var nullValue []*gits.GitWebHookArguments
	return nullValue
}

func EqSliceOfPtrToGitsGitWebHookArguments(value []*gits.GitWebHookArguments) []*gits.GitWebHookArguments {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue []*gits.GitWebHookArguments
	return nullValue
}
