package main

func main() {
	// 初始化
	context := NewContext()
	for i := 0; i < 5; i++ {
		temp := i
		context.add(func(ctx *Context) {
			// 注释ctx.Next()后  0 1 2 3 4
			// 不注释为链式(深度优先)调用    4 3 2 1 0
			// ctx.Next()
			print(temp)
		})
	}
	// 入口
	context.Next()
}

func NewContext() *Context {
	return &Context{
		funcIndex:   -1,
		FuncHandles: make([]FuncHandle, 0),
	}
}

type Context struct {
	funcIndex   int
	FuncHandles []FuncHandle
}

type FuncHandle func(ctx *Context)

// 调用
func (this *Context) Next() {
	this.funcIndex++
	for ; this.funcIndex < len(this.FuncHandles); this.funcIndex++ {
		this.FuncHandles[this.funcIndex](this)
	}
}

// 添加Handle
func (this *Context) add(handle FuncHandle) {
	this.FuncHandles = append(this.FuncHandles, handle)
}
