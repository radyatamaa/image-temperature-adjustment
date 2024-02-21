package response

import (
	"errors"

	"github.com/beego/i18n"
)

/*
Rules penulisan error code

Format : Format XXXX-XXX-XXX
- 4 digit pertama adalah nama service / aplikasi.
- 3 digit selanjutnya adalah sub service / module tersebut.
- 3 digit terakhir adalah kode unik dari error tersebut.

- Contoh : CORE-AUTH-001
		   CORE-API-001
           CORE-KDM-001
*/

const (
	ApiKeyNotRegisteredCodeError    = "ERROR-AUTH-001"
	MissingApiKeyCodeError          = "ERROR-AUTH-002"
	InvalidApiKeyCodeError          = "ERROR-AUTH-003"
	UnauthorizedCodeError           = "ERROR-AUTH-004"
	RequestForbiddenCodeError       = "ERROR-API-001"
	ResourceNotFoundCodeError       = "ERROR-API-002"
	RequestTimeoutCodeError         = "ERROR-API-003"
	ApiValidationCodeError          = "ERROR-API-004"
	DataNotFoundCodeError           = "ERROR-API-005"
	InvalidCredentialCodeError      = "ERROR-API-007"
	InvalidTokenCodeError           = "ERROR-API-008"
	ExpiredTokenCodeError           = "ERROR-API-009"
	MissingTokenCodeError           = "ERROR-API-010"
	AuthElseWhereCodeError          = "ERROR-API-011"
	NotAllowedTransaction           = "ERROR-API-012"
	TransactionAlreadyExist         = "ERROR-API-013"
	TransactionRejected             = "ERROR-API-014"
	TransactionNotFound             = "ERROR-API-015"
	InsufficientLimit               = "ERROR-API-016"
	InvalidReturnAmount             = "ERROR-API-017"
	DataAlreadyExistCodeError       = "ERROR-API-018"
	InvalidMinMax                   = "ERROR-API-019"
	InvalidActiveDate               = "ERROR-API-020"
	CustomerStatusNotFoundErrorCode = "ERROR-API-021"
	LimitStatusNotFoundErrorCode    = "ERROR-API-022"
	CustomerIDNotFoundErrorCode     = "ERROR-API-023"
	TenorIDNotFoundErrorCode        = "ERROR-API-024"
	InvalidActiveEndDate            = "ERROR-API-025"
	QueryParamInvalidCode           = "ERROR-API-026"
	PathParamInvalidCode            = "ERROR-API-027"
	ServerErrorCode                 = "ERROR-API-999"


	InvalidCustomerAndAdminAccount = "ERROR-API-028"

	InvalidCustomerAccount  = "ERROR-API-029"
	InvalidAdminAccount = "ERROR-API-030"
	InvalidEmailPassword = "ERROR-API-031"
	CustomerNotMatchWithOrder = "ERROR-API-032"
	CreateOrderProductIdRequired = "ERROR-API-033"
	ConfirmOrderAlreadyCompleted = "ERROR-API-034"

	InvalidFormatFileJpegErrorCode = "ERROR-API-035"
	RequiredFileErrorCode = "ERROR-API-036"
)

var (
	//query param invalid
	ErrQueryParamInvalid = errors.New("query param is invalid")


	ErrInvalidCustomerAndAdminAccount         = errors.New("token is not customer or admin")
	ErrInvalidCustomerAccount         = errors.New("token is not customer")
	ErrInvalidAdminAccount         = errors.New("token is not admin")


	ErrInvalidEmailPassword     = errors.New("invalid Email and Password")
	ErrCustomerNotMatchWithOrder = errors.New("the order should with customer correct")

	ErrCreateOrderProductIdRequired = errors.New("Product Ids is required")
	ErrConfirmOrderAlreadyCompleted = errors.New("the order is already completed")

	ErrInvalidFormatFileJpeg = errors.New("file format must be JPEG image")
	ErrRequiredFile = errors.New("file required")
)

