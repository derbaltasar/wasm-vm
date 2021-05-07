package host

import (
	"math/big"

	mock "github.com/ElrondNetwork/arwen-wasm-vm/mock/context"
)

func execESDTTransferAndCallParentMock(instanceMock *mock.InstanceMock, config interface{}) {
	testConfig := config.(directCallGasTestConfig)
	instanceMock.AddMockMethod("execESDTTransferAndCall", func() *mock.InstanceMock {
		host := instanceMock.Host
		instance := mock.GetMockInstance(host)
		host.Metering().UseGas(testConfig.gasUsedByParent)

		arguments := host.Runtime().Arguments()
		if len(arguments) != 3 {
			host.Runtime().SignalUserError("need 3 arguments")
			return instance
		}

		input := DefaultTestContractCallInput()
		input.CallerAddr = host.Runtime().GetSCAddress()
		input.GasProvided = testConfig.gasProvidedToChild
		input.Arguments = [][]byte{
			ESDTTestTokenName,
			big.NewInt(int64(testConfig.ESDTTokensToTransfer)).Bytes(),
			arguments[2],
		}
		input.RecipientAddr = arguments[0]
		input.Function = string(arguments[1])

		_, _, err := host.ExecuteOnDestContext(input)
		if err != nil {
			host.Runtime().FailExecution(err)
		}

		return instance
	})
}