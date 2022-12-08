package iSolarCloud

import (
	"GoSungrow/iSolarCloud/AppService/getDeviceList"
	"GoSungrow/iSolarCloud/AppService/getDeviceModelInfoList"
	"GoSungrow/iSolarCloud/AppService/getPowerDevicePointNames"
	"GoSungrow/iSolarCloud/AppService/queryDeviceList"
	"GoSungrow/iSolarCloud/api/GoStruct/valueTypes"
	"errors"
	"fmt"
	"github.com/MickMake/GoUnify/Only"
	datatable "go.pennock.tech/tabular/auto"
	"sort"
	"strings"
	"time"
)


// GetDeviceList - AppService.getDeviceList
func (sg *SunGrow) GetDeviceList(psIds ...string) ([]getDeviceList.Device, error) {
	var ret []getDeviceList.Device

	for range Only.Once {
		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, psId := range pids {
			ep := sg.GetByStruct(getDeviceList.EndPointName,
				getDeviceList.RequestData {
					PsId: psId,
				},
				time.Hour * 24,
			)
			if sg.IsError() {
				break
			}

			data := getDeviceList.Assert(ep)
			ret = append(ret, data.Response.ResultData.PageList...)
		}
	}

	return ret, sg.Error
}

// QueryDeviceList - AppService.queryDeviceList
func (sg *SunGrow) QueryDeviceList(psIds ...string) ([]queryDeviceList.Device, error) {
	var ret []queryDeviceList.Device
	for range Only.Once {
		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}

		for _, psId := range pids {
			ep := sg.GetByStruct(queryDeviceList.EndPointName,
				queryDeviceList.RequestData {
					PsId: psId,
				},
				time.Hour*24,
			)
			if sg.IsError() {
				break
			}

			data := queryDeviceList.Assert(ep)
			ret = append(ret, data.Response.ResultData.PageList...)
		}
	}

	return ret, sg.Error
}


func (sg *SunGrow) Devices(psIds ...string) (string, error) {
	var ret string

	for range Only.Once {
		var devices getDeviceList.Devices
		devices, sg.Error = sg.GetDeviceList()
		if sg.Error != nil {
			break
		}

		ret += fmt.Sprintf("# Devices on ps_id %s:\n", strings.Join(psIds, ", "))
		table := datatable.New("utf8-heavy")
		table.AddHeaders("Ps Key", "Ps Id", "Device Type", "Device Code", "Channel Id", "Serial #", "Factory Name", "Device Model")
		for _, device := range devices {
			table.AddRowItems(device.PsKey, device.PsId, device.DeviceType, device.DeviceCode, device.ChannelId, device.Sn, device.FactoryName, device.DeviceModel)
		}

		ret, sg.Error = table.Render()
		if sg.Error != nil {
			break
		}
	}

	return ret, sg.Error
}


// DeviceData -
func (sg *SunGrow) DeviceData(deviceType string, startDate string, endDate string, interval string) error {
	for range Only.Once {
		if deviceType == "" {
			sg.Error = errors.New("no template defined")
			break
		}

		fmt.Println("NOT YET IMPLEMENTED.")

		// data := sg.QueryUserCurveDeviceData(deviceType)
		// if sg.IsError() {
		// 	break
		// }
		//
		// var points []string
		// for an := range data.PointsData.Devices {
		// 	// fmt.Println(an)
		// 	for _, b := range data.PointsData.Devices[an].Points {
		// 		points = append(points, b.PointId.Full())
		// 		// fmt.Println(bn)
		// 		// fmt.Printf("%v\n", b)
		// 	}
		// }
		//
		// sg.PointData(startDate, endDate, interval, points...)
		// if sg.Error != nil {
		// 	break
		// }
	}

	return sg.Error
}


