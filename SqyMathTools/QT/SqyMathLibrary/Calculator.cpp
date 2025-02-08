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

		//判断是否为负数
		if (s.size() > 1 && s[0] == '-') i++;

		for (; i < s.size(); i++) {

			if (s[i] == '.') {
				if (++dot > 1) return false;  //小数点超过1返回false
				continue;
			}
			if (s[i] < '0' || s[i] > '9') {  //非数字返回false
				return false;
			}
			
			if (dot == 0) digital++;  //计算小数点前数位
		}

		if (digital > MAX_DIGITAL) return false;  //超过最大数位限制，返回false
		
		return true;
	}

	bool Calculator::IsOperator(const std::string& s) const{
		if (GET_MATH_OPERATOR(s) == NULL) return false;  //未注册，非法运算符
		return true;
	}

	OPERAND Calculator::GetNum(const std::string& s) const{
		if (s == "e") return E;
		if (s == "pi") return PI;
		if (s == "zero") return 0; //zero为不可见0，是方便计算类似于-3+（-5）这样的表达式设计的
		if (s == "INF") return INF;
		OPERAND res =  stod(s);  //调用标准库函数
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
		this->SetResult(true);  //初始化运算结果
		if (this->m_MathOpts.empty() == true) {
			this->SetResult(false, CALC_ERROR_LACK_OPT);
			return;
		}

		MathOperator* opt = this->m_MathOpts.top();
		this->m_MathOpts.pop();

		//此时运算符号栈中不应该有左括号
		if (opt->GetSymbol() == "(") { 
			this->SetResult(false, CALC_ERROR_LACK_BRA);
			return;
		}

		//运算数不符合运算符目数
		if (this->m_Nums.size() < opt->GetCount()) {
			this->SetResult(false, CALC_ERROR_LACK_NUM);
			return;
		}

		OPERAND temp[2] = { 0 };
		//按运算符的目数取出相应数量的数字
		for (int i = 0; i < opt->GetCount(); i++) {
			temp[i] = this->m_Nums.top();
			this->m_Nums.pop();
		}

		//运算，结果保存到数字栈中
		OPERAND res = opt->Operate(temp[1], temp[0]);
		if (res > INF) res = INF;
		else if (res < -INF) res = -INF;
		this->m_Nums.push(res);

		//设置运算结果
		this->SetResult(opt->IsSuccess(), opt->GetError());
	}

	void Calculator::OptTackle(MathOperator* opt) {
		this->SetResult(true);  //初始化运算结果
		if (opt->GetSymbol() == "(") { //左括号直接压栈
			this->m_MathOpts.push(opt);
			return;
		}

		if (opt->GetSymbol() == ")") {  //右括号

			//取出所有运算符运算，直到左括号
			while (!this->m_MathOpts.empty() && this->m_MathOpts.top()->GetSymbol() != "(") {
				this->Operate();
				if (this->IsSuccess() == false) return;
			}

			if (this->m_MathOpts.empty()) {  //栈空，没有对应左括号
				this->SetResult(false, CALC_ERROR_LACK_BRA);
				return;
			}

			m_MathOpts.pop();  //消去左括号

		}
		else { //其他运算符

			//栈中运算符优先级>=当前运算符的优先计算
			while (!this->m_MathOpts.empty() && opt->GetPriority() <= this->m_MathOpts.top()->GetPriority()) {
				if (this->m_MathOpts.top()->GetSymbol() == "(") break; //遇到左括号，无需计算，直接退出
				
				this->Operate();
				if (this->IsSuccess() == false) return;
			}

			//压栈
			this->m_MathOpts.push(opt);
		}

	}


	OPERAND Calculator::Calculate(const MathExpression& expression) {
		//检查表达式是否为空
		if (expression.empty()) {
			this->SetResult(false, CALC_ERROR_EMPTY_EXPR);
			return 0;
		}

		this->SetResult(true);  //初始化运算结果
		//遍历
		for (int i = 0; i < expression.size(); i++) {

			//不允许空括号存在，否则会导致计算bug
			if (i > 0 && expression[i - 1] == "(" && expression[i] == ")") {
				this->SetResult(false, CALC_ERROR_EMPTY_BRA);
				this->Reset();
				return 0;
			}

			if (this->IsNum(expression[i]) == true) {  //数字
				this->m_Nums.push(this->GetNum(expression[i])); //直接压栈
			}
			else if (this->IsOperator(expression[i]) == true) { //运算符
				MathOperator* opt = this->GetMathOperator(expression[i]);
				this->OptTackle(opt);
				if (this->IsSuccess() == false) {  //运算符处理失败，清空计算器并退出
					this->Reset();
					return 0;
				}
			}
			else {  //非数字非运算符
				this->SetResult(false, CALC_ERROR_EXPRESSION);
				this->Reset();
				return 0;
			}

		}

		//清算运算符，计算最终结果
		while (!this->m_MathOpts.empty()) {
			this->Operate();
			if (this->IsSuccess() == false) {
				this->Reset();
				return 0;
			}
		}

		//数字栈空间最终应只剩一个数字
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
