package tests

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/DeleteElf/zero-core/utils/iohelper"
)

func TestIOPath(t *testing.T) {
	paths := []string{"test", "http://test2", "/test3"}

	fmt.Printf("直接合并的结果：%s\r\n", filepath.Join(paths...))

	fmt.Printf("CombinePath合并的结果：%s\r\n", iohelper.CombinePath(paths...))

	fmt.Printf("CombinePath合并测试http地址的结果：%s\r\n", iohelper.CombinePath(paths[0], paths[1]))

	paths = []string{"test", "http://test2", "test3/", "test4", "test5"}
	fmt.Printf("CombinePath合并测试http地址后端合并结果：%s\r\n", iohelper.CombinePath(paths...))

	paths = []string{"test", "https://test2", "test3/", "test4", "test5", "?q=1"}
	fmt.Printf("CombinePath合并测试https地址后端合并结果：%s\r\n", iohelper.CombinePath(paths...))

	paths = []string{"test", "wss://test2", "test3/", "test4", "test5"}
	fmt.Printf("CombinePath合并测试wss地址后端合并结果：%s\r\n", iohelper.CombinePath(paths...))
	paths = []string{"test", "ftp://test2", "test3/", "test4", "test5"}
	fmt.Printf("CombinePath合并测试ftp地址后端合并结果：%s\r\n", iohelper.CombinePath(paths...))
}

func TestUrlQueryString(t *testing.T) {
	paths := []string{"test=1", "test2=2&", "&test3=3"}

	fmt.Printf("CombineQueryString合并的结果：%s\r\n", iohelper.CombineQueryString(paths...))

	paths = []string{"test=1", "test2=2&", "&test3=3", "?id=1"}
	fmt.Printf("CombineQueryString合并的结果：%s\r\n", iohelper.CombineQueryString(paths...))

	paths = []string{"test=1", "test2=2&", "&test3=3", "?id=1", "&test=test&"}
	fmt.Printf("CombineQueryString合并的结果：%s\r\n", iohelper.CombineQueryString(paths...))

	paths = []string{"?test=1", "test2=2&", "&test3=3", "?id=1", "&test=test&", "#saaa"}
	fmt.Printf("CombineQueryString合并的结果：%s\r\n", iohelper.CombineQueryString(paths...))

	paths = []string{"#test", "test2=2&", "&test3=3", "?id=1", "&test=test&", "#saaa"}
	fmt.Printf("CombineQueryString合并的结果：%s\r\n", iohelper.CombineQueryString(paths...))
}
