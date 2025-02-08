/***************************************
     �ļ�����Calculator.h
     ���ߣ�������
     ������������ѧ���ʽ����
     �������ڣ�2023-02-06
     �汾��v1.0
     ����޸����ڣ�2022-02-06
***************************************/

#ifndef _CALCULATOR_H_
#define _CALCULATOR_H_

#include <string>
#include <stack>
#include <vector>
#include "MathOperatorMap.h"

namespace SqyMathLibrary {

    typedef std::vector<std::string> MathExpression;

    class Calculator {

    public:  //��Ĭ�Ϸ���
        Calculator();

        //���º����ɱ������Զ����ɼ���
        //~Calculator() {};
        //Calculator(const Calculator&) {};
        //Calculator& operator=(const Calculator&) {};

    public:  //�����ӿ�
        //�����ۺϱ��ʽ������Ϊ�������������������ַ������飬����������
        OPERAND Calculate(const MathExpression &expression);

        //�����ۺϱ��ʽ������Ϊ���ʽ�ַ������Ҹ�Ԫ��֮���ԡ�#���Ÿ���������������
        //��3+(5*4)������ʽ�ַ���Ӧ��Ϊ3#+#(#5#*#4#)
        OPERAND Calculate(const std::string expression);

        bool IsSuccess() const;  //�����Ƿ�ɹ�
        std::string GetError() const;  //��ȡ����ʧ��ԭ��

    private:  //˽�нӿ�
        bool IsNum( const std::string &s) const;  //�Ƿ�Ϊ����
        bool IsOperator(const std::string &s) const;  //�Ƿ�Ϊ�����

        OPERAND GetNum(const std::string &s) const;  //���ַ���ת��Ϊ����
        MathOperator* GetMathOperator(const std::string &s) const;  //��ȡ����������

        void SetResult(bool res, std::string reason = "");  //���ñ��ʽ����ɹ����
        void Reset(); //��ռ�����ջ������
        
        void Operate();  //���е�����������㣬���������ջ��
        void OptTackle( MathOperator* opt); //ʶ�������������Ӧ����
        

        
    private:
        std::stack<OPERAND> m_Nums;  //����ջ
        std::stack<MathOperator*> m_MathOpts;  //�����ջ
        bool m_Success;  //��һ������Ľ��
        std::string m_Error;  //����һ������ʧ�ܣ��ᱣ����һ������ʧ�ܵ�ԭ��
    
    };

}




#endif


