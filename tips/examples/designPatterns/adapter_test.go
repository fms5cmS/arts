package designPatterns

import (
	"fmt"
	"testing"
)

type Computer interface {
	InsertIntoLightningPort()
}

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine")
}

type Windows struct{}

func (w *Windows) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into windows machine")
}

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("client inserts lightning connector into computer.")
	com.InsertIntoLightningPort()
}

// 适配器
type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB")
	w.windowsMachine.InsertIntoLightningPort()
}

func TestAdapter(t *testing.T) {
	client := &Client{}

	mac := &Mac{}
	client.InsertLightningConnectorIntoComputer(mac)

	win := &Windows{}
	winAdapter := &WindowsAdapter{windowsMachine: win}
	client.InsertLightningConnectorIntoComputer(winAdapter)
}
