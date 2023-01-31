package split_string

import (
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		//测试用例失败了
		t.Errorf("want:%v but got:%v\n", want, ret)
	}
} //测试用例一

func Test2Split(t *testing.T) {
	ret := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want:%v but get:%v\n", want, ret)
	}
} //测试用例二

// 一次测试多个
func Test3Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := []testCase{
		testCase{"babcbef", "b", []string{"", "a", "c", "efd"}}, //~~~~~~~~~~~~~人工设置错误
		testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		testCase{"abcdef", "bc", []string{"a", "def"}},
		testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}
	for _, test := range testGroup {
		got := Split(test.str, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want:%#v got:%#v\n", test.want, got)
		}
	}
}

// 子测试
func Test4Split(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := map[string]testCase{
		"case 1": testCase{"babcbef", "b", []string{"", "a", "c", "ef"}},
		"case 2": testCase{"a:b:c", ":", []string{"a", "b", "c"}},
		"case 3": testCase{"abcdef", "bc", []string{"a", "def"}},
		"case 4": testCase{"沙河有沙又有河", "有", []string{"沙河", "沙又", "河"}},
	}

	for name, test := range testGroup {
		t.Run(name, func(t *testing.T) {
			got := Split(test.str, test.sep)
			if !reflect.DeepEqual(got, test.want) {
				t.Fatalf("want:%#v got:%#v\n", test.want, got)
			}
		})
	}
} //会把每个map中的样例试结果都打印出来

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}

// 网络测试
func TestHelloHandler(t *testing.T) {
	ln, err := net.Listen("tcp", "localhost:0")
	handleError(t, err)
	defer ln.Close()

	http.HandleFunc("/hello", HelloHandler)
	go http.Serve(ln, nil)

	resp, err := http.Get("http://" + ln.Addr().String() + "hello")
	handleError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	if string(body) != "hello world!" {
		t.Fatal("expected hello world,but got", string(body))
	}
}
