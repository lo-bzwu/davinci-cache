package main

import (
	"time"
)

type DavinciResponse struct {
	About struct {
		ETag            string `json:"eTag"`
		SchemaVersion   string `json:"schemaVersion"`
		Server          string `json:"server"`
		ServerTimeStamp string `json:"serverTimeStamp"`
		ServerVersion   string `json:"serverVersion"`
	} `json:"about"`
	Result struct {
		Buildings []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
			ID          string `json:"id"`
		} `json:"buildings"`
		ClassAbsenceReasons []any `json:"classAbsenceReasons"`
		ClassAbsences       []struct {
			ClassRef  string `json:"classRef"`
			EndDate   string `json:"endDate"`
			EndTime   string `json:"endTime"`
			ID        string `json:"id"`
			StartDate string `json:"startDate"`
			StartTime string `json:"startTime"`
		} `json:"classAbsences"`
		Classes []struct {
			Code          string   `json:"code"`
			Color         string   `json:"color,omitempty"`
			Description   string   `json:"description"`
			ID            string   `json:"id"`
			TeamRefs      []string `json:"teamRefs"`
			TimeframeCode string   `json:"timeframeCode,omitempty"`
		} `json:"classes"`
		Courses []struct {
			CourseNr    string `json:"courseNr,omitempty"`
			Description string `json:"description,omitempty"`
			ID          string `json:"id"`
			Remarks     string `json:"remarks,omitempty"`
			SubjectRef  string `json:"subjectRef,omitempty"`
			Title       string `json:"title,omitempty"`
		} `json:"courses"`
		DisplaySchedule struct {
			Display struct {
				AbsReasonCaption       int    `json:"absReasonCaption"`
				AbsenceColor           string `json:"absenceColor"`
				AbsenceReasons         int    `json:"absenceReasons"`
				AdditionalLessonColor  string `json:"additionalLessonColor"`
				AlwaysLessonTime       int    `json:"alwaysLessonTime"`
				BackgroundColor        string `json:"backgroundColor"`
				EventColor             string `json:"eventColor"`
				HeaderBgColor          string `json:"headerBgColor"`
				LessonChangeColor      string `json:"lessonChangeColor"`
				LessonColor            int    `json:"lessonColor"`
				LessonGradient         int    `json:"lessonGradient"`
				LessonLayout           int    `json:"lessonLayout"`
				LessonWeeks            string `json:"lessonWeeks"`
				MessageColor           string `json:"messageColor"`
				PosLabel               int    `json:"posLabel"`
				PublishDays            int    `json:"publishDays"`
				PublishSubstMessage    bool   `json:"publishSubstMessage"`
				RoomChangeColor        string `json:"roomChangeColor"`
				SupervisionChangeColor string `json:"supervisionChangeColor"`
				SupervisionColor       string `json:"supervisionColor"`
				TodayColor             string `json:"todayColor"`
				TodayHeaderColor       string `json:"todayHeaderColor"`
			} `json:"display"`
			Effectivity struct {
				EndDate   string `json:"endDate"`
				StartDate string `json:"startDate"`
			} `json:"effectivity"`
			EventTimes []struct {
				EndDate      string `json:"endDate"`
				EndTime      string `json:"endTime"`
				EventCaption string `json:"eventCaption"`
				EventRef     string `json:"eventRef"`
				EventStatus  int    `json:"eventStatus"`
				NoLessons    bool   `json:"noLessons"`
				StartDate    string `json:"startDate"`
				StartTime    string `json:"startTime"`
				WholeDay     bool   `json:"wholeDay"`
			} `json:"eventTimes"`
			LessonTimes []struct {
				BuildingCodes []string `json:"buildingCodes,omitempty"`
				Changes       *struct {
					AbsentClassCodes   []string `json:"absentClassCodes,omitempty"`
					AbsentRoomCodes    []string `json:"absentRoomCodes,omitempty"`
					AbsentTeacherCodes []string `json:"absentTeacherCodes,omitempty"`
					Caption            string   `json:"caption,omitempty"`
					ChangeType         int      `json:"changeType"`
					Information        string   `json:"information,omitempty"`
					LessonTitle        string   `json:"lessonTitle,omitempty"`
					Message            string   `json:"message,omitempty"`
					Modified           string   `json:"modified,omitempty"`
					NewRoomCodes       []string `json:"newRoomCodes,omitempty"`
					ReasonType         string   `json:"reasonType,omitempty"`
				} `json:"changes,omitempty"`
				ClassCodes   []string `json:"classCodes,omitempty"`
				CourseRef    string   `json:"courseRef,omitempty"`
				CourseTitle  string   `json:"courseTitle,omitempty"`
				Dates        []string `json:"dates"`
				EndTime      string   `json:"endTime"`
				LessonBlock  string   `json:"lessonBlock,omitempty"`
				LessonRef    string   `json:"lessonRef,omitempty"`
				RoomCodes    []string `json:"roomCodes,omitempty"`
				StartTime    string   `json:"startTime"`
				SubjectCode  string   `json:"subjectCode,omitempty"`
				TeacherCodes []string `json:"teacherCodes,omitempty"`
			} `json:"lessonTimes"`
			ScheduleDescription string `json:"scheduleDescription"`
			ScheduleID          string `json:"scheduleID"`
			Session             struct {
				EndDate   string `json:"endDate"`
				StartDate string `json:"startDate"`
			} `json:"session"`
			SupervisionTimes []any `json:"supervisionTimes"`
			Weekspan         struct {
				WeekdayEnd   int `json:"weekdayEnd"`
				WeekdayStart int `json:"weekdayStart"`
			} `json:"weekspan"`
		} `json:"displaySchedule"`
		FirstLesson time.Time `json:"firstLesson"`
		Resources   []struct {
			Code        string `json:"code"`
			Color       string `json:"color"`
			Description string `json:"description"`
			ID          string `json:"id"`
		} `json:"resources"`
		RoomAbsenceReasons []any `json:"roomAbsenceReasons"`
		RoomAbsences       []any `json:"roomAbsences"`
		Rooms              []struct {
			BuildingRef string   `json:"buildingRef,omitempty"`
			Code        string   `json:"code"`
			Description string   `json:"description"`
			ID          string   `json:"id"`
			TeamRefs    []string `json:"teamRefs,omitempty"`
		} `json:"rooms"`
		Subjects []struct {
			Code        string   `json:"code"`
			Color       string   `json:"color,omitempty"`
			Description string   `json:"description,omitempty"`
			ID          string   `json:"id"`
			TeamRefs    []string `json:"teamRefs,omitempty"`
		} `json:"subjects"`
		TeacherAbsenceReasons []any `json:"teacherAbsenceReasons"`
		TeacherAbsences       []struct {
			EndDate    string `json:"endDate"`
			EndTime    string `json:"endTime"`
			ID         string `json:"id"`
			Note       string `json:"note,omitempty"`
			StartDate  string `json:"startDate"`
			StartTime  string `json:"startTime"`
			TeacherRef string `json:"teacherRef"`
		} `json:"teacherAbsences"`
		Teachers []struct {
			Code          string   `json:"code"`
			FirstName     string   `json:"firstName"`
			ID            string   `json:"id"`
			LastName      string   `json:"lastName"`
			TeamRefs      []string `json:"teamRefs,omitempty"`
			TimeframeCode string   `json:"timeframeCode,omitempty"`
		} `json:"teachers"`
		Teams []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
			ID          string `json:"id"`
		} `json:"teams"`
		Timeframes []struct {
			Code                  string `json:"code"`
			TimeslotFragmentation int    `json:"timeslotFragmentation"`
			Timeslots             []struct {
				Color     string `json:"color,omitempty"`
				EndTime   string `json:"endTime"`
				Label     string `json:"label"`
				StartTime string `json:"startTime"`
			} `json:"timeslots"`
		} `json:"timeframes"`
	} `json:"result"`
	User struct {
		Policy  int    `json:"policy"`
		Profile string `json:"profile"`
	} `json:"user"`
}
