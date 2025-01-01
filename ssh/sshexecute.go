package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

// 每个执行机ip调用一次
func Sshonnect_execute(ip, command string) (string, error) {
	// 配置SSH客户端的认证信息
	config := &ssh.ClientConfig{
		User: "root", // 用户名
		Auth: []ssh.AuthMethod{
			ssh.Password("HWbesano@1"), // 密码
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 忽略主机密钥检查
	}

	// 连接到远程主机
	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		return "", fmt.Errorf("Failed to connect: %v", err)
	}
	defer client.Close()

	// 在远程主机上执行命令
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create session: %v", err)
	}
	defer session.Close()

	// 执行命令
	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("Failed to run command: %v", err)
	}

	return string(output), nil
}
