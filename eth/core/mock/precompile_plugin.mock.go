// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"pkg.berachain.dev/polaris/eth/core"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	"sync"
)

// Ensure, that PrecompilePluginMock does implement core.PrecompilePlugin.
// If this is not the case, regenerate this file with moq.
var _ core.PrecompilePlugin = &PrecompilePluginMock{}

// PrecompilePluginMock is a mock implementation of core.PrecompilePlugin.
//
//	func TestSomethingThatUsesPrecompilePlugin(t *testing.T) {
//
//		// make and configure a mocked core.PrecompilePlugin
//		mockedPrecompilePlugin := &PrecompilePluginMock{
//			DisableReentrancyFunc: func(precompileEVM vm.PrecompileEVM)  {
//				panic("mock out the DisableReentrancy method")
//			},
//			EnableReentrancyFunc: func(precompileEVM vm.PrecompileEVM)  {
//				panic("mock out the EnableReentrancy method")
//			},
//			GetFunc: func(addr common.Address) vm.PrecompiledContract {
//				panic("mock out the Get method")
//			},
//			GetActiveFunc: func(rules *params.Rules) []common.Address {
//				panic("mock out the GetActive method")
//			},
//			GetPrecompilesFunc: func(rules *params.Rules) []precompile.Registrable {
//				panic("mock out the GetPrecompiles method")
//			},
//			HasFunc: func(addr common.Address) bool {
//				panic("mock out the Has method")
//			},
//			RegisterFunc: func(precompiledContract vm.PrecompiledContract) error {
//				panic("mock out the Register method")
//			},
//			RunFunc: func(evm vm.PrecompileEVM, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error) {
//				panic("mock out the Run method")
//			},
//		}
//
//		// use mockedPrecompilePlugin in code that requires core.PrecompilePlugin
//		// and then make assertions.
//
//	}
type PrecompilePluginMock struct {
	// DisableReentrancyFunc mocks the DisableReentrancy method.
	DisableReentrancyFunc func(precompileEVM vm.PrecompileEVM)

	// EnableReentrancyFunc mocks the EnableReentrancy method.
	EnableReentrancyFunc func(precompileEVM vm.PrecompileEVM)

	// GetFunc mocks the Get method.
	GetFunc func(addr common.Address) vm.PrecompiledContract

	// GetActiveFunc mocks the GetActive method.
	GetActiveFunc func(rules *params.Rules) []common.Address

	// GetPrecompilesFunc mocks the GetPrecompiles method.
	GetPrecompilesFunc func(rules *params.Rules) []precompile.Registrable

	// HasFunc mocks the Has method.
	HasFunc func(addr common.Address) bool

	// RegisterFunc mocks the Register method.
	RegisterFunc func(precompiledContract vm.PrecompiledContract) error

	// RunFunc mocks the Run method.
	RunFunc func(evm vm.PrecompileEVM, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// DisableReentrancy holds details about calls to the DisableReentrancy method.
		DisableReentrancy []struct {
			// PrecompileEVM is the precompileEVM argument value.
			PrecompileEVM vm.PrecompileEVM
		}
		// EnableReentrancy holds details about calls to the EnableReentrancy method.
		EnableReentrancy []struct {
			// PrecompileEVM is the precompileEVM argument value.
			PrecompileEVM vm.PrecompileEVM
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Addr is the addr argument value.
			Addr common.Address
		}
		// GetActive holds details about calls to the GetActive method.
		GetActive []struct {
			// Rules is the rules argument value.
			Rules *params.Rules
		}
		// GetPrecompiles holds details about calls to the GetPrecompiles method.
		GetPrecompiles []struct {
			// Rules is the rules argument value.
			Rules *params.Rules
		}
		// Has holds details about calls to the Has method.
		Has []struct {
			// Addr is the addr argument value.
			Addr common.Address
		}
		// Register holds details about calls to the Register method.
		Register []struct {
			// PrecompiledContract is the precompiledContract argument value.
			PrecompiledContract vm.PrecompiledContract
		}
		// Run holds details about calls to the Run method.
		Run []struct {
			// Evm is the evm argument value.
			Evm vm.PrecompileEVM
			// P is the p argument value.
			P vm.PrecompiledContract
			// Input is the input argument value.
			Input []byte
			// Caller is the caller argument value.
			Caller common.Address
			// Value is the value argument value.
			Value *big.Int
			// SuppliedGas is the suppliedGas argument value.
			SuppliedGas uint64
			// Readonly is the readonly argument value.
			Readonly bool
		}
	}
	lockDisableReentrancy sync.RWMutex
	lockEnableReentrancy  sync.RWMutex
	lockGet               sync.RWMutex
	lockGetActive         sync.RWMutex
	lockGetPrecompiles    sync.RWMutex
	lockHas               sync.RWMutex
	lockRegister          sync.RWMutex
	lockRun               sync.RWMutex
}

