/***************************************
     文件名：QFunctionDlg.h
     作者：苏青羽
     功能描述：函数表达式编辑对话框
     创建日期：2023-05-09
     版本：v1.0
     最后修改日期：2023-05-09
***************************************/
#ifndef QFUNCTIONDLG_H
#define QFUNCTIONDLG_H

#include <QDialog>
#include "Calculator.h"
namespace SML = SqyMathLibrary;  //命名空间重命名

namespace Ui {
class QFunctionDlg;
}

class QFunctionDlg : public QDialog
{
    Q_OBJECT

public:
    explicit QFunctionDlg(char dv, SML::MathExpression &expression, QString &editStr,QWidget *parent = nullptr);
    ~QFunctionDlg();

private slots:
    void on_m_Num_0_clicked();

    void on_m_Num_1_clicked();

    void on_m_Num_2_clicked();

    void on_m_Num_3_clicked();

    void on_m_Num_4_clicked();

    void on_m_Num_5_clicked();

    void on_m_Num_6_clicked();

    void on_m_Num_7_clicked();

    void on_m_Num_8_clicked();

    void on_m_Num9_clicked();

    void on_m_Dot_clicked();

    void on_m_Plus_clicked();

    void on_m_Minus_clicked();

    void on_m_Multiply_clicked();

    void on_m_Divide_clicked();

    void on_m_Pow_clicked();

    void on_m_Mod_clicked();

    void on_m_Sin_clicked();

    void on_m_Cos_clicked();

    void on_m_Tan_clicked();

    void on_m_Sqrt_clicked();

    void on_m_Log_clicked();

    void on_m_Ln_clicked();

    void on_m_Reciproal_clicked();

    void on_m_Factorial_clicked();

    void on_m_Negative_clicked();

    void on_m_E_clicked();

    void on_m_Pi_clicked();

    void on_m_Left_clicked();

    void on_m_Right_clicked();

    void on_m_Abs_clicked();

    void on_m_AllClear_clicked();

    void on_m_Delete_clicked();

    void on_m_DV_clicked();

    void on_buttonBox_accepted();

    void on_buttonBox_rejected();

private:
    Ui::QFunctionDlg *ui;

public:
    SML::MathExpression GetExpression();  //获得输入后的表达式
    QString GetEditStr();  //获得输入后的编辑框文本

private:
    bool IsNewNum();  //是否为新数字，该函数仅被数字button使用

private:
    char m_DV;  //自变量
    SML::MathExpression m_Expression;  //表达式
    QString m_EditStr;  //编辑框
};

#endif // QFUNCTIONDLG_H
