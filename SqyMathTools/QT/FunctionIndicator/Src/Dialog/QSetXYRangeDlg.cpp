#include "QSetXYRangeDlg.h"
#include "ui_QSetXYRangeDlg.h"

QSetXYRangeDlg::QSetXYRangeDlg(double minX, double maxX,
                               double minY, double maxY,
                               QWidget *parent) :
    QDialog(parent),
    ui(new Ui::QSetXYRangeDlg)
{
    ui->setupUi(this);
    ui -> m_Edit_Min_X -> setText(QString::number(minX));
    ui -> m_Edit_Max_X -> setText(QString::number(maxX));
    ui -> m_Edit_Min_Y -> setText(QString::number(minY));
    ui -> m_Edit_Max_Y -> setText(QString::number(maxY));
}

QSetXYRangeDlg::~QSetXYRangeDlg()
{
    delete ui;
}

double QSetXYRangeDlg::GetMinX() {
    return ui -> m_Edit_Min_X -> text().toDouble();
}

double QSetXYRangeDlg::GetMaxX() {
    return ui -> m_Edit_Max_X -> text().toDouble();
}

double QSetXYRangeDlg::GetMinY() {
    return ui -> m_Edit_Min_Y -> text().toDouble();
}

double QSetXYRangeDlg::GetMaxY() {
    return ui -> m_Edit_Max_Y -> text().toDouble();
}

void QSetXYRangeDlg::on_buttonBox_accepted()
{
    done(Accepted);
}

void QSetXYRangeDlg::on_buttonBox_rejected()
{
    done(Rejected);
}
