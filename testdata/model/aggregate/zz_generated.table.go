/*
Package aggregate GENERATED BY gengo:table 
DON'T EDIT THIS FILE
*/
package aggregate

import (
	datatypes "github.com/octohelm/storage/pkg/datatypes"
	sqlbuilder "github.com/octohelm/storage/pkg/sqlbuilder"
	modelscoped "github.com/octohelm/storage/pkg/sqlbuilder/modelscoped"
	model "github.com/octohelm/storage/testdata/model"
)

func (CountedUser) TableName() string {
	return "t_counted_user"
}

func (tableCountedUser) New() sqlbuilder.Model {
	return &CountedUser{}
}

type tableCountedUser struct {
	modelscoped.Table[CountedUser]

	I indexesOfCountedUser

	// 用户ID
	ID modelscoped.TypedColumn[CountedUser, model.UserID]
	// 姓名
	Name modelscoped.TypedColumn[CountedUser, string]
	// 昵称
	Nickname modelscoped.TypedColumn[CountedUser, string]
	// 用户名
	Username modelscoped.TypedColumn[CountedUser, string]

	Gender modelscoped.TypedColumn[CountedUser, model.Gender]
	// 年龄
	Age modelscoped.TypedColumn[CountedUser, int64]

	CreatedAt modelscoped.TypedColumn[CountedUser, datatypes.Datetime]

	UpdatedAt modelscoped.TypedColumn[CountedUser, int64]

	DeletedAt modelscoped.TypedColumn[CountedUser, int64]

	Count modelscoped.TypedColumn[CountedUser, int]
}

type indexesOfCountedUser struct {
}

var CountedUserT = &tableCountedUser{
	Table: modelscoped.FromModel[CountedUser](),

	ID:        modelscoped.CastTypedColumn[CountedUser, model.UserID](modelscoped.FromModel[CountedUser]().F("ID")),
	Name:      modelscoped.CastTypedColumn[CountedUser, string](modelscoped.FromModel[CountedUser]().F("Name")),
	Nickname:  modelscoped.CastTypedColumn[CountedUser, string](modelscoped.FromModel[CountedUser]().F("Nickname")),
	Username:  modelscoped.CastTypedColumn[CountedUser, string](modelscoped.FromModel[CountedUser]().F("Username")),
	Gender:    modelscoped.CastTypedColumn[CountedUser, model.Gender](modelscoped.FromModel[CountedUser]().F("Gender")),
	Age:       modelscoped.CastTypedColumn[CountedUser, int64](modelscoped.FromModel[CountedUser]().F("Age")),
	CreatedAt: modelscoped.CastTypedColumn[CountedUser, datatypes.Datetime](modelscoped.FromModel[CountedUser]().F("CreatedAt")),
	UpdatedAt: modelscoped.CastTypedColumn[CountedUser, int64](modelscoped.FromModel[CountedUser]().F("UpdatedAt")),
	DeletedAt: modelscoped.CastTypedColumn[CountedUser, int64](modelscoped.FromModel[CountedUser]().F("DeletedAt")),
	Count:     modelscoped.CastTypedColumn[CountedUser, int](modelscoped.FromModel[CountedUser]().F("Count")),

	I: indexesOfCountedUser{},
}
