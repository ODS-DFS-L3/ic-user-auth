// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	traceability "authenticator-backend/domain/model/traceability"

	mock "github.com/stretchr/testify/mock"
)

// IPlantUsecase is an autogenerated mock type for the IPlantUsecase type
type IPlantUsecase struct {
	mock.Mock
}

// ListPlants provides a mock function with given fields: getPlantModel
func (_m *IPlantUsecase) ListPlants(getPlantModel traceability.GetPlantModel) ([]traceability.PlantModel, error) {
	ret := _m.Called(getPlantModel)

	if len(ret) == 0 {
		panic("no return value specified for ListPlants")
	}

	var r0 []traceability.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.GetPlantModel) ([]traceability.PlantModel, error)); ok {
		return rf(getPlantModel)
	}
	if rf, ok := ret.Get(0).(func(traceability.GetPlantModel) []traceability.PlantModel); ok {
		r0 = rf(getPlantModel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]traceability.PlantModel)
		}
	}

	if rf, ok := ret.Get(1).(func(traceability.GetPlantModel) error); ok {
		r1 = rf(getPlantModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PutPlant provides a mock function with given fields: plantModel
func (_m *IPlantUsecase) PutPlant(plantModel traceability.PlantModel) (traceability.PlantModel, error) {
	ret := _m.Called(plantModel)

	if len(ret) == 0 {
		panic("no return value specified for PutPlant")
	}

	var r0 traceability.PlantModel
	var r1 error
	if rf, ok := ret.Get(0).(func(traceability.PlantModel) (traceability.PlantModel, error)); ok {
		return rf(plantModel)
	}
	if rf, ok := ret.Get(0).(func(traceability.PlantModel) traceability.PlantModel); ok {
		r0 = rf(plantModel)
	} else {
		r0 = ret.Get(0).(traceability.PlantModel)
	}

	if rf, ok := ret.Get(1).(func(traceability.PlantModel) error); ok {
		r1 = rf(plantModel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIPlantUsecase creates a new instance of IPlantUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPlantUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPlantUsecase {
	mock := &IPlantUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
