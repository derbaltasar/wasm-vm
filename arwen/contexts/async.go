package contexts

import (
	"encoding/json"
	"math/big"

	"github.com/ElrondNetwork/arwen-wasm-vm/arwen"
	extramath "github.com/ElrondNetwork/arwen-wasm-vm/math"
	"github.com/ElrondNetwork/elrond-go/core/vmcommon"
)

var _ arwen.AsyncContext = (*asyncContext)(nil)

type asyncContext struct {
	host       arwen.VMHost
	stateStack []*asyncContext

	CallerAddr      []byte
	GasPrice        uint64
	ReturnData      []byte
	AsyncCallGroups []*arwen.AsyncCallGroup
}

// NewAsyncContext creates a new asyncContext.
func NewAsyncContext(host arwen.VMHost) *asyncContext {
	return &asyncContext{
		host:            host,
		stateStack:      nil,
		CallerAddr:      nil,
		ReturnData:      nil,
		AsyncCallGroups: make([]*arwen.AsyncCallGroup, 0),
	}
}

// InitState initializes the internal state of the AsyncContext.
func (context *asyncContext) InitState() {
	context.CallerAddr = make([]byte, 0)
	context.ReturnData = make([]byte, 0)
	context.AsyncCallGroups = make([]*arwen.AsyncCallGroup, 0)
}

// InitStateFromInput initializes the internal state of the AsyncContext with
// information provided by a ContractCallInput.
func (context *asyncContext) InitStateFromInput(input *vmcommon.ContractCallInput) {
	context.InitState()
	context.SetCaller(input.CallerAddr)
	context.SetGasPrice(input.GasPrice)
}

// SetCaller sets the address of the original caller.
func (context *asyncContext) SetCaller(caller []byte) {
	context.CallerAddr = caller
}

// SetGasPrice sets the gas price.
func (context *asyncContext) SetGasPrice(gasPrice uint64) {
	context.GasPrice = gasPrice
}

// PushState creates a deep clone of the internal state and pushes it onto the
// internal state stack.
func (context *asyncContext) PushState() {
	newState := &asyncContext{
		CallerAddr:      context.CallerAddr,
		GasPrice:        context.GasPrice,
		ReturnData:      context.ReturnData,
		AsyncCallGroups: context.cloneCallGroups(),
	}

	context.stateStack = append(context.stateStack, newState)
}

func (context *asyncContext) cloneCallGroups() []*arwen.AsyncCallGroup {
	groupCount := len(context.AsyncCallGroups)
	clonedGroups := make([]*arwen.AsyncCallGroup, groupCount)

	for i := 0; i < groupCount; i++ {
		clonedGroups[i] = context.AsyncCallGroups[i].Clone()
	}

	return clonedGroups
}

// PopDiscard is a no-operation for the AsyncContext.
func (context *asyncContext) PopDiscard() {
}

// PopSetActiveState pops the state found at the top of the internal state
// stack and sets it as the 'active' state of the AsyncContext.
func (context *asyncContext) PopSetActiveState() {
	stateStackLen := len(context.stateStack)
	if stateStackLen == 0 {
		return
	}

	prevState := context.stateStack[stateStackLen-1]
	context.stateStack = context.stateStack[:stateStackLen-1]

	context.CallerAddr = prevState.CallerAddr
	context.GasPrice = prevState.GasPrice
	context.ReturnData = prevState.ReturnData
	context.AsyncCallGroups = prevState.AsyncCallGroups
}

// PopMergeActiveState is a no-operation for the AsyncContext.
func (context *asyncContext) PopMergeActiveState() {
}

// ClearStateStack deletes all the states stored on the internal state stack.
func (context *asyncContext) ClearStateStack() {
	context.stateStack = make([]*asyncContext, 0)
}

// GetCallerAddress returns the address of the original caller.
func (context *asyncContext) GetCallerAddress() []byte {
	return context.CallerAddr
}

// GetReturnData returns the data to be sent back to the original caller.
func (context *asyncContext) GetReturnData() []byte {
	return context.ReturnData
}

// GetCallGroup retrieves an AsyncCallGroup by its ID.
func (context *asyncContext) GetCallGroup(groupID string) (*arwen.AsyncCallGroup, bool) {
	index, ok := context.findGroupByID(groupID)
	if ok {
		return context.AsyncCallGroups[index], true
	}
	return nil, false
}

