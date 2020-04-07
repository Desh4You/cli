// Code generated by counterfeiter. DO NOT EDIT.
package sharedactionfakes

import (
	"context"
	"sync"
	"time"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	client "code.cloudfoundry.org/go-log-cache"
	"code.cloudfoundry.org/go-loggregator/rpc/loggregator_v2"
)

type FakeLogCacheClient struct {
	ReadStub        func(context.Context, string, time.Time, ...client.ReadOption) ([]*loggregator_v2.Envelope, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 time.Time
		arg4 []client.ReadOption
	}
	readReturns struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogCacheClient) Read(arg1 context.Context, arg2 string, arg3 time.Time, arg4 ...client.ReadOption) ([]*loggregator_v2.Envelope, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 time.Time
		arg4 []client.ReadOption
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Read", []interface{}{arg1, arg2, arg3, arg4})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeLogCacheClient) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeLogCacheClient) ReadCalls(stub func(context.Context, string, time.Time, ...client.ReadOption) ([]*loggregator_v2.Envelope, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeLogCacheClient) ReadArgsForCall(i int) (context.Context, string, time.Time, []client.ReadOption) {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeLogCacheClient) ReadReturns(result1 []*loggregator_v2.Envelope, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) ReadReturnsOnCall(i int, result1 []*loggregator_v2.Envelope, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 []*loggregator_v2.Envelope
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 []*loggregator_v2.Envelope
		result2 error
	}{result1, result2}
}

func (fake *FakeLogCacheClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogCacheClient) recordInvocation(key string, args []interface{}) {
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

var _ sharedaction.LogCacheClient = new(FakeLogCacheClient)
