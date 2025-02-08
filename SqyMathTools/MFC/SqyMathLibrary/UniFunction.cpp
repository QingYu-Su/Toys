#include "./include/UniFunction.h"


namespace SqyMathLibrary {

	FunctionExpression::FunctionExpression(MathExpression& expression, char iv, char dv, OPERAND left, OPERAND right)
		:m_Expression(expression), m_IV(iv), m_DV(dv)
	{
		if (right > INF) right = INF;  //���Ҷ˵㲻�ɳ���INF
		if (left < -INF) left = -INF;
		this->m_Left = left;
		this->m_Right = right;
	}

	FunctionTool::FunctionTool() {
		this->m_Error = "";
		this->m_Success = true;
	}

	MathExpression FunctionTool::TranslateExpression(FunctionExpression& fe, OPERAND value) {
		MathExpression res = fe.m_Expression;
		char key = fe.m_DV;

		for (int i = 0; i < res.size(); i++) {
			//��valueת��Ϊ�ַ��������滻���е��Ա���
			if (res[i].size() == 1 && res[i][0] == key) res[i] = std::to_string(value);
		}

		this->SetResult(true);
		return res;
	}

	OPERAND FunctionTool::GetValue(FunctionExpression& fe, OPERAND parameter) {

		//���뺯�����ʽ������
		MathExpression me = this->TranslateExpression(fe, parameter);
		OPERAND res = m_Calc.Calculate(me);
		
		//���������Ƿ񱨴�
		if (m_Calc.IsSuccess() == false) {
			std::string error = m_Calc.GetError();

			if (error == CALC_ERROR_TAN || error == CALC_ERROR_DIV0) {
				this->SetResult(true);  //tan���0���󲻻᷵�ش��󣬶��Ƿ���һ��INF��ʾ�õ�ȡֵ�����
				return INF;
			}
			else {
				this->SetResult(false, error);  //��������ֱ�ӷ���INV
				return INV;
			}

		}

		//���������󣬵�����������ֱ�ӷ���INV
		if (parameter < fe.m_Left || parameter > fe.m_Right) {
			this->SetResult(false, FUNC_ERROR_DEF_RANGE);
			return INV;
		}

		return res;
	}

	bool FunctionTool::IsSuccess() {
		return this->m_Success;
	}

	std::string  FunctionTool::GetError() {
		return this->m_Error;
	}

	void FunctionTool::SetResult(bool res, std::string reason) {
		this->m_Success = res;
		this->m_Error = reason;
	}
}