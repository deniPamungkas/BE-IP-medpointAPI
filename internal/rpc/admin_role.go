package rpc

// import (
// 	"github.com/sev-2/raiden"
// 	"ipmedpointsistem/internal/models"
// )

// type AdminRoleParams struct {
// }
// type AdminRoleResult bool

// type AdminRole struct {
// 	raiden.RpcBase
// 	Params   *AdminRoleParams `json:"-"`
// 	Return   AdminRoleResult `json:"-"`
// }

// func (r *AdminRole) GetName() string {
// 	return "admin_role"
// }

// func (r *AdminRole) GetSecurity() raiden.RpcSecurityType {
// 	return raiden.RpcSecurityTypeDefiner
// }

// func (r *AdminRole) GetReturnType() raiden.RpcReturnDataType {
// 	return raiden.RpcReturnDataTypeBoolean
// }

// func (r *AdminRole) BindModels() {
// 	r.BindModel(models.PublicUsers{}, "p")
// }

// func (r *AdminRole) GetRawDefinition() string {
// 	return `begin return exists ( select 1 from :p where (select auth.uid()) = id and role = 'admin' ); end;`
// }