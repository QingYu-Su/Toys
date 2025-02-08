#ifndef QNORMALFUNCDLG_H
#define QNORMALFUNCDLG_H

#include <QDialog>
#include "MathFunction.h"
namespace SML = SqyMathLibrary;  //命名空间重命名

namespace Ui {
class QNormalFuncDlg;
}

class QNormalFuncDlg : public QDialog
{
    Q_OBJECT

public:
    explicit QNormalFuncDlg(QString title, SML::MathExpression expression,
                            QString expressionStr, double left, double right, size_t precision = 1000,
                        int lineWidth = 1, int lineType = Qt::SolidLine, QColor color = QColor(0, 0, 0),QWidget *parent = nullptr);
    ~QNormalFuncDlg();

private:
    Ui::QNormalFuncDlg *ui;

public:
    SML::FunctionExpression GetFunctionExpression();  //获得函数表达式
    SML::MathFunction* GetMathFunction();  //获得函数类对象
    QString GetExpressionStr();  //返回函数表达式字符串
    size_t GetPrecision();  //获得函数精度
    double GetMin();  //获得函数定义域左端点
    double GetMax();  //获得函数定义域右端点
    int GetLineWidth();  //获得线宽
    int GetLineType();  //获得线型
    QColor GetLineColor();  //获得线颜色

private slots:
    void on_m_Button_EditFunc_clicked();

    void on_m_Button_LineColor_clicked();

    void on_buttonBox_accepted();

    void on_buttonBox_rejected();

private:
    SML::MathExpression m_Expression;  //函数表达式
    QString m_ExpressionStr;  //函数表达式字符串
    QColor m_LineColor;  //线颜色
};

#endif // QNORMALFUNCDLG_H
