/***************************************
     �ļ�����UniFunction.h
     ���ߣ�������
     ����������ͨ�ú���
     �������ڣ�2023-02-14
     �汾��v1.0
     ����޸����ڣ�2022-02-14
***************************************/

#ifndef _UNIFUNCTION_H_
#define _UNIFUNCTION_H_

#include <string>
#include "Calculator.h"
#include <vector>

namespace SqyMathLibrary {

    //ͨ�ú������ʽ
    struct FunctionExpression {
        MathExpression m_Expression;  //������ʽ�����к����Ա������ţ��޷�ֱ�Ӽ���
        char m_IV;  //����� 
        char m_DV;  //�Ա���
        OPERAND m_Left, m_Right;  //�Ա������������䣬���Ҷ˵㲻�ɳ���INF��Ĭ��Ϊ-INF~INF
        FunctionExpression(MathExpression& expression, char iv, char dv, OPERAND left = -INF, OPERAND right = INF);
    };

    //���������࣬�ṩһϵ�еķ���ר�����Դ��������ʽ
    class FunctionTool {
    public:  //��Ĭ�Ϸ���
        FunctionTool();

        //���º����ɱ������Զ����ɼ���
        //~FunctionTool() {};
        //FunctionTool(const FunctionTool&) {};
        //FunctionTool& operator=(const FunctionTool&) {};
    public:
        //���뺯�����ʽ�������е��Ա��������ض�ֵ������һ���������ѧ���ʽ
        MathExpression TranslateExpression(FunctionExpression& fe, OPERAND value);

        //���ض�ֵ���뺯�����ʽ�����㣬���ؽ��ֵ��INV��ʾ����ʧ��
        //ע�⣬����ֵINF��һ����Чֵ������ʾ����󣬶����Ǵ˴μ���ʧ��
        OPERAND GetValue(FunctionExpression& fe, OPERAND parameter);

        bool IsSuccess(); //�����Ƿ�ɹ�
        std::string GetError();  //��������ʧ��ԭ��
    private:
        void SetResult(bool res, std::string reason = ""); //���ò����ɹ����
    private:
        bool m_Success;  //��һ�β����Ľ��
        std::string m_Error;   //����һ�β���ʧ�ܣ��ᱣ����һ�β���ʧ�ܵ�ԭ��
        Calculator m_Calc;  //������
    }; 

}


#endif
