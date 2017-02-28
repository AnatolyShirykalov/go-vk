package apivk

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v0"
	"strconv"
)

const root = "https://api.vk.com/method/"

func Init(tkn string) func(string, map[string]string) (*resty.Response, error) {
	return func(method string, pars map[string]string) (*resty.Response, error) {
		if _, ok := pars["access_token"]; !ok {
			pars["access_token"] = tkn
		}
		if _, ok := pars["version"]; !ok {
			pars["v"] = "5.62"
		}
		return resty.R().
			SetQueryParams(pars).
			SetHeader("Accept-Charset", "utf-8").
			Get(fmt.Sprintf("%s%s", root, method))
	}
}

func Woodpecker(tkn string) func(string, map[string]string) []interface{} {
	get := Init(tkn)
	return func(method string, params map[string]string) []interface{} {
		done := false
		ret := make([]interface{}, 20)
		for !done {
			resp, err := get(method, params)
			if err != nil {
				panic(err)
			}
			data := make(map[string]interface{})
			err = json.Unmarshal(resp.Body(), &data)
			if err != nil {
				panic(err)
			}
			if _, ok := data["response"]; ok {
				done = true
				res := data["response"].(map[string]interface{})
				count := int(res["count"].(float64))
				_, has_count := params["count"]
				if count > 0 && has_count {
					items := res["items"].([]interface{})
					ret = append(ret, items...)
					if count > len(items) {
						var offset int
						if _, has_offset := params["offset"]; !has_offset {
							offset = 0
						} else {
							v := params["offset"]
							offset, err = strconv.Atoi(v)
						}
						if offset+len(items) < count {
							offset += len(items)
							done = false
						}
					}
					fmt.Println(len(items), count)
				}
			} else {
				vk_err := data["error"].(map[string]interface{})
				if msg := vk_err["error_msg"]; msg == "Too many requests per second" {
					fmt.Println("error", msg, "Try again")
				} else {
					fmt.Println("Unknown error", msg)
					panic(nil)
				}
			}
		}
		return ret
	}
}

func Run(token string) {
	get := Woodpecker(token)
	targets := []string{
		"Авиамоторная",
		"Автозаводская",
		"Академическая",
		"Александровский сад",
		"Алексеевская",
		"Алма-Атинская",
		"Алтуфьево",
		"Аннино",
		"Арбатская",
		"Аэропорт",
		"Бабушкинская",
		"Багратионовская",
		"Баррикадная",
		"Бауманская",
		"Беговая",
		"Белорусская",
		"Беляево",
		"Бибирево",
		"Библиотека имени Ленина",
		"Борисово",
		"Боровицкая",
		"Ботанический сад",
		"Братиславская",
		"Бульвар адмирала Ушакова",
		"Бульвар Дмитрия Донского",
		"Бульвар Рокоссовского",
		"Бунинская аллея",
		"Варшавская",
		"ВДНХ",
		"Владыкино",
		"Водный стадион",
		"Войковская",
		"Волгоградский проспект",
		"Волжская",
		"Волоколамская",
		"Воробьевы горы",
		"Выставочная",
		"Выхино",
		"Деловой центр",
		"Динамо",
		"Дмитровская",
		"Добрынинская",
		"Домодедовская",
		"Достоевская",
		"Дубровка",
		"Жулебино",
		"Зябликово",
		"Измайловская",
		"Калужская",
		"Кантемировская",
		"Каховская",
		"Каширская",
		"Киевская",
		"Китай-город",
		"Кожуховская",
		"Коломенская",
		"Комсомольская",
		"Коньково",
		"Красногвардейская",
		"Краснопресненская",
		"Красносельская",
		"Красные ворота",
		"Крестьянская застава",
		"Кропоткинская",
		"Крылатское",
		"Кузнецкий мост",
		"Кузьминки",
		"Кунцевская",
		"Курская",
		"Кутузовская",
		"Ленинский проспект",
		"Лермонтовский проспект",
		"Лубянка",
		"Люблино",
		"Марксистская",
		"Марьина роща",
		"Марьино",
		"Маяковская",
		"Медведково",
		"Международная",
		"Менделеевская",
		"Митино",
		"Молодежная",
		"Мякинино",
		"Нагатинская",
		"Нагорная",
		"Нахимовский проспект",
		"Новогиреево",
		"Новокосино",
		"Новокузнецкая",
		"Новослободская",
		"Новоясеневская",
		"Новые Черемушки",
		"Октябрьская",
		"Октябрьское поле",
		"Орехово",
		"Отрадное",
		"Охотныйряд",
		"Павелецкая",
		"Парк культуры",
		"Парк Победы",
		"Партизанская",
		"Первомайская",
		"Перово",
		"Петровско-Разумовская",
		"Печатники",
		"Пионерская",
		"Планерная",
		"Площадь Ильича",
		"Площадь Революции",
		"Полежаевская",
		"Полянка",
		"Пражская",
		"Преображенская площадь",
		"Пролетарская",
		"Проспект Вернадского",
		"Проспект Мира",
		"Профсоюзная",
		"Пушкинская",
		"Пятницкое шоссе",
		"Речной вокзал",
		"Рижская",
		"Римская",
		"Рязанский проспект",
		"Савеловская",
		"Свиблово",
		"Севастопольская",
		"Семеновская",
		"Серпуховская",
		"Славянский бульвар",
		"Сокол",
		"Сокольники",
		"Спартак",
		"Спортивная",
		"Сретенский бульвар",
		"Строгино",
		"Студенческая",
		"Сухаревская",
		"Сходненская",
		"Таганская",
		"Тверская",
		"Театральная",
		"Текстильщики",
		"Теплый стан",
		"Тимирязевская",
		"Третьяковская",
		"Тропарево",
		"Трубная",
		"Тульская",
		"Тургеневская",
		"Тушинская",
		"Улица Академика Янгеля",
		"Улица Горчакова",
		"Улица Скобелевская",
		"Улица Старокачаловская",
		"Улица 1905 года",
		"Университет",
		"Филевский парк",
		"Фили",
		"Фрунзенская",
		"Царицыно",
		"Цветной бульвар",
		"Черкизовская",
		"Чертановская",
		"Чеховская",
		"Чистые пруды",
		"Чкаловская",
		"Шаболовская",
		"Шипиловская",
		"Шоссе Энтузиастов",
		"Щелковская",
		"Щукинская",
		"Электрозаводская",
		"Юго-Западная",
		"Южная",
		"Ясенево",
	}
	for _, station := range targets {
		data := get("groups.search", map[string]string{
			"q":     fmt.Sprintf("подслушано %s", station),
			"count": "20",
		})
		data = data
	}
}
