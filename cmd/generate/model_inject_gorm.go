package generate

import (
	"context"
	"fmt"
	"strings"
)

func getModelInjectGormFileName(dir string) string {
	fullname := fmt.Sprintf("%s/model/gormx/model/model.go", dir)
	return fullname
}

func insertModelInjectGorm(ctx context.Context, dir, name string) error {
	fullName := getModelInjectGormFileName(dir)

	injectContent := fmt.Sprintf("%sSet,", name)
	injectStart := 0
	insertFn := func(line string) (data string, flag int, ok bool) {
		if injectStart == 0 && strings.Contains(line, "var ModelSet = wire.NewSet(") {
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

	err := insertContent(fullName, insertFn)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullName)

	return execGoFmt(fullName)
}
