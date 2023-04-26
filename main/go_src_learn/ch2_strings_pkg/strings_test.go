package ch2_strings_pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestStringsCompare(t *testing.T) {
	compare1 := strings.Compare("23423", "34234")
	compare2 := strings.Compare("1233", "1233")
	compare3 := strings.Compare("yuxiangjin", "1232")
	t.Log(compare1)
	t.Log(compare2)
	t.Log(compare3)
}

func TestStringsJoin(t *testing.T) {
	joinString1 := strings.Join([]string{"12321", "jklgja", "wang"}, ",")
	joinString2 := strings.Join([]string{"12321", "jklgja", "wang"}, "")
	joinString3 := strings.Join([]string{"12321", "jklgja", "wang"}, "=")
	joinString4 := strings.Join([]string{"12321", "jklgja", "wang"}, "|")
	t.Log(joinString1)
	t.Log(joinString2)
	t.Log(joinString3)
	t.Log(joinString4)
}

func TestStringsSplit(t *testing.T) {
	joinString1 := strings.Split("jkljkljkljkl,gewgew", ",")
	t.Log(joinString1)
}

func TestStringsReader(t *testing.T) {
	reader := strings.NewReader("123213213gejwkrfljaew中文试试")
	bytes := make([]byte, reader.Len())
	n, err := reader.Read(bytes)
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(string(bytes))
	bytes = append(bytes[:1], bytes[2:]...)
	copy(bytes, "重新赋值后")
	fmt.Println(string(bytes))
}

func TestStringsReplacer(t *testing.T) {
	replacer := strings.NewReplacer("old string", "jklfew", "new", "lsjkljkl")
	replace := replacer.Replace("new string")
	fmt.Println(replace)
	buffer := &bytes.Buffer{}
	n, err := replacer.WriteString(buffer, "greagrag")
	fmt.Println(n)
	fmt.Println(err)
	fmt.Println(buffer.String())
}

func TestBytesBuffer(t *testing.T) {
	buffer := &bytes.Buffer{}
	fmt.Println(buffer.Len())
	fmt.Println(buffer.Cap())
	buffer.Grow(100)
	buffer.WriteString("fewfewwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwwww")
	all, err := ioutil.ReadAll(buffer)
	fmt.Println(string(all))
	fmt.Println(err)
	fmt.Println(buffer.Len())
	fmt.Println(buffer.Cap())
	i2 := make([]byte, 0, 20)
	fmt.Println(len(i2))
	i2 = i2[:10]
	fmt.Println(len(i2))
}

func TestSliceGrow(t *testing.T) {
	intList := make([]int32, 0, 16)
	//intList[1] = 12
	//fmt.Println(intList)
	int32s := intList[:4]
	int32s[1] = 123
	fmt.Println(int32s)
	fmt.Println(len(int32s))
	fmt.Println(cap(int32s))
	subList1 := intList[:0]
	fmt.Println(len(subList1))
	fmt.Println(cap(subList1))
	fmt.Println(subList1)

	subList2 := intList[1:4]
	fmt.Println(len(subList2))
	fmt.Println(cap(subList2))
	fmt.Println(subList2)

	subList3 := intList[1:15]
	subList3[13] = 342
	fmt.Println(len(subList3))
	fmt.Println(cap(subList3))
	fmt.Println(subList3)
}

func TestEmptySlice(t *testing.T) {
	i := make([]string, 0)
	fmt.Println(i)

}
