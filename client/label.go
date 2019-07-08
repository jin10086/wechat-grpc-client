package client

import (
	"encoding/json"
)

const (
	CmdGetLabelList = 639
	CmdAddLabel     = 635
	CmdSetUserLabel = 638
)

/**
 * 获取标签列表
 */
func GetLabels(wxUser string) interface{} {
	var (
		result []byte
		err    error
	)
	FastShortRequest(wxUser, CmdGetLabelList, []byte{}, "获取标签列表成功", &result, &err)
	if err != nil {
		return nil
	}
	var responseData interface{}
	json.Unmarshal(result, &responseData)
	return responseData
}

/**
 * 新增标签
 */
func AddLabel(wxUser string, name string) interface{} {
	var (
		result []byte
		err    error
	)
	FastShortRequest(wxUser, CmdAddLabel, []byte(`{
		"LabelName": "`+name+`"
	}`), "新增标签列表成功", &result, &err)
	if err != nil {
		return nil
	}
	var responseData interface{}
	json.Unmarshal(result, &responseData)
	return responseData
}

//不可用
func SetUserLabel(wxUser string, user string, labelIds string) bool {
	var (
		result []byte
		err    error
	)
	FastShortRequest(wxUser, CmdSetUserLabel, []byte(`{
		"Username": "`+user+`",
		"Labelids": "`+labelIds+`"
	}`), "设置用户标签成功", &result, &err)
	if err != nil {
		return false
	}
	return true
}
