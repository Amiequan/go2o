package app

import (
	"context"
	"fmt"
	"github.com/ixre/go2o/core/domain/interface/registry"
	"github.com/ixre/go2o/core/service"
	"github.com/ixre/go2o/core/service/proto"
	"github.com/ixre/gof/types"
	"log"
	"os"
)

/**
 * Copyright (C) 2007-2020 56X.NET,All rights reserved.
 *
 * name : domain
 * author : jarrysix (jarrysix#gmail.com)
 * date : 2020-10-09 17:18
 * description :
 * history :
 */

func reFn(cli proto.RegistryServiceClient, k string) string {
	ret, _ := cli.GetValue(context.TODO(), &proto.String{Value: k})
	return ret.Value
}

func getDomain(prefixKey string) string {
	trans, cli, err := service.RegistryServiceClient()
	if err != nil {
		log.Println("[ Go2o][ Fatal]: ", err.Error())
		os.Exit(0)
	}
	protocol := reFn(cli, registry.HttpProtocols)
	domain := reFn(cli, registry.Domain)
	prefix := reFn(cli, prefixKey)
	trans.Close()
	return fmt.Sprintf("%s://%s%s", protocol, prefix, domain)
}

func GetDomain() string {
	trans, cli, _ := service.RegistryServiceClient()
	domain := reFn(cli, registry.Domain)
	trans.Close()
	return domain
}

// GetImageURL 获取资源URI
func GetImageURL() string {
	return getDomain(registry.DomainPrefixImage)
}

// GetMemberURL 获取会员URL
func GetMemberURL(mobile bool) string {
	key := types.StringCond(mobile,
		registry.DomainPrefixMobileMember,
		registry.DomainPrefixMember)
	return getDomain(key)
}

// GetPassportURL 获取通行证URL
func GetPassportURL(mobile bool) string {
	key := types.StringCond(mobile,
		registry.DomainPrefixMobilePassport,
		registry.DomainPrefixPassport)
	return getDomain(key)
}

// GetPortalURL 获取门户URL
func GetPortalURL(mobile bool) string {
	key := types.StringCond(mobile,
		registry.DomainPrefixMobilePortal,
		registry.DomainPrefixPortal)
	return getDomain(key)
}

func GetHApiURL() string {
	return getDomain(registry.DomainPrefixHApi)
}

func GetWholesaleURL() string {
	return getDomain(registry.DomainPrefixWholesalePortal)
}