// DevicePoints - Return all points associated a device_type.
func (sg *SunGrow) DevicePoints(deviceTypes ...string) string {
	var ret string

	for range Only.Once {
		if len(deviceTypes) == 0 {
			break
		}

		table := datatable.New("utf8-heavy")
		table.AddHeaders("Device Type", "Point Id", "Name", "Cal Type")

		var points []getPowerDevicePointNames.Point
		for _, deviceType := range deviceTypes {
			ep := sg.GetByStruct(getPowerDevicePointNames.EndPointName,
				getPowerDevicePointNames.RequestData{ DeviceType: valueTypes.SetIntegerString(deviceType) },
				DefaultCacheTimeout,
			)
			if sg.IsError() {
				break
			}
			data := getPowerDevicePointNames.Assert(ep)
			points = append(points, data.Response.ResultData...)

			// Sort table based on PointId
			pn := map[string]int{}
			for index, point := range points {
				pn[point.PointId.String()] = index
			}
			var names []string
			for point := range pn {
				names = append(names, point)
			}
			sort.Strings(names)

			for _, name := range names {
				index := pn[name]
				point := points[index]
				table.AddRowItems(deviceType, point.PointId, point.PointName, point.PointCalType)
			}
		}
		if sg.IsError() {
			break
		}

		var r string
		r, sg.Error = table.Render()
		if sg.Error != nil {
			break
		}
		ret += fmt.Sprintln("# Available points:")
		ret += r
	}

	return ret
}


// DeviceTypeList - Return all device_types.
func (sg *SunGrow) DeviceTypeList(psIds ...string) string {
	var ret string

	for range Only.Once {
		pids := sg.SetPsIds(psIds...)
		if sg.Error != nil {
			break
		}
		if len(pids) == 0 {
			break
		}


		// data := sg.NewSunGrowData()
		// data.SetEndpoints(queryDeviceList.EndPointName)
		// data.SetArgs(
		// 	"PsId:" + pids.Strings()[0],
		// )
		// sg.Error = data.GetData()
		// if sg.Error != nil {
		// 	break
		// }
		//
		// sg.Error = data.OutputDataTables()
		// if sg.Error != nil {
		// 	break
		// }


		table := datatable.New("utf8-heavy")
		table.AddHeaders("Device Type", "Name")

		ep := sg.GetByStruct(queryDeviceList.EndPointName,
			queryDeviceList.RequestData{ PsId: pids[0] },
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}
		data := queryDeviceList.Assert(ep)

		// Sort table based on PointId
		var names []string
		for name := range data.Response.ResultData.DevTypeDefinition {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			table.AddRowItems(name, data.Response.ResultData.DevTypeDefinition[name])
		}

		var r string
		r, sg.Error = table.Render()
		if sg.Error != nil {
			break
		}
		ret += fmt.Sprintln("# Available points:")
		ret += r
	}

	return ret
}


func (sg *SunGrow) GetPowerDevicePointNames(device valueTypes.Integer) ([]getPowerDevicePointNames.Point, error) {
	var ret []getPowerDevicePointNames.Point
	for range Only.Once {
		ep := sg.GetByStruct(getPowerDevicePointNames.EndPointName,
			getPowerDevicePointNames.RequestData{DeviceType: device},
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}

		data := getPowerDevicePointNames.Assert(ep)
		ret = data.Response.ResultData
	}

	return ret, sg.Error
}


func (sg *SunGrow) DeviceModelInfoList() error {
	for range Only.Once {
		ep := sg.GetByStruct(getDeviceModelInfoList.EndPointName,
			getDeviceModelInfoList.RequestData{},
			DefaultCacheTimeout,
		)
		if sg.IsError() {
			break
		}

		data := getDeviceModelInfoList.Assert(ep)
		table := data.GetPointDataTable()
		if table.Error != nil {
			sg.Error = table.Error
			break
		}

		table.SetTitle("Models")
		table.SetFilePrefix(data.SetFilenamePrefix(""))
		table.SetGraphFilter("")
		table.SetSaveFile(sg.SaveAsFile)
		table.OutputType = sg.OutputType
		sg.Error = table.Output()
		if sg.IsError() {
			break
		}
	}

	return sg.Error
}


