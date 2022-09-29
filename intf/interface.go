// Package shared contains shared data between the host and plugins.
package intf

type SM interface{
	Send(phone string,text string) Res
}

type Res struct{
	Reslut bool
	Msg string
}