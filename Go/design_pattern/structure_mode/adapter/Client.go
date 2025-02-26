package main

type Client struct{}

func (c *Client) InsertUSBtoPC(pc PC) {
	pc.InsertUSB()
}
