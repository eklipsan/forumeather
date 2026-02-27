package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"
)

// Помощник serverError записывает сообщение об ошибке в errorLog и
// затем отправляет пользователю ответ 500 "Внутренняя ошибка сервера".
func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Помощник clientError отправляет определенный код состояния и соответствующее описание
// пользователю. Мы будем использовать это в следующий уроках, чтобы отправлять ответы вроде 400 "Bad
// Request", когда есть проблема с пользовательским запросом.
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Мы также реализуем помощник notFound. Это просто
// удобная оболочка вокруг clientError, которая отправляет пользователю ответ "404 Страница не найдена".
func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}


// Проверка времени на срок действия (12 часов)
func IsTimeExpired(t time.Time, hoursToExpire int) bool {
	return time.Since(t) > hoursToExpireCurrent * time.Hour
}


func GetTimeLabel(t time.Time) string {
	date := t.Format("02.01.2006")
	weekday := getRussianWeekday(t)

	return fmt.Sprintf("%s (%s)", date, weekday)
}

func getRussianWeekday(t time.Time) string {
    weekdays := map[time.Weekday]string{
        time.Sunday:    "воскресенье",
        time.Monday:    "понедельник",
        time.Tuesday:   "вторник",
        time.Wednesday: "среда",
        time.Thursday:  "четверг",
        time.Friday:    "пятница",
        time.Saturday:  "суббота",
    }
    return weekdays[t.Weekday()]
}