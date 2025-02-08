#include "QFuncNumDlg.h"
#include "ui_QFuncNumDlg.h"

QFuncNumDlg::QFuncNumDlg(QString title, QString tip, QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QFuncNumDlg)
{
    ui->setupUi(this);
    setWindowTitle(title);
    ui -> m_Tip -> setText(tip);
    ui -> m_Edit_Num -> setText(QString::number(1));
    setWindowFlags(Qt::MSWindowsFixedSizeDialogHint);  //设置窗口不可更改大小
}

QFuncNumDlg::~QFuncNumDlg()
{
    delete ui;
}

int QFuncNumDlg::GetNum() {
    return ui -> m_Edit_Num -> text().toInt();
}

void QFuncNumDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QFuncNumDlg::on_buttonBox_rejected()
{
    done(Rejected);
}
