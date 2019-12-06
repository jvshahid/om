// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type ProductDiffService struct {
	ProductDiffStub        func(string) (api.ProductDiff, error)
	productDiffMutex       sync.RWMutex
	productDiffArgsForCall []struct {
		arg1 string
	}
	productDiffReturns struct {
		result1 api.ProductDiff
		result2 error
	}
	productDiffReturnsOnCall map[int]struct {
		result1 api.ProductDiff
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProductDiffService) ProductDiff(arg1 string) (api.ProductDiff, error) {
	fake.productDiffMutex.Lock()
	ret, specificReturn := fake.productDiffReturnsOnCall[len(fake.productDiffArgsForCall)]
	fake.productDiffArgsForCall = append(fake.productDiffArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ProductDiff", []interface{}{arg1})
	fake.productDiffMutex.Unlock()
	if fake.ProductDiffStub != nil {
		return fake.ProductDiffStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.productDiffReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ProductDiffService) ProductDiffCallCount() int {
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	return len(fake.productDiffArgsForCall)
}

func (fake *ProductDiffService) ProductDiffCalls(stub func(string) (api.ProductDiff, error)) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = stub
}

func (fake *ProductDiffService) ProductDiffArgsForCall(i int) string {
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	argsForCall := fake.productDiffArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ProductDiffService) ProductDiffReturns(result1 api.ProductDiff, result2 error) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = nil
	fake.productDiffReturns = struct {
		result1 api.ProductDiff
		result2 error
	}{result1, result2}
}

func (fake *ProductDiffService) ProductDiffReturnsOnCall(i int, result1 api.ProductDiff, result2 error) {
	fake.productDiffMutex.Lock()
	defer fake.productDiffMutex.Unlock()
	fake.ProductDiffStub = nil
	if fake.productDiffReturnsOnCall == nil {
		fake.productDiffReturnsOnCall = make(map[int]struct {
			result1 api.ProductDiff
			result2 error
		})
	}
	fake.productDiffReturnsOnCall[i] = struct {
		result1 api.ProductDiff
		result2 error
	}{result1, result2}
}

func (fake *ProductDiffService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.productDiffMutex.RLock()
	defer fake.productDiffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProductDiffService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
