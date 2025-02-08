#ifndef QFUNCNUMDLG_H
#define QFUNCNUMDLG_H

#include <QDialog>

namespace Ui {
class QFuncNumDlg;
}

class QFuncNumDlg : public QDialog
{
    Q_OBJECT

public:
    explicit QFuncNumDlg(QString title, QString tip,QWidget *parent = nullptr);
    ~QFuncNumDlg();
    int GetNum();

private slots:
    void on_buttonBox_accepted();

    void on_buttonBox_rejected();

private:
    Ui::QFuncNumDlg *ui;
};

#endif // QFUNCNUMDLG_H
