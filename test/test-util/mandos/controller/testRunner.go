package mandoscontroller

import (
	fr "github.com/ElrondNetwork/arwen-wasm-vm/test/test-util/mandos/json/fileresolver"
	mj "github.com/ElrondNetwork/arwen-wasm-vm/test/test-util/mandos/json/model"
	mjparse "github.com/ElrondNetwork/arwen-wasm-vm/test/test-util/mandos/json/parse"
)

// TestExecutor describes a component that can run a VM test.
type TestExecutor interface {
	// ExecuteTest executes the test and checks if it passed. Failure is signaled by returning an error.
	ExecuteTest(*mj.Test) error
}

// TestRunner is a component that can run tests, using a provided executor.
type TestRunner struct {
	Executor TestExecutor
	Parser   mjparse.Parser
}

// NewTestRunner creates new TestRunner instance.
func NewTestRunner(executor TestExecutor, fileResolver fr.FileResolver) *TestRunner {
	return &TestRunner{
		Executor: executor,
		Parser:   mjparse.NewParser(fileResolver),
	}
}
