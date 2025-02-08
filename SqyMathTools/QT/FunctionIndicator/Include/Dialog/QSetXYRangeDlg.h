#ifndef QSETXYRANGEDLG_H
#define QSETXYRANGEDLG_H

#include <QDialog>

namespace Ui {
class QSetXYRangeDlg;
}

class QSetXYRangeDlg : public QDialog
{
    Q_OBJECT

public:
    explicit QSetXYRangeDlg(double minX, double maxX,
                            double minY, double maxY,
                            QWidget *parent = nullptr);
    ~QSetXYRangeDlg();

private:
    Ui::QSetXYRangeDlg *ui;

public:
    double GetMinX();
    double GetMaxX();
    double GetMinY();
    double GetMaxY();
private slots:
    void on_buttonBox_accepted();
    void on_buttonBox_rejected();
};

#endif // QSETXYRANGEDLG_H
