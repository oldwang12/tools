package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	b64 "encoding/base64"
)

type Base64MsgRequest struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

type Base64MsgResponse struct {
	Data string `json:"data"`
}

func Base64Msg(c *gin.Context) {
	var b Base64MsgRequest
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, &Base64MsgResponse{Data: RequestErr})
		return
	}

	if b.Data == "" {
		if b.Type == "encode" { // 编码
			c.JSON(http.StatusBadRequest, &Base64MsgResponse{Data: Base64BadRequestEncode})
		} else if b.Type == "decode" { // 解码
			c.JSON(http.StatusBadRequest, &Base64MsgResponse{Data: Base64BadRequestDecode})
		}
		return
	}

	var resData string
	if b.Type == "encode" {
		resData = b64.URLEncoding.EncodeToString([]byte(b.Data))
	} else if b.Type == "decode" { // 解码
		body, _ := b64.URLEncoding.DecodeString(b.Data)
		resData = string(body)
	} else {
		c.JSON(http.StatusBadRequest, &Base64MsgResponse{Data: RequestErr})
		return
	}

	fmt.Println("resData: ", resData)

	c.JSON(http.StatusOK, &Base64MsgResponse{Data: resData})
}
