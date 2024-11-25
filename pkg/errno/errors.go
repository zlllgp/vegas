package errno

import (
	"github.com/zlllgp/ecode"
	"github.com/zlllgp/vegas/kitex_gen/api"
)

var (
	Success                = ecode.NewErrNo(int64(api.Err_Success), "success")
	ActivityNotFoundErr    = ecode.NewErrNo(int64(api.Err_ActivityNotFoundRrr), "activity not found")
	CSRFErr                = ecode.NewErrNo(int64(api.Err_CSRFErr), "csrf exception")
	RegisterErr            = ecode.NewErrNo(int64(api.Err_RegisterErr), "user already exists")
	LoginErr               = ecode.NewErrNo(int64(api.Err_LoginErr), "wrong username or password")
	AuthorizationFailedErr = ecode.NewErrNo(int64(api.Err_AuthorizationFailedErr), "authorization failed")
	DataBaseErr            = ecode.NewErrNo(int64(api.Err_DataBaseErr), "database error")
	DrawErr                = ecode.NewErrNo(int64(api.Err_DrawErr), "vegas draw error")
)
