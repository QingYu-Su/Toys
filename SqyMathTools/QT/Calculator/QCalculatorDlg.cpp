#include "QCalculatorDlg.h"
#include "ui_QCalculatorDlg.h"

QCalculatorDlg::QCalculatorDlg(QWidget *parent)
    : QDialog(parent)
    , ui(new Ui::QCalculatorDlg)
{
    ui->setupUi(this);
    setWindowFlags(Qt::MSWindowsFixedSizeDialogHint);  //设置窗口不可更改大小
}

QCalculatorDlg::~QCalculatorDlg()
{
    delete ui;
}

bool QCalculatorDlg::IsNewNum() {
    if (this->m_Expression.empty()) return true;  //空表达式
    if (this->m_Expression.back() == "INF") return true;  //前一个为INF，无法为其添加新数位
    if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL) return true; //前一个占位符为运算符
    return false;
}

void QCalculatorDlg::on_m_Num_0_clicked()
{
    this->m_CurText += "0";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("0");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "0";
    }

}

void QCalculatorDlg::on_m_Num_1_clicked()
{
    this->m_CurText += "1";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("1");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "1";
    }
}

void QCalculatorDlg::on_m_Num_2_clicked()
{
    this->m_CurText += "2";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("2");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "2";
    }
}

void QCalculatorDlg::on_m_Num_3_clicked()
{
    this->m_CurText += "3";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("3");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "3";
    }
}

void QCalculatorDlg::on_m_Num_4_clicked()
{
    this->m_CurText += "4";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("4");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "4";
    }
}

void QCalculatorDlg::on_m_Num_5_clicked()
{
    this->m_CurText += "5";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("5");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "5";
    }
}

void QCalculatorDlg::on_m_Num_6_clicked()
{
    this->m_CurText += "6";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("6");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "6";
    }
}

void QCalculatorDlg::on_m_Num_7_clicked()
{
    this->m_CurText += "7";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("7");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "7";
    }
}

void QCalculatorDlg::on_m_Num_8_clicked()
{
    this->m_CurText += "8";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("8");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "8";
    }
}

void QCalculatorDlg::on_m_Num9_clicked()
{
    this->m_CurText += "9";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back("9");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += "9";
    }
}

void QCalculatorDlg::on_m_Dot_clicked()
{
    this->m_CurText += ".";  //当前编辑框添加字符
    ui -> m_Current -> setText(this->m_CurText);
    if (IsNewNum() == true) {  //新数字-在表达式数组添加新元素
       this->m_Expression.push_back(".");
    }
    else {  //非新数字-在表达式数组最后一个元素添加数字
       this->m_Expression[m_Expression.size() - 1] += ".";
    }
}

