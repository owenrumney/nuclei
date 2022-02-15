// Code generated by MockGen. DO NOT EDIT.
// Source: dbsql/querier.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dbsql "github.com/projectdiscovery/nuclei/v2/pkg/rest/db/dbsql"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// AddIssue mocks base method.
func (m *MockQuerier) AddIssue(ctx context.Context, arg dbsql.AddIssueParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddIssue", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddIssue indicates an expected call of AddIssue.
func (mr *MockQuerierMockRecorder) AddIssue(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddIssue", reflect.TypeOf((*MockQuerier)(nil).AddIssue), ctx, arg)
}

// AddScan mocks base method.
func (m *MockQuerier) AddScan(ctx context.Context, arg dbsql.AddScanParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddScan", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddScan indicates an expected call of AddScan.
func (mr *MockQuerierMockRecorder) AddScan(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddScan", reflect.TypeOf((*MockQuerier)(nil).AddScan), ctx, arg)
}

// AddTarget mocks base method.
func (m *MockQuerier) AddTarget(ctx context.Context, arg dbsql.AddTargetParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTarget", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTarget indicates an expected call of AddTarget.
func (mr *MockQuerierMockRecorder) AddTarget(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTarget", reflect.TypeOf((*MockQuerier)(nil).AddTarget), ctx, arg)
}

// AddTemplate mocks base method.
func (m *MockQuerier) AddTemplate(ctx context.Context, arg dbsql.AddTemplateParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTemplate", ctx, arg)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTemplate indicates an expected call of AddTemplate.
func (mr *MockQuerierMockRecorder) AddTemplate(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTemplate", reflect.TypeOf((*MockQuerier)(nil).AddTemplate), ctx, arg)
}

// DeleteIssue mocks base method.
func (m *MockQuerier) DeleteIssue(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIssue", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIssue indicates an expected call of DeleteIssue.
func (mr *MockQuerierMockRecorder) DeleteIssue(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssue", reflect.TypeOf((*MockQuerier)(nil).DeleteIssue), ctx, id)
}

// DeleteIssueByScanID mocks base method.
func (m *MockQuerier) DeleteIssueByScanID(ctx context.Context, scanid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteIssueByScanID", ctx, scanid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteIssueByScanID indicates an expected call of DeleteIssueByScanID.
func (mr *MockQuerierMockRecorder) DeleteIssueByScanID(ctx, scanid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteIssueByScanID", reflect.TypeOf((*MockQuerier)(nil).DeleteIssueByScanID), ctx, scanid)
}

// DeleteScan mocks base method.
func (m *MockQuerier) DeleteScan(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScan", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScan indicates an expected call of DeleteScan.
func (mr *MockQuerierMockRecorder) DeleteScan(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScan", reflect.TypeOf((*MockQuerier)(nil).DeleteScan), ctx, id)
}

// DeleteTarget mocks base method.
func (m *MockQuerier) DeleteTarget(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTarget", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTarget indicates an expected call of DeleteTarget.
func (mr *MockQuerierMockRecorder) DeleteTarget(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTarget", reflect.TypeOf((*MockQuerier)(nil).DeleteTarget), ctx, id)
}

// DeleteTemplate mocks base method.
func (m *MockQuerier) DeleteTemplate(ctx context.Context, path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTemplate", ctx, path)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTemplate indicates an expected call of DeleteTemplate.
func (mr *MockQuerierMockRecorder) DeleteTemplate(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTemplate", reflect.TypeOf((*MockQuerier)(nil).DeleteTemplate), ctx, path)
}

// GetIssue mocks base method.
func (m *MockQuerier) GetIssue(ctx context.Context, id int64) (dbsql.GetIssueRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssue", ctx, id)
	ret0, _ := ret[0].(dbsql.GetIssueRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssue indicates an expected call of GetIssue.
func (mr *MockQuerierMockRecorder) GetIssue(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssue", reflect.TypeOf((*MockQuerier)(nil).GetIssue), ctx, id)
}

// GetIssues mocks base method.
func (m *MockQuerier) GetIssues(ctx context.Context, arg dbsql.GetIssuesParams) ([]dbsql.GetIssuesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssues", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetIssuesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssues indicates an expected call of GetIssues.
func (mr *MockQuerierMockRecorder) GetIssues(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssues", reflect.TypeOf((*MockQuerier)(nil).GetIssues), ctx, arg)
}

// GetIssuesByScanID mocks base method.
func (m *MockQuerier) GetIssuesByScanID(ctx context.Context, arg dbsql.GetIssuesByScanIDParams) ([]dbsql.GetIssuesByScanIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesByScanID", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetIssuesByScanIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesByScanID indicates an expected call of GetIssuesByScanID.
func (mr *MockQuerierMockRecorder) GetIssuesByScanID(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesByScanID", reflect.TypeOf((*MockQuerier)(nil).GetIssuesByScanID), ctx, arg)
}

// GetIssuesMatches mocks base method.
func (m *MockQuerier) GetIssuesMatches(ctx context.Context, arg dbsql.GetIssuesMatchesParams) ([]dbsql.GetIssuesMatchesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIssuesMatches", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetIssuesMatchesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIssuesMatches indicates an expected call of GetIssuesMatches.
func (mr *MockQuerierMockRecorder) GetIssuesMatches(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIssuesMatches", reflect.TypeOf((*MockQuerier)(nil).GetIssuesMatches), ctx, arg)
}

// GetScan mocks base method.
func (m *MockQuerier) GetScan(ctx context.Context, id int64) (dbsql.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScan", ctx, id)
	ret0, _ := ret[0].(dbsql.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScan indicates an expected call of GetScan.
func (mr *MockQuerierMockRecorder) GetScan(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScan", reflect.TypeOf((*MockQuerier)(nil).GetScan), ctx, id)
}

// GetScans mocks base method.
func (m *MockQuerier) GetScans(ctx context.Context, arg dbsql.GetScansParams) ([]dbsql.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScans", ctx, arg)
	ret0, _ := ret[0].([]dbsql.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScans indicates an expected call of GetScans.
func (mr *MockQuerierMockRecorder) GetScans(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScans", reflect.TypeOf((*MockQuerier)(nil).GetScans), ctx, arg)
}

// GetScansBySearchKey mocks base method.
func (m *MockQuerier) GetScansBySearchKey(ctx context.Context, dollar_1 sql.NullString) ([]dbsql.Scan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScansBySearchKey", ctx, dollar_1)
	ret0, _ := ret[0].([]dbsql.Scan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScansBySearchKey indicates an expected call of GetScansBySearchKey.
func (mr *MockQuerierMockRecorder) GetScansBySearchKey(ctx, dollar_1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScansBySearchKey", reflect.TypeOf((*MockQuerier)(nil).GetScansBySearchKey), ctx, dollar_1)
}

// GetScansForSchedule mocks base method.
func (m *MockQuerier) GetScansForSchedule(ctx context.Context, scheduleoccurence sql.NullString) ([]dbsql.GetScansForScheduleRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScansForSchedule", ctx, scheduleoccurence)
	ret0, _ := ret[0].([]dbsql.GetScansForScheduleRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetScansForSchedule indicates an expected call of GetScansForSchedule.
func (mr *MockQuerierMockRecorder) GetScansForSchedule(ctx, scheduleoccurence interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScansForSchedule", reflect.TypeOf((*MockQuerier)(nil).GetScansForSchedule), ctx, scheduleoccurence)
}

// GetSettingByName mocks base method.
func (m *MockQuerier) GetSettingByName(ctx context.Context, name string) (dbsql.GetSettingByNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettingByName", ctx, name)
	ret0, _ := ret[0].(dbsql.GetSettingByNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettingByName indicates an expected call of GetSettingByName.
func (mr *MockQuerierMockRecorder) GetSettingByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettingByName", reflect.TypeOf((*MockQuerier)(nil).GetSettingByName), ctx, name)
}

// GetSettings mocks base method.
func (m *MockQuerier) GetSettings(ctx context.Context) ([]dbsql.Setting, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", ctx)
	ret0, _ := ret[0].([]dbsql.Setting)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockQuerierMockRecorder) GetSettings(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockQuerier)(nil).GetSettings), ctx)
}

// GetTarget mocks base method.
func (m *MockQuerier) GetTarget(ctx context.Context, id int64) (dbsql.GetTargetRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTarget", ctx, id)
	ret0, _ := ret[0].(dbsql.GetTargetRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTarget indicates an expected call of GetTarget.
func (mr *MockQuerierMockRecorder) GetTarget(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTarget", reflect.TypeOf((*MockQuerier)(nil).GetTarget), ctx, id)
}

// GetTargetByName mocks base method.
func (m *MockQuerier) GetTargetByName(ctx context.Context, name string) (dbsql.GetTargetByNameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetByName", ctx, name)
	ret0, _ := ret[0].(dbsql.GetTargetByNameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetByName indicates an expected call of GetTargetByName.
func (mr *MockQuerierMockRecorder) GetTargetByName(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetByName", reflect.TypeOf((*MockQuerier)(nil).GetTargetByName), ctx, name)
}

// GetTargets mocks base method.
func (m *MockQuerier) GetTargets(ctx context.Context, arg dbsql.GetTargetsParams) ([]dbsql.GetTargetsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargets", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetTargetsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargets indicates an expected call of GetTargets.
func (mr *MockQuerierMockRecorder) GetTargets(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargets", reflect.TypeOf((*MockQuerier)(nil).GetTargets), ctx, arg)
}

// GetTargetsForSearch mocks base method.
func (m *MockQuerier) GetTargetsForSearch(ctx context.Context, arg dbsql.GetTargetsForSearchParams) ([]dbsql.GetTargetsForSearchRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTargetsForSearch", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetTargetsForSearchRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTargetsForSearch indicates an expected call of GetTargetsForSearch.
func (mr *MockQuerierMockRecorder) GetTargetsForSearch(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTargetsForSearch", reflect.TypeOf((*MockQuerier)(nil).GetTargetsForSearch), ctx, arg)
}

// GetTemplateContents mocks base method.
func (m *MockQuerier) GetTemplateContents(ctx context.Context, path string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplateContents", ctx, path)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplateContents indicates an expected call of GetTemplateContents.
func (mr *MockQuerierMockRecorder) GetTemplateContents(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplateContents", reflect.TypeOf((*MockQuerier)(nil).GetTemplateContents), ctx, path)
}

// GetTemplates mocks base method.
func (m *MockQuerier) GetTemplates(ctx context.Context, arg dbsql.GetTemplatesParams) ([]dbsql.GetTemplatesRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplates", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetTemplatesRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplates indicates an expected call of GetTemplates.
func (mr *MockQuerierMockRecorder) GetTemplates(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplates", reflect.TypeOf((*MockQuerier)(nil).GetTemplates), ctx, arg)
}

// GetTemplatesByFolder mocks base method.
func (m *MockQuerier) GetTemplatesByFolder(ctx context.Context, folder string) ([]dbsql.GetTemplatesByFolderRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplatesByFolder", ctx, folder)
	ret0, _ := ret[0].([]dbsql.GetTemplatesByFolderRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplatesByFolder indicates an expected call of GetTemplatesByFolder.
func (mr *MockQuerierMockRecorder) GetTemplatesByFolder(ctx, folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplatesByFolder", reflect.TypeOf((*MockQuerier)(nil).GetTemplatesByFolder), ctx, folder)
}

// GetTemplatesByFolderOne mocks base method.
func (m *MockQuerier) GetTemplatesByFolderOne(ctx context.Context, folder string) (dbsql.GetTemplatesByFolderOneRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplatesByFolderOne", ctx, folder)
	ret0, _ := ret[0].(dbsql.GetTemplatesByFolderOneRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplatesByFolderOne indicates an expected call of GetTemplatesByFolderOne.
func (mr *MockQuerierMockRecorder) GetTemplatesByFolderOne(ctx, folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplatesByFolderOne", reflect.TypeOf((*MockQuerier)(nil).GetTemplatesByFolderOne), ctx, folder)
}

// GetTemplatesBySearchKey mocks base method.
func (m *MockQuerier) GetTemplatesBySearchKey(ctx context.Context, arg dbsql.GetTemplatesBySearchKeyParams) ([]dbsql.GetTemplatesBySearchKeyRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplatesBySearchKey", ctx, arg)
	ret0, _ := ret[0].([]dbsql.GetTemplatesBySearchKeyRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplatesBySearchKey indicates an expected call of GetTemplatesBySearchKey.
func (mr *MockQuerierMockRecorder) GetTemplatesBySearchKey(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplatesBySearchKey", reflect.TypeOf((*MockQuerier)(nil).GetTemplatesBySearchKey), ctx, arg)
}

// GetTemplatesForScan mocks base method.
func (m *MockQuerier) GetTemplatesForScan(ctx context.Context, folder string) ([]dbsql.GetTemplatesForScanRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTemplatesForScan", ctx, folder)
	ret0, _ := ret[0].([]dbsql.GetTemplatesForScanRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTemplatesForScan indicates an expected call of GetTemplatesForScan.
func (mr *MockQuerierMockRecorder) GetTemplatesForScan(ctx, folder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTemplatesForScan", reflect.TypeOf((*MockQuerier)(nil).GetTemplatesForScan), ctx, folder)
}

// SetSettings mocks base method.
func (m *MockQuerier) SetSettings(ctx context.Context, arg dbsql.SetSettingsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSettings", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSettings indicates an expected call of SetSettings.
func (mr *MockQuerierMockRecorder) SetSettings(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSettings", reflect.TypeOf((*MockQuerier)(nil).SetSettings), ctx, arg)
}

// UpdateIssue mocks base method.
func (m *MockQuerier) UpdateIssue(ctx context.Context, arg dbsql.UpdateIssueParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIssue", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIssue indicates an expected call of UpdateIssue.
func (mr *MockQuerierMockRecorder) UpdateIssue(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIssue", reflect.TypeOf((*MockQuerier)(nil).UpdateIssue), ctx, arg)
}

// UpdateScanState mocks base method.
func (m *MockQuerier) UpdateScanState(ctx context.Context, arg dbsql.UpdateScanStateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateScanState", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateScanState indicates an expected call of UpdateScanState.
func (mr *MockQuerierMockRecorder) UpdateScanState(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateScanState", reflect.TypeOf((*MockQuerier)(nil).UpdateScanState), ctx, arg)
}

// UpdateSettings mocks base method.
func (m *MockQuerier) UpdateSettings(ctx context.Context, arg dbsql.UpdateSettingsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSettings", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings.
func (mr *MockQuerierMockRecorder) UpdateSettings(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockQuerier)(nil).UpdateSettings), ctx, arg)
}

// UpdateTargetMetadata mocks base method.
func (m *MockQuerier) UpdateTargetMetadata(ctx context.Context, arg dbsql.UpdateTargetMetadataParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTargetMetadata", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTargetMetadata indicates an expected call of UpdateTargetMetadata.
func (mr *MockQuerierMockRecorder) UpdateTargetMetadata(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTargetMetadata", reflect.TypeOf((*MockQuerier)(nil).UpdateTargetMetadata), ctx, arg)
}

// UpdateTemplate mocks base method.
func (m *MockQuerier) UpdateTemplate(ctx context.Context, arg dbsql.UpdateTemplateParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTemplate", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTemplate indicates an expected call of UpdateTemplate.
func (mr *MockQuerierMockRecorder) UpdateTemplate(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTemplate", reflect.TypeOf((*MockQuerier)(nil).UpdateTemplate), ctx, arg)
}
