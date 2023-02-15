package mini

import (
	"context"
	"fmt"

	"github.com/medreams/wechat/common"
)

type VehicleLicenseData struct {
	common.WxCommonResponse
	VehicleType    string `json:"vehicle_type"`    //车辆类型
	Owner          string `json:"owner"`           //所有人
	Addr           string `json:"add"`             //住址
	UseCharacter   string `json:"use_character"`   //使用性质
	Model          string `json:"model"`           //品牌型号
	Vin            string `json:"vin"`             //车辆识别代号
	EngineNum      string `json:"engine_num"`      //发动机号码
	RegisterDate   string `json:"register_date"`   //注册日期
	IssueDate      string `json:"issue_date"`      //发证日期
	PlateNumB      string `json:"plate_num_b"`     //车牌号码
	Record         string `json:"record"`          //号牌
	PassengersNum  string `json:"passengers_num"`  //核定载人数
	TotalQuality   string `json:"total_quality"`   //总质量
	PrepareQuality string `json:"prepare_quality"` //整备质量
}

// 行驾证识别 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr/ocr/vehicleLicenseOCR.html
func (sdk *SDK) OcrVehicleLicense(ctx context.Context, imgUrl string) (req *VehicleLicenseData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap["img_url"] = imgUrl

	req = &VehicleLicenseData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/driving?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, req); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	return req, nil
}

type BankCardData struct {
	common.WxCommonResponse
	Number string `json:"number"` //银行卡号
}

// 银行卡识别 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr/ocr/bankCardOCR.html
func (sdk *SDK) OcrBankCard(ctx context.Context, imgUrl string) (bc *BankCardData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("img_url", imgUrl)

	bc = &BankCardData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/bankcard?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, bc); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	return bc, nil
}

type BusinessLicenseData struct {
	common.WxCommonResponse
	RegNum              string `json:"regnum"`               //注册号
	Serial              string `json:"serial"`               //编号
	LegalRepresentative string `json:"legal_representative"` //法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      //企业名称
	TypeOfOrganization  string `json:"type_of_organization"` //组成形式
	Address             string `json:"address"`              //经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"`   //公司类型
	BusinessScope       string `json:"business_scope"`       //经营范围
	RegisteredCapital   string `json:"registered_capital"`   //注册资本
	PaidInCapital       string `json:"paid_in_capital"`      //实收资本
	ValidPeriod         string `json:"valid_period"`         //营业期限
	RegisteredDate      string `json:"registered_date"`      //注册日期/成立日期
}

func (sdk *SDK) OcrBusinessLicense(ctx context.Context, imgUrl string) (bl *BusinessLicenseData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("img_url", imgUrl)

	bl = &BusinessLicenseData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/bizlicense?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, bl); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	return bl, nil
}

type DriverLicenseData struct {
	common.WxCommonResponse
	IdNum        string `json:"id_num"`        //证号
	Name         string `json:"name"`          //姓名
	Sex          string `json:"sex"`           //性别
	Address      string `json:"address"`       //地址
	BirthDate    string `json:"birth_date"`    //出生日期
	IssueDate    string `json:"issue_date"`    //初次领证日期
	CarClass     string `json:"car_class"`     //准驾车型
	ValidFrom    string `json:"valid_from"`    //有效期限起始日
	ValidTo      string `json:"valid_to"`      //有效期限终止日
	OfficialSeal string `json:"official_seal"` //印章文构
}

// 驾驶证识别 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr/ocr/driverLicenseOCR.html
func (sdk *SDK) OcrDriverLicense(ctx context.Context, imgUrl string) (dl *DriverLicenseData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("img_url", imgUrl)

	dl = &DriverLicenseData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/bizlicense?access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, dl); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	return dl, nil
}

type IdCardData struct {
	common.WxCommonResponse
	Type        string `json:"type"`        //正面或背面，Front / Back
	Name        string `json:"name"`        //正面返回，姓名
	Id          string `json:"id"`          //正面返回，身份证号
	ValidDate   string `json:"valid_date"`  //背面返回，有效期
	Addr        string `json:"addr"`        //正面返回，地址
	Gender      string `json:"gender"`      //正面返回，性别
	Nationality string `json:"nationality"` //正面返回，民族
}

// 身份证识别 https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/img-ocr/ocr/idCardOCR.html
func (sdk *SDK) OcrIdCard(ctx context.Context, imgUrl string) (id *IdCardData, err error) {

	bodyMap := make(common.BodyMap)
	bodyMap.Set("img_url", imgUrl)

	id = &IdCardData{}
	uri := fmt.Sprintf("https://api.weixin.qq.com/cv/ocr/idcard?type=photo&access_token=%s", sdk.AccessToken)

	if err = common.DoRequestPost(ctx, uri, bodyMap, id); err != nil {
		return nil, fmt.Errorf("do request get phone: %w", err)
	}

	return id, nil
}
