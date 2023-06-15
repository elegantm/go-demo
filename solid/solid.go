/*
@Time : 2022/7/21 10:57
@Author : elegantm
@File : solid
//All rights reserved.
功能介绍：
修订历史:
*/
package solid
import "fmt"


type ICar interface {
	carRun()
}

type IDriver interface{
	Drive(car ICar)
}

type Driver struct {

}

func (Driver)Drive(car ICar)  {
	car.carRun()
}

type Benz struct {}

func (B Benz) carRun(){
	fmt.Println("奔驰运行")
}

type BMW struct {}

func (B BMW) carRun() {
	fmt.Println("宝马运行")
}



