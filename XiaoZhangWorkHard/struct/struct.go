//type后面跟的是类型
package m
type myInt int //自定义类型,%T输出为main.myInt
type yourInt=int//类型别名，只是在代码编写过程中起作用
// var n myInt    
// n=100

//结构体的定义
//type 类型名 struct{
//字段名 字段类型
//字段名 字段类型
//……
//}
//结构体占用一块连续的内存空间
type person struct{
name string
age int
hobby []string
gender string
}
var p person //{var p=new(person)
p.name="周林"
p.age=9000
p.gender="男"
p.hobby= []string{"篮球","足球","双色球"},


p2:=person{
"小王子",
"男",
}//var p2=perseon{}
//匿名结构体(多用于一些临时场景)
var s struct{
name string
age int
}
s.x="嘿嘿嘿"
s.y=100

//构造函数初始化结构体
type person struct{
name string,
age int
}


//如果结构体占内存比较大，可以让函数返回结构体地址来节省内存
//约定俗成构造函数以new开头命名
func newPerson(name string,age int)person{
return person{
name:name,
age:age,
}
}

func main(){
p1:=newPerson("元帅",18)
p2:=newPerson("周林",9000)
fmt.Println(p1,p2)
}

package dog
import "fmt"
//标识符：变量名 函数名 类型名 方法名
//Go语言中如果标识符首字母是大写的，就表示对外部可见（暴露的，公有的）


//方法和接收者
//func (接收者变量 接收者类型)方法名（参数列表）（返回参数）
//只能给自己定义的类型添加方法
//不能给别的包里面的类型添加方法，只能给自己包里的类型添加方法
type dog struct{
 name string
}
func newDog(name string)dog{
  return dog{
  name:name,
}
}
func (d dog)wang(){
  fmt.Printf("%s汪汪汪\n",d.name)
}
//方法是作用于特定类型的函数，只能被dog调用
//接受者表示的是调用方法的具体类型变量，多用类型名首字母小写表示
func main(){
d1:=newDog("zhoulin")
d1.wang()
}

type myInt int
func(m myInt)hello(){
 fmt.Println("我是一个int")
}
func main(){
m:=myInt(100) //var m myInt=100
m.hello()
}

//匿名字段
type person struct{
  string
  int
}

func main(){
  var a = struct{
  	x int
  	y int
  }(10,20)
  p1:=person{
  "周林",
  90000,
  }
  fmt.Println(p1.string) //打印周林
}

//嵌套结构体
type address struct{
	province string
	city string
}
type person struct{
	name string
	age int
	addr address
}
type company struct{
	name string
	addr address
}
func main(){
	p1:=person{
    	name:"周林",
    	age:9000,
    	addr:address{
		province:"山东",
		city:"威海",
 		},
 }
 fmt.Println(p1.addr.city)//周林居住的城市
}


//匿名嵌套结构体
type person struct{
	name string
	age int
	address
}
func main(){
	p1:=person{
    name:"周林",
    age:9000,
    address:address{
		province:"山东",
		city:"威海"，
	}，
	}
 fmt.Println(p1)
 fmt.Println(p1.name,p1.address.city)
 fmt.Println(p1,city)
 //先在自己结构体找到这个字段，找不到再去匿名结构体中查找
 //此外，嵌套结构体外层可以使用内层的方法
}

//结构体与JSON
//1.序列化：把Go语言中的结构体变量-->json格式的字符串
//2.反序列化：json格式的字符串-->Go语言中能够识别的结构体变量
//json文件是前端能够识别和使用的数据
import "encoding/json"
type person struct {
	Name string `json:"name"` //json模式下用name解析
	Age  int    `json:"age"`
} //变量首字母必须大写，不然json拿不到
func main() {
	p1 := person{
		Name: "周林",
		Age:  9000,
	}
	//序列化
	b, err := json.Marshal(p1) //将结构体转换为json
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Println(string(b))
	//反序列化
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能在函数内部修改p2的值
	fmt.Printf("%v\n", p2)
}

