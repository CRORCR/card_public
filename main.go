package main

import(
	"fmt"
	"./server/db"
	"./server/mainserver"
)


func main(){
	var rpcserver mainserver.RPCServer
	if nil != db.InitDB() {
		panic("数据库打开失败...")
	}
	fmt.Println("数据库打开成功...")
	rpcserver.Start()
}

/*

ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCzfWVDgTbe+JTiyYwkfDRlWnqS7zty+GBG+nDIlF70nBOcrvQ/AS2pOcd8gT+pnV/8zyKz0D/ePGnleNKpMDdnc9VkRUZ8EXDnSFFaF8s/DiYGLh66W4XOqzU0alsMCg9Ctaqpb7In4LLn7f/mNCNvlzwQp3A4f4ToV1vg4265psUwa4JSxu72fZ0Cx6bN/iZLstfhM4OU+wxTT2QB+UqDqfdwnbGkdh0dhzp/xqCylgmzov+enwh299iAaMbmYpbFwP4aSKlPyiJEEQvG+6uTUHvWTCVNXOq99ZbEjStO98hBiNZ4HKVPugM9PEDtyKsuO+w87QVkcigAGlozW6s/rJz3/XW1ADNjtExJfKi+ePyxdB5riVVUYglpRKZZzG2lLXqOVx4oXCNpNWXBty1YGj8nxLSwUKk86E8/S9QkJYubTVTtkCdbL7atMep6QVVcN91ko81jVVSib6Q29gXUdOZWsijQ4ICRc/T3QhDbhZytnH1IQSCQoU4Zv+aYsih2Ii3R78SIr88X0RIR3Fz1Yyze3FgIpTx3vop2URNCwlTftafZWhqqaTvTJ6adiiTjiEGBuPveoqGrKAl8s8Ad8HrWWwW24j3+xBtNeAxcf6MTzed7fWcpQdSeAThf1YG79mlxO0fWl2keCno1TESMiNyLR1eDi0PI/OceWyvc4w== usershane


*/
