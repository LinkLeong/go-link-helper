package a_consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/hashicorp/go-uuid"
)


type ConsulOpt struct {
	//consul agent的地址 例如：192.168.200.147:8500
	ConsulAddress  string
	//当前客户端ip，注册到consul的IP地址
	LocalIp string
	//当前程序启用的端口
	LocalPort int
	//服务名称
	Name string
	//标签
	Tags[] string
}

//服务注册与发现
func Setup( op ConsulOpt) error {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = op.ConsulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		return err
	}

	id,_:=uuid.GenerateUUID()
	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = id
	registration.Name = op.Name
	registration.Port = op.LocalPort
	registration.Tags = op.Tags
	registration.Address = op.LocalIp

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d/get/check", registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
	return err
}


