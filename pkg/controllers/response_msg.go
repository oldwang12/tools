package controllers

const (
	RequestErr = "错误的请求"

	// base64
	Base64BadRequestEncode = "加密内容不能为空"
	Base64BadRequestDecode = "解密内容不能为空"

	// ip
	IPResponseContains   = "包含"
	IPUnResponseContains = "不包含"
	CidrFormatErr        = "网段格式有误"
	IPFormatErr          = "IP格式有误"
)
