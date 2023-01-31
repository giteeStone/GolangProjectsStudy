//手动抛出异常
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获：", err)
		}
	}()
	panic("出现异常！") // 手动抛出异常
	// 捕获： 出现异常！
}

//系统异常
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获：", err)
		}
	}()

	nums := []int{1, 2, 3}
	fmt.Println(nums[4]) // 系统抛出异常
	// 捕获： runtime error: index out of range [4] with length 3
}

//返回异常
func getCircleArea(radius float32) (area float32, err error) {
	if radius < 0 {
		// 构建个异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 * radius * radius
	return
}

func main() {
	area, err := getCircleArea(-5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(area)
	}
}

//自定义异常
type PathError struct {
	path       string
	op         string
	createTime string
	message    string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s",
		p.path, p.op, p.createTime, p.message)
}

func Open(filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return &PathError{
			path:       filename,
			op:         "read",
			message:    err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}

	defer file.Close()
	return nil
}

func main() {
	err := Open("test.txt")
	switch v := err.(type) {
	case *PathError:
		fmt.Println("get path error,", v)
	default:
	}
}
