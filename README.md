# 🏃 Fitness Tracker

Go-модуль для обработки и расчёта показателей физической активности пользователя.

Проект построен на структурах и методах Go и использует интерфейсы для унификации работы с различными типами тренировок.

## Возможности

- расчёт пройденной дистанции и средней скорости;
- расчёт энергозатрат во время тренировок;
- обработка данных пользователя;
- работа с тренировками и дневной активностью;
- форматирование информации о тренировках;
- единый интерфейс для получения информации о различных видах активности.

## Структура проекта

```text
cmd/
└── tracker/
    └── main.go

internal/
├── actioninfo/
│   ├── actioninfo.go
│   └── actioninfo_test.go
├── daysteps/
│   ├── daysteps.go
│   └── daysteps_test.go
├── personaldata/
│   ├── personaldata.go
│   └── personaldata_test.go
├── spentenergy/
│   ├── spentenergy.go
│   └── spentenergy_test.go
└── trainings/
    ├── trainings.go
    └── trainings_test.go