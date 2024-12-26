package rpc

import (
	"github.com/sev-2/raiden"
)

type CreateNewUserWhenSignupParams struct {
}
type CreateNewUserWhenSignupResult interface{}

type CreateNewUserWhenSignup struct {
	raiden.RpcBase
	Params   *CreateNewUserWhenSignupParams `json:"-"`
	Return   CreateNewUserWhenSignupResult `json:"-"`
}

func (r *CreateNewUserWhenSignup) GetName() string {
	return "create_new_user_when_signup"
}

func (r *CreateNewUserWhenSignup) GetSecurity() raiden.RpcSecurityType {
	return raiden.RpcSecurityTypeDefiner
}

func (r *CreateNewUserWhenSignup) GetReturnType() raiden.RpcReturnDataType {
	return raiden.RpcReturnDataTypeTrigger
}

func (r *CreateNewUserWhenSignup) GetRawDefinition() string {
	return `begin insert into public.public_users (id, username, role, email) values (new.id, new.email, 'authenticated', new.email); return new; end;`
}