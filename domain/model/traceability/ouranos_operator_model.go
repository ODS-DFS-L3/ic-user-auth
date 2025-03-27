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

var DataModelType = "test1"

// OperatorModel
// Summary: This is structure which defines OperatorModel.
// Service: Dataspace
// Router: [PUT] /api/v2/authinfo/operator
// Usage: output
type OperatorModel struct {
	DataModelType string                  `json:"dataModelType"`
	Attribute     OperatorAttributeDetail `json:"attribute"`
}

type OperatorAttributeDetail struct {
	OperatorID        uuid.UUID         `json:"operatorId"`
	OperatorName      string            `json:"operatorName"`
	OperatorAddress   string            `json:"operatorAddress"`
	OpenOperatorID    string            `json:"openOperatorId"`
	OperatorAttribute OperatorAttribute `json:"operatorAttribute"`
}

// OperatorAttribute
// Summary: This is structure which defines OperatorAttribute.
type OperatorAttribute struct {
	GlobalOperatorID *string `json:"globalOperatorId,omitempty"`
}

// OperatorModels
// Summary: This is structure which defines OperatorModels.
// Service: Dataspace
// Router: [GET] /api/v2/authinfo/operator
// Usage: output
type OperatorModels []*OperatorModel

type OperatorModelInput struct {
	DataModelType string                        `json:"dataModelType"`
	Attribute     *OperatorAttributeInputDetail `json:"attribute"`
}

// PutOperatorInput
// Summary: This is structure which defines PutOperatorInput.
// Service: Dataspace
// Router: [PUT] /api/v2/authinfo/operator
// Usage: input
type OperatorAttributeInputDetail struct {
	OperatorID             string                  `json:"operatorId"`
	OperatorName           string                  `json:"operatorName"`
	OperatorAddress        string                  `json:"operatorAddress"`
	OpenOperatorID         string                  `json:"openOperatorId"`
	OperatorAttributeInput *OperatorAttributeInput `json:"operatorAttribute"`
}

// OperatorAttributeInput
// Summary: This is structure which defines OperatorAttributeInput.
type OperatorAttributeInput struct {
	GlobalOperatorID *string `json:"globalOperatorId"`
}

