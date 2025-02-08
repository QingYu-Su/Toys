#include "Dialog/QNormalFuncDlg.h"
#include "ui_QNormalFuncDlg.h"
#include "QFunctionDlg.h"
#include <QColorDialog>

QNormalFuncDlg::QNormalFuncDlg(QString title, SML::MathExpression expression,
                               QString expressionStr, double left, double right, size_t precision,
                           int lineWidth, int lineType, QColor color,QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QNormalFuncDlg),
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

QNormalFuncDlg::~QNormalFuncDlg()
{
    delete ui;
}

SML::FunctionExpression QNormalFuncDlg::GetFunctionExpression() {
    //返回一个封装好了的函数表达式
    return SML::FunctionExpression(this->m_Expression, 'y', 'x', GetMin(), GetMax());
}

SML::MathFunction* QNormalFuncDlg::GetMathFunction() {
    SML::FunctionExpression exp = this->GetFunctionExpression();
    return new SML::NormalFunction(exp); //返回普通函数对象
}

QString QNormalFuncDlg::GetExpressionStr() {
    return m_ExpressionStr;
}

size_t QNormalFuncDlg::GetPrecision() {
    return ui -> m_Edit_Precision -> text().toUInt();
}

double QNormalFuncDlg::GetMin() {
    return ui -> m_Edit_Min -> text().toDouble();
}

double QNormalFuncDlg::GetMax() {
    return ui -> m_Edit_Max -> text().toDouble();
}

int QNormalFuncDlg::GetLineWidth() {
    return ui -> m_Edit_LineWidth -> text().toInt();
}

int QNormalFuncDlg::GetLineType() {
    return ui -> m_Combox_LineType -> currentIndex() + 1;  //下拉框值索引从0开始，需要加1
}

QColor QNormalFuncDlg::GetLineColor() {
    return this->m_LineColor;
}

void QNormalFuncDlg::on_m_Button_EditFunc_clicked()
{
    //弹出函数表达式编辑对话框
    QFunctionDlg dlg('x', this->m_Expression, this -> m_ExpressionStr );
    if ( dlg.exec() == Accepted ){
        this->m_Expression = dlg.GetExpression();  //同步函数表达式
        this->m_ExpressionStr = dlg.GetEditStr();  //同步函数表达式字符串
        ui -> m_Edit_Function -> setText(dlg.GetEditStr()); //显示
    }
}

void QNormalFuncDlg::on_m_Button_LineColor_clicked()
{
    m_LineColor = QColorDialog::getColor();
}

void QNormalFuncDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QNormalFuncDlg::on_buttonBox_rejected()
{
    done(Rejected);
}
