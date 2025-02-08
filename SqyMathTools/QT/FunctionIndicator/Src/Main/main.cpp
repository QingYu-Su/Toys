#include "Main/QFunctionIndicatorWnd.h"

#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QFunctionIndicatorWnd w;
    w.show();
    return a.exec();
}
