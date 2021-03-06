// This file was generated by counterfeiter
package dbfakes

import (
	"sync"

	"github.com/concourse/atc/db"
)

type FakePipelinesDB struct {
	GetAllPublicPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPublicPipelinesMutex       sync.RWMutex
	getAllPublicPipelinesArgsForCall []struct{}
	getAllPublicPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePipelinesDB) GetAllPublicPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPublicPipelinesMutex.Lock()
	fake.getAllPublicPipelinesArgsForCall = append(fake.getAllPublicPipelinesArgsForCall, struct{}{})
	fake.recordInvocation("GetAllPublicPipelines", []interface{}{})
	fake.getAllPublicPipelinesMutex.Unlock()
	if fake.GetAllPublicPipelinesStub != nil {
		return fake.GetAllPublicPipelinesStub()
	} else {
		return fake.getAllPublicPipelinesReturns.result1, fake.getAllPublicPipelinesReturns.result2
	}
}

func (fake *FakePipelinesDB) GetAllPublicPipelinesCallCount() int {
	fake.getAllPublicPipelinesMutex.RLock()
	defer fake.getAllPublicPipelinesMutex.RUnlock()
	return len(fake.getAllPublicPipelinesArgsForCall)
}

func (fake *FakePipelinesDB) GetAllPublicPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPublicPipelinesStub = nil
	fake.getAllPublicPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakePipelinesDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAllPublicPipelinesMutex.RLock()
	defer fake.getAllPublicPipelinesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePipelinesDB) recordInvocation(key string, args []interface{}) {
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

var _ db.PipelinesDB = new(FakePipelinesDB)