// DisableReentrancy calls DisableReentrancyFunc.
func (mock *PrecompilePluginMock) DisableReentrancy(precompileEVM vm.PrecompileEVM) {
	if mock.DisableReentrancyFunc == nil {
		panic("PrecompilePluginMock.DisableReentrancyFunc: method is nil but PrecompilePlugin.DisableReentrancy was just called")
	}
	callInfo := struct {
		PrecompileEVM vm.PrecompileEVM
	}{
		PrecompileEVM: precompileEVM,
	}
	mock.lockDisableReentrancy.Lock()
	mock.calls.DisableReentrancy = append(mock.calls.DisableReentrancy, callInfo)
	mock.lockDisableReentrancy.Unlock()
	mock.DisableReentrancyFunc(precompileEVM)
}

// DisableReentrancyCalls gets all the calls that were made to DisableReentrancy.
// Check the length with:
//
//	len(mockedPrecompilePlugin.DisableReentrancyCalls())
func (mock *PrecompilePluginMock) DisableReentrancyCalls() []struct {
	PrecompileEVM vm.PrecompileEVM
} {
	var calls []struct {
		PrecompileEVM vm.PrecompileEVM
	}
	mock.lockDisableReentrancy.RLock()
	calls = mock.calls.DisableReentrancy
	mock.lockDisableReentrancy.RUnlock()
	return calls
}

// EnableReentrancy calls EnableReentrancyFunc.
func (mock *PrecompilePluginMock) EnableReentrancy(precompileEVM vm.PrecompileEVM) {
	if mock.EnableReentrancyFunc == nil {
		panic("PrecompilePluginMock.EnableReentrancyFunc: method is nil but PrecompilePlugin.EnableReentrancy was just called")
	}
	callInfo := struct {
		PrecompileEVM vm.PrecompileEVM
	}{
		PrecompileEVM: precompileEVM,
	}
	mock.lockEnableReentrancy.Lock()
	mock.calls.EnableReentrancy = append(mock.calls.EnableReentrancy, callInfo)
	mock.lockEnableReentrancy.Unlock()
	mock.EnableReentrancyFunc(precompileEVM)
}

