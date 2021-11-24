package postgres

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-apk/database/dbmodels"
	"ottosfa-api-apk/models"
	"time"
)

// AttendancePostWithAccurationPhoto ..
func AttendancePostWithAccurationPhoto(req models.AttendancePostReq) error {
	fmt.Println(">>> AttendancePost - AttendancePostWithAccurationPhoto - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("INSERT INTO attendances "+
		" (sales_id, sales_phone, sales_name, attend_category, attend_category_type, type_attendance, clocktime_server, clocktime_local, selfie, location, longitude, latitude, notes, created_at, updated_at ) "+
		" VALUES ( ? , ? , ? , ? , ? , ? , now(), ? , ? , ? , ? , ? , ?, now(), now());",
		req.SalesID, req.SalesPhone, req.SalesName, req.AttendanceCategory, req.AttendanceCategoryType, req.TypeAttendance, req.Time, req.PhotoSelfie, req.Location, req.Longitude, req.Latitude, req.Notes).Error
	if err != nil {
		sugarLogger.Error("Failed Post Attendance")
		return err
	}
	return nil
}

// AttendancePostWithAccurationPhotoV23 ..
func AttendancePostWithAccurationPhotoV23(req models.AttendancePostReq, status int, today time.Time) error {
	fmt.Println(">>> AttendancePost - AttendancePostWithAccurationPhotoV23 - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	err := Dbcon.Exec("INSERT INTO attendances "+
		" (sales_id, sales_phone, sales_name, attend_category, attend_category_type, type_attendance, clocktime_server, clocktime_local, selfie, location, longitude, latitude, notes, status, created_at, updated_at ) "+
		" VALUES ( ? , ? , ? , ? , ? , ? , now(), ? , ? , ? , ? , ? , ?, ?, now(), now());",
		req.SalesID, req.SalesPhone, req.SalesName, req.AttendanceCategory, req.AttendanceCategoryType, req.TypeAttendance, req.Time, req.PhotoSelfie, req.Location, req.Longitude, req.Latitude, req.Notes, status).Error
	if err != nil {
		sugarLogger.Error("Failed Post Attendance")
		return err
	}
	return nil
}

// CheckStatusAbsenHarian ..
func CheckStatusAbsenHarian(req models.AttendancePostReq, today string) (res []dbmodels.Attendance, err error) {
	fmt.Println(">>> AttendancePost - CheckStatusAbsenHarian - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	var query = " select * from attendances where sales_id = " + fmt.Sprintf("%d",req.SalesID) + " and attend_category = 'Absen Harian' and" +
		" (clocktime_server between '"+today+" 00:00:00' and '"+today+" 23:59:59') and status in (1,2,3) "
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed Get Attendance")
		return res, err
	}
	return res, nil
}

// CheckStatusEvent ..
func CheckStatusEvent(req models.AttendancePostReq, today string) (res []dbmodels.Attendance, err error) {
	fmt.Println(">>> AttendancePost - CheckStatusEvent - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	time1 := today + ottoutils.GetEnv("ATTENDANCE_EVENT_TIME_1", "")
	time2 := today + ottoutils.GetEnv("ATTENDANCE_EVENT_TIME_2", "")

	var query = " select * from attendances where sales_id = " + fmt.Sprintf("%d",req.SalesID) + " and attend_category = 'Event' and" +
		" (clocktime_local between '"+time1+"' and '"+time2+" ') and status in (1,2,3) "
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Error("Failed Get Attendance")
		return res, err
	}
	return res, nil
}