// addCallGroup adds the provided AsyncCallGroup to the AsyncContext, if it does not exist already.
func (context *asyncContext) addCallGroup(group *arwen.AsyncCallGroup) error {
	_, exists := context.findGroupByID(group.Identifier)
	if exists {
		return arwen.ErrAsyncCallGroupExistsAlready
	}

	context.AsyncCallGroups = append(context.AsyncCallGroups, group)
	return nil
}

// SetGroupCallback registers the name of the callback method to be called upon the completion of the specified AsyncCallGroup.
func (context *asyncContext) SetGroupCallback(groupID string, callbackName string, data []byte, gas uint64) error {
	group, exists := context.GetCallGroup(groupID)
	if !exists {
		return arwen.ErrAsyncCallGroupDoesNotExist
	}

	err := context.host.Runtime().ValidateCallbackName(callbackName)
	if err != nil {
		return err
	}

	metering := context.host.Metering()
	gasToLock := metering.ComputeGasLockedForAsync() + gas
	err = metering.UseGasBounded(gasToLock)
	if err != nil {
		return err
	}

	group.Callback = callbackName
	group.GasLocked = gasToLock
	group.CallbackData = data

	return nil
}

func (context *asyncContext) deleteCallGroupByID(groupID string) {
	index, ok := context.findGroupByID(groupID)
	if !ok {
		return
	}

	context.deleteCallGroup(index)
}

func (context *asyncContext) deleteCallGroup(index int) {
	groups := context.AsyncCallGroups
	if len(groups) == 0 {
		return
	}

	last := len(groups) - 1
	if index < 0 || index > last {
		return
	}

	groups[index] = groups[last]
	groups = groups[:last]
	context.AsyncCallGroups = groups
}

// AddCall adds the provided AsyncCall to the specified AsyncCallGroup
func (context *asyncContext) AddCall(groupID string, call *arwen.AsyncCall) error {
	if context.host.IsBuiltinFunctionName(call.SuccessCallback) {
		return arwen.ErrCannotUseBuiltinAsCallback
	}
	if context.host.IsBuiltinFunctionName(call.ErrorCallback) {
		return arwen.ErrCannotUseBuiltinAsCallback
	}

	group, ok := context.GetCallGroup(groupID)
	if !ok {
		group = arwen.NewAsyncCallGroup(groupID)
		err := context.addCallGroup(group)
		if err != nil {
			return err
		}
	}

	// TODO lock gas for callback

	execMode, err := context.determineExecutionMode(call.Destination, call.Data)
	if err != nil {
		return err
	}

	call.ExecutionMode = execMode
	group.AddAsyncCall(call)

	return nil
}

func (context *asyncContext) isValidCallbackName(callback string) bool {
	if callback == arwen.InitFunctionName {
		return false
	}
	if context.host.IsBuiltinFunctionName(callback) {
		return false
	}

	return true
}

func (context *asyncContext) findCall(destination []byte) (string, int, error) {
	for _, group := range context.AsyncCallGroups {
		callIndex, ok := group.FindByDestination(destination)
		if ok {
			return group.Identifier, callIndex, nil
		}
	}

	return "", -1, arwen.ErrAsyncCallNotFound
}

// UpdateCurrentCallStatus detects the AsyncCall returning as callback,
// extracts the ReturnCode from data provided by the destination call, and updates
// the status of the AsyncCall with its value.
func (context *asyncContext) UpdateCurrentCallStatus() (*arwen.AsyncCall, error) {
	vmInput := context.host.Runtime().GetVMInput()
	if vmInput.CallType != vmcommon.AsynchronousCallBack {
		return nil, nil
	}

	if len(vmInput.Arguments) == 0 {
		return nil, arwen.ErrCannotInterpretCallbackArgs
	}

	// The first argument of the callback is the return code of the destination call
	destReturnCode := big.NewInt(0).SetBytes(vmInput.Arguments[0]).Uint64()

	groupID, index, err := context.findCall(vmInput.CallerAddr)
	if err != nil {
		return nil, err
	}

	group, _ := context.GetCallGroup(groupID)
	call := group.AsyncCalls[index]
	call.UpdateStatus(vmcommon.ReturnCode(destReturnCode))

	return call, nil
}

