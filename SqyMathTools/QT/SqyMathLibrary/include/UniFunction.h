/***************************************
     文件名：UniFunction.h
     作者：苏青羽
     功能描述：通用函数
     创建日期：2023-02-14
     版本：v1.0
     最后修改日期：2022-02-14
***************************************/

#ifndef _UNIFUNCTION_H_
#define _UNIFUNCTION_H_

#include <string>
#include "Calculator.h"
#include <vector>

namespace SqyMathLibrary {

    //通用函数表达式
    struct FunctionExpression {
        MathExpression m_Expression;  //具体表达式，其中含有自变量符号，无法直接计算
        char m_IV;  //因变量 
        char m_DV;  //自变量
        OPERAND m_Left, m_Right;  //自变量定义域区间，左右端点不可超过INF，默认为-INF~INF
        FunctionExpression(MathExpression& expression, char iv, char dv, OPERAND left = -INF, OPERAND right = INF);
    };

    //函数工具类，提供一系列的方法专门用以处理函数表达式
    class FunctionTool {
    public:  //类默认方法
        FunctionTool();

        //以下函数由编译器自动生成即可
        //~FunctionTool() {};
        //FunctionTool(const FunctionTool&) {};
        //FunctionTool& operator=(const FunctionTool&) {};
    public:
        //翻译函数表达式，将其中的自变量赋予特定值并返回一个具体的数学表达式
        MathExpression TranslateExpression(FunctionExpression& fe, OPERAND value);

        //将特定值代入函数表达式并计算，返回结果值，INV表示计算失败
        //注意，返回值INF是一个有效值，他表示无穷大，而不是此次计算失败
        OPERAND GetValue(FunctionExpression& fe, OPERAND parameter);

        bool IsSuccess(); //操作是否成功
        std::string GetError();  //操作运算失败原因
    private:
        void SetResult(bool res, std::string reason = ""); //设置操作成功与否
    private:
        bool m_Success;  //上一次操作的结果
        std::string m_Error;   //若上一次操作失败，会保存上一次操作失败的原因
        Calculator m_Calc;  //计算器
    }; 

}


#endif
