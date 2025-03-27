package traceability

import (
	"fmt"
	"time"

	"authenticator-backend/domain/common"
	"authenticator-backend/extension/logger"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlantModel struct {
	DataModelType string               `json:"dataModelType"`
	Attribute     PlantAttributeDetail `json:"attribute"`
}

// PlantAttributeDetail
// Summary: This is structure which defines PlantAttributeDetail.
// Service: Dataspace
// Router: [PUT] /api/v2/authinfo/plant
// Usage: output
type PlantAttributeDetail struct {
	PlantID        uuid.UUID      `json:"plantId"`
	OperatorID     uuid.UUID      `json:"operatorId"`
	PlantName      string         `json:"plantName"`
	PlantAddress   string         `json:"plantAddress"`
	OpenPlantID    string         `json:"openPlantId"`
	PlantAttribute PlantAttribute `json:"plantAttribute"`
}

// PlantAttribute
// Summary: This is structure which defines PlantAttribute.
type PlantAttribute struct {
	GlobalPlantID *string `json:"globalPlantId,omitempty"`
}

// PlantEntityModel
// Summary: This is structure which defines PlantEntityModel.
// DBName: plants
type PlantEntityModel struct {
	PlantID       uuid.UUID      `json:"plantId" gorm:"type:uuid"`
	OperatorID    uuid.UUID      `json:"operatorId" gorm:"type:uuid;not null"`
	PlantName     string         `json:"plantName" gorm:"type:string"`
	PlantAddress  string         `json:"plantAddress" gorm:"type:string"`
	OpenPlantID   string         `json:"openPlantId" gorm:"type:string"`
	GlobalPlantID *string        `json:"globalPlantId" gorm:"type:string"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
	CreatedAt     time.Time      `json:"createdAt" gorm:"<-:create "`
	CreatedUserID string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	UpdatedUserID string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// PlantModels
// Summary: This is structure which defines PlantModels.
// Service: Dataspace
// Router: [GET] /api/v2/authinfo/plant
// Usage: output
type PlantModels []PlantModel

// PlantEntityModels
// Summary: This is structure which defines PlantEntityModels.
type PlantEntityModels []PlantEntityModel

// GetPlantModel
// Summary: This is structure which defines GetPlantModel.
// Service: Dataspace
// Router: [GET] /api/v2/authinfo/plant
// Usage: input
type GetPlantModel struct {
	OperatorID uuid.UUID
}

// NewPlantEntityModel
// Summary: This is function which make PlantEntityModel.
// input: operatorID(uuid.UUID) UUID of operatorID
// input: plantName(string) value of plantName
// input: plantAddress(string) value of plantAddress
// input: openPlantID(string) value of openPlantID
// input: globalPlantID(string) pointer of globalPlantID
// output: (PlantEntityModel) PlantEntityModel object
func NewPlantEntityModel(
	operatorID uuid.UUID,
	plantName string,
	plantAddress string,
	openPlantID string,
	globalPlantID *string,
) PlantEntityModel {
	t := time.Now()
	e := PlantEntityModel{
		PlantID:       uuid.New(),
		OperatorID:    operatorID,
		PlantName:     plantName,
		PlantAddress:  plantAddress,
		OpenPlantID:   openPlantID,
		GlobalPlantID: globalPlantID,
		CreatedAt:     t,
		DeletedAt:     gorm.DeletedAt{},
		CreatedUserID: operatorID.String(),
		UpdatedAt:     t,
		UpdatedUserID: operatorID.String(),
	}
	return e
}

type PlantModelInput struct {
	DataModelType string                     `json:"dataModelType"`
	Attribute     *PlantAttributeInputDetail `json:"attribute"`
}

// PlantAttributeInputDetail
// Summary: This is structure which defines PlantAttributeInputDetail.
// Service: Dataspace
// Router: [PUT] /api/v2/authinfo/plant
// Usage: input
type PlantAttributeInputDetail struct {
	PlantID        *string              `json:"plantId"`
	OperatorID     string               `json:"operatorId"`
	PlantName      string               `json:"plantName"`
	PlantAddress   string               `json:"plantAddress"`
	OpenPlantID    *string              `json:"openPlantId"`
	PlantAttribute *PlantAttributeInput `json:"plantAttribute"`
}

// PlantAttributeInput
// Summary: This is structure which defines PlantAttributeInput.
type PlantAttributeInput struct {
	GlobalPlantID *string `json:"globalPlantId"`
}

// Validate
// Summary: This is function which validate value of PutPlantInput.
// output: (error) error object
func (i PlantAttributeInput) Validate() error {
	if err := i.validate(); err != nil {
		logger.Set(nil).Warnf(err.Error())

		return err
	}

	return nil
}

// validate
// Summary: This is function which validate value of PutPlantInput.
// output: (error) error object
func (i PlantModelInput) validate() error {
	errors := []error{}
	err := validation.ValidateStruct(&i,
		validation.Field(
			&i.DataModelType,
			validation.Required,
		),
	)
	if err != nil {
		errors = append(errors, err)
	}
	attribute := *i.Attribute
	attributeDetailErr := validation.ValidateStruct(&attribute,
		validation.Field(
			&attribute.PlantID,
			validation.By(common.StringPtrNilOrUUIDValid),
		),
		validation.Field(
			&attribute.OperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&attribute.PlantName,
			validation.Required,
			validation.RuneLength(1, 256),
		),
		validation.Field(
			&attribute.PlantAddress,
			validation.Required,
			validation.RuneLength(1, 256),
		),
		validation.Field(
			&attribute.OpenPlantID,
			validation.RuneLength(0, 26),
			validation.NotNil,
			validation.By(common.StringPtrLast6CharsNumeric),
		),
		validation.Field(
			&attribute.PlantAttribute,
			validation.Required,
		),
	)

	if attributeDetailErr != nil {
		errors = append(errors, attributeDetailErr)
	}

	var attributeErr error
	if attribute.PlantAttribute != nil {
		if attributeErr = attribute.PlantAttribute.validate(); attributeErr != nil {
			attributeErr = fmt.Errorf("plantAttribute: (%v)", attributeErr)
			errors = append(errors, attributeErr)
		}
	}

	if len(errors) > 0 {
		if attributeErr != nil {
			return common.JoinErrors(errors)
		} else if attributeDetailErr != nil {
			return attributeDetailErr
		} else {
			return err
		}
	}

	return nil
}

// validate
// Summary: This is function which validate value of PlantAttributeInput.
// output: (error) error object
func (i *PlantAttributeInput) validate() error {
	return validation.ValidateStruct(i,
		validation.Field(
			&i.GlobalPlantID,
			validation.RuneLength(0, 256),
		),
	)
}

// Validate
// Summary: This is function which validate value of PutOperatorInput.
// output: (error) error object
func (i PlantModelInput) Validate() error {
	if err := i.validate(); err != nil {
		logger.Set(nil).Warnf(err.Error())

		return err
	}

	return nil
}

// ToModel
// Summary: This is function which convert PutPlantInput to PlantModel.
// output: (PlantModel) PlantModel object
// output: (error) error object
func (i PlantModelInput) ToModel() (PlantModel, error) {
	var m PlantModel

	if i.Attribute.PlantID != nil {
		plantID, err := uuid.Parse(*i.Attribute.PlantID)
		if err != nil {
			logger.Set(nil).Warnf(err.Error())

			return PlantModel{}, fmt.Errorf(common.InvalidUUIDError("plantId"))
		}
		m.Attribute.PlantID = plantID
	}

	operatorID, err := uuid.Parse(i.Attribute.OperatorID)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return PlantModel{}, fmt.Errorf(common.InvalidUUIDError("operatorId"))
	}
	m.Attribute.OperatorID = operatorID

	m.Attribute.PlantName = i.Attribute.PlantName
	m.Attribute.PlantAddress = i.Attribute.PlantAddress
	m.Attribute.OpenPlantID = *i.Attribute.OpenPlantID
	PlantAttribute := PlantAttribute{
		GlobalPlantID: i.Attribute.PlantAttribute.GlobalPlantID,
	}
	m.Attribute.PlantAttribute = PlantAttribute

	return m, nil
}

// Update
// Summary: This is function which update value of Plant.
// input: operatorID(uuid.UUID) UUID of operatorID
// input: plantName(string) value of plantName
// input: plantAddress(string) value of plantAddress
// input: openPlantID(string) value of openPlantID
// input: globalPlantID(string) pointer of globalPlantID
func (e *PlantEntityModel) Update(
	operatorID uuid.UUID,
	plantName string,
	plantAddress string,
	openPlantID string,
	globalPlantID *string,
) {
	e.PlantName = plantName
	e.PlantAddress = plantAddress
	e.OpenPlantID = openPlantID
	e.GlobalPlantID = globalPlantID
	e.UpdatedAt = time.Now()
	e.UpdatedUserID = operatorID.String()
}

func (e PlantEntityModel) ToModel() PlantModel {
	plantAttribute := PlantAttribute{
		GlobalPlantID: e.GlobalPlantID,
	}
	return PlantModel{
		DataModelType: DataModelType,
		Attribute: PlantAttributeDetail{
			PlantID:        e.PlantID,
			OperatorID:     e.OperatorID,
			PlantName:      e.PlantName,
			PlantAddress:   e.PlantAddress,
			OpenPlantID:    e.OpenPlantID,
			PlantAttribute: plantAttribute,
		},
	}
}

// Update
// Summary: This is function which convert PlantEntityModels to array of PlantModel.
// output: globalPlantID(PlantModel) Array of PlantModel
func (es PlantEntityModels) ToModels() []PlantModel {
	ms := []PlantModel{}
	for _, e := range es {
		ms = append(ms, e.ToModel())
	}
	return ms
}
