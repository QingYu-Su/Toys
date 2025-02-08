#include "./include/Calculator.h"

namespace SqyMathLibrary {

	Calculator::Calculator() {
		this->m_Success = true;
		this->m_Error = "";
	}

	bool Calculator::IsSuccess() const {
		return this->m_Success;
	}

	std::string Calculator::GetError() const {
		return this->m_Error;
	}

	bool Calculator::IsNum(const std::string& s) const{
		if (s == "e" || s == "pi" || s == "zero" || s == "INF") return true;
		if (s == ".") return false;

		int i = 0;
		int dot = 0;
		int digital = 0;

		//�ж��Ƿ�Ϊ����
		if (s.size() > 1 && s[0] == '-') i++;

		for (; i < s.size(); i++) {

			if (s[i] == '.') {
				if (++dot > 1) return false;  //С���㳬��1����false
				continue;
			}
			if (s[i] < '0' || s[i] > '9') {  //�����ַ���false
				return false;
			}
			
			if (dot == 0) digital++;  //����С����ǰ��λ
		}

		if (digital > MAX_DIGITAL) return false;  //���������λ���ƣ�����false
		
		return true;
	}

	bool Calculator::IsOperator(const std::string& s) const{
		if (GET_MATH_OPERATOR(s) == NULL) return false;  //δע�ᣬ�Ƿ������
		return true;
	}

	OPERAND Calculator::GetNum(const std::string& s) const{
		if (s == "e") return E;
		if (s == "pi") return PI;
		if (s == "zero") return 0; //zeroΪ���ɼ�0���Ƿ������������-3+��-5�������ı��ʽ��Ƶ�
		if (s == "INF") return INF;
		OPERAND res =  stod(s);  //���ñ�׼�⺯��
		return res;
	}

	MathOperator* Calculator::GetMathOperator(const std::string& s) const{
		return GET_MATH_OPERATOR(s); 
	}

	void Calculator::SetResult(bool res, std::string reason) {
		this->m_Success = res;
		this->m_Error = reason;
	}

	void Calculator::Reset() {
		while (!this->m_Nums.empty()) this->m_Nums.pop();
		while (!this->m_MathOpts.empty()) this->m_MathOpts.pop();
	}

	void Calculator::Operate() {
		this->SetResult(true);  //��ʼ��������
		if (this->m_MathOpts.empty() == true) {
			this->SetResult(false, CALC_ERROR_LACK_OPT);
			return;
		}

		MathOperator* opt = this->m_MathOpts.top();
		this->m_MathOpts.pop();

		//��ʱ�������ջ�в�Ӧ����������
		if (opt->GetSymbol() == "(") { 
			this->SetResult(false, CALC_ERROR_LACK_BRA);
			return;
		}

		//�����������������Ŀ��
		if (this->m_Nums.size() < opt->GetCount()) {
			this->SetResult(false, CALC_ERROR_LACK_NUM);
			return;
		}

		OPERAND temp[2] = { 0 };
		//���������Ŀ��ȡ����Ӧ����������
		for (int i = 0; i < opt->GetCount(); i++) {
			temp[i] = this->m_Nums.top();
			this->m_Nums.pop();
		}

		//���㣬������浽����ջ��
		OPERAND res = opt->Operate(temp[1], temp[0]);
		if (res > INF) res = INF;
		else if (res < -INF) res = -INF;
		this->m_Nums.push(res);

		//����������
		this->SetResult(opt->IsSuccess(), opt->GetError());
	}

	void Calculator::OptTackle(MathOperator* opt) {
		this->SetResult(true);  //��ʼ��������
		if (opt->GetSymbol() == "(") { //������ֱ��ѹջ
			this->m_MathOpts.push(opt);
			return;
		}

		if (opt->GetSymbol() == ")") {  //������

			//ȡ��������������㣬ֱ��������
			while (!this->m_MathOpts.empty() && this->m_MathOpts.top()->GetSymbol() != "(") {
				this->Operate();
				if (this->IsSuccess() == false) return;
			}

			if (this->m_MathOpts.empty()) {  //ջ�գ�û�ж�Ӧ������
				this->SetResult(false, CALC_ERROR_LACK_BRA);
				return;
			}

			m_MathOpts.pop();  //��ȥ������

		}
		else { //���������

			//ջ����������ȼ�>=��ǰ����������ȼ���
			while (!this->m_MathOpts.empty() && opt->GetPriority() <= this->m_MathOpts.top()->GetPriority()) {
				if (this->m_MathOpts.top()->GetSymbol() == "(") break; //���������ţ�������㣬ֱ���˳�
				
				this->Operate();
				if (this->IsSuccess() == false) return;
			}

			//ѹջ
			this->m_MathOpts.push(opt);
		}

	}


	OPERAND Calculator::Calculate(const MathExpression& expression) {
		//�����ʽ�Ƿ�Ϊ��
		if (expression.empty()) {
			this->SetResult(false, CALC_ERROR_EMPTY_EXPR);
			return 0;
		}

		this->SetResult(true);  //��ʼ��������
		//����
		for (int i = 0; i < expression.size(); i++) {

			//����������Ŵ��ڣ�����ᵼ�¼���bug
			if (i > 0 && expression[i - 1] == "(" && expression[i] == ")") {
				this->SetResult(false, CALC_ERROR_EMPTY_BRA);
				this->Reset();
				return 0;
			}

			if (this->IsNum(expression[i]) == true) {  //����
				this->m_Nums.push(this->GetNum(expression[i])); //ֱ��ѹջ
			}
			else if (this->IsOperator(expression[i]) == true) { //�����
				MathOperator* opt = this->GetMathOperator(expression[i]);
				this->OptTackle(opt);
				if (this->IsSuccess() == false) {  //���������ʧ�ܣ���ռ��������˳�
					this->Reset();
					return 0;
				}
			}
			else {  //�����ַ������
				this->SetResult(false, CALC_ERROR_EXPRESSION);
				this->Reset();
				return 0;
			}

		}

		//������������������ս��
		while (!this->m_MathOpts.empty()) {
			this->Operate();
			if (this->IsSuccess() == false) {
				this->Reset();
				return 0;
			}
		}

		//����ջ�ռ�����Ӧֻʣһ������
		if (this->m_Nums.size() != 1) {
			this->SetResult(false, CALC_ERROR_EXCE_NUM);
			this->Reset();
			return 0;
		}

		OPERAND res = this->m_Nums.top();
		this->m_Nums.pop();

		return res;
	}

	OPERAND Calculator::Calculate(const std::string expression) {
		MathExpression splitExprssion;
		int i = 0, j = 0;
		while (j < expression.size()) {
			if (expression[j] == '#') {
				splitExprssion.push_back(expression.substr(i, j - i));
				i = j + 1;
			}
			j++;
		}
		splitExprssion.push_back(expression.substr(i, j - i));
		return this->Calculate(splitExprssion);
	}
}