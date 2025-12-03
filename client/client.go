// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("umeng-push"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据消息ID取消发送
//
// @param request - CancelByMsgIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelByMsgIdResponse
func (client *Client) CancelByMsgIdWithOptions(request *CancelByMsgIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelByMsgIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelByMsgId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/CancelByMsgId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelByMsgIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据消息ID取消发送
//
// @param request - CancelByMsgIdRequest
//
// @return CancelByMsgIdResponse
func (client *Client) CancelByMsgId(request *CancelByMsgIdRequest) (_result *CancelByMsgIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelByMsgIdResponse{}
	_body, _err := client.CancelByMsgIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 消息状态查询
//
// @param request - QueryMsgStatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryMsgStatResponse
func (client *Client) QueryMsgStatWithOptions(request *QueryMsgStatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryMsgStatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MsgId) {
		body["MsgId"] = request.MsgId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryMsgStat"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/QueryMsgStat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryMsgStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 消息状态查询
//
// @param request - QueryMsgStatRequest
//
// @return QueryMsgStatResponse
func (client *Client) QueryMsgStat(request *QueryMsgStatRequest) (_result *QueryMsgStatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMsgStatResponse{}
	_body, _err := client.QueryMsgStatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定别名发送
//
// @param tmpReq - SendByAliasRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasResponse
func (client *Client) SendByAliasWithOptions(tmpReq *SendByAliasRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAliasResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAliasShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.AliasType) {
		body["AliasType"] = request.AliasType
	}

	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByAlias"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByAlias"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAliasResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名发送
//
// @param request - SendByAliasRequest
//
// @return SendByAliasResponse
func (client *Client) SendByAlias(request *SendByAliasRequest) (_result *SendByAliasResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAliasResponse{}
	_body, _err := client.SendByAliasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定别名文件发送
//
// @param tmpReq - SendByAliasFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAliasFileIdResponse
func (client *Client) SendByAliasFileIdWithOptions(tmpReq *SendByAliasFileIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAliasFileIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAliasFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AliasType) {
		body["AliasType"] = request.AliasType
	}

	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByAliasFileId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByAliasFileId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAliasFileIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定别名文件发送
//
// @param request - SendByAliasFileIdRequest
//
// @return SendByAliasFileIdResponse
func (client *Client) SendByAliasFileId(request *SendByAliasFileIdRequest) (_result *SendByAliasFileIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAliasFileIdResponse{}
	_body, _err := client.SendByAliasFileIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 广播
//
// @param tmpReq - SendByAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByAppResponse
func (client *Client) SendByAppWithOptions(tmpReq *SendByAppRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByAppResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByAppShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByApp"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByApp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 广播
//
// @param request - SendByAppRequest
//
// @return SendByAppResponse
func (client *Client) SendByApp(request *SendByAppRequest) (_result *SendByAppResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByAppResponse{}
	_body, _err := client.SendByAppWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定设备发送
//
// @param tmpReq - SendByDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceResponse
func (client *Client) SendByDeviceWithOptions(tmpReq *SendByDeviceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByDeviceShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.DeviceTokens) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByDevice"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备发送
//
// @param request - SendByDeviceRequest
//
// @return SendByDeviceResponse
func (client *Client) SendByDevice(request *SendByDeviceRequest) (_result *SendByDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByDeviceResponse{}
	_body, _err := client.SendByDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 指定设备文件发送
//
// @param tmpReq - SendByDeviceFileIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByDeviceFileIdResponse
func (client *Client) SendByDeviceFileIdWithOptions(tmpReq *SendByDeviceFileIdRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByDeviceFileIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByDeviceFileIdShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.AndroidShortPayload) {
		request.AndroidShortPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidShortPayload, dara.String("AndroidShortPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayloadShrink) {
		body["AndroidShortPayload"] = request.AndroidShortPayloadShrink
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileId) {
		body["FileId"] = request.FileId
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByDeviceFileId"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByDeviceFileId"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByDeviceFileIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 指定设备文件发送
//
// @param request - SendByDeviceFileIdRequest
//
// @return SendByDeviceFileIdResponse
func (client *Client) SendByDeviceFileId(request *SendByDeviceFileIdRequest) (_result *SendByDeviceFileIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByDeviceFileIdResponse{}
	_body, _err := client.SendByDeviceFileIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据筛选条件发送
//
// @param tmpReq - SendByFilterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendByFilterResponse
func (client *Client) SendByFilterWithOptions(tmpReq *SendByFilterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendByFilterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &SendByFilterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AndroidPayload) {
		request.AndroidPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AndroidPayload, dara.String("AndroidPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ChannelProperties) {
		request.ChannelPropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ChannelProperties, dara.String("ChannelProperties"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.HarmonyPayload) {
		request.HarmonyPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HarmonyPayload, dara.String("HarmonyPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.IosPayload) {
		request.IosPayloadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IosPayload, dara.String("IosPayload"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Policy) {
		request.PolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Policy, dara.String("Policy"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AndroidPayloadShrink) {
		body["AndroidPayload"] = request.AndroidPayloadShrink
	}

	if !dara.IsNil(request.AndroidShortPayload) {
		body["AndroidShortPayload"] = request.AndroidShortPayload
	}

	if !dara.IsNil(request.ChannelPropertiesShrink) {
		body["ChannelProperties"] = request.ChannelPropertiesShrink
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Filter) {
		body["Filter"] = request.Filter
	}

	if !dara.IsNil(request.HarmonyPayloadShrink) {
		body["HarmonyPayload"] = request.HarmonyPayloadShrink
	}

	if !dara.IsNil(request.IosPayloadShrink) {
		body["IosPayload"] = request.IosPayloadShrink
	}

	if !dara.IsNil(request.PolicyShrink) {
		body["Policy"] = request.PolicyShrink
	}

	if !dara.IsNil(request.ProductionMode) {
		body["ProductionMode"] = request.ProductionMode
	}

	if !dara.IsNil(request.ReceiptType) {
		body["ReceiptType"] = request.ReceiptType
	}

	if !dara.IsNil(request.ReceiptUrl) {
		body["ReceiptUrl"] = request.ReceiptUrl
	}

	if !dara.IsNil(request.ThirdPartyId) {
		body["ThirdPartyId"] = request.ThirdPartyId
	}

	if !dara.IsNil(request.CallbackParams) {
		body["callbackParams"] = request.CallbackParams
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendByFilter"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/SendByFilter"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据筛选条件发送
//
// @param request - SendByFilterRequest
//
// @return SendByFilterResponse
func (client *Client) SendByFilter(request *SendByFilterRequest) (_result *SendByFilterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendByFilterResponse{}
	_body, _err := client.SendByFilterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传设备列表创建设备文件
//
// @param request - UploadDeviceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDeviceResponse
func (client *Client) UploadDeviceWithOptions(request *UploadDeviceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadDeviceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceTokens) {
		body["DeviceTokens"] = request.DeviceTokens
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDevice"),
		Version:     dara.String("2022-02-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/UploadDevice"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDeviceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传设备列表创建设备文件
//
// @param request - UploadDeviceRequest
//
// @return UploadDeviceResponse
func (client *Client) UploadDevice(request *UploadDeviceRequest) (_result *UploadDeviceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadDeviceResponse{}
	_body, _err := client.UploadDeviceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
