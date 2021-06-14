package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// 调用远程函数就像本地调用函数一样
// 给本地方法加上http请求

func main() {
	h := &hello{
		endpoint: "http://localhost:8080/",
	}
	//resq, err := h.SayHello("golang")
	//if err != nil {
	//	fmt.Printf("%s", err)
	//}

	//fmt.Printf("%s\n", resq)

	SetFunField(h)
	field, _ := h.FuncField("golang")
	fmt.Printf(field)

}


type hello struct {
	endpoint  string
	FuncField func(name string) (string, error)

	// 对结构体里面的方法进行篡改
}

func SetFunField(val interface{}) {

	v := reflect.ValueOf(val) // 反射出来 指向该对象的指针
	// dynamic type 就是实现了该接口的对象
	elem := v.Elem()  // 拿到指针指向的结构体
	t1 := elem.Type() // 拿到结构体的类型信息

	num := t1.NumField() // 拿到结构体里面的属性
	for i := 0; i < num; i++ {
		f := elem.Field(i)
		//fmt.Println(elem.Field(i))
		if f.CanSet() {
			// 我们对函数进行修改，就是将整个函数传入要篡改的函数
			// 先找到函数的入参 出参 不管内部实现

			fn := func(args []reflect.Value) (results []reflect.Value) {
				name := args[0].Interface().(string)
				// setp 1 ： 客户端发送请求

				serviceName := val.(Service).ServiceName()
				endpoint:=CfgMap[serviceName].Endpoint

				client := http.Client{}
				resp, err := client.Get(endpoint + name)
				if err != nil {
					fmt.Println(err.Error())
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				data, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("%+v", err)
					return []reflect.Value{reflect.ValueOf(""), reflect.ValueOf(err)}
				}
				return []reflect.Value{reflect.ValueOf(string(data)), reflect.Zero(reflect.TypeOf(new(error)).Elem())}
			}
			f.Set(reflect.MakeFunc(f.Type(), fn))
		}
	}
}



// 反射 ： 获取程序本身的信息

// 原始情况：我们在我们的本地方法上面加上了http通信
// 这样 需要给每一个方法里面都加上http通信
// 改造后：利用反射 为每个方法设置http通信
// go 是静态语言 在运行时就确定数据的类型 所以我们就可以获得数据的类型进行反射，
// 获取到数据本身的属性 元编程 编程语言对自身的操作

// 方法是通过反射去篡改掉原方法 希望反射去为每个方法加上http通信
// 如果不通过这样，我们就需要修改每个已经写好的方法，给他们加上http通信

// go里面的机制 可以去看看java中的机制
// rpc 本来是调用本地方法，但是方法中去请求了服务器，服务器拿到了数据，再返回数据。

type Service interface {
	ServiceName()string
}

func (h *hello)ServiceName() string {
	return "hello"
}