// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type UpdateSSLCertificateService struct {
	UpdateSSLCertificateStub        func(api.SSLCertificateInput) error
	updateSSLCertificateMutex       sync.RWMutex
	updateSSLCertificateArgsForCall []struct {
		arg1 api.SSLCertificateInput
	}
	updateSSLCertificateReturns struct {
		result1 error
	}
	updateSSLCertificateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UpdateSSLCertificateService) UpdateSSLCertificate(arg1 api.SSLCertificateInput) error {
	fake.updateSSLCertificateMutex.Lock()
	ret, specificReturn := fake.updateSSLCertificateReturnsOnCall[len(fake.updateSSLCertificateArgsForCall)]
	fake.updateSSLCertificateArgsForCall = append(fake.updateSSLCertificateArgsForCall, struct {
		arg1 api.SSLCertificateInput
	}{arg1})
	fake.recordInvocation("UpdateSSLCertificate", []interface{}{arg1})
	fake.updateSSLCertificateMutex.Unlock()
	if fake.UpdateSSLCertificateStub != nil {
		return fake.UpdateSSLCertificateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateSSLCertificateReturns.result1
}

func (fake *UpdateSSLCertificateService) UpdateSSLCertificateCallCount() int {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	return len(fake.updateSSLCertificateArgsForCall)
}

func (fake *UpdateSSLCertificateService) UpdateSSLCertificateArgsForCall(i int) api.SSLCertificateInput {
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	return fake.updateSSLCertificateArgsForCall[i].arg1
}

func (fake *UpdateSSLCertificateService) UpdateSSLCertificateReturns(result1 error) {
	fake.UpdateSSLCertificateStub = nil
	fake.updateSSLCertificateReturns = struct {
		result1 error
	}{result1}
}

func (fake *UpdateSSLCertificateService) UpdateSSLCertificateReturnsOnCall(i int, result1 error) {
	fake.UpdateSSLCertificateStub = nil
	if fake.updateSSLCertificateReturnsOnCall == nil {
		fake.updateSSLCertificateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateSSLCertificateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *UpdateSSLCertificateService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.updateSSLCertificateMutex.RLock()
	defer fake.updateSSLCertificateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *UpdateSSLCertificateService) recordInvocation(key string, args []interface{}) {
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
