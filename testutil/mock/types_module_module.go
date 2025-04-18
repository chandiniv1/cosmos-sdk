// Code generated by MockGen. DO NOT EDIT.
// Source: types/module/module.go
//
// Generated by this command:
//
//	mockgen -source=types/module/module.go -package mock -destination testutil/mock/types_module_module.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	registry "cosmossdk.io/core/registry"
	client "github.com/cosmos/cosmos-sdk/client"
	types "github.com/cosmos/cosmos-sdk/types"
	module "github.com/cosmos/cosmos-sdk/types/module"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockAppModuleBasic is a mock of AppModuleBasic interface.
type MockAppModuleBasic struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleBasicMockRecorder
	isgomock struct{}
}

// MockAppModuleBasicMockRecorder is the mock recorder for MockAppModuleBasic.
type MockAppModuleBasicMockRecorder struct {
	mock *MockAppModuleBasic
}

// NewMockAppModuleBasic creates a new mock instance.
func NewMockAppModuleBasic(ctrl *gomock.Controller) *MockAppModuleBasic {
	mock := &MockAppModuleBasic{ctrl: ctrl}
	mock.recorder = &MockAppModuleBasicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModuleBasic) EXPECT() *MockAppModuleBasicMockRecorder {
	return m.recorder
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockAppModuleBasic) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockAppModuleBasicMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockAppModuleBasic) RegisterLegacyAminoCodec(arg0 registry.AminoRegistrar) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockAppModuleBasicMockRecorder) RegisterLegacyAminoCodec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterLegacyAminoCodec), arg0)
}

// MockAppModule is a mock of AppModule interface.
type MockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleMockRecorder
	isgomock struct{}
}

// MockAppModuleMockRecorder is the mock recorder for MockAppModule.
type MockAppModuleMockRecorder struct {
	mock *MockAppModule
}

// NewMockAppModule creates a new mock instance.
func NewMockAppModule(ctrl *gomock.Controller) *MockAppModule {
	mock := &MockAppModule{ctrl: ctrl}
	mock.recorder = &MockAppModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppModule) EXPECT() *MockAppModuleMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockAppModule) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockAppModuleMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockAppModule)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockAppModule) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockAppModuleMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockAppModule)(nil).IsOnePerModuleType))
}

// Name mocks base method.
func (m *MockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModule)(nil).Name))
}

// MockHasAminoCodec is a mock of HasAminoCodec interface.
type MockHasAminoCodec struct {
	ctrl     *gomock.Controller
	recorder *MockHasAminoCodecMockRecorder
	isgomock struct{}
}

// MockHasAminoCodecMockRecorder is the mock recorder for MockHasAminoCodec.
type MockHasAminoCodecMockRecorder struct {
	mock *MockHasAminoCodec
}

// NewMockHasAminoCodec creates a new mock instance.
func NewMockHasAminoCodec(ctrl *gomock.Controller) *MockHasAminoCodec {
	mock := &MockHasAminoCodec{ctrl: ctrl}
	mock.recorder = &MockHasAminoCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasAminoCodec) EXPECT() *MockHasAminoCodecMockRecorder {
	return m.recorder
}

// RegisterLegacyAminoCodec mocks base method.
func (m *MockHasAminoCodec) RegisterLegacyAminoCodec(arg0 registry.AminoRegistrar) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}

// RegisterLegacyAminoCodec indicates an expected call of RegisterLegacyAminoCodec.
func (mr *MockHasAminoCodecMockRecorder) RegisterLegacyAminoCodec(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockHasAminoCodec)(nil).RegisterLegacyAminoCodec), arg0)
}

// MockHasGRPCGateway is a mock of HasGRPCGateway interface.
type MockHasGRPCGateway struct {
	ctrl     *gomock.Controller
	recorder *MockHasGRPCGatewayMockRecorder
	isgomock struct{}
}

// MockHasGRPCGatewayMockRecorder is the mock recorder for MockHasGRPCGateway.
type MockHasGRPCGatewayMockRecorder struct {
	mock *MockHasGRPCGateway
}

// NewMockHasGRPCGateway creates a new mock instance.
func NewMockHasGRPCGateway(ctrl *gomock.Controller) *MockHasGRPCGateway {
	mock := &MockHasGRPCGateway{ctrl: ctrl}
	mock.recorder = &MockHasGRPCGatewayMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasGRPCGateway) EXPECT() *MockHasGRPCGatewayMockRecorder {
	return m.recorder
}

// RegisterGRPCGatewayRoutes mocks base method.
func (m *MockHasGRPCGateway) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}

// RegisterGRPCGatewayRoutes indicates an expected call of RegisterGRPCGatewayRoutes.
func (mr *MockHasGRPCGatewayMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockHasGRPCGateway)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}

// MockHasInvariants is a mock of HasInvariants interface.
type MockHasInvariants struct {
	ctrl     *gomock.Controller
	recorder *MockHasInvariantsMockRecorder
	isgomock struct{}
}

// MockHasInvariantsMockRecorder is the mock recorder for MockHasInvariants.
type MockHasInvariantsMockRecorder struct {
	mock *MockHasInvariants
}

