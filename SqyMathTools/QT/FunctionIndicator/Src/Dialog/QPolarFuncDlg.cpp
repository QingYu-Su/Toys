#include "QPolarFuncDlg.h"
#include "ui_QPolarFuncDlg.h"
#include "Dialog/QFunctionDlg.h"
#include <QColorDialog>

QPolarFuncDlg::QPolarFuncDlg(QString title, SML::MathExpression expression,
                             QString expressionStr, double left, double right, size_t precision,
                         int lineWidth, int lineType, QColor color,QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QPolarFuncDlg),
    m_Expression(expression),
    m_ExpressionStr(expressionStr),
    m_LineColor(color)
{
    ui->setupUi(this);
    setWindowTitle(title);
    ui -> m_Edit_Function -> setText(expressionStr);

    ui -> m_Edit_Min->setText(QString::number(left));
    ui -> m_Edit_Max->setText(QString::number(right));
    ui -> m_Edit_Precision -> setText(QString::number(precision));
    ui -> m_Edit_LineWidth -> setText(QString::number(lineWidth));

    //下来框索引从0开始，需要对Qt::PenStyle值减一
    ui -> m_Combox_LineType -> setCurrentIndex(lineType - 1);
}

QPolarFuncDlg::~QPolarFuncDlg()
{
    delete ui;
}

SML::FunctionExpression QPolarFuncDlg::GetFunctionExpression() {
    //返回一个封装好了的函数表达式
    return SML::FunctionExpression(this->m_Expression, 'r', 'a', GetMin(), GetMax());
}

SML::MathFunction* QPolarFuncDlg::GetMathFunction() {
    SML::FunctionExpression exp = this->GetFunctionExpression();
    return new SML::PolarFunction(exp); //返回极坐标函数对象
}

QString QPolarFuncDlg::GetExpressionStr() {
    return m_ExpressionStr;
}

size_t QPolarFuncDlg::GetPrecision() {
    return ui -> m_Edit_Precision -> text().toUInt();
}

double QPolarFuncDlg::GetMin() {
    return ui -> m_Edit_Min -> text().toDouble();
}

double QPolarFuncDlg::GetMax() {
    return ui -> m_Edit_Max -> text().toDouble();
}

int QPolarFuncDlg::GetLineWidth() {
    return ui -> m_Edit_LineWidth -> text().toInt();
}

int QPolarFuncDlg::GetLineType() {
    return ui -> m_Combox_LineType -> currentIndex() + 1;  //下拉框值索引从0开始，需要加1
}

QColor QPolarFuncDlg::GetLineColor() {
    return this->m_LineColor;
}

void QPolarFuncDlg::on_m_Button_EditFunc_clicked()
{
    //弹出函数表达式编辑对话框
    QFunctionDlg dlg('a', this->m_Expression, this -> m_ExpressionStr );
    if ( dlg.exec() == Accepted ){
        this->m_Expression = dlg.GetExpression();  //同步函数表达式
        this->m_ExpressionStr = dlg.GetEditStr();  //同步函数表达式字符串
        ui -> m_Edit_Function -> setText(dlg.GetEditStr()); //显示
    }
}

void QPolarFuncDlg::on_m_Button_LineColor_clicked()
{
    m_LineColor = QColorDialog::getColor();
}

void QPolarFuncDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QPolarFuncDlg::on_buttonBox_rejected()
{
    done(Rejected);
}