// PrepareLegacyAsyncCall builds an AsyncCall struct from its arguments, sets it as
// the default async call and informs Wasmer to stop contract execution with BreakpointAsyncCall
func (context *asyncContext) PrepareLegacyAsyncCall(address []byte, data []byte, value []byte) error {
	legacyGroupID := arwen.LegacyAsyncCallGroupID

	_, exists := context.GetCallGroup(legacyGroupID)
	if exists {
		return arwen.ErrOnlyOneLegacyAsyncCallAllowed
	}

	gasToLock, err := context.prepareGasForLegacyAsyncCall()
	if err != nil {
		return err
	}

	metering := context.host.Metering()
	gas := metering.GasLeft()

	err = context.AddCall(legacyGroupID, &arwen.AsyncCall{
		Status:          arwen.AsyncCallPending,
		Destination:     address,
		Data:            data,
		ValueBytes:      value,
		SuccessCallback: arwen.CallbackFunctionName,
		ErrorCallback:   arwen.CallbackFunctionName,
		ProvidedGas:     gas,
		GasLocked:       gasToLock,
	})
	if err != nil {
		return err
	}

	context.host.Runtime().SetRuntimeBreakpointValue(arwen.BreakpointAsyncCall)

	return nil
}

// Execute is the entry-point of the async calling mechanism; it is called by
// host.ExecuteOnDestContext() and host.callSCMethod(). When Execute()
// finishes, there should be no remaining AsyncCalls that can be executed
// synchronously, and all AsyncCalls that require asynchronous execution must
// already have corresponding entries in vmOutput.OutputAccounts, to be
// dispatched across shards.
//
// Execute() does NOT handle the callbacks of cross-shard AsyncCalls. See
// PostprocessCrossShardCallback() for that.
//
// Note that Execute() is mutually recursive with host.ExecuteOnDestContext(),
// because synchronous AsyncCalls are executed with
// host.ExecuteOnDestContext(), which, in turn, calls asyncContext.Execute() to
// resolve AsyncCalls generated by the AsyncCalls, and so on.

// Moreover, host.ExecuteOnDestContext() will push the state stack of the
// AsyncContext and work with a clean state before calling Execute(), making
// Execute() and host.ExecuteOnDestContext() mutually reentrant.
func (context *asyncContext) Execute() error {
	if context.IsComplete() {
		return nil
	}

	// Step 1: execute all AsyncCalls that can be executed synchronously
	// (includes smart contracts and built-in functions in the same shard)
	err := context.setupAsyncCallsGas()
	if err != nil {
		return err
	}

	err = context.executeSynchronousCalls()
	if err != nil {
		return err
	}

	// Step 2: redistribute unspent gas; then, in one combined step, do the
	// following:
	// * locally execute built-in functions with cross-shard
	//   destinations, whereby the cross-shard OutputAccount entries are generated
	// * call host.sendAsyncCallCrossShard() for each pending AsyncCall, to
	//   generate the corresponding cross-shard OutputAccount entries
	err = context.setupAsyncCallsGas()
	if err != nil {
		return err
	}

	for _, group := range context.AsyncCallGroups {
		for _, call := range group.AsyncCalls {
			err = context.executeAsyncCall(call)
			if err != nil {
				return err
			}
		}
	}

	context.deleteCallGroupByID(arwen.LegacyAsyncCallGroupID)

	err = context.Save()
	if err != nil {
		return err
	}

	return nil
}

func (context *asyncContext) executeAsyncCall(asyncCall *arwen.AsyncCall) error {
	if asyncCall.ExecutionMode == arwen.AsyncBuiltinFunc {
		err := context.executeSyncHalfOfBuiltinFunction(asyncCall)
		if err != nil {
			return err
		}
	}

	return context.sendAsyncCallCrossShard(asyncCall)
}

func (context *asyncContext) prepareGasForLegacyAsyncCall() (uint64, error) {
	metering := context.host.Metering()
	err := metering.UseGasForAsyncStep()
	if err != nil {
		return 0, err
	}

	var shouldLockGas bool

	if !context.host.IsDynamicGasLockingEnabled() {
		// Legacy mode: static gas locking, always enabled
		shouldLockGas = true
	} else {
		// Dynamic mode: lock only if callBack() exists
		shouldLockGas = context.host.Runtime().HasCallbackMethod()
	}

	gasToLock := uint64(0)
	if shouldLockGas {
		gasToLock = metering.ComputeGasLockedForAsync()
		err = metering.UseGasBounded(gasToLock)
		if err != nil {
			return 0, err
		}
	}

	return gasToLock, nil
}

