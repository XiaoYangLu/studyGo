package mysql

/*
 * mysql语句
 */
func StudyMysqlStatement() {
	//Some words you need to know before learning
	//varchar：可变长字符串		create：创建
	//database：数据库			show：查看
	//use：使用					drop：减少
	//table：表					describe：描述
	//alter：改变				change：改变
	//modify：修改				add：添加
	//insert：插入				info：消息
	//update：更新				where：...情况下
	//set：设置，放				from：从...起
	//delete：删除				select：选择
	//group：组					limit：限度
	//primary：主要的			unique：唯一的
	//index：索引				default：默认
	//foreign：外围的			references：涉及
	//union：联合				cross：交叉
	//join：参加					inner：内部的
	//right：右					left：左
	//outer：外面的				natural：自然的

	//创建数据库					 		创建表
	//CREATE DTATBASE 数据库名;			CREATE TABLEE 表名(
	//查看所有数据库							字段  属性,
	//SHOW DATABASE;						...
	//									);
	//查看已经创建的数据库				查看表
	//SHOW CREATE DATABASE;				SHOW CREATE TABLE 表名 [\G];  #\G格式化查看
	//切换数据库							查看表列信息
	//USE 数据库名;						DESCRIBE 表名;
	//删除数据库							修改表字段
	//DROP DATABASE 数据库名;			ALTER TABLE 旧表名 [TO] 新表名;
	//									修改表属性
	//									ALTER TABLE 表名 MODTFY 字段名 新属性;
	//									添加字段
	//									ALTER TABLE 表名 ADD 新字段 属性;
	//									指定|批量列插入数据
	//									INSERT INFO 表名[(字段...)] VALUES(值...);
	//									不指定列插入数据
	//									INSERT INFO 表名 VALUES(值...);
	//									删除字段
	//									ALTER TABLE 表名 DROP 字段;
	//									修改字段排列
	//									ALTER TABLE 表名 MODTFY 字段1 属性 [FIRST|AFTER] 字段2;
	//									删除表
	//									DROP TABLE 表名;
	//									更新所有|部分数据
	//									UPDATE 表名 SET 字段=值 [...] [WHERE 条件];
	//									删除数据
	//									DELETE FROM 表名;
	//									删除所有数据
	//									TRUNCATE TABLE 表名;
	//									查询所有|指定字段
	//									SELECT * [字段1，字段2...] FROM 表名;  #: *通配符
	//									条件查询
	//									SELECT 字段... FROM 表名 WHERE 条件1 AND|OR  条件2;
	//									AND：需满足所有条件； OR：满足条件即可
	//									判断某个字段是否在集合中
	//									SELECT 字段... FROM 表名 WHERE 字段 [NOF] IN (元素...);
	//									判断字段是否为空
	//									SELECT 字段... FROM  表名 WHERE 字段 IS [NOT] NULL;
	//									判断字段值是否在某个范围
	//									SELECT 字段... FROM 表名 WHERE 字段 [NOT] BETWHERE 值1 AND 值2;
	//									判读字段是否匹配某个字符
	//									SELECT 字段... FROM 表名 WHERE 字段 [NOT] LIKE '匹配字符串‘;
	//									去除重复查询
	//									SELECT DISTINCT 字段 FROM 表名;
	//									排序查询
	//									SELECT 字段... FROM 表名 ORDER BY 字段1[ASC|DESC],字段2[ASC|DESC]...:
	//									ASC：升序    DESC：降序
	//									分组查询【对数据进行过滤】
	//									SELECT 字段... FROM 表名 GROUP BY 字段... [HAVING 条件];
	//									查询结果的数量
	//									SELECT 字段... FROM 表名 LIMIT [m,]n  #: m开始，n为M+1后的n条
	//									主键约束
	//									CREATE TABLE 表名 (字段 属性 PRIMAPY KEY,...);
	//									ALTER TABLE 表名 ADD PRIMAPY KEY (列名);
	//									唯一约束
	//									ALTER TABLE 表名 ADD UNIQUE(列名);
	//									CREATE TABLE 表名 (字段 属性 UNIQUE,...);
	//									表字段值自动增加
	//									CREATE TABLE 表名(字段 属性 AUTO_INCREMENT,...);
	//									ALTER TABLE 表名 MODIFY 字段 属性 PRIMAPY KEY AUTO_INCREMENT;
	//									索引
	//									CREATE TABLE 表名 (字段 属性,...,INDEX [索引名](字段[(长度)]));
	//									CREATE INDEX 索引名 ON 表名(字段([长度]));
	//									唯一索引
	//									CREATE TABLE 表名 (字段 属性,...,UNIQUE INDEX [索引名](字段([长度]));
	//									CREATE UNIQUE INDEX 索引名 ON 表名(字段([长度]));
	//									非空约束
	//									CREATE TABLE 表名(字段 属性 NOT NULL,...);
	//									ALTER TABLE 表名 MODIFY 字段 属性 NOT NULL;
	//									默认值约束
	//									CREATE TABLE 表名(字段 属性 DEFAULT 默认值,...);
	//									ALTER TABLE 表名 MODIFT 字段 相似 DEFAUlT 默认值;
	//									CREATE TABLE 表名(字段 属性,...,FOREING KEY (外键字段名) REFERENCES 主表表名(主键字段名));
	//									删除外键约束
	//									ALTER TABLE 表名 DROP FOREING KEY 外键名;
	//									一对多、多对多的关系中可用联接查询
	//									合并结果集(不过滤重复数据)
	//									SELECT * FROM 表名1 UNION [ALL] SELECT * FROM 表名2;
	//									笛卡尔积
	//									SELECT 查询字段 FROM 表1 CROSS JOIN 表2;
	//									内连接
	//									SELECT 查询字段 FROM 表1 [INNER] JOIN 表2 ON 表1.关系字段 = 表2.关系字段 WHERE 条件;
	//									左外连接
	//									SELECt 查询字段 FROM 表1 LEFT [OUTER} JOIN 表2 ON 表1.关系字段 = 表2.关系字段 WHERE 条件;
	//									右外连接
	//									SELECt 查询字段 FROM 表1 RIGHT [OUTER} JOIN 表2 ON 表1.关系字段 = 表2.关系字段 WHERE 条件;
	//									多表连接
	//									SELECT 查询字段 FROM 表1 [别名] JOIN 表2 [别名] ON 表1.关系字段 = 表2.关系字段 JOIN ...;
	//									自然连接(左|右连接)
	//									SELECT 查询字段 FROM 表1 [别名] HATURAL [LEFT|RIGHT] JOIN 表2 [别名];
	//									自连接
	//									SELECT 查询字段 FROM 表名 [别名2],表名 [别名2] WHERE 查询条件;
	//									子查询即是嵌套查询，即在SELECT中包含SELECT。
}