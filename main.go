package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type ClassLessonOccurrence struct {
	SubjectCode   string
	BuildingCodes []string
	RoomCodes     []string
	TeacherCodes  []string
	Dates         []string
	StartTime     string
	EndTime       string
	ChangeType    *string
}

const (
	ChangeRoom             = "room_change"
	ChangeAdditionalLesson = "lesson_change"
	ChangeTeacherAbsence   = "teacher_absence"
	ChangeClassAbsence     = "class_absence"
	ChangeUnknown          = "unknown"
)

type DescriptedCodeEntity struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type DescriptedTeacherEntity struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Code      string `json:"code"`
}

type TeacherLessonOccurrence struct {
	SubjectCode       string
	ClassCodes        []string
	RoomCodes         []string
	Dates             []string
	BuildingCodes     []string
	OtherTeacherCodes []string
	StartTime         string
	EndTime           string
	ChangeType        *string
}

type SimplifiedTeacher struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CacheState struct {
	ETag                string
	ScheduleDescription string
	Timeslots           []CacheTimeslot
	ClassMap            map[string]string
	SubjectMap          map[string]string
	RoomMap             map[string]string
	TeacherMap          map[string]SimplifiedTeacher
	ClassLessonMap      map[string][]ClassLessonOccurrence
	TeacherLessonMap    map[string][]TeacherLessonOccurrence
}

type ClassesResponse struct {
	Teachers  map[string]SimplifiedTeacher `json:"teachers"`
	Subjects  map[string]string            `json:"subjects"`
	Rooms     map[string]string            `json:"rooms"`
	Timeslots []CacheTimeslot              `json:"timeslots"`
	Result    []ResponseEntity             `json:"result"`
}

type CacheTimeslot struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type ResponseClassOccurrence struct {
	SubjectCode  string   `json:"subject"`
	TeacherCodes []string `json:"teachers"`
	RoomCodes    []string `json:"rooms"`
	ClassCodes   []string `json:"classes"`
	StartTime    string   `json:"start_time"`
	EndTime      string   `json:"end_time"`
	Dates        []string `json:"dates"`
	ChangeType   *string  `json:"change"`
}

type ResponseEntity struct {
	Code        string                    `json:"code"`
	Label       string                    `json:"label"`
	Occurrences []ResponseClassOccurrence `json:"occurrences"`
}

func getCaches(davinciURL string) (CacheState, error) {
	cache := CacheState{}
	resp, err := http.Get(davinciURL)
	if err != nil {
		return cache, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return cache, err
	}

	response := DavinciResponse{}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return cache, err
	}

	cache.ETag = response.About.ETag
	cache.ScheduleDescription = response.Result.DisplaySchedule.ScheduleDescription
	cache.ClassMap = map[string]string{}
	cache.SubjectMap = map[string]string{}
	cache.RoomMap = map[string]string{}
	cache.TeacherMap = map[string]SimplifiedTeacher{}
	cache.ClassLessonMap = map[string][]ClassLessonOccurrence{}
	cache.TeacherLessonMap = map[string][]TeacherLessonOccurrence{}

	for _, class := range response.Result.Classes {
		cache.ClassMap[class.Code] = class.Description
	}
	for _, subject := range response.Result.Subjects {
		cache.SubjectMap[subject.Code] = subject.Description
	}
	for _, teacher := range response.Result.Teachers {
		cache.TeacherMap[teacher.Code] = SimplifiedTeacher{FirstName: teacher.FirstName, LastName: teacher.LastName}
	}
	for _, room := range response.Result.Rooms {
		cache.RoomMap[room.Code] = room.Description
	}
	for _, timeslot := range response.Result.Timeframes[0].Timeslots {
		cache.Timeslots = append(cache.Timeslots, CacheTimeslot{
			StartTime: timeslot.StartTime,
			EndTime:   timeslot.EndTime,
		})
	}
	for _, lessonTime := range response.Result.DisplaySchedule.LessonTimes {
		for _, classCode := range lessonTime.ClassCodes {
			cache.ClassLessonMap[classCode] = append(cache.ClassLessonMap[classCode], ClassLessonOccurrence{
				SubjectCode:   lessonTime.SubjectCode,
				BuildingCodes: lessonTime.BuildingCodes,
				RoomCodes:     lessonTime.RoomCodes,
				Dates:         lessonTime.Dates,
				StartTime:     lessonTime.StartTime,
				EndTime:       lessonTime.EndTime,
				TeacherCodes:  lessonTime.TeacherCodes,
			})
		}
		for i, teacherCode := range lessonTime.TeacherCodes {
			otherTeacherCodes := make([]string, 0)
			for ii, code := range lessonTime.TeacherCodes {
				if i == ii {
					continue
				}
				otherTeacherCodes = append(otherTeacherCodes, code)
			}
			cache.TeacherLessonMap[teacherCode] = append(cache.TeacherLessonMap[teacherCode], TeacherLessonOccurrence{
				SubjectCode:       lessonTime.SubjectCode,
				ClassCodes:        lessonTime.ClassCodes,
				RoomCodes:         lessonTime.RoomCodes,
				Dates:             lessonTime.Dates,
				BuildingCodes:     lessonTime.BuildingCodes,
				OtherTeacherCodes: otherTeacherCodes,
				StartTime:         lessonTime.StartTime,
				EndTime:           lessonTime.EndTime,
			})
		}
	}
	return cache, nil

}

