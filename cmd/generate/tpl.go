package generate

type TplItem struct {
	StructName string         `yaml:"name"`    // 结构体名称
	Comment    string         `yaml:"comment"` // 注释
	Fields     []TplFieldItem `yaml:"fields"`  // 字段项
}

func (t TplItem) toSchemaFields() []schemaField {
	var items []schemaField
	for _, f := range t.Fields {
		items = append(items, schemaField{
			Name:           f.StructFieldName,
			Comment:        f.Comment,
			Type:           f.StructFieldType,
			IsRequired:     f.StructFieldRequired,
			BindingOptions: f.BindingOptions,
		})
	}
	return items
}

func (t TplItem) toEntityGormFields() []entityGormField {
	var items []entityGormField
	for _, f := range t.Fields {
		items = append(items, entityGormField{
			Name:        f.StructFieldName,
			Comment:     f.Comment,
			Type:        f.StructFieldType,
			GormOptions: f.GormOptions,
		})
	}
	return items
}

func (t TplItem) toEntityMongoFields() []entityMongoField {
	var items []entityMongoField
	for _, f := range t.Fields {
		items = append(items, entityMongoField{
			Name:    f.StructFieldName,
			Comment: f.Comment,
			Type:    f.StructFieldType,
		})
	}
	return items
}

type TplFieldItem struct {
	StructFieldName     string `yaml:"name"`
	StructFieldRequired bool   `yaml:"required"`
	Comment             string `yaml:"comment"`         // 注释
	StructFieldType     string `yaml:"type"`            // 结构体字段类型
	GormOptions         string `yaml:"gorm_options"`    // gorm配置项
	BindingOptions      string `yaml:"binding_options"` // binding配置项
}
