package cmd

import "fmt"

const DateFormat = "02.01.2006г."

// Default values
const (
	DefaultOccupation    = "Гражданин"
	DefaultTimeOrdinance = "17 часов 30 минут"
	DefaultTimeAccident  = "11 часов 30 минут"
	DefaultDecision      = "Предупреждения"
)

var Placeholders = [][]interface{}{
	{NumberOfProtocol, "Введите № протокола: "},
	{DateOfProtocol, fmt.Sprintf("Введите дату регистрации протокола, в следующем формате - %s: ", DateFormat)},
	{FullName, "Введите полное имя - Магомадов Магомед Магомедович: "},
	{Birthday, "Введите дату рождения, в следующем формате - 30.12.1954г.: "},
	{PlaceOfBirth, "Введите место рождения(как в паспорте): "},
	{OfficialAddress, "Введите адрес регистрации: "},
	{ActualAddress, "Введите фактический адрес (по умолчанию - адрес регистрации): "},
	{Identifier, "Введите документ, удостоверяющий личность - 'Паспорт серия 9610 № 224309 выдан Отделом УФМС России по ЧР в Гудермесском районе 20.12.2012г.: "},
	{DateOfAccident, fmt.Sprintf("Введите дату происшествия, в следующем формате - %s (по умолчанию - дата регистрации протокола): ", DateFormat)},
	{TimeOfAccident, "Введите время происшествия, в следующем формате - 10 40 (по умолчанию - 11 30): "},
	{Occupation, "Введите должность - \"Директор\" || \"Гражданин\" (по умолчанию Гражданин): "},
	{NumberOfOrdinance, "Введите № постановления: "},
	{DateOfOrdinance, fmt.Sprintf("Введите дату рассмотрения протокола (Дата регистрации постановления), в следующем формате - %s  (по умолчанию следующий день от даты регистрации протокола): ", DateFormat)},
	{TimeOfOrdinance, "Введите время рассмотрения протокола, в следующем формате - 10 40  (по умолчанию - 17 30): "},
	{Decision, "Введите решение постановления в родительноме падеже, например - \"штрафа в размере 5000 (ПЯТЬ ТЫСЯЧ РУБЛЕЙ)\" || \"Предупреждения\"  (по умолчанию - Предупреждения): "},
}