// Code generated by counterfeiter. DO NOT EDIT.
package pushactionfakes

import (
	sync "sync"

	pushaction "code.cloudfoundry.org/cli/actor/pushaction"
)

type FakeRandomWordGenerator struct {
	BabbleStub        func() string
	babbleMutex       sync.RWMutex
	babbleArgsForCall []struct {
	}
	babbleReturns struct {
		result1 string
	}
	babbleReturnsOnCall map[int]struct {
		result1 string
	}
	RandomAdjectiveStub        func() string
	randomAdjectiveMutex       sync.RWMutex
	randomAdjectiveArgsForCall []struct {
	}
	randomAdjectiveReturns struct {
		result1 string
	}
	randomAdjectiveReturnsOnCall map[int]struct {
		result1 string
	}
	RandomNounStub        func() string
	randomNounMutex       sync.RWMutex
	randomNounArgsForCall []struct {
	}
	randomNounReturns struct {
		result1 string
	}
	randomNounReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRandomWordGenerator) Babble() string {
	fake.babbleMutex.Lock()
	ret, specificReturn := fake.babbleReturnsOnCall[len(fake.babbleArgsForCall)]
	fake.babbleArgsForCall = append(fake.babbleArgsForCall, struct {
	}{})
	fake.recordInvocation("Babble", []interface{}{})
	fake.babbleMutex.Unlock()
	if fake.BabbleStub != nil {
		return fake.BabbleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.babbleReturns
	return fakeReturns.result1
}

func (fake *FakeRandomWordGenerator) BabbleCallCount() int {
	fake.babbleMutex.RLock()
	defer fake.babbleMutex.RUnlock()
	return len(fake.babbleArgsForCall)
}

func (fake *FakeRandomWordGenerator) BabbleCalls(stub func() string) {
	fake.babbleMutex.Lock()
	defer fake.babbleMutex.Unlock()
	fake.BabbleStub = stub
}

func (fake *FakeRandomWordGenerator) BabbleReturns(result1 string) {
	fake.babbleMutex.Lock()
	defer fake.babbleMutex.Unlock()
	fake.BabbleStub = nil
	fake.babbleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) BabbleReturnsOnCall(i int, result1 string) {
	fake.babbleMutex.Lock()
	defer fake.babbleMutex.Unlock()
	fake.BabbleStub = nil
	if fake.babbleReturnsOnCall == nil {
		fake.babbleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.babbleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) RandomAdjective() string {
	fake.randomAdjectiveMutex.Lock()
	ret, specificReturn := fake.randomAdjectiveReturnsOnCall[len(fake.randomAdjectiveArgsForCall)]
	fake.randomAdjectiveArgsForCall = append(fake.randomAdjectiveArgsForCall, struct {
	}{})
	fake.recordInvocation("RandomAdjective", []interface{}{})
	fake.randomAdjectiveMutex.Unlock()
	if fake.RandomAdjectiveStub != nil {
		return fake.RandomAdjectiveStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.randomAdjectiveReturns
	return fakeReturns.result1
}

func (fake *FakeRandomWordGenerator) RandomAdjectiveCallCount() int {
	fake.randomAdjectiveMutex.RLock()
	defer fake.randomAdjectiveMutex.RUnlock()
	return len(fake.randomAdjectiveArgsForCall)
}

func (fake *FakeRandomWordGenerator) RandomAdjectiveCalls(stub func() string) {
	fake.randomAdjectiveMutex.Lock()
	defer fake.randomAdjectiveMutex.Unlock()
	fake.RandomAdjectiveStub = stub
}

func (fake *FakeRandomWordGenerator) RandomAdjectiveReturns(result1 string) {
	fake.randomAdjectiveMutex.Lock()
	defer fake.randomAdjectiveMutex.Unlock()
	fake.RandomAdjectiveStub = nil
	fake.randomAdjectiveReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) RandomAdjectiveReturnsOnCall(i int, result1 string) {
	fake.randomAdjectiveMutex.Lock()
	defer fake.randomAdjectiveMutex.Unlock()
	fake.RandomAdjectiveStub = nil
	if fake.randomAdjectiveReturnsOnCall == nil {
		fake.randomAdjectiveReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.randomAdjectiveReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) RandomNoun() string {
	fake.randomNounMutex.Lock()
	ret, specificReturn := fake.randomNounReturnsOnCall[len(fake.randomNounArgsForCall)]
	fake.randomNounArgsForCall = append(fake.randomNounArgsForCall, struct {
	}{})
	fake.recordInvocation("RandomNoun", []interface{}{})
	fake.randomNounMutex.Unlock()
	if fake.RandomNounStub != nil {
		return fake.RandomNounStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.randomNounReturns
	return fakeReturns.result1
}

func (fake *FakeRandomWordGenerator) RandomNounCallCount() int {
	fake.randomNounMutex.RLock()
	defer fake.randomNounMutex.RUnlock()
	return len(fake.randomNounArgsForCall)
}

func (fake *FakeRandomWordGenerator) RandomNounCalls(stub func() string) {
	fake.randomNounMutex.Lock()
	defer fake.randomNounMutex.Unlock()
	fake.RandomNounStub = stub
}

func (fake *FakeRandomWordGenerator) RandomNounReturns(result1 string) {
	fake.randomNounMutex.Lock()
	defer fake.randomNounMutex.Unlock()
	fake.RandomNounStub = nil
	fake.randomNounReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) RandomNounReturnsOnCall(i int, result1 string) {
	fake.randomNounMutex.Lock()
	defer fake.randomNounMutex.Unlock()
	fake.RandomNounStub = nil
	if fake.randomNounReturnsOnCall == nil {
		fake.randomNounReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.randomNounReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRandomWordGenerator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.babbleMutex.RLock()
	defer fake.babbleMutex.RUnlock()
	fake.randomAdjectiveMutex.RLock()
	defer fake.randomAdjectiveMutex.RUnlock()
	fake.randomNounMutex.RLock()
	defer fake.randomNounMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRandomWordGenerator) recordInvocation(key string, args []interface{}) {
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

var _ pushaction.RandomWordGenerator = new(FakeRandomWordGenerator)
