#ifndef QTWOFUNCDLG_H
#define QTWOFUNCDLG_H

#include <QDialog>
#include "MathFunction.h"
namespace SML = SqyMathLibrary;  //命名空间重命名

namespace Ui {
class QTwoFuncDlg;
}

class QTwoFuncDlg : public QDialog
{
    Q_OBJECT

public:
    explicit QTwoFuncDlg(QString title, SML::MathExpression expressionX, QString expressionStrX,
                         SML::MathExpression expressionY,QString expressionStrY, double left,
                         double right, size_t precision = 1000,int lineWidth = 1,
                         int lineType = Qt::SolidLine, QColor color = QColor(0, 0, 0),
                         QWidget *parent = nullptr);
    ~QTwoFuncDlg();

private:
    Ui::QTwoFuncDlg *ui;
public:
    SML::MathFunction* GetMathFunction();  //获得函数类对象
    QString GetExpressionStrX();  //返回X(t)函数表达式字符串
    QString GetExpressionStrY();  //返回Y(t)函数表达式字符串
    size_t GetPrecision();  //获得函数精度
    double GetMin();  //获得函数定义域左端点
    double GetMax();  //获得函数定义域右端点
    int GetLineWidth();  //获得线宽
    int GetLineType();  //获得线型
    QColor GetLineColor();  //获得线颜色

private slots:

    void on_m_Button_LineColor_clicked();

    void on_buttonBox_accepted();

    void on_buttonBox_rejected();

    void on_m_Button_EditFunc_X_clicked();

    void on_m_Button_EditFunc_Y_clicked();

private:
    SML::MathExpression m_ExpressionX;  //X(t)函数表达式
    SML::MathExpression m_ExpressionY;  //Y(t)函数表达式
    QString m_ExpressionStrX;  //X(t)函数表达式字符串
    QString m_ExpressionStrY;  //Y(t)函数表达式字符串
    QColor m_LineColor;  //线颜色
};

#endif // QTWOFUNCDLG_H
