/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	io "io"

	volumesnapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v7/apis/volumesnapshot/v1"
	mock "github.com/stretchr/testify/mock"

	"github.com/vmware-tanzu/velero/internal/volume"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	itemoperation "github.com/vmware-tanzu/velero/pkg/itemoperation"
	"github.com/vmware-tanzu/velero/pkg/persistence"
	"github.com/vmware-tanzu/velero/pkg/util/results"
)

// BackupStore is an autogenerated mock type for the BackupStore type
type BackupStore struct {
	mock.Mock
}

// BackupExists provides a mock function with given fields: bucket, backupName
func (_m *BackupStore) BackupExists(bucket string, backupName string) (bool, error) {
	ret := _m.Called(bucket, backupName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(bucket, backupName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(bucket, backupName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBackup provides a mock function with given fields: name
func (_m *BackupStore) DeleteBackup(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRestore provides a mock function with given fields: name
func (_m *BackupStore) DeleteRestore(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBackupContents provides a mock function with given fields: name
func (_m *BackupStore) GetBackupContents(name string) (io.ReadCloser, error) {
	ret := _m.Called(name)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupItemOperations provides a mock function with given fields: name
func (_m *BackupStore) GetBackupItemOperations(name string) ([]*itemoperation.BackupOperation, error) {
	ret := _m.Called(name)

	var r0 []*itemoperation.BackupOperation
	if rf, ok := ret.Get(0).(func(string) []*itemoperation.BackupOperation); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*itemoperation.BackupOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupMetadata provides a mock function with given fields: name
func (_m *BackupStore) GetBackupMetadata(name string) (*v1.Backup, error) {
	ret := _m.Called(name)

	var r0 *v1.Backup
	if rf, ok := ret.Get(0).(func(string) *v1.Backup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Backup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBackupVolumeSnapshots provides a mock function with given fields: name
func (_m *BackupStore) GetBackupVolumeSnapshots(name string) ([]*volume.Snapshot, error) {
	ret := _m.Called(name)

	var r0 []*volume.Snapshot
	if rf, ok := ret.Get(0).(func(string) []*volume.Snapshot); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volume.Snapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCSIVolumeSnapshotClasses provides a mock function with given fields: name
func (_m *BackupStore) GetCSIVolumeSnapshotClasses(name string) ([]*volumesnapshotv1.VolumeSnapshotClass, error) {
	ret := _m.Called(name)

	var r0 []*volumesnapshotv1.VolumeSnapshotClass
	if rf, ok := ret.Get(0).(func(string) []*volumesnapshotv1.VolumeSnapshotClass); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volumesnapshotv1.VolumeSnapshotClass)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCSIVolumeSnapshotContents provides a mock function with given fields: name
func (_m *BackupStore) GetCSIVolumeSnapshotContents(name string) ([]*volumesnapshotv1.VolumeSnapshotContent, error) {
	ret := _m.Called(name)

	var r0 []*volumesnapshotv1.VolumeSnapshotContent
	if rf, ok := ret.Get(0).(func(string) []*volumesnapshotv1.VolumeSnapshotContent); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volumesnapshotv1.VolumeSnapshotContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCSIVolumeSnapshots provides a mock function with given fields: name
func (_m *BackupStore) GetCSIVolumeSnapshots(name string) ([]*volumesnapshotv1.VolumeSnapshot, error) {
	ret := _m.Called(name)

	var r0 []*volumesnapshotv1.VolumeSnapshot
	if rf, ok := ret.Get(0).(func(string) []*volumesnapshotv1.VolumeSnapshot); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volumesnapshotv1.VolumeSnapshot)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDownloadURL provides a mock function with given fields: target
func (_m *BackupStore) GetDownloadURL(target v1.DownloadTarget) (string, map[string][]string ,error) {
	ret := _m.Called(target)

	var r0 string
	if rf, ok := ret.Get(0).(func(v1.DownloadTarget) string); ok {
		r0 = rf(target)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v1.DownloadTarget) error); ok {
		r1 = rf(target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, make(map[string][]string), r1
}

// GetPodVolumeBackups provides a mock function with given fields: name
func (_m *BackupStore) GetPodVolumeBackups(name string) ([]*v1.PodVolumeBackup, error) {
	ret := _m.Called(name)

	var r0 []*v1.PodVolumeBackup
	if rf, ok := ret.Get(0).(func(string) []*v1.PodVolumeBackup); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.PodVolumeBackup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemOperations provides a mock function with given fields: name
func (_m *BackupStore) GetRestoreItemOperations(name string) ([]*itemoperation.RestoreOperation, error) {
	ret := _m.Called(name)

	var r0 []*itemoperation.RestoreOperation
	if rf, ok := ret.Get(0).(func(string) []*itemoperation.RestoreOperation); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*itemoperation.RestoreOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoreItemOperations provides a mock function with given fields: name
func (_m *BackupStore) GetBackupVolumeInfos(name string) ([]*volume.BackupVolumeInfo, error) {
	ret := _m.Called(name)

	var r0 []*volume.BackupVolumeInfo
	if rf, ok := ret.Get(0).(func(string) []*volume.BackupVolumeInfo); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*volume.BackupVolumeInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBackupVolumeInfos provides a mock function with given fields: name, volumeInfo
func (_m *BackupStore) PutBackupVolumeInfos(name string, volumeInfo io.Reader) error {
	ret := _m.Called(name, volumeInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(name, volumeInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetRestoreResults provides a mock function with given fields: name
func (_m *BackupStore) GetRestoreResults(name string) (map[string]results.Result, error) {
	ret := _m.Called(name)

	r0 := make(map[string]results.Result)
	if rf, ok := ret.Get(0).(func(string) map[string]results.Result); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]results.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRestoredResourceList provides a mock function with given fields: name
func (_m *BackupStore) GetRestoredResourceList(name string) (map[string][]string, error) {
	ret := _m.Called(name)

	r0 := make(map[string][]string)
	if rf, ok := ret.Get(0).(func(string) map[string][]string); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string][]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsValid provides a mock function with given fields:
func (_m *BackupStore) IsValid() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListBackups provides a mock function with given fields:
func (_m *BackupStore) ListBackups() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutBackup provides a mock function with given fields: info
func (_m *BackupStore) PutBackup(info persistence.BackupInfo) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(persistence.BackupInfo) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutBackupContents provides a mock function with given fields: backup, backupContents
func (_m *BackupStore) PutBackupContents(backup string, backupContents io.Reader) error {
	ret := _m.Called(backup, backupContents)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(backup, backupContents)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutBackupItemOperations provides a mock function with given fields: backup, backupItemOperations
func (_m *BackupStore) PutBackupItemOperations(backup string, backupItemOperations io.Reader) error {
	ret := _m.Called(backup, backupItemOperations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(backup, backupItemOperations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutBackupMetadata provides a mock function with given fields: backup, backupMetadata
func (_m *BackupStore) PutBackupMetadata(backup string, backupMetadata io.Reader) error {
	ret := _m.Called(backup, backupMetadata)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(backup, backupMetadata)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreItemOperations provides a mock function with given fields: restore, restoreItemOperations
func (_m *BackupStore) PutRestoreItemOperations(restore string, restoreItemOperations io.Reader) error {
	ret := _m.Called(restore, restoreItemOperations)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(restore, restoreItemOperations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreLog provides a mock function with given fields: backup, restore, log
func (_m *BackupStore) PutRestoreLog(backup string, restore string, log io.Reader) error {
	ret := _m.Called(backup, restore, log)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(backup, restore, log)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreResults provides a mock function with given fields: backup, restore, results
func (_m *BackupStore) PutRestoreResults(backup string, restore string, results io.Reader) error {
	ret := _m.Called(backup, restore, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) error); ok {
		r0 = rf(backup, restore, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoredResourceList provides a mock function with given fields: restore, results
func (_m *BackupStore) PutRestoredResourceList(restore string, results io.Reader) error {
	ret := _m.Called(restore, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(restore, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PutRestoreVolumeInfo provides a mock function with given fields: restore, results
func (_m *BackupStore) 	PutRestoreVolumeInfo(restore string, results io.Reader) error {
	ret := _m.Called(restore, results)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, io.Reader) error); ok {
		r0 = rf(restore, results)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
type mockConstructorTestingTNewBackupStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewBackupStore creates a new instance of BackupStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBackupStore(t mockConstructorTestingTNewBackupStore) *BackupStore {
	mock := &BackupStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
