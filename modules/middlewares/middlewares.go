package middlewares

type middlewareHandlersErrCode string

const (
	RouterCheckErr middlewareHandlersErrCode = "middlware-001"
	JwtAuthErr     middlewareHandlersErrCode = "middlware-002"
	ParamsCheckErr middlewareHandlersErrCode = "middlware-003"
	AuthorizeErr   middlewareHandlersErrCode = "middlware-004"
	ApiKeyErr      middlewareHandlersErrCode = "middlware-005"
)