func main() {

	davinciURL := os.Getenv("DAVINCI_URL")
	if davinciURL == "" {
		panic("no DAVINCI_URL environment variable provided")
	}

	cs, err := getCaches(davinciURL)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			newCacheState, err := getCaches(davinciURL)
			if err != nil {
				fmt.Println("Failed to refresh cache:", err.Error())
			} else {
				cs = newCacheState
				fmt.Println("Cache successfully updated.")
			}

			// Wait for an hour
			time.Sleep(time.Hour)
		}
	}()

	http.HandleFunc("/classes", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		for class, _ := range cs.ClassMap {
			w.Write([]byte(class + "\n"))
		}
	})

	http.HandleFunc("/lessons", func(w http.ResponseWriter, r *http.Request) {
		queryClasses := r.URL.Query().Get("classes")
		queryTeachers := r.URL.Query().Get("teachers")

		fmt.Println("received query for", queryClasses, queryTeachers, r.Header.Get("If-None-Match"))

		if len(queryClasses) == 0 && len(queryTeachers) == 0 {
			http.Error(w, "provide classes with ?classes=xxx,xxx&teachers=xxx,xxx", http.StatusBadRequest)
			return
		}

		header := w.Header()
		header.Set("access-control-allow-origin", "*")

		classes := strings.Split(string(queryClasses), ",")
		teachers := strings.Split(string(queryTeachers), ",")

		hash := sha1.Sum([]byte("c" + queryClasses + "t" + queryTeachers + cs.ETag))
		etag := "\"" + hex.EncodeToString(hash[:]) + "\""
		if r.Header.Get("If-None-Match") == etag {
			header.Set("ETag", etag)
			w.WriteHeader(http.StatusNotModified)
			return
		}

		response := ClassesResponse{}
		response.Rooms = map[string]string{}
		response.Subjects = map[string]string{}
		response.Teachers = map[string]SimplifiedTeacher{}
		response.Timeslots = cs.Timeslots

		for _, class := range classes {
			occurrences, hasMap := cs.ClassLessonMap[class]
			if !hasMap {
				continue
			}

			currentClass := ResponseEntity{}
			currentClass.Code = class
			currentClass.Label = cs.ClassMap[class]

			for _, occurrence := range occurrences {
				for _, roomCode := range occurrence.RoomCodes {
					response.Rooms[roomCode] = cs.RoomMap[roomCode]
				}
				for _, teacherCodes := range occurrence.TeacherCodes {
					response.Teachers[teacherCodes] = cs.TeacherMap[teacherCodes]
				}
				response.Subjects[occurrence.SubjectCode] = cs.SubjectMap[occurrence.SubjectCode]

				currentClass.Occurrences = append(currentClass.Occurrences, ResponseClassOccurrence{
					SubjectCode:  occurrence.SubjectCode,
					TeacherCodes: occurrence.TeacherCodes,
					RoomCodes:    occurrence.RoomCodes,
					StartTime:    occurrence.StartTime,
					EndTime:      occurrence.EndTime,
					ChangeType:   occurrence.ChangeType,
					Dates:        occurrence.Dates,
				})
			}
			response.Result = append(response.Result, currentClass)
		}

		for _, teacherCode := range teachers {
			occurrences, hasMap := cs.TeacherLessonMap[teacherCode]
			if !hasMap {
				continue
			}

			currentTeacher := ResponseEntity{}
			currentTeacher.Code = teacherCode
			currentTeacher.Label = cs.TeacherMap[teacherCode].FirstName + " " + cs.TeacherMap[teacherCode].LastName

			for _, occurrence := range occurrences {
				for _, roomCode := range occurrence.RoomCodes {
					response.Rooms[roomCode] = cs.RoomMap[roomCode]
				}
				for _, teacherCodes := range occurrence.OtherTeacherCodes {
					response.Teachers[teacherCodes] = cs.TeacherMap[teacherCodes]
				}
				response.Subjects[occurrence.SubjectCode] = cs.SubjectMap[occurrence.SubjectCode]

				currentTeacher.Occurrences = append(currentTeacher.Occurrences, ResponseClassOccurrence{
					SubjectCode:  occurrence.SubjectCode,
					TeacherCodes: occurrence.OtherTeacherCodes,
					RoomCodes:    occurrence.RoomCodes,
					StartTime:    occurrence.StartTime,
					ClassCodes:   occurrence.ClassCodes,
					EndTime:      occurrence.EndTime,
					ChangeType:   occurrence.ChangeType,
					Dates:        occurrence.Dates,
				})
			}
			response.Result = append(response.Result, currentTeacher)
		}

		encoded, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusBadRequest)
			return
		}

		header.Set("content-type", "application/json")
		header.Set("etag", etag)
		w.Write(encoded)
	})

	fmt.Println("Server listening on port :8000")
	http.ListenAndServe("0.0.0.0:8000", http.DefaultServeMux)

}
