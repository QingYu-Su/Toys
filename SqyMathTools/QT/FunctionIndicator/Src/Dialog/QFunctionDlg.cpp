#include "Dialog/QFunctionDlg.h"
#include "ui_QFunctionDlg.h"

QFunctionDlg::QFunctionDlg(char dv, SML::MathExpression &expression, QString &editStr,QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QFunctionDlg),
    m_DV(dv),
    m_Expression(expression),
    m_EditStr(editStr)
{
    ui->setupUi(this);
    setWindowFlags(Qt::MSWindowsFixedSizeDialogHint);  //设置窗口不可更改大小
    ui -> m_EditFunction -> setText(this->m_EditStr);  //显示编辑框
}

QFunctionDlg::~QFunctionDlg()
{
    delete ui;
}


SML::MathExpression QFunctionDlg::GetExpression(){
    return this -> m_Expression;
}

QString QFunctionDlg::GetEditStr(){
    return this->m_EditStr;
}

bool QFunctionDlg::IsNewNum() {
    if (this->m_Expression.empty()) return true;  //空表达式
    if (this->m_Expression.back() == "INF") return true;  //前一个为INF，无法为其添加新数位
    if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL) return true; //前一个占位符为运算符
    return false;
}

void QFunctionDlg::on_m_Num_0_clicked()
{
    this->m_EditStr += "0";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("0");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "0";
    }

}

void QFunctionDlg::on_m_Num_1_clicked()
{
    this->m_EditStr += "1";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("1");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "1";
    }
}

void QFunctionDlg::on_m_Num_2_clicked()
{
    this->m_EditStr += "2";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("2");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "2";
    }
}

void QFunctionDlg::on_m_Num_3_clicked()
{
    this->m_EditStr += "3";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("3");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "3";
    }
}

void QFunctionDlg::on_m_Num_4_clicked()
{
    this->m_EditStr += "4";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("4");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "4";
    }
}

void QFunctionDlg::on_m_Num_5_clicked()
{
    this->m_EditStr += "5";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("5");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "5";
    }
}

void QFunctionDlg::on_m_Num_6_clicked()
{
    this->m_EditStr += "6";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("6");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "6";
    }
}

void QFunctionDlg::on_m_Num_7_clicked()
{
    this->m_EditStr += "7";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("7");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "7";
    }
}

void QFunctionDlg::on_m_Num_8_clicked()
{
    this->m_EditStr += "8";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("8");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "8";
    }
}

void QFunctionDlg::on_m_Num9_clicked()
{
    this->m_EditStr += "9";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("9");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "9";
    }
}

void QFunctionDlg::on_m_Dot_clicked()
{
    this->m_EditStr += ".";  //当前编辑框添加字符
    ui -> m_EditFunction -> setText(this->m_EditStr);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back(".");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += ".";
    }
}

void QFunctionDlg::on_m_Plus_clicked()
{
   this->m_EditStr += "+";
   this->m_Expression.push_back("+");
   ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Minus_clicked()
{
    //若减号前无元素或为括号，则需要在减号前加不可见0“zero”，方便如-3+（-7）之类的计算
    if (this->m_Expression.empty() || this->m_Expression.back() == "(") {
       this->m_Expression.push_back("zero");   //不可见0不会显示在编辑框中
    }
    this->m_EditStr += "-";
    this->m_Expression.push_back("-");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Multiply_clicked()
{
    this->m_EditStr += "*";
    this->m_Expression.push_back("*");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Divide_clicked()
{
    this->m_EditStr += "/";
    this->m_Expression.push_back("/");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Pow_clicked()
{
    this->m_EditStr += "^";
    this->m_Expression.push_back("^");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Mod_clicked()
{
    this->m_EditStr += "mod";
    this->m_Expression.push_back("mod");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Sin_clicked()
{
    this->m_EditStr += "sin(";
    this->m_Expression.push_back("sin");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Cos_clicked()
{
    this->m_EditStr += "cos(";
    this->m_Expression.push_back("cos");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Tan_clicked()
{
    this->m_EditStr += "tan(";
    this->m_Expression.push_back("tan");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Sqrt_clicked()
{
    this->m_EditStr += "sqrt(";
    this->m_Expression.push_back("sqrt");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Log_clicked()
{
    this->m_EditStr += "log(";
    this->m_Expression.push_back("log");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Ln_clicked()
{
    this->m_EditStr += "ln(";
    this->m_Expression.push_back("ln");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Reciproal_clicked()
{
    this->m_EditStr += "rec(";
    this->m_Expression.push_back("rec");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Factorial_clicked()
{
    this->m_EditStr += "fact(";
    this->m_Expression.push_back("fact");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Negative_clicked()
{
    this->m_EditStr += "neg(";
    this->m_Expression.push_back("neg");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_E_clicked()
{
    this->m_EditStr += "e";
    this->m_Expression.push_back("e");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Pi_clicked()
{
    this->m_EditStr += "pi";
    this->m_Expression.push_back("pi");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Left_clicked()
{
    this->m_EditStr += "(";
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Right_clicked()
{
    this->m_EditStr += ")";
    this->m_Expression.push_back(")");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Abs_clicked()
{
    this->m_EditStr += "abs(";
    this->m_Expression.push_back("abs");
    this->m_Expression.push_back("(");
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_AllClear_clicked()
{
    this->m_EditStr.clear();
    this->m_Expression.clear();
    ui -> m_EditFunction -> setText(this->m_EditStr);
}

void QFunctionDlg::on_m_Delete_clicked()
{
    if (this->m_EditStr.isEmpty() == true) return; //当前编辑框为空，直接返回

        int length = 0;  //当前编辑框新长度
        if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL || m_Expression.back() == "pi" || m_Expression.back() == "INF") {
            //表达式最后一个元素为运算符或者“pi”，删除该元素，当前编辑框长度更新
            length = m_EditStr.size() - this->m_Expression.back().size();
            this->m_Expression.pop_back();
        }
        else {
            //表达式最后一个元素为数字，删除该数字的最后一位，当前编辑框长度更新
            length = m_EditStr.size() - 1;
            this->m_Expression[m_Expression.size() - 1].pop_back();
            //表达式最后一个元素内容被清空，表达式需要删除该元素
            if (m_Expression[m_Expression.size() - 1].empty()) this->m_Expression.pop_back();;
        }

        //进行删除操作后，若剩下不可见0“zero”，则必须删除，保持表达式和编辑框的同步
        if ( !this->m_Expression.empty() && this->m_Expression.back() == "zero") {
            this->m_Expression.pop_back();
        }

        this->m_EditStr = this->m_EditStr.left(length);  //当前编辑框截取相应长度，从左边算起
        ui -> m_EditFunction -> setText(this->m_EditStr);  //显示编辑框
}

void QFunctionDlg::on_m_DV_clicked()
{
    std::string dv;
    dv.push_back(this->m_DV);
    this->m_EditStr += dv.c_str();
    this->m_Expression.push_back(dv);
    ui -> m_EditFunction -> setText(this->m_EditStr);  //显示编辑框
}

void QFunctionDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QFunctionDlg::on_buttonBox_rejected()
{
    done(Rejected);
}
