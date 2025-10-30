package main

import (
    "flag"
    "log"
    "strconv"
    "fmt"

    service "github.com/OmineDev/flowers-for-machines/std_server/service/src"
    "github.com/pterm/pterm"
)

var (
	rentalServerCode     *string
	rentalServerPasscode *string
	authServerAddress    *string
	authServerToken      *string
	standardServerPort   *int
	consoleDimensionID   *int
	consoleCenterX       *int
	consoleCenterY       *int
	consoleCenterZ       *int
)

func init() {
	rentalServerCode = flag.String("rsn", "", "The rental server number.")
	rentalServerPasscode = flag.String("rsp", "", "The pass code of the rental server.")
	authServerAddress = flag.String("asa", "", "The auth server address.")
	authServerToken = flag.String("ast", "", "The auth server token.")
	standardServerPort = flag.Int("ssp", 0, "The server port to running.")
	consoleDimensionID = flag.Int("cdi", 0, "The dimension ID of the console. (e.g. overworld = 0, nether = 1, end = 2, dmT = T, etc.)")
	consoleCenterX = flag.Int("ccx", 0, "The X position of the center of the console.")
	consoleCenterY = flag.Int("ccy", 0, "The Y position of the center of the console.")
	consoleCenterZ = flag.Int("ccz", 0, "The Z position of the center of the console.")

	flag.Parse()
    if len(*rentalServerCode) == 0 || len(*authServerAddress) == 0 || *standardServerPort == 0 {
        pterm.DefaultSection.Println("交互式启动向导（缺少参数）")

        // 向导：必填参数
        if len(*rentalServerCode) == 0 {
            code, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入租赁服务器号 (rsn)").Show()
            *rentalServerCode = code
        }
        if len(*authServerAddress) == 0 {
            addr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入认证服务器地址 (asa), 例如 https://nv1.nethard.pro").Show()
            *authServerAddress = addr
        }
        if *standardServerPort == 0 {
            portStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("请输入标准服务器端口 (ssp), 例如 8080").Show()
            if v, err := strconv.Atoi(portStr); err == nil {
                *standardServerPort = v
            } else {
                log.Fatalln("Invalid port provided in wizard")
            }
        }

        // 向导：可选参数
        if len(*rentalServerPasscode) == 0 {
            passcode, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）请输入租赁服务器密码 (rsp)，留空则不设置").Show()
            *rentalServerPasscode = passcode
        }
        if len(*authServerToken) == 0 {
            token, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）请输入认证令牌 (ast)，留空则不设置").Show()
            *authServerToken = token
        }

        // 控制台坐标与维度
        if *consoleDimensionID == 0 {
            dimStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）控制台维度ID (cdi)，默认0").Show()
            if dimStr != "" {
                if v, err := strconv.Atoi(dimStr); err == nil { *consoleDimensionID = v }
            }
        }
        if *consoleCenterX == 0 {
            xStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）控制台中心X (ccx)，默认0").Show()
            if xStr != "" { if v, err := strconv.Atoi(xStr); err == nil { *consoleCenterX = v } }
        }
        if *consoleCenterY == 0 {
            yStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）控制台中心Y (ccy)，默认0").Show()
            if yStr != "" { if v, err := strconv.Atoi(yStr); err == nil { *consoleCenterY = v } }
        }
        if *consoleCenterZ == 0 {
            zStr, _ := pterm.DefaultInteractiveTextInput.WithDefaultText("（可选）控制台中心Z (ccz)，默认0").Show()
            if zStr != "" { if v, err := strconv.Atoi(zStr); err == nil { *consoleCenterZ = v } }
        }

        // 最终校验
        if len(*rentalServerCode) == 0 {
            log.Fatalln("Please provide your rental server number.\n\te.g. -rsn=\"123456\"")
        }
        if len(*authServerAddress) == 0 {
            log.Fatalln("Please provide your auth server address.\n\te.g. -asa=\"http://127.0.0.1\"")
        }
        if *standardServerPort == 0 {
            log.Fatalln("Please provide the server port to running.\n\te.g. -ssp=0")
        }
    }
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("后端运行异常（标准服务器启动失败）！")
            fmt.Printf("错误信息：%v\n", r)
            fmt.Println("请检查启动参数、token和网络连接，并阅读 README 帮助排查。如仍有疑问，请将上述错误信息截图后咨询技术支持。\n")
        }
    }()
    service.RunServer(
        *rentalServerCode,
        *rentalServerPasscode,
        *authServerAddress,
        *authServerToken,
        *standardServerPort,
        *consoleDimensionID,
        *consoleCenterX,
        *consoleCenterY,
        *consoleCenterZ,
    )
}
