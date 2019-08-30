// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	commands "github.com/pivotal-cf/om/commands"
)

type StemcellArtifacter struct {
	SlugStub        func() string
	slugMutex       sync.RWMutex
	slugArgsForCall []struct {
	}
	slugReturns struct {
		result1 string
	}
	slugReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StemcellArtifacter) Slug() string {
	fake.slugMutex.Lock()
	ret, specificReturn := fake.slugReturnsOnCall[len(fake.slugArgsForCall)]
	fake.slugArgsForCall = append(fake.slugArgsForCall, struct {
	}{})
	fake.recordInvocation("Slug", []interface{}{})
	fake.slugMutex.Unlock()
	if fake.SlugStub != nil {
		return fake.SlugStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.slugReturns
	return fakeReturns.result1
}

func (fake *StemcellArtifacter) SlugCallCount() int {
	fake.slugMutex.RLock()
	defer fake.slugMutex.RUnlock()
	return len(fake.slugArgsForCall)
}

func (fake *StemcellArtifacter) SlugCalls(stub func() string) {
	fake.slugMutex.Lock()
	defer fake.slugMutex.Unlock()
	fake.SlugStub = stub
}

func (fake *StemcellArtifacter) SlugReturns(result1 string) {
	fake.slugMutex.Lock()
	defer fake.slugMutex.Unlock()
	fake.SlugStub = nil
	fake.slugReturns = struct {
		result1 string
	}{result1}
}

func (fake *StemcellArtifacter) SlugReturnsOnCall(i int, result1 string) {
	fake.slugMutex.Lock()
	defer fake.slugMutex.Unlock()
	fake.SlugStub = nil
	if fake.slugReturnsOnCall == nil {
		fake.slugReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.slugReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *StemcellArtifacter) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *StemcellArtifacter) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *StemcellArtifacter) VersionCalls(stub func() string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *StemcellArtifacter) VersionReturns(result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *StemcellArtifacter) VersionReturnsOnCall(i int, result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *StemcellArtifacter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.slugMutex.RLock()
	defer fake.slugMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StemcellArtifacter) recordInvocation(key string, args []interface{}) {
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

var _ commands.StemcellArtifacter = new(StemcellArtifacter)
