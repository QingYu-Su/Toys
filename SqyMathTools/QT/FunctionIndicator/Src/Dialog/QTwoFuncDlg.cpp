#include "QTwoFuncDlg.h"
#include "ui_QTwoFuncDlg.h"
#include "QFunctionDlg.h"
#include <QColorDialog>

QTwoFuncDlg::QTwoFuncDlg(QString title, SML::MathExpression expressionX, QString expressionStrX,
                         SML::MathExpression expressionY,QString expressionStrY, double left,
                         double right, size_t precision,int lineWidth,
                         int lineType, QColor color,QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QTwoFuncDlg),
    m_ExpressionX(expressionX),
    m_ExpressionStrX(expressionStrX),
    m_ExpressionY(expressionY),
    m_ExpressionStrY(expressionStrY),
    m_LineColor(color)
{
    ui->setupUi(this);
    setWindowTitle(title);
    ui -> m_Edit_Function_X -> setText(expressionStrX);
    ui -> m_Edit_Function_Y -> setText(expressionStrY);

    ui -> m_Edit_Min->setText(QString::number(left));
    ui -> m_Edit_Max->setText(QString::number(right));
    ui -> m_Edit_Precision -> setText(QString::number(precision));
    ui -> m_Edit_LineWidth -> setText(QString::number(lineWidth));

    //下来框索引从0开始，需要对Qt::PenStyle值减一
    ui -> m_Combox_LineType -> setCurrentIndex(lineType - 1);
}

QTwoFuncDlg::~QTwoFuncDlg()
{
    delete ui;
}

SML::MathFunction* QTwoFuncDlg::GetMathFunction() {
    //创建X(t)函数表达式
    SML::FunctionExpression funcExpressionX(this->m_ExpressionX, 'x', 't', GetMin(), GetMax());

    //创建Y(t)函数表达式
    SML::FunctionExpression funcExpressionY(this->m_ExpressionY, 'y', 't', GetMin(), GetMax());

    return new SML::TwoFunction(funcExpressionX, funcExpressionY); //返回参数方程函数对象
}

QString QTwoFuncDlg::GetExpressionStrX() {
    return m_ExpressionStrX;
}

QString QTwoFuncDlg::GetExpressionStrY() {
    return m_ExpressionStrY;
}

size_t QTwoFuncDlg::GetPrecision() {
    return ui -> m_Edit_Precision -> text().toUInt();
}

double QTwoFuncDlg::GetMin() {
    return ui -> m_Edit_Min -> text().toDouble();
}

double QTwoFuncDlg::GetMax() {
    return ui -> m_Edit_Max -> text().toDouble();
}

int QTwoFuncDlg::GetLineWidth() {
    return ui -> m_Edit_LineWidth -> text().toInt();
}

int QTwoFuncDlg::GetLineType() {
    return ui -> m_Combox_LineType -> currentIndex() + 1;  //下拉框值索引从0开始，需要加1
}

QColor QTwoFuncDlg::GetLineColor() {
    return this->m_LineColor;
}

void QTwoFuncDlg::on_m_Button_LineColor_clicked()
{
    m_LineColor = QColorDialog::getColor();
}

void QTwoFuncDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QTwoFuncDlg::on_buttonBox_rejected()
{
    done(Rejected);
}

void QTwoFuncDlg::on_m_Button_EditFunc_X_clicked()
{
    //弹出函数表达式编辑对话框
    QFunctionDlg dlg('t', this->m_ExpressionX, this -> m_ExpressionStrX );
    if ( dlg.exec() == Accepted ){
        this->m_ExpressionX = dlg.GetExpression();  //同步函数表达式
        this->m_ExpressionStrX = dlg.GetEditStr();  //同步函数表达式字符串
        ui -> m_Edit_Function_X -> setText(dlg.GetEditStr()); //显示
    }
}

void QTwoFuncDlg::on_m_Button_EditFunc_Y_clicked()
{
    //弹出函数表达式编辑对话框
    QFunctionDlg dlg('t', this->m_ExpressionY, this -> m_ExpressionStrY );
    if ( dlg.exec() == Accepted ){
        this->m_ExpressionY = dlg.GetExpression();  //同步函数表达式
        this->m_ExpressionStrY = dlg.GetEditStr();  //同步函数表达式字符串
        ui -> m_Edit_Function_Y -> setText(dlg.GetEditStr()); //显示
    }
}