func ErrorCodeText(code, locale string, args ...interface{}) string {
	switch code {
	case ApiKeyNotRegisteredCodeError:
		return i18n.Tr(locale, "message.errorApiKeyNotRegistered", args)
	case MissingApiKeyCodeError:
		return i18n.Tr(locale, "message.errorMissingApiKey", args)
	case ApiValidationCodeError:
		return i18n.Tr(locale, "message.errorValidation", args)
	case InvalidApiKeyCodeError:
		return i18n.Tr(locale, "message.errorInvalidApiKey", args)
	case UnauthorizedCodeError:
		return i18n.Tr(locale, "message.errorUnauthorized", args)
	case RequestForbiddenCodeError:
		return i18n.Tr(locale, "message.errorRequestForbidden", args)
	case ResourceNotFoundCodeError:
		return i18n.Tr(locale, "message.errorResourceNotFound", args)
	case ServerErrorCode:
		return i18n.Tr(locale, "message.errorServerError", args)
	case RequestTimeoutCodeError:
		return i18n.Tr(locale, "message.errorRequestTimeout", args)
	case InvalidCredentialCodeError:
		return i18n.Tr(locale, "message.errorInvalidCredential", args)
	case DataNotFoundCodeError:
		return i18n.Tr(locale, "message.errorDataNotFound", args)
	case InvalidTokenCodeError:
		return i18n.Tr(locale, "message.errorInvalidToken", args)
	case ExpiredTokenCodeError:
		return i18n.Tr(locale, "message.errorExpiredToken", args)
	case MissingTokenCodeError:
		return i18n.Tr(locale, "message.errorMissingToken", args)
	case AuthElseWhereCodeError:
		return i18n.Tr(locale, "message.errorAuthElseWhere", args)
	case NotAllowedTransaction:
		return i18n.Tr(locale, "message.errorNotAllowedTransaction", args)
	case TransactionAlreadyExist:
		return i18n.Tr(locale, "message.errorTransactionAlreadyExist", args)
	case TransactionRejected:
		return i18n.Tr(locale, "message.errorTransactionRejected", args)
	case TransactionNotFound:
		return i18n.Tr(locale, "message.errorTransactionNotFound", args)
	case InsufficientLimit:
		return i18n.Tr(locale, "message.errorInsufficientLimit", args)
	case InvalidReturnAmount:
		return i18n.Tr(locale, "message.errorInvalidReturnAmount", args)
	case DataAlreadyExistCodeError:
		return i18n.Tr(locale, "message.errorDataAlreadyExist", args)
	case InvalidMinMax:
		return i18n.Tr(locale, "message.errorInvalidMinMax", args)
	case InvalidActiveDate:
		return i18n.Tr(locale, "message.errorActiveMoreThanExpired", args)
	case InvalidActiveEndDate:
		return i18n.Tr(locale, "message.errorActiveMoreThanEnd", args)
	case CustomerStatusNotFoundErrorCode:
		return i18n.Tr(locale, "message.errorCustomerStatusNotFound", args)
	case LimitStatusNotFoundErrorCode:
		return i18n.Tr(locale, "message.errorLimitStatusNotFound", args)
	case CustomerIDNotFoundErrorCode:
		return i18n.Tr(locale, "message.errorCustomerIDNotFound", args)
	case TenorIDNotFoundErrorCode:
		return i18n.Tr(locale, "message.errorTenorIDNotFound", args)
	case QueryParamInvalidCode:
		return i18n.Tr(locale, "message.errorQueryParamInvalid", args)
	case PathParamInvalidCode:
		return i18n.Tr(locale, "message.errorPathParamInvalid", args)
	case InvalidCustomerAndAdminAccount:
		return i18n.Tr(locale, "message.errorInvalidCustomerAndAdminAccount", args)
	case InvalidCustomerAccount:
		return i18n.Tr(locale, "message.errorInvalidCustomerAccount", args)
	case InvalidAdminAccount:
		return i18n.Tr(locale, "message.errorInvalidAdminAccount", args)
	case InvalidEmailPassword:
		return i18n.Tr(locale, "message.errorInvalidEmailPassword", args)
	case CustomerNotMatchWithOrder:
		return i18n.Tr(locale, "message.errorCustomerNotMatchWithOrder", args)
	case CreateOrderProductIdRequired:
		return i18n.Tr(locale, "message.errorCreateOrderProductIdRequired", args)
	case ConfirmOrderAlreadyCompleted:
		return i18n.Tr(locale, "message.errorConfirmOrderAlreadyCompleted", args)
	case InvalidFormatFileJpegErrorCode:
		return i18n.Tr(locale, "message.errorInvalidFormatFileJpeg", args)
	case RequiredFileErrorCode:
		return i18n.Tr(locale, "message.errorRequiredFile", args)
	default:
		return ""
	}
}
