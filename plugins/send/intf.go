package send

// intf is interface


type SM interface{
	Send(phone string,text string) Res
}

type Res struct{
	Reslut bool
	Msg string
}