// EnableReentrancyCalls gets all the calls that were made to EnableReentrancy.
// Check the length with:
//
//	len(mockedPrecompilePlugin.EnableReentrancyCalls())
func (mock *PrecompilePluginMock) EnableReentrancyCalls() []struct {
	PrecompileEVM vm.PrecompileEVM
} {
	var calls []struct {
		PrecompileEVM vm.PrecompileEVM
	}
	mock.lockEnableReentrancy.RLock()
	calls = mock.calls.EnableReentrancy
	mock.lockEnableReentrancy.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *PrecompilePluginMock) Get(addr common.Address) vm.PrecompiledContract {
	if mock.GetFunc == nil {
		panic("PrecompilePluginMock.GetFunc: method is nil but PrecompilePlugin.Get was just called")
	}
	callInfo := struct {
		Addr common.Address
	}{
		Addr: addr,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(addr)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedPrecompilePlugin.GetCalls())
func (mock *PrecompilePluginMock) GetCalls() []struct {
	Addr common.Address
} {
	var calls []struct {
		Addr common.Address
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// GetActive calls GetActiveFunc.
func (mock *PrecompilePluginMock) GetActive(rules *params.Rules) []common.Address {
	if mock.GetActiveFunc == nil {
		panic("PrecompilePluginMock.GetActiveFunc: method is nil but PrecompilePlugin.GetActive was just called")
	}
	callInfo := struct {
		Rules *params.Rules
	}{
		Rules: rules,
	}
	mock.lockGetActive.Lock()
	mock.calls.GetActive = append(mock.calls.GetActive, callInfo)
	mock.lockGetActive.Unlock()
	return mock.GetActiveFunc(rules)
}

// GetActiveCalls gets all the calls that were made to GetActive.
// Check the length with:
//
//	len(mockedPrecompilePlugin.GetActiveCalls())
func (mock *PrecompilePluginMock) GetActiveCalls() []struct {
	Rules *params.Rules
} {
	var calls []struct {
		Rules *params.Rules
	}
	mock.lockGetActive.RLock()
	calls = mock.calls.GetActive
	mock.lockGetActive.RUnlock()
	return calls
}

// GetPrecompiles calls GetPrecompilesFunc.
func (mock *PrecompilePluginMock) GetPrecompiles(rules *params.Rules) []precompile.Registrable {
	if mock.GetPrecompilesFunc == nil {
		panic("PrecompilePluginMock.GetPrecompilesFunc: method is nil but PrecompilePlugin.GetPrecompiles was just called")
	}
	callInfo := struct {
		Rules *params.Rules
	}{
		Rules: rules,
	}
	mock.lockGetPrecompiles.Lock()
	mock.calls.GetPrecompiles = append(mock.calls.GetPrecompiles, callInfo)
	mock.lockGetPrecompiles.Unlock()
	return mock.GetPrecompilesFunc(rules)
}

// GetPrecompilesCalls gets all the calls that were made to GetPrecompiles.
// Check the length with:
//
//	len(mockedPrecompilePlugin.GetPrecompilesCalls())
func (mock *PrecompilePluginMock) GetPrecompilesCalls() []struct {
	Rules *params.Rules
} {
	var calls []struct {
		Rules *params.Rules
	}
	mock.lockGetPrecompiles.RLock()
	calls = mock.calls.GetPrecompiles
	mock.lockGetPrecompiles.RUnlock()
	return calls
}

// Has calls HasFunc.
func (mock *PrecompilePluginMock) Has(addr common.Address) bool {
	if mock.HasFunc == nil {
		panic("PrecompilePluginMock.HasFunc: method is nil but PrecompilePlugin.Has was just called")
	}
	callInfo := struct {
		Addr common.Address
	}{
		Addr: addr,
	}
	mock.lockHas.Lock()
	mock.calls.Has = append(mock.calls.Has, callInfo)
	mock.lockHas.Unlock()
	return mock.HasFunc(addr)
}

// HasCalls gets all the calls that were made to Has.
// Check the length with:
//
//	len(mockedPrecompilePlugin.HasCalls())
func (mock *PrecompilePluginMock) HasCalls() []struct {
	Addr common.Address
} {
	var calls []struct {
		Addr common.Address
	}
	mock.lockHas.RLock()
	calls = mock.calls.Has
	mock.lockHas.RUnlock()
	return calls
}

// Register calls RegisterFunc.
func (mock *PrecompilePluginMock) Register(precompiledContract vm.PrecompiledContract) error {
	if mock.RegisterFunc == nil {
		panic("PrecompilePluginMock.RegisterFunc: method is nil but PrecompilePlugin.Register was just called")
	}
	callInfo := struct {
		PrecompiledContract vm.PrecompiledContract
	}{
		PrecompiledContract: precompiledContract,
	}
	mock.lockRegister.Lock()
	mock.calls.Register = append(mock.calls.Register, callInfo)
	mock.lockRegister.Unlock()
	return mock.RegisterFunc(precompiledContract)
}

// RegisterCalls gets all the calls that were made to Register.
// Check the length with:
//
//	len(mockedPrecompilePlugin.RegisterCalls())
func (mock *PrecompilePluginMock) RegisterCalls() []struct {
	PrecompiledContract vm.PrecompiledContract
} {
	var calls []struct {
		PrecompiledContract vm.PrecompiledContract
	}
	mock.lockRegister.RLock()
	calls = mock.calls.Register
	mock.lockRegister.RUnlock()
	return calls
}

// Run calls RunFunc.
func (mock *PrecompilePluginMock) Run(evm vm.PrecompileEVM, p vm.PrecompiledContract, input []byte, caller common.Address, value *big.Int, suppliedGas uint64, readonly bool) ([]byte, uint64, error) {
	if mock.RunFunc == nil {
		panic("PrecompilePluginMock.RunFunc: method is nil but PrecompilePlugin.Run was just called")
	}
	callInfo := struct {
		Evm         vm.PrecompileEVM
		P           vm.PrecompiledContract
		Input       []byte
		Caller      common.Address
		Value       *big.Int
		SuppliedGas uint64
		Readonly    bool
	}{
		Evm:         evm,
		P:           p,
		Input:       input,
		Caller:      caller,
		Value:       value,
		SuppliedGas: suppliedGas,
		Readonly:    readonly,
	}
	mock.lockRun.Lock()
	mock.calls.Run = append(mock.calls.Run, callInfo)
	mock.lockRun.Unlock()
	return mock.RunFunc(evm, p, input, caller, value, suppliedGas, readonly)
}

// RunCalls gets all the calls that were made to Run.
// Check the length with:
//
//	len(mockedPrecompilePlugin.RunCalls())
func (mock *PrecompilePluginMock) RunCalls() []struct {
	Evm         vm.PrecompileEVM
	P           vm.PrecompiledContract
	Input       []byte
	Caller      common.Address
	Value       *big.Int
	SuppliedGas uint64
	Readonly    bool
} {
	var calls []struct {
		Evm         vm.PrecompileEVM
		P           vm.PrecompiledContract
		Input       []byte
		Caller      common.Address
		Value       *big.Int
		SuppliedGas uint64
		Readonly    bool
	}
	mock.lockRun.RLock()
	calls = mock.calls.Run
	mock.lockRun.RUnlock()
	return calls
}