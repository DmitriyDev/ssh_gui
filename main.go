package main

import (
	"bufio"
	"os"
	ssh "github.com/DmitriyDev/ssh_client"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
)

var application App

type Client struct {
	reader     *bufio.Reader
	connection *ssh.Connection
}

func (client *Client) Execute(cmdString string) string {
	err, out := client.connection.RunCmd(cmdString)
	if err != nil {
		return err.Error()
	}

	return out
}

func Connect(ip string, user string, password string, port int) Client {
	client := &ssh.SSHClient{
		ip, user, password, port,
	}

	reader := bufio.NewReader(os.Stdin)
	connection := client.Connect(ssh.CERT_PASSWORD)

	return Client{
		reader:     reader,
		connection: &connection,
	}
}
func (client *Client) Close() {
	client.connection.Close()
}

func appMain(driver gxui.Driver) {

	theme := AppTheme{}
	theme.init(driver)

	application = App{}
	application.Theme = theme
	application.createWindow(1024, 768, "SSH GUI client")
	application.buildWindow()

}

func main() {
	gl.StartDriver(appMain)
}