void QCalculatorDlg::on_m_Plus_clicked()
{
   this->m_CurText += "+";
   this->m_Expression.push_back("+");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Minus_clicked()
{
    //若减号前无元素或为括号，则需要在减号前加不可见0“zero”，方便如-3+（-7）之类的计算
    if (this->m_Expression.empty() || this->m_Expression.back() == "(") {
       this->m_Expression.push_back("zero");   //不可见0不会显示在编辑框中
    }
    this->m_CurText += "-";
    this->m_Expression.push_back("-");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Multiply_clicked()
{
    this->m_CurText += "*";
    this->m_Expression.push_back("*");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Divide_clicked()
{
    this->m_CurText += "/";
    this->m_Expression.push_back("/");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Pow_clicked()
{
    this->m_CurText += "^";
    this->m_Expression.push_back("^");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Mod_clicked()
{
    this->m_CurText += "mod";
    this->m_Expression.push_back("mod");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Sin_clicked()
{
    this->m_CurText += "sin(";
    this->m_Expression.push_back("sin");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Cos_clicked()
{
    this->m_CurText += "cos(";
    this->m_Expression.push_back("cos");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Tan_clicked()
{
    this->m_CurText += "tan(";
    this->m_Expression.push_back("tan");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Sqrt_clicked()
{
    this->m_CurText += "sqrt(";
    this->m_Expression.push_back("sqrt");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Log_clicked()
{
    this->m_CurText += "log(";
    this->m_Expression.push_back("log");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Ln_clicked()
{
    this->m_CurText += "ln(";
    this->m_Expression.push_back("ln");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Reciproal_clicked()
{
    this->m_CurText += "rec(";
    this->m_Expression.push_back("rec");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Factorial_clicked()
{
    this->m_CurText += "fact(";
    this->m_Expression.push_back("fact");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Negative_clicked()
{
    this->m_CurText += "neg(";
    this->m_Expression.push_back("neg");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_E_clicked()
{
    this->m_CurText += "e";
    this->m_Expression.push_back("e");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Pi_clicked()
{
    this->m_CurText += "pi";
    this->m_Expression.push_back("pi");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Left_clicked()
{
    this->m_CurText += "(";
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Right_clicked()
{
    this->m_CurText += ")";
    this->m_Expression.push_back(")");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Abs_clicked()
{
    this->m_CurText += "abs(";
    this->m_Expression.push_back("abs");
    this->m_Expression.push_back("(");
    ui -> m_Current -> setText(this->m_CurText);
}

void QCalculatorDlg::on_m_Equal_clicked()
{
    double res = m_Calculator.Calculate(this->m_Expression);  //计算
        if (m_Calculator.IsSuccess() == false) { //计算失败弹窗
            QMessageBox msg(QMessageBox::Critical, "计算器", m_Calculator.GetError().c_str());
            msg.exec();
            return;
        }

        this->m_CurText += "=";
        this->m_HisText = this->m_CurText;  //历史编辑框赋值

        this->m_Expression.clear();  //表达式清空
        this->m_CurText.clear(); //当前编辑框清空

        //为当前编辑框赋值
        if (res == INF) {
            this->m_CurText = "INF";
        }
        else if (res == -INF) {
            this->m_CurText = "-INF";
        }
        else {
            this->m_CurText = std::to_string(res).c_str();
        }

        //为表达式赋值
        if (res < 0) {  //结果为负数需要添加zero
            this->m_Expression.push_back("zero");
            this->m_Expression.push_back("-");
            res = -res;
        }

        if (res == INF) this->m_Expression.push_back("INF");
        else this->m_Expression.push_back(std::to_string(res));

        //显示编辑框文本内容
        ui -> m_Current -> setText(this->m_CurText);
        ui -> m_Histroy -> setText(this->m_HisText);
}

void QCalculatorDlg::on_m_AllClear_clicked()
{
    this->m_CurText.clear();
    this->m_HisText.clear();
    this->m_Expression.clear();
    ui -> m_Current -> setText(this->m_CurText);
    ui -> m_Histroy -> setText(this->m_HisText);
}

void QCalculatorDlg::on_m_Delete_clicked()
{
    if (this->m_CurText.isEmpty() == true) return; //当前编辑框为空，直接返回

        int length = 0;  //当前编辑框新长度
        if (SML::GET_MATH_OPERATOR(this->m_Expression.back()) != NULL || m_Expression.back() == "pi" || m_Expression.back() == "INF") {
            //表达式最后一个元素为运算符或者“pi”，删除该元素，当前编辑框长度更新
            length = m_CurText.size() - this->m_Expression.back().size();
            this->m_Expression.pop_back();
        }
        else {
            //表达式最后一个元素为数字，删除该数字的最后一位，当前编辑框长度更新
            length = m_CurText.size() - 1;
            this->m_Expression[m_Expression.size() - 1].pop_back();
            //表达式最后一个元素内容被清空，表达式需要删除该元素
            if (m_Expression[m_Expression.size() - 1].empty()) this->m_Expression.pop_back();;
        }

        //进行删除操作后，若剩下不可见0“zero”，则必须删除，保持表达式和编辑框的同步
        if ( !this->m_Expression.empty() && this->m_Expression.back() == "zero") {
            this->m_Expression.pop_back();
        }

        this->m_CurText = this->m_CurText.left(length);  //当前编辑框截取相应长度，从左边算起
        ui -> m_Current -> setText(this->m_CurText);  //显示编辑框
}