// OperatorEntityModel
// Summary: This is structure which defines OperatorEntityModel.
// DBName: operators
type OperatorEntityModel struct {
	OperatorID        uuid.UUID      `json:"operatorId" gorm:"type:uuid;not null"`
	OperatorName      string         `json:"operatorName" gorm:"type:string"`
	OperatorAddress   string         `json:"operatorAddress" gorm:"type:string"`
	OpenOperatorID    string         `json:"openOperatorId" gorm:"type:string"`
	GlobalOperatorID  *string        `json:"globalOperatorId" gorm:"type:string"`
	DeletedAt         gorm.DeletedAt `json:"deletedAt"`
	CreatedAt         time.Time      `json:"createdAt" gorm:"<-:create "`
	CreatedOperatorID string         `json:"createdUserId" gorm:"type:varchar(256);not null; <-:create"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	UpdatedOperatorID string         `json:"updatedUserId" gorm:"type:varchar(256);not null"`
}

// OperatorEntityModels
// Summary: This is structure which defines OperatorEntityModels.
type OperatorEntityModels []*OperatorEntityModel

// GetOperatorInput
// Summary: This is structure which defines GetOperatorInput.
// Service: Dataspace
// Router: [GET] /api/v2/authinfo/operator
// Usage: input
type GetOperatorInput struct {
	OperatorID     string  `json:"operatorId"`
	OpenOperatorID *string `json:"openOperatorId"`
}

// GetOperatorsInput
// Summary: This is structure which defines GetOperatorsInput.
// Service: Dataspace
// Router: [GET] /api/v2/authinfo/operator
// Usage: input
type GetOperatorsInput struct {
	OperatorIDs []uuid.UUID `json:"operatorIds"`
}

// validate
// Summary: This is function which validate value of PutOperatorInput.
// output: (error) error object
func (i OperatorModelInput) validate() error {
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
			&attribute.OperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&attribute.OperatorName,
			validation.Required,
			validation.RuneLength(1, 256),
		),
		validation.Field(
			&attribute.OperatorAddress,
			validation.Required,
			validation.RuneLength(1, 256),
		),
		validation.Field(
			&attribute.OpenOperatorID,
			validation.Required,
			validation.RuneLength(1, 20),
		),
		validation.Field(
			&attribute.OperatorAttributeInput,
			validation.Required,
		),
	)
	if attributeDetailErr != nil {
		errors = append(errors, attributeDetailErr)
	}

	var attributeErr error
	if attribute.OperatorAttributeInput != nil {
		if attributeErr = attribute.OperatorAttributeInput.validate(); attributeErr != nil {
			attributeErr = fmt.Errorf(common.ValidateStructureError("operatorAttribute", attributeErr))
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
// Summary: This is function which validate value of OperatorAttributeInput.
// output: (error) error object
func (i *OperatorAttributeInput) validate() error {
	return validation.ValidateStruct(i,
		validation.Field(
			&i.GlobalOperatorID,
			validation.RuneLength(0, 256),
		),
	)
}

// Validate
// Summary: This is function which validate value of PutOperatorInput.
// output: (error) error object
func (i OperatorModelInput) Validate() error {
	if err := i.validate(); err != nil {
		logger.Set(nil).Warnf(err.Error())

		return err
	}

	return nil
}

// ToModel
// Summary: This is function which convert PutOperatorInput to OperatorModel.
// output: (OperatorModel) OperatorModel object
// output: (error) error object
func (i OperatorModelInput) ToModel() (OperatorModel, error) {
	operatorID, err := uuid.Parse(i.Attribute.OperatorID)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return OperatorModel{}, fmt.Errorf(common.InvalidUUIDError("operatorId"))
	}

	operatorAttribute := OperatorAttribute{
		GlobalOperatorID: i.Attribute.OperatorAttributeInput.GlobalOperatorID,
	}
	m := OperatorModel{
		DataModelType: i.DataModelType,
		Attribute: struct {
			OperatorID        uuid.UUID         `json:"operatorId"`
			OperatorName      string            `json:"operatorName"`
			OperatorAddress   string            `json:"operatorAddress"`
			OpenOperatorID    string            `json:"openOperatorId"`
			OperatorAttribute OperatorAttribute `json:"operatorAttribute"`
		}{
			OperatorID:        operatorID,
			OperatorName:      i.Attribute.OperatorName,
			OperatorAddress:   i.Attribute.OperatorAddress,
			OpenOperatorID:    i.Attribute.OpenOperatorID,
			OperatorAttribute: operatorAttribute,
		},
	}
	return m, nil
}

// ToModel
// Summary: This is function which convert OperatorEntityModel to OperatorModel.
// output: (OperatorModel) OperatorModel object
func (e *OperatorEntityModel) ToModel() OperatorModel {
	OperatorAttribute := OperatorAttribute{
		GlobalOperatorID: e.GlobalOperatorID,
	}

	return OperatorModel{
		DataModelType: DataModelType,
		Attribute: OperatorAttributeDetail{
			OperatorID:        e.OperatorID,
			OperatorName:      e.OperatorName,
			OperatorAddress:   e.OperatorAddress,
			OpenOperatorID:    e.OpenOperatorID,
			OperatorAttribute: OperatorAttribute,
		},
	}
}

// ToModels
// Summary: This is function which convert OperatorEntityModels to []OperatorModel.
// output: ([]OperatorModel) slice of OperatorModel object
func (es OperatorEntityModels) ToModels() []OperatorModel {
	ms := make([]OperatorModel, len(es))
	for i, e := range es {
		m := e.ToModel()
		ms[i] = m
	}
	return ms
}

// Update
// Summary: This is function which update value of Operator.
// input: operatorModel(OperatorModel) OperatorModel object
// output: (error) error object
func (e *OperatorEntityModel) Update(operatorModel OperatorModel) error {
	if e.OpenOperatorID != operatorModel.Attribute.OpenOperatorID {
		err := fmt.Errorf(common.FieldIsImutable("openOperatorId"))
		logger.Set(nil).Warnf(err.Error())

		return common.NewCustomError(common.CustomErrorCode400, err.Error(), nil, common.HTTPErrorSourceAuth)
	}

	e.OperatorName = operatorModel.Attribute.OperatorName
	e.OperatorAddress = operatorModel.Attribute.OperatorAddress
	e.GlobalOperatorID = operatorModel.Attribute.OperatorAttribute.GlobalOperatorID
	e.UpdatedAt = time.Now()

	return nil
}

// Validate
// Summary: This is function which validate value of GetOperatorInput.
// output: (error) error object
func (i GetOperatorInput) Validate() error {
	if err := i.validate(); err != nil {
		logger.Set(nil).Warnf(err.Error())

		return err
	}
	return nil
}

// validate
// Summary: This is function which validate value of GetOperatorInput.
// output: (error) error object
func (i GetOperatorInput) validate() error {
	// Corporate number is 13 digits, verified only if not nil
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.OperatorID,
			validation.By(common.StringUUIDValid),
		),
		validation.Field(
			&i.OpenOperatorID,
			validation.When(i.OpenOperatorID != nil, validation.RuneLength(1, 256)),
		),
	)
}
