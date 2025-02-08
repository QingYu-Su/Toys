#include "QCalculatorDlg.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QCalculatorDlg w;
    w.show();
    return a.exec();
}