// func (sg *SunGrow) GetDevices(psId valueTypes.PsId) (getDeviceList.Devices, error) {
// 	var ret getDeviceList.Devices
//
// 	for range Only.Once {
// 		// ret = append(ret, getDeviceList.Device{
// 		// 	Vendor:        valueTypes.SetStringValue(""),
// 		// 	PsId:          psId.PsId,
// 		// 	PsKey:         valueTypes.SetPsKeyString(psId.PsId.String()),
// 		// 	DeviceName:    psId.PsName,
// 		// 	DeviceProSn:   psId.PsShortName,
// 		// 	DeviceModel:   valueTypes.SetStringValue(""),
// 		// 	DeviceType:    psId.PsType,
// 		// 	DeviceCode:    valueTypes.SetIntegerValue(0),
// 		// 	ChannelId:     valueTypes.SetIntegerValue(0),
// 		// 	DeviceModelId: valueTypes.SetIntegerValue(0),
// 		// 	TypeName:      valueTypes.SetStringValue("Ps Id"),
// 		// 	DeviceState:   psId.PsHealthStatus,
// 		// 	DevStatus:     psId.PsStatus,
// 		// 	Uuid:          valueTypes.SetIntegerValue(0),
// 		//
// 		// 	// 			PsFaultStatus:  d.PsFaultStatus,
// 		// 	//			PsHealthStatus: d.PsHealthStatus,
// 		// 	//			PsHolder:       d.PsHolder,
// 		// 	//			PsId:           d.PsId,
// 		// 	//			PsName:         d.PsName,
// 		// 	//			PsShortName:    d.PsShortName,
// 		// 	//			PsStatus:       d.PsStatus,
// 		// 	//			PsType:         d.PsType,
// 		// })
//
// 		ep := sg.GetByStruct(
// 			"AppService.getDeviceList",
// 			getDeviceList.RequestData{PsId: psId},
// 			DefaultCacheTimeout,
// 		)
// 		if sg.IsError() {
// 			break
// 		}
//
// 		data := getDeviceList.Assert(ep)
// 		ret = data.GetDevices()
// 	}
//
// 	return ret, sg.Error
// }

// func (sg *SunGrow) GetDeviceList(psIds ...string) error {
// 	for range Only.Once {
// 		data := sg.NewSunGrowData()
// 		data.SetPsIds(psIds...)
// 		data.SetEndpoints("getDeviceList")
//
// 		sg.Error = data.GetData()
// 		if sg.IsError() {
// 			break
// 		}
//
// 		sg.Error = data.GetOutput()
// 		if sg.IsError() {
// 			break
// 		}
//
// 		// var ret getDeviceList.Devices
// 		// for _, psId := range psIds {
// 		// 	ep := sg.GetByStruct(
// 		// 		"AppService.getDeviceList",
// 		// 		// getDeviceList.RequestData{PsId: strconv.FormatInt(psId, 10)},
// 		// 		getDeviceList.RequestData{PsId: psId},
// 		// 		DefaultCacheTimeout,
// 		// 	)
// 		// 	if sg.Error != nil {
// 		// 		break
// 		// 	}
// 		//
// 		// 	data := getDeviceList.Assert(ep)
// 		// 	ret = append(ret, data.GetDevices()...)
// 		// }
// 		//
// 		// table := getDeviceList.GetDevicesTable(ret)
// 		// table.SetTitle("All Devices")
// 		// table.SetFilePrefix("")
// 		// table.SetGraphFilter("")
// 		// table.SetSaveFile(sg.SaveAsFile)
// 		// table.OutputType = sg.OutputType
// 		// sg.Error = table.Output()
// 		// if sg.Error != nil {
// 		// 	break
// 		// }
// 	}
//
// 	return sg.Error
// }
