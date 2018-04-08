# Mysql的索引优化和语句优化工具

## Introduction

该工具可以实现如下两种SQL语句的优化：
 * 自动分析可以添加的索引(已完成)
 * 自动分析可以改写的方式（未完成）
  
## Prerequisites

本工具需要Go 1.8版本或以上版本


## Function
* 支持Mysql的Select语法
* 基于NFA实现的SQL词法分析  （已完成）
* 基于Goyacc实现的SQL语法分析 （已完成）
* 基于Visitor模式的AST （已完成）
* SQL的预处理,包括合法性验证,获取元数据 (已完成)
* 绑定SQL中的的column,alias的作用域 (已完成）
* 对所有的表达式实现constant folding (已完成)
* 对所有的表达式实现Constant Condition Removal (已完成)
* 参照Mysql选取索引的规则,检测所有的Join条件,Where条件,Order条件中，给出索引优化建议 (已完成)
* 实现逻辑执行计划的生成 (未完成)
* 实现子查询的优化(未完成)





