/***************************************
     文件名：MathDef.h
     作者：苏青羽
     功能描述：数学库预定义
     创建日期：2023-02-08
     版本：v1.0
     最后修改日期：2022-02-08
***************************************/

#ifndef _MATHDEF_H_
#define _MATHDEF_H_

#define PI atan(1.0) * 4
#define E 2.718281828459
#define INF 1e+15  //无穷大
#define INV 1e+16 //无效值
#define MAX_DIGITAL 14  //外部输入的最大数位
typedef double OPERAND;  //定义运算数类型

#define CALC_ERROR_TAN  "tan取值为无穷"
#define CALC_ERROR_DIV0  "不能除以0"
#define CALC_ERROR_EMPTY_EXPR  "空表达式"
#define CALC_ERROR_LACK_OPT "缺少运算符"
#define CALC_ERROR_LACK_BRA "括号数不匹配"
#define CALC_ERROR_LACK_NUM "缺少运算数"
#define CALC_ERROR_EXCE_NUM "存在多余运算数"
#define CALC_ERROR_EMPTY_BRA "括号里不允许为空"
#define CALC_ERROR_EXPRESSION "表达式违法"

#define FUNC_ERROR_DV "自变量符号错误"
#define FUNC_ERROR_IV "因变量符号错误"
#define FUNC_ERROR_DEF_RANGE "函数定义域或计算范围错误"




#endif
