// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	api "github.com/pivotal-cf/om/api"
)

type CurlService struct {
	CurlStub        func(api.RequestServiceCurlInput) (api.RequestServiceCurlOutput, error)
	curlMutex       sync.RWMutex
	curlArgsForCall []struct {
		arg1 api.RequestServiceCurlInput
	}
	curlReturns struct {
		result1 api.RequestServiceCurlOutput
		result2 error
	}
	curlReturnsOnCall map[int]struct {
		result1 api.RequestServiceCurlOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CurlService) Curl(arg1 api.RequestServiceCurlInput) (api.RequestServiceCurlOutput, error) {
	fake.curlMutex.Lock()
	ret, specificReturn := fake.curlReturnsOnCall[len(fake.curlArgsForCall)]
	fake.curlArgsForCall = append(fake.curlArgsForCall, struct {
		arg1 api.RequestServiceCurlInput
	}{arg1})
	fake.recordInvocation("Curl", []interface{}{arg1})
	fake.curlMutex.Unlock()
	if fake.CurlStub != nil {
		return fake.CurlStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.curlReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CurlService) CurlCallCount() int {
	fake.curlMutex.RLock()
	defer fake.curlMutex.RUnlock()
	return len(fake.curlArgsForCall)
}

func (fake *CurlService) CurlCalls(stub func(api.RequestServiceCurlInput) (api.RequestServiceCurlOutput, error)) {
	fake.curlMutex.Lock()
	defer fake.curlMutex.Unlock()
	fake.CurlStub = stub
}

func (fake *CurlService) CurlArgsForCall(i int) api.RequestServiceCurlInput {
	fake.curlMutex.RLock()
	defer fake.curlMutex.RUnlock()
	argsForCall := fake.curlArgsForCall[i]
	return argsForCall.arg1
}

func (fake *CurlService) CurlReturns(result1 api.RequestServiceCurlOutput, result2 error) {
	fake.curlMutex.Lock()
	defer fake.curlMutex.Unlock()
	fake.CurlStub = nil
	fake.curlReturns = struct {
		result1 api.RequestServiceCurlOutput
		result2 error
	}{result1, result2}
}

func (fake *CurlService) CurlReturnsOnCall(i int, result1 api.RequestServiceCurlOutput, result2 error) {
	fake.curlMutex.Lock()
	defer fake.curlMutex.Unlock()
	fake.CurlStub = nil
	if fake.curlReturnsOnCall == nil {
		fake.curlReturnsOnCall = make(map[int]struct {
			result1 api.RequestServiceCurlOutput
			result2 error
		})
	}
	fake.curlReturnsOnCall[i] = struct {
		result1 api.RequestServiceCurlOutput
		result2 error
	}{result1, result2}
}

func (fake *CurlService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.curlMutex.RLock()
	defer fake.curlMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CurlService) recordInvocation(key string, args []interface{}) {
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
