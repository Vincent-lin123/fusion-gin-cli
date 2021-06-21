package generate

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v2"
	"path/filepath"
	"strings"
)

type Config struct {
	Dir     string
	PkgName string
	Name    string
	Comment string
	File    string
	Storage string
	Modules string
}

type Command struct {
	cfg *Config
}

func Exec(cfg Config) error {
	cmd := &Command{cfg: &cfg}
	return cmd.Exec()
}

func (a *Command) hasModule(m string) bool {
	if v := a.cfg.Modules; v == "" || v == "all" {
		return true
	}

	for _, s := range strings.Split(a.cfg.Modules, ",") {
		if s == m {
			return true
		}
	}

	return false
}

func (a *Command) handleError(err error, desc string) {
	if err != nil {
		fmt.Printf("%s:%s", desc, err.Error())
	}
}

func (a *Command) Exec() error {
	var item TplItem

	if a.cfg.File != "" {
		b, err := readFile(a.cfg.File)
		if err != nil {
			return err
		}

		err = yaml.Unmarshal(b.Bytes(), &item)
		if err != nil {
			return err
		}
	} else {
		item.StructName = a.cfg.Name
		item.Comment = a.cfg.Comment
	}

	dir, err := filepath.Abs(a.cfg.Dir)
	if err != nil {
		return err
	}

	pkgName := a.cfg.PkgName
	ctx := context.Background()

	if a.hasModule("schema") {
		err = genSchema(ctx, pkgName, dir, item.StructName, item.Comment, item.toSchemaFields()...)
		if err != nil {
			return err
		}
	}

	if a.hasModule("model") {
		switch a.cfg.Storage {
		case "mongo":
			err = genMongoEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityMongoFields()...)
			a.handleError(err, "Generate mongo entity")

			err = insertEntityInjectMongo(ctx, dir, item.StructName)
			a.handleError(err, "Insert mongo entity inject")

			err = genModelImplMongo(ctx, pkgName, dir, item.StructName, item.Comment)
			a.handleError(err, "Generate mongo model")

			err = insertModelInjectMongo(ctx, dir, item.StructName)
			a.handleError(err, "Insert mongo model inject")
		default:
			err = genGormEntity(ctx, pkgName, dir, item.StructName, item.Comment, item.toEntityGormFields()...)
			a.handleError(err, "Generate gorm entity")

			err = insertEntityInjectGorm(ctx, dir, item.StructName)
			a.handleError(err, "Insert gorm entity inject")

			err = genModelImplGorm(ctx, pkgName, dir, item.StructName, item.Comment)
			a.handleError(err, "Generate gorm model")

			err = insertModelInjectGorm(ctx, dir, item.StructName)
			a.handleError(err, "Insert gorm model inject")
		}
	}

	if a.hasModule("service") {
		//err = genBll(ctx, pkgName, dir, item.StructName, item.Comment)
		//a.handleError(err, "Generate bll interface")

		err = genSrvImpl(ctx, pkgName, dir, item.StructName, item.Comment)
		a.handleError(err, "Generate srv impl")

		err = insertSrvInject(ctx, dir, item.StructName)
		a.handleError(err, "Insert srv inject")
	}

	if a.hasModule("api") {
		err = genAPI(ctx, pkgName, dir, item.StructName, item.Comment)
		a.handleError(err, "Generate api")

		err = insertAPIInject(ctx, dir, item.StructName)
		a.handleError(err, "Insert api inject")
	}

	if a.hasModule("mock") {
		err = genAPIMock(ctx, pkgName, dir, item.StructName, item.Comment)
		a.handleError(err, "Generate api mock")

		err = insertAPIMockInject(ctx, dir, item.StructName)
		a.handleError(err, "Insert api mock inject")
	}

	if a.hasModule("router") {
		err = insertRouterAPI(ctx, dir, item.StructName)
		a.handleError(err, "Insert router api")

		err = insertRouterInject(ctx, dir, item.StructName)
		a.handleError(err, "Insert router inject")
	}

	return nil
}
