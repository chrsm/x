package main

import "github.com/chrsm/winapi/user"

func main() {
	user.SendMessage(user.HwndBroadcast, user.WmSysCommand, user.ScMonitorPower, 0x0002)
}
