package main

//使用类型分支判断接口类型
import "fmt"


type CantainCanUseFaceID interface {
	CanUseFaceID()
}

type CantinStolen interface {
	Stolen()
}

type Alipay struct {

}

type Cash struct {

}

func (a *Alipay) CanUseFaceID()  {
	fmt.Println("====================")
}

func (c *Cash) Stolen ()  {

}


func Print(paymethod interface{})  {
	switch paymethod.(type) {
	case CantainCanUseFaceID:
		fmt.Printf("%T can use faceID \n",paymethod )
	case CantinStolen:
		fmt.Printf("%T may be stolen",paymethod)
	}
}

func main()  {
	var a CantainCanUseFaceID
	a = new(Alipay)
	a.CanUseFaceID()

	Print(new(Alipay))
	Print(new(Cash))
}