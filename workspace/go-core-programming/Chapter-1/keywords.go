// 引导程序整体结构的8 个关键字
package        //定义包名的关键字
import         //导入包名关键字
const          //常量声明关键字
var            //变量声明关键字
func           //函数定义关键字
defer          //延迟执行关键字
go             //并发语法糖关键字
return         //函数返回关键字

// 声明复合数据结构的4 个关键字
struct         //定义结构类型关键字
interface      //定义接口类型关键字
map            //声明或创建map 类型关键字
chan           //声明或创建遥远类型关键字

// 控制程序结构的13 个关键宇
if else                                        //if else 语句关键字
for range break continue                       //for 循环使用的关键字
switch select type case default fallthrough    //switch 和select 语句使用的关键字
goto                                           //goto 跳转语句关键字


// 内置数据类型标识符（ 20 个）
"""
数值（ 16 个）
	整型（ 12 个）
		byte int int8 int16 int32 int64
		uint unint8 uint16 uint32 uint64 uintprt
	浮点型（ 2 个）
		float32 float64
	复数型（ 2 个）
		complex64 complex128
字符和字符串型（ 2 个）
	str 工ng rune
接口型（ 1 个）
	error
布尔型（ 1 个）
	bool
"""

//内置函数（ 15 个）
make new len cap append copy delete panic recover close complex real image Print Println

// 常量值标识符（ 4 个）
true false // true 和false 表示bool 类型的两常量值：具和假
iota       // 用在连续的枚举类型的声明中
nil        // 指针／引用型的变量的默认值就是nil


//空白标识符（ 1 个）
_