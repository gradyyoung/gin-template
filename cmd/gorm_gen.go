package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DBDSN = "root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	cfg := gen.Config{
		ModelPkgPath: "./internal/model", // 结构体生成目录
		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true,

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false,

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false,
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: false,
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	}
	// 处理表名
	cfg.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		// 需要忽略的表
		//if strings.EqualFold(tableName, "pay_credit") ||
		//	strings.EqualFold(tableName, "pay_log") ||
		//	strings.EqualFold(tableName, "space") ||
		//	strings.EqualFold(tableName, "space_account") {
		//	return ""
		//}
		return tableName
	})
	// 处理 model名
	cfg.WithModelNameStrategy(func(tableName string) (targetTableName string) {
		s := tableName
		//if strings.HasPrefix(tableName, "conf_") {
		//	s = strings.TrimPrefix(tableName, "conf_")
		//}
		ns := schema.NamingStrategy{SingularTable: true}
		return ns.SchemaName(s)
		return tableName
	})
	// 处理文件名
	cfg.WithFileNameStrategy(func(tableName string) (targetTableName string) {
		//if strings.HasPrefix(tableName, "conf_") {
		//	return strings.TrimPrefix(tableName, "conf_")
		//}
		return tableName
	})
	g := gen.NewGenerator(cfg)

	gormdb, _ := gorm.Open(mysql.Open(DBDSN))

	g.UseDB(gormdb) // reuse your gorm db

	// 自定义字段的数据类型
	dataMap := map[string]func(detailType gorm.ColumnType) (dataType string){
		"tinyint":   func(detailType gorm.ColumnType) (dataType string) { return "bool" },
		"smallint":  func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"mediumint": func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"bigint":    func(detailType gorm.ColumnType) (dataType string) { return "int" },
		"int":       func(detailType gorm.ColumnType) (dataType string) { return "int" },
		// 这里将数据库中的`datetime`映射为`config.Time`，方便我们序列化和反序列化
		"datetime": func(detailType gorm.ColumnType) (dataType string) { return "config.Time" },
	}
	// 要先于`ApplyBasic`执行
	g.WithDataTypeMap(dataMap)
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		return tag.Set("autoUpdateTime")
	})
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		return tag.Set("autoCreateTime")
	})

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{autoUpdateTimeField, autoCreateTimeField}

	// 创建 全部模型文件 , 并覆盖前面创建的同名模型
	g.GenerateAllTable(fieldOpts...)

	g.Execute()
}