// PostprocessCrossShardCallback() is called by host.callSCMethod() after it
// has locally executed the callback of a returning cross-shard AsyncCall,
// which means that the AsyncContext corresponding to the original transaction
// must be loaded from storage, and then the corresponding AsyncCall must be
// deleted from the current AsyncContext.
func (context *asyncContext) PostprocessCrossShardCallback() error {
	runtime := context.host.Runtime()
	if runtime.Function() == arwen.CallbackFunctionName {
		// Legacy callbacks do not require postprocessing.
		return nil
	}

	// TODO FindAsyncCallByDestination() only returns the first matched AsyncCall
	// by destination, but there could be multiple matches in an AsyncContext.
	vmInput := runtime.GetVMInput()
	currentGroupID, asyncCallIndex, err := context.findCall(vmInput.CallerAddr)
	if err != nil {
		return err
	}

	currentCallGroup, ok := context.GetCallGroup(currentGroupID)
	if !ok {
		return arwen.ErrCallBackFuncNotExpected
	}

	currentCallGroup.DeleteAsyncCall(asyncCallIndex)
	if currentCallGroup.HasPendingCalls() {
		return nil
	}

	// The current group expects no more callbacks, so its own callback can be
	// executed now.
	context.executeCallGroupCallback(currentCallGroup)
	context.deleteCallGroupByID(currentGroupID)
	// Are we still waiting for callbacks to return?
	if context.HasPendingCallGroups() {
		return nil
	}

	// There are no more callbacks to return from other shards. The context can
	// be deleted from storage.
	err = context.Delete()
	if err != nil {
		return err
	}

	return context.executeContextCallback()
}

// HasPendingCallGroups returns true if the AsyncContext still contains AsyncCallGroup.
func (context *asyncContext) HasPendingCallGroups() bool {
	return len(context.AsyncCallGroups) > 0
}

// IsComplete returns true if there are no more AsyncCallGroups contained in the AsyncContext.
func (context *asyncContext) IsComplete() bool {
	return len(context.AsyncCallGroups) == 0
}

// Save serializes and saves the AsyncContext to the storage of the contract, under a protected key.
func (context *asyncContext) Save() error {
	if len(context.AsyncCallGroups) == 0 {
		return nil
	}

	storage := context.host.Storage()
	runtime := context.host.Runtime()

	storageKey := arwen.CustomStorageKey(arwen.AsyncDataPrefix, runtime.GetPrevTxHash())
	data, err := context.serialize()
	if err != nil {
		return err
	}

	_, err = storage.SetStorage(storageKey, data)
	if err != nil {
		return err
	}

	return nil
}

// Load restores the internal state of the AsyncContext from the storage of the contract.
func (context *asyncContext) Load() error {
	runtime := context.host.Runtime()
	storage := context.host.Storage()

	storageKey := arwen.CustomStorageKey(arwen.AsyncDataPrefix, runtime.GetPrevTxHash())
	data := storage.GetStorage(storageKey)
	if len(data) == 0 {
		return arwen.ErrNoStoredAsyncContextFound
	}

	loadedContext, err := context.deserialize(data)
	if err != nil {
		return err
	}

	context.CallerAddr = loadedContext.CallerAddr
	context.ReturnData = loadedContext.ReturnData
	context.AsyncCallGroups = loadedContext.AsyncCallGroups

	return nil
}

// Delete deletes the persisted state of the AsyncContext from the contract storage.
func (context *asyncContext) Delete() error {
	runtime := context.host.Runtime()
	storage := context.host.Storage()

	storageKey := arwen.CustomStorageKey(arwen.AsyncDataPrefix, runtime.GetPrevTxHash())
	_, err := storage.SetStorage(storageKey, nil)
	return err
}

func (context *asyncContext) determineExecutionMode(destination []byte, data []byte) (arwen.AsyncCallExecutionMode, error) {
	runtime := context.host.Runtime()
	blockchain := context.host.Blockchain()

	// If ArgParser cannot read the Data field, then this is neither a SC call,
	// nor a built-in function call.
	functionName, _, err := context.host.CallArgsParser().ParseData(string(data))
	if err != nil {
		return arwen.AsyncUnknown, err
	}

	shardOfSC := blockchain.GetShardOfAddress(runtime.GetSCAddress())
	shardOfDest := blockchain.GetShardOfAddress(destination)
	sameShard := shardOfSC == shardOfDest

	if sameShard {
		return arwen.SyncExecution, nil
	}

	if context.host.IsBuiltinFunctionName(functionName) {
		return arwen.AsyncBuiltinFunc, nil
	}

	return arwen.AsyncUnknown, nil
}

