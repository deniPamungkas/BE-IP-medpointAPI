package rpc

import (
	"github.com/sev-2/raiden"
	"ipmedpointsistem/internal/models"
)

type IsAdminParams struct {
}
type IsAdminResult bool

type IsAdmin struct {
	raiden.RpcBase
	Params   *IsAdminParams `json:"-"`
	Return   IsAdminResult `json:"-"`
}

func (r *IsAdmin) GetName() string {
	return "is_admin"
}

func (r *IsAdmin) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeBoolean
}

func (r *IsAdmin) BindModels() {
	r.BindModel(models.PublicUsers{}, "p")
}

func (r *IsAdmin) GetRawDefinition() string {
	return `declare user_role text; begin select role into user_role from :p where id = auth.uid(); if user_role = 'admin' then return true; else return false; end if; end;`
}