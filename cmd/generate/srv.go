package generate

import (
	"context"
	"fmt"
	"fusion-gin-cli/util"
)

const srvTpl = `
package service

import (
	"context"

	"{{.PkgName}}/schema"
)

// I{{.Name}} {{.Comment}}业务逻辑接口
type I{{.Name}} interface {
	// 查询数据
	Query(ctx context.Context, params schema.{{.Name}}QueryParam, opts ...schema.{{.Name}}QueryOptions) (*schema.{{.Name}}QueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, id string, opts ...schema.{{.Name}}GetOptions) (*schema.{{.Name}}, error)
	// 创建数据
	Create(ctx context.Context, item schema.{{.Name}}) (*schema.IDResult, error)
	// 更新数据
	Update(ctx context.Context, id string, item schema.{{.Name}}) error
	// 删除数据
	Delete(ctx context.Context, id string) error
}

`

func getSrvFileName(dir, name string) string {
	fullname := fmt.Sprintf("%s/service/%s.go", dir, util.ToLowerUnderlinedNamer(name))
	return fullname
}

// 生成service文件
func genSrv(ctx context.Context, pkgName, dir, name, comment string) error {
	data := map[string]interface{}{
		"PkgName": pkgName,
		"Name":    name,
		"Comment": comment,
	}

	buf, err := execParseTpl(srvTpl, data)
	if err != nil {
		return err
	}

	fullname := getSrvFileName(dir, name)
	err = createFile(ctx, fullname, buf)
	if err != nil {
		return err
	}

	fmt.Printf("文件[%s]写入成功\n", fullname)

	return execGoFmt(fullname)
}
