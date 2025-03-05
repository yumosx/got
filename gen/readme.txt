# 必选参数
--table    指定数据库表名（默认用模型名小写复数）
--fields   字段定义（格式：name:type,age:int）
--driver   数据库驱动类型（mysql/postgresql/sqlite）

# 可选参数
--output   自定义输出目录（默认./repositories）
--dry-run  预览生成效果不写入文件
--force    强制覆盖已存在文件