// TODO decide whether this function is necessary at all, because unspent gas should
// be accummulated into the AsyncContext, then refunded to the original caller.
// Redistribution of gas among calls in the same group should not be necessary.
func (context *asyncContext) setupAsyncCallsGas() error {
	gasLeft := context.host.Metering().GasLeft()
	gasNeeded := uint64(0)
	callsWithZeroGas := uint64(0)

	for _, group := range context.AsyncCallGroups {
		for _, asyncCall := range group.AsyncCalls {
			var err error
			gasNeeded, err = extramath.AddUint64(gasNeeded, asyncCall.ProvidedGas)
			if err != nil {
				return err
			}

			if gasNeeded > gasLeft {
				return arwen.ErrNotEnoughGas
			}

			if asyncCall.ProvidedGas == 0 {
				callsWithZeroGas++
				continue
			}

			asyncCall.GasLimit = asyncCall.ProvidedGas
		}
	}

	if callsWithZeroGas == 0 {
		return nil
	}

	if gasLeft <= gasNeeded {
		return arwen.ErrNotEnoughGas
	}

	gasShare := (gasLeft - gasNeeded) / callsWithZeroGas
	for _, group := range context.AsyncCallGroups {
		for _, asyncCall := range group.AsyncCalls {
			if asyncCall.ProvidedGas == 0 {
				asyncCall.GasLimit = gasShare
			}
		}
	}

	return nil
}

func (context *asyncContext) sendAsyncCallCrossShard(asyncCall arwen.AsyncCallHandler) error {
	host := context.host
	runtime := host.Runtime()
	output := host.Output()

	err := output.Transfer(
		asyncCall.GetDestination(),
		runtime.GetSCAddress(),
		asyncCall.GetGasLimit(),
		asyncCall.GetGasLocked(),
		big.NewInt(0).SetBytes(asyncCall.GetValue()),
		asyncCall.GetData(),
		vmcommon.AsynchronousCall,
	)
	if err != nil {
		metering := host.Metering()
		metering.UseGas(metering.GasLeft())
		runtime.FailExecution(err)
		return err
	}

	return nil
}

// executeAsyncContextCallback will either execute a sync call (in-shard) to
// the original caller by invoking its callback directly, or will dispatch a
// cross-shard callback to it.
func (context *asyncContext) executeContextCallback() error {
	execMode, err := context.determineExecutionMode(context.CallerAddr, context.ReturnData)
	if err != nil {
		return err
	}

	if execMode != arwen.SyncExecution {
		return context.sendContextCallbackToOriginalCaller()
	}

	// The caller is in the same shard, execute its callback
	context.executeSyncContextCallback()

	return nil
}

func (context *asyncContext) sendContextCallbackToOriginalCaller() error {
	host := context.host
	runtime := host.Runtime()
	output := host.Output()
	metering := host.Metering()
	currentCall := runtime.GetVMInput()

	err := output.Transfer(
		context.CallerAddr,
		runtime.GetSCAddress(),
		metering.GasLeft(),
		0,
		currentCall.CallValue,
		context.ReturnData,
		vmcommon.AsynchronousCallBack,
	)
	if err != nil {
		metering.UseGas(metering.GasLeft())
		runtime.FailExecution(err)
		return err
	}

	return nil
}

func (context *asyncContext) serialize() ([]byte, error) {
	serializableContext := &asyncContext{
		host:            nil,
		stateStack:      nil,
		CallerAddr:      context.CallerAddr,
		ReturnData:      context.ReturnData,
		AsyncCallGroups: context.AsyncCallGroups,
	}
	return json.Marshal(serializableContext)
}

func (context *asyncContext) deserialize(data []byte) (*asyncContext, error) {
	deserializedContext := &asyncContext{}
	err := json.Unmarshal(data, deserializedContext)
	if err != nil {
		return nil, err
	}

	return deserializedContext, nil
}

func (context *asyncContext) findGroupByID(groupID string) (int, bool) {
	for index, group := range context.AsyncCallGroups {
		if group.Identifier == groupID {
			return index, true
		}
	}
	return -1, false
}

func computeDataLengthFromArguments(function string, arguments [][]byte) int {
	// Calculate what length would the Data field have, were it of the
	// form "callback@arg1@arg4...

	// TODO this needs tests, especially for the case when the arguments slice
	// contains an empty []byte
	numSeparators := len(arguments)
	dataLength := len(function) + numSeparators
	for _, element := range arguments {
		dataLength += len(element)
	}

	return dataLength
}