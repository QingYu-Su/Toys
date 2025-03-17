# 项目说明

项目教学地址：[7天用Go从零实现ORM框架GeeORM](https://geektutu.com/post/geeorm.html)

## 第1天
1. 实现简单的log库，对外开放4个接口：Error，Errorf，Info，Infof，并且可以通过SetLevel设置日志级别，Info表示显示所有日志，Error表示只显示错误日志，Disable表示不显示日志。
2. 实现Session，封装对数据库的SQL操作。SQL操作可以让用户输入sql语句和"?"占位符，占位符可以被特定的输入变量所填充，这也就可以灵活输入sql语句了
3. 实现Engine，封装与用户的交互。

## 第2天
1. 为适配不同的数据库，映射数据类型和特定的 SQL 语句，创建 Dialect 层屏蔽数据库差异。
2. 设计 Schema，利用反射(reflect)完成结构体和数据库表结构的映射，包括表名、字段名、字段类型、字段 tag 等
3. 为session构造创建(create)、删除(drop)、存在性(table exists) 的 SQL 语句完成数据库表的基本操作

## 第3天
1. 设计clause，实现各个子句的生成规则，并且提供Set方法设置子句，Build方法按特定子句顺序构造完整的SQL语句
2. schema实现了RecordValues方法，输入一个对象，将其字段值结构为切片中的元素值，方便输入到SQL语句中
3. session实现了Insert和Find方法。Insert方法可以输入多个对象，然后解析为SQL语句直接插入到表中。Find方法输入一个空对象切片，然后自动查询表，并将表中每个查询结果构造成对象，组成对象切片返回。这些都利用了之前的clause子句

## 第4天
1. 添加update、delete和count对应的clause子句
2. 实现session中的update、delete和count方法
3. 实现session中的where、limit和order by方法，这些都是查询语句，如果在返回值中将整个session指针返回，那么就可以支持链式调用，如`s.Where("Age > 18").Limit(3).Find(&users)`
4. 实现session中的first方法，只返回一条记录，底层使用了limit和find结合链式调用。

## 第5天
1. 定义钩子，钩子函数由结构体实现，一般为数据库表类的结构体实现，可以在操作的前后自定义操作
2. 钩子函数不使用interface，而是使用反射实现，比较反常规。
3. 在find、insert方法内部调用钩子函数

## 第6天
1. 由于go使用tx提供对事务的支持，所以我们在 Session 结构体中新增成员变量` tx *sql.Tx`。具体来说，构造CommonDB作为tx和db的共有接口，调用DB函数时，当 tx 不为空时，则使用 tx 执行 SQL 语句，否则使用 db 执行 SQL 语句。这样既兼容了原有的执行方式，又提供了对事务的支持。
2. 在session中封装事务的 Begin、Commit 和 Rollback 三个接口
3. 在 `geeorm.go `中为用户提供傻瓜式/一键式使用的接口：Transaction。传入一个函数，Transaction会保证其中的事务执行，并且底层函数一旦执行失败，会自动回滚。

## 第7天
1. 在geeorm中实现数据库表的迁移，即Migrate函数，比如对原有的表增加字段和删除字段。
2. 增加字段可以直接使用SQL语句。
3. 删除字段可以通过复制新表时只复制想要的列，然后删除旧表，最后重命名新表即可。