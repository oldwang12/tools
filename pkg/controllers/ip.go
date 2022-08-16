package controllers

import (
	"net"
	"net/http"
	"net/netip"
	"strings"

	"github.com/apparentlymart/go-cidr/cidr"
	"github.com/gin-gonic/gin"
)

type QueryCidrResponse struct {
	Contains   string `json:"contains"`
	IpLen      uint64 `json:"ip_len"`
	StartEndIP string `json:"start_end_ip"`
	Error      string `json:"error"`
}

func QueryCidr(c *gin.Context) {
	qCidr := c.Query("cidr")
	ip := c.Query("ip")

	cidrs := strings.Split(qCidr, "/")
	if len(cidrs) < 2 {
		c.JSON(http.StatusBadRequest, &QueryCidrResponse{Error: CidrFormatErr})
		return
	}

	if !isIP(cidrs[0]) {
		c.JSON(http.StatusBadRequest, &QueryCidrResponse{Error: CidrFormatErr})
		return
	}

	if !isIP(ip) {
		c.JSON(http.StatusBadRequest, &QueryCidrResponse{Error: IPFormatErr})
		return
	}

	var contains string
	if isCidrContainsIP(qCidr, ip) {
		contains = IPResponseContains
	} else {
		contains = IPUnResponseContains
	}

	_, cidrnet, err := net.ParseCIDR(qCidr)
	if err != nil {
		c.JSON(http.StatusBadRequest, &QueryCidrResponse{Error: CidrFormatErr})
		return
	}

	start, end := cidr.AddressRange(cidrnet)

	c.JSON(http.StatusOK, &QueryCidrResponse{
		Contains:   contains,
		IpLen:      cidr.AddressCount(cidrnet),
		StartEndIP: start.String() + "/" + end.String(),
	})
}

func isCidrContainsIP(cidr, ip string) bool {
	p, err := netip.ParsePrefix(cidr)
	if err != nil {
		return false
	}
	a, err := netip.ParseAddr(ip)
	if err != nil {
		return false
	}
	return p.Contains(a)
}

func isIP(s string) bool {
	return net.ParseIP(s) != nil
}
