package getInstallerInfoByDealerOrgCodeOrId

import (
	"GoSungrow/iSolarCloud/api"
	"GoSungrow/iSolarCloud/api/apiReflect"
	"GoSungrow/iSolarCloud/api/valueTypes"
	"github.com/MickMake/GoUnify/Only"
	"fmt"
)

const Url = "/v1/orgService/getInstallerInfoByDealerOrgCodeOrId"
const Disabled = false

type RequestData struct {
	DealerOrgCode valueTypes.String `json:"dealer_org_code"` // required:"true"`
	OrgId         valueTypes.String `json:"org_id"`          // required:"true"`
}

func (rd RequestData) IsValid() error {
	return apiReflect.VerifyOptionsRequired(rd)
}

func (rd RequestData) Help() string {
	ret := fmt.Sprintf("")
	return ret
}


type ResultData struct {
	UserInfoList []struct {
		DealerOrgCode     string      `json:"dealer_org_code"`
		Email             string      `json:"email"`
		Installer         string      `json:"installer"`
		InstallerEmail    string      `json:"installer_email"`
		InstallerPhone    string      `json:"installer_phone"`
		MobleTel          interface{} `json:"moble_tel"`
		OrgId             valueTypes.Integer `json:"org_id"`
		OrgName           string      `json:"org_name"`
		UserId            valueTypes.Integer `json:"user_id"`
		UserName          string      `json:"user_name"`
		UserTelNationCode interface{} `json:"user_tel_nation_code"`
	} `json:"user_info_list" DataTable:"true"`
}

func (e *ResultData) IsValid() error {
	var err error
	// switch {
	// case e.Dummy == "":
	// 	break
	// default:
	// 	err = errors.New(fmt.Sprintf("unknown error '%s'", e.Dummy))
	// }
	return err
}

//type DecodeResultData ResultData
//
//func (e *ResultData) UnmarshalJSON(data []byte) error {
//	var err error
//
//	for range Only.Once {
//		if len(data) == 0 {
//			break
//		}
//		var pd DecodeResultData
//
//		// Store ResultData
//		_ = json.Unmarshal(data, &pd)
//		e.Dummy = pd.Dummy
//	}
//
//	return err
//}

func (e *EndPoint) GetData() api.DataMap {
	entries := api.NewDataMap()

	for range Only.Once {
		entries.StructToDataMap(*e, "system", nil)
	}

	return entries
}
