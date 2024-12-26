package rpc

import (
	"github.com/sev-2/raiden"
)

type CreateUserRoleParams struct {
}
type CreateUserRoleResult interface{}

type CreateUserRole struct {
	raiden.RpcBase
	Params   *CreateUserRoleParams `json:"-"`
	Return   CreateUserRoleResult `json:"-"`
}

func (r *CreateUserRole) GetName() string {
	return "create_user_role"
}

func (r *CreateUserRole) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeTrigger
}

func (r *CreateUserRole) GetRawDefinition() string {
	return `begin insert into public.user_role (user_id) values (new.id); return new; end;`
}