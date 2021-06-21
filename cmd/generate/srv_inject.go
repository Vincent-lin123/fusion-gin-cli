package generate

import (
	"context"
	"fmt"
	"strings"
)

func getSrvInjectFileName(dir string) string {
	fullname := fmt.Sprintf("%s/service/service.go", dir)
	return fullname
}

// 插入srv注入文件
func insertSrvInject(ctx context.Context, dir, name string) error {
	fullname := getSrvInjectFileName(dir)

	injectContent := fmt.Sprintf("%sSet,", name)
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		if injectStart == 0 && strings.Contains(line, "var ServiceSet = wire.NewSet(") {
			injectStart = 1
			return
		}

		if injectStart == 1 && strings.Contains(line, ")") {
			injectStart = -1
			data = injectContent
			flag = -1
			ok = true
			return
		}

		return "", 0, false
	}

	err := insertContent(fullname, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}
