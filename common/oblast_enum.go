package common

// Oblast represents Ukrainian regions as an enum with manual numbering.
type Oblast int

// Define the enum values manually to match the specific IDs.
const (
	Khmelnytska             Oblast = 3
	Vinnytska               Oblast = 4
	Rivnenska               Oblast = 5
	Volynska                Oblast = 8
	Dnipropetrovska         Oblast = 9
	Zhytomyrska             Oblast = 10
	Zakarpatska             Oblast = 11
	Zaporizka               Oblast = 12
	IvanoFrankivska         Oblast = 13
	Kyivska                 Oblast = 14
	Kirovohradska           Oblast = 15
	Luhanska                Oblast = 16
	Mykolaivska             Oblast = 17
	Odeska                  Oblast = 18
	Poltavska               Oblast = 19
	Sumska                  Oblast = 20
	Ternopilska             Oblast = 21
	Kharkivska              Oblast = 22
	Khersonska              Oblast = 23
	Cherkaska               Oblast = 24
	Chernihivska            Oblast = 25
	Chernivetska            Oblast = 26
	Lvivska                 Oblast = 27
	Donetska                Oblast = 28
	AvtonomnaRespublikaKrym Oblast = 29
	Sevastopol              Oblast = 30
	Kyiv                    Oblast = 31
)

// Map of Oblast to their Ukrainian names
var OblastNames = map[Oblast]string{
	Khmelnytska:             "Хмельницька область",
	Vinnytska:               "Вінницька область",
	Rivnenska:               "Рівненська область",
	Volynska:                "Волинська область",
	Dnipropetrovska:         "Дніпропетровська область",
	Zhytomyrska:             "Житомирська область",
	Zakarpatska:             "Закарпатська область",
	Zaporizka:               "Запорізька область",
	IvanoFrankivska:         "Івано-Франківська область",
	Kyivska:                 "Київська область",
	Kirovohradska:           "Кіровоградська область",
	Luhanska:                "Луганська область",
	Mykolaivska:             "Миколаївська область",
	Odeska:                  "Одеська область",
	Poltavska:               "Полтавська область",
	Sumska:                  "Сумська область",
	Ternopilska:             "Тернопільська область",
	Kharkivska:              "Харківська область",
	Khersonska:              "Херсонська область",
	Cherkaska:               "Черкаська область",
	Chernihivska:            "Чернігівська область",
	Chernivetska:            "Чернівецька область",
	Lvivska:                 "Львівська область",
	Donetska:                "Донецька область",
	AvtonomnaRespublikaKrym: "Автономна Республіка Крим",
	Sevastopol:              "м. Севастополь",
	Kyiv:                    "м. Київ",
}

// GetOblastNameByNumber returns the name of the oblast given its number.
func GetOblastNameByNumber(num int) string {
	oblast := Oblast(num)
	if name, exists := OblastNames[oblast]; exists {
		return name
	}
	return "Unknown Oblast"
}