// NewMockHasInvariants creates a new mock instance.
func NewMockHasInvariants(ctrl *gomock.Controller) *MockHasInvariants {
	mock := &MockHasInvariants{ctrl: ctrl}
	mock.recorder = &MockHasInvariantsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasInvariants) EXPECT() *MockHasInvariantsMockRecorder {
	return m.recorder
}

// RegisterInvariants mocks base method.
func (m *MockHasInvariants) RegisterInvariants(arg0 types.InvariantRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInvariants", arg0)
}

// RegisterInvariants indicates an expected call of RegisterInvariants.
func (mr *MockHasInvariantsMockRecorder) RegisterInvariants(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInvariants", reflect.TypeOf((*MockHasInvariants)(nil).RegisterInvariants), arg0)
}

// MockHasServices is a mock of HasServices interface.
type MockHasServices struct {
	ctrl     *gomock.Controller
	recorder *MockHasServicesMockRecorder
	isgomock struct{}
}

// MockHasServicesMockRecorder is the mock recorder for MockHasServices.
type MockHasServicesMockRecorder struct {
	mock *MockHasServices
}

// NewMockHasServices creates a new mock instance.
func NewMockHasServices(ctrl *gomock.Controller) *MockHasServices {
	mock := &MockHasServices{ctrl: ctrl}
	mock.recorder = &MockHasServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasServices) EXPECT() *MockHasServicesMockRecorder {
	return m.recorder
}

// RegisterServices mocks base method.
func (m *MockHasServices) RegisterServices(arg0 module.Configurator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterServices", arg0)
}

// RegisterServices indicates an expected call of RegisterServices.
func (mr *MockHasServicesMockRecorder) RegisterServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockHasServices)(nil).RegisterServices), arg0)
}

// MockHasRegisterServices is a mock of HasRegisterServices interface.
type MockHasRegisterServices struct {
	ctrl     *gomock.Controller
	recorder *MockHasRegisterServicesMockRecorder
	isgomock struct{}
}

// MockHasRegisterServicesMockRecorder is the mock recorder for MockHasRegisterServices.
type MockHasRegisterServicesMockRecorder struct {
	mock *MockHasRegisterServices
}

// NewMockHasRegisterServices creates a new mock instance.
func NewMockHasRegisterServices(ctrl *gomock.Controller) *MockHasRegisterServices {
	mock := &MockHasRegisterServices{ctrl: ctrl}
	mock.recorder = &MockHasRegisterServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasRegisterServices) EXPECT() *MockHasRegisterServicesMockRecorder {
	return m.recorder
}

// IsAppModule mocks base method.
func (m *MockHasRegisterServices) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasRegisterServicesMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasRegisterServices)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasRegisterServices) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasRegisterServicesMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasRegisterServices)(nil).IsOnePerModuleType))
}

// RegisterServices mocks base method.
func (m *MockHasRegisterServices) RegisterServices(arg0 grpc.ServiceRegistrar) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterServices", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterServices indicates an expected call of RegisterServices.
func (mr *MockHasRegisterServicesMockRecorder) RegisterServices(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockHasRegisterServices)(nil).RegisterServices), arg0)
}

// MockHasABCIEndBlock is a mock of HasABCIEndBlock interface.
type MockHasABCIEndBlock struct {
	ctrl     *gomock.Controller
	recorder *MockHasABCIEndBlockMockRecorder
	isgomock struct{}
}

// MockHasABCIEndBlockMockRecorder is the mock recorder for MockHasABCIEndBlock.
type MockHasABCIEndBlockMockRecorder struct {
	mock *MockHasABCIEndBlock
}

// NewMockHasABCIEndBlock creates a new mock instance.
func NewMockHasABCIEndBlock(ctrl *gomock.Controller) *MockHasABCIEndBlock {
	mock := &MockHasABCIEndBlock{ctrl: ctrl}
	mock.recorder = &MockHasABCIEndBlockMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHasABCIEndBlock) EXPECT() *MockHasABCIEndBlockMockRecorder {
	return m.recorder
}

// EndBlock mocks base method.
func (m *MockHasABCIEndBlock) EndBlock(arg0 context.Context) ([]module.ValidatorUpdate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0)
	ret0, _ := ret[0].([]module.ValidatorUpdate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndBlock indicates an expected call of EndBlock.
func (mr *MockHasABCIEndBlockMockRecorder) EndBlock(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockHasABCIEndBlock)(nil).EndBlock), arg0)
}

// IsAppModule mocks base method.
func (m *MockHasABCIEndBlock) IsAppModule() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsAppModule")
}

// IsAppModule indicates an expected call of IsAppModule.
func (mr *MockHasABCIEndBlockMockRecorder) IsAppModule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAppModule", reflect.TypeOf((*MockHasABCIEndBlock)(nil).IsAppModule))
}

// IsOnePerModuleType mocks base method.
func (m *MockHasABCIEndBlock) IsOnePerModuleType() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IsOnePerModuleType")
}

// IsOnePerModuleType indicates an expected call of IsOnePerModuleType.
func (mr *MockHasABCIEndBlockMockRecorder) IsOnePerModuleType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOnePerModuleType", reflect.TypeOf((*MockHasABCIEndBlock)(nil).IsOnePerModuleType))
}

// Name mocks base method.
func (m *MockHasABCIEndBlock) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockHasABCIEndBlockMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockHasABCIEndBlock)(nil).Name))
}
