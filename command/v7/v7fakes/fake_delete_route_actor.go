// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDeleteRouteActor struct {
	DeleteRouteStub        func(string, string, string) (v7action.Warnings, error)
	deleteRouteMutex       sync.RWMutex
	deleteRouteArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	deleteRouteReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	deleteRouteReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteRouteActor) DeleteRoute(arg1 string, arg2 string, arg3 string) (v7action.Warnings, error) {
	fake.deleteRouteMutex.Lock()
	ret, specificReturn := fake.deleteRouteReturnsOnCall[len(fake.deleteRouteArgsForCall)]
	fake.deleteRouteArgsForCall = append(fake.deleteRouteArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("DeleteRoute", []interface{}{arg1, arg2, arg3})
	fake.deleteRouteMutex.Unlock()
	if fake.DeleteRouteStub != nil {
		return fake.DeleteRouteStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteRouteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDeleteRouteActor) DeleteRouteCallCount() int {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	return len(fake.deleteRouteArgsForCall)
}

func (fake *FakeDeleteRouteActor) DeleteRouteCalls(stub func(string, string, string) (v7action.Warnings, error)) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = stub
}

func (fake *FakeDeleteRouteActor) DeleteRouteArgsForCall(i int) (string, string, string) {
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	argsForCall := fake.deleteRouteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeDeleteRouteActor) DeleteRouteReturns(result1 v7action.Warnings, result2 error) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = nil
	fake.deleteRouteReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteRouteActor) DeleteRouteReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.deleteRouteMutex.Lock()
	defer fake.deleteRouteMutex.Unlock()
	fake.DeleteRouteStub = nil
	if fake.deleteRouteReturnsOnCall == nil {
		fake.deleteRouteReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.deleteRouteReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeDeleteRouteActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteRouteMutex.RLock()
	defer fake.deleteRouteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteRouteActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DeleteRouteActor = new(FakeDeleteRouteActor)
