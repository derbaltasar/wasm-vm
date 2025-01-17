package contracts

import (
	"errors"
	"fmt"
	"math/big"

	mock "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/mock/context"
	test "github.com/ElrondNetwork/arwen-wasm-vm/v1_4/testcommon"
	"github.com/ElrondNetwork/elrond-vm-common/txDataBuilder"
)

// WasteGasChildMock is an exposed mock contract method
func WasteGasChildMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(GasTestConfig)
	instanceMock.AddMockMethod("wasteGas", test.SimpleWasteGasMockMethod(instanceMock, testConfig.GetGasUsedByChild()))
}

// FailChildMock is an exposed mock contract method
func FailChildMock(instanceMock *mock.InstanceMock, config interface{}) {
	instanceMock.AddMockMethod("fail", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Runtime().FailExecution(errors.New("forced fail"))
		return instance
	})
}

// FailChildAndBurnESDTMock is an exposed mock contract method
func FailChildAndBurnESDTMock(instanceMock *mock.InstanceMock, config interface{}) {
	instanceMock.AddMockMethod("failAndBurn", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)

		runtime := host.Runtime()

		input := test.DefaultTestContractCallInput()
		input.CallerAddr = runtime.GetSCAddress()
		input.GasProvided = runtime.GetVMInput().GasProvided / 2
		input.Arguments = [][]byte{
			test.ESDTTestTokenName,
			runtime.Arguments()[0],
		}
		input.RecipientAddr = host.Runtime().GetSCAddress()
		input.Function = "ESDTLocalBurn"

		returnValue := ExecuteOnDestContextInMockContracts(host, input)
		if returnValue != 0 {
			host.Runtime().FailExecution(fmt.Errorf("Return value %d", returnValue))
		}

		host.Runtime().FailExecution(errors.New("forced fail"))
		return instance
	})
}

// ExecOnSameCtxParentMock is an exposed mock contract method
func ExecOnSameCtxParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(DirectCallGasTestConfig)
	instanceMock.AddMockMethod("execOnSameCtx", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Metering().UseGas(testConfig.GasUsedByParent)

		argsPerCall := 3
		arguments := host.Runtime().Arguments()
		if len(arguments)%argsPerCall != 0 {
			host.Runtime().SignalUserError("need 3 arguments per individual call")
			return instance
		}

		input := test.DefaultTestContractCallInput()
		input.GasProvided = testConfig.GasProvidedToChild
		input.CallerAddr = instance.Address

		for callIndex := 0; callIndex < len(arguments); callIndex += argsPerCall {
			input.RecipientAddr = arguments[callIndex+0]
			input.Function = string(arguments[callIndex+1])
			numCalls := big.NewInt(0).SetBytes(arguments[callIndex+2]).Uint64()

			for i := uint64(0); i < numCalls; i++ {
				returnValue := ExecuteOnSameContextInMockContracts(host, input)
				if returnValue != 0 {
					host.Runtime().FailExecution(fmt.Errorf("Return value %d", returnValue))
				}
			}
		}

		return instance
	})
}

// ExecOnDestCtxParentMock is an exposed mock contract method
func ExecOnDestCtxParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(DirectCallGasTestConfig)
	instanceMock.AddMockMethod("execOnDestCtx", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Metering().UseGas(testConfig.GasUsedByParent)

		argsPerCall := 3
		arguments := host.Runtime().Arguments()
		if len(arguments)%argsPerCall != 0 {
			host.Runtime().SignalUserError("need 3 arguments per individual call")
			return instance
		}

		input := test.DefaultTestContractCallInput()
		input.GasProvided = testConfig.GasProvidedToChild
		input.CallerAddr = instance.Address

		for callIndex := 0; callIndex < len(arguments); callIndex += argsPerCall {
			input.RecipientAddr = arguments[callIndex+0]
			input.Function = string(arguments[callIndex+1])
			numCalls := big.NewInt(0).SetBytes(arguments[callIndex+2]).Uint64()

			for i := uint64(0); i < numCalls; i++ {
				returnValue := ExecuteOnDestContextInMockContracts(host, input)
				if returnValue != 0 {
					host.Runtime().FailExecution(fmt.Errorf("Return value %d", returnValue))
				}
			}
		}

		return instance
	})
}

// ExecOnDestCtxSingleCallParentMock is an exposed mock contract method
func ExecOnDestCtxSingleCallParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(DirectCallGasTestConfig)
	instanceMock.AddMockMethod("execOnDestCtxSingleCall", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Metering().UseGas(testConfig.GasUsedByParent)

		arguments := host.Runtime().Arguments()
		if len(arguments) != 2 {
			host.Runtime().SignalUserError("need 2 arguments")
			return instance
		}

		input := test.DefaultTestContractCallInput()
		input.GasProvided = testConfig.GasProvidedToChild
		input.CallerAddr = instance.Address

		input.RecipientAddr = arguments[0]
		input.Function = string(arguments[1])

		returnValue := ExecuteOnDestContextInMockContracts(host, input)
		if returnValue != 0 {
			host.Runtime().FailExecution(fmt.Errorf("Return value %d", returnValue))
		}

		return instance
	})
}

// WasteGasParentMock is an exposed mock contract method
func WasteGasParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(DirectCallGasTestConfig)
	instanceMock.AddMockMethod("wasteGas", test.SimpleWasteGasMockMethod(instanceMock, testConfig.GasUsedByParent))
}

const (
	esdtOnCallbackSuccess int = iota
	esdtOnCallbackWrongNumOfArgs
	esdtOnCallbackFail
)

// ESDTTransferToParentMock is an exposed mock contract method
func ESDTTransferToParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	esdtTransferToParentMock(instanceMock, config, esdtOnCallbackSuccess)
}

// ESDTTransferToParentWrongESDTArgsNumberMock is an exposed mock contract method
func ESDTTransferToParentWrongESDTArgsNumberMock(instanceMock *mock.InstanceMock, config interface{}) {
	esdtTransferToParentMock(instanceMock, config, esdtOnCallbackWrongNumOfArgs)
}

// ESDTTransferToParentCallbackWillFail is an exposed mock contract method
func ESDTTransferToParentCallbackWillFail(instanceMock *mock.InstanceMock, config interface{}) {
	esdtTransferToParentMock(instanceMock, config, esdtOnCallbackFail)
}

func esdtTransferToParentMock(instanceMock *mock.InstanceMock, config interface{}, behavior int) {
	testConfig := config.(*AsyncCallTestConfig)
	instanceMock.AddMockMethod("transferESDTToParent", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Metering().UseGas(testConfig.GasUsedByParent)

		callData := txDataBuilder.NewBuilder()
		callData.Func(string("ESDTTransfer"))
		callData.Bytes(test.ESDTTestTokenName)
		callData.Bytes(big.NewInt(int64(testConfig.CallbackESDTTokensToTransfer)).Bytes())

		switch behavior {
		case esdtOnCallbackSuccess:
			host.Output().Finish([]byte("success"))
		case esdtOnCallbackWrongNumOfArgs:
			callData.Bytes([]byte{})
		case esdtOnCallbackFail:
			host.Output().Finish([]byte("fail"))
		}

		value := big.NewInt(0).Bytes()

		err := host.Runtime().ExecuteAsyncCall(test.ParentAddress, callData.ToBytes(), value)

		if err != nil {
			host.Runtime().FailExecution(err)
		}

		return instance
	})
}
