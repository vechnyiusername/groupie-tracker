# 🚀 **GROUPIE TRACKERS — ПОЛНЫЙ ROADMAP ОТ НУЛЯ ДО ГОТОВОГО ПРОЕКТА**

---

# **ЭТАП 0 — Подготовка (10 минут)**

### ✔️ Установи:

- Go
- Git

### ✔️ Проверь:

```bash
go version
git --version
```

### ✔️ Создай рабочую папку:

Например:

```
C:\projects\groupie
```

---

# **ЭТАП 1 — Создание проекта (10 минут)**

В терминале:

```bash
mkdir groupie-trackers
cd groupie-trackers
go mod init groupie-trackers
```

Готово: теперь Go понимает, что это твой модуль.

---

# **ЭТАП 2 — Структура проекта (10 минут)**

Создай такую структуру:

```
groupie-trackers/
│   main.go
│   go.mod
│
├── data/       ← JSON-файлы API (artists, locations, dates...)
├── templates/  ← HTML-страницы
└── static/     ← CSS и JS
```

Минимально нужны:

```
templates/
    index.html
    artist.html
static/
    style.css
    script.js
data/
    artists.json
    locations.json
    dates.json
    relations.json
```

---

# **ЭТАП 3 — Backend: запуск сервера (20 минут)**

В `main.go`:

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

Теперь:

```bash
go run .
```

Открой браузер и зайди:

```
http://localhost:8080
```

---

# **ЭТАП 4 — Загрузка данных API (1–1.5 часа)**

### Создай папку `models/` и файл `models.go`

Там будут структуры для JSON:

```go
package models

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
```

Также сделай структуры для:

- locations
- dates
- relation

Затем — писaть функцию загрузки JSON:

```go
func LoadArtists() []Artist {
	file, _ := os.ReadFile("data/artists.json")
	var artists []Artist
	json.Unmarshal(file, &artists)
	return artists
}
```

Ты сделаешь 4 такие функции.

Потом в `main.go` загружаешь данные:

```go
var Artists []models.Artist

func init() {
	Artists = models.LoadArtists()
}
```

---

# **ЭТАП 5 — Главная страница: список артистов (1 час)**

В `homeHandler`:

```go
t := template.Must(template.ParseFiles("templates/index.html"))
t.Execute(w, Artists)
```

В `templates/index.html`:

```html
{{range .}}
<div class="card">
  <img src="{{.Image}}" />
  <h3>{{.Name}}</h3>
  <p>Основана: {{.CreationDate}}</p>
  <a href="/artist?id={{.ID}}">Подробнее</a>
</div>
{{end}}
```

Теперь главная страница показывает список артистов.

---

# **ЭТАП 6 — Страница артиста (1–2 часа)**

### Создай хендлер:

```go
func artistHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for _, a := range Artists {
		if a.ID == id {
			t := template.Must(template.ParseFiles("templates/artist.html"))
			t.Execute(w, a)
			return
		}
	}

	http.Error(w, "Artist not found", 404)
}
```

### Добавь маршрут:

```go
http.HandleFunc("/artist", artistHandler)
```

### В шаблоне `artist.html`:

```html
<h1>{{.Name}}</h1>
<img src="{{.Image}}" />
<ul>
  {{range .Members}}
  <li>{{.}}</li>
  {{end}}
</ul>
```

---

# **ЭТАП 7 — Связь с датами и локациями (1–2 часа)**

Ты должен:

1. загрузить ещё `relations.json`
2. найти для артиста:

   - его города
   - его даты концертов

3. передать всё в HTML-шаблон

Это делается похожим кодом, только через ID.

---

# **ЭТАП 8 — Реализация EVENT / ACTION (1–1.5 часа)**

Самое важное для сдачи.

Самый простой вариант:

### 📌 Кнопка «Показать концерты» делает запрос на сервер

`artist.html`:

```html
<button onclick="loadConcerts({{.ID}})">Показать концерты</button>
<div id="concerts"></div>
<script src="/static/script.js"></script>
```

`static/script.js`:

```js
function loadConcerts(id) {
  fetch("/api/concerts?id=" + id)
    .then((r) => r.json())
    .then((data) => {
      let html = "";
      for (let city in data) {
        html += `<p><b>${city}:</b> ${data[city].join(", ")}</p>`;
      }
      document.getElementById("concerts").innerHTML = html;
    });
}
```

### В Go:

```go
func concertsAPI(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	rel := FindRelation(id) // твоя функция поиска связи

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rel.DatesLocations)
}
```

Добавь маршрут:

```go
http.HandleFunc("/api/concerts", concertsAPI)
```

Теперь у тебя есть **событие**:
клик → fetch → сервер → JSON → обновление страницы.

---

# **ЭТАП 9 — Ошибки (30 минут)**

Добавь обработку:

- `/404` — страница не найдена
- `/500` — ошибка сервера
- неправильный ID

Можно сделать:

```go
func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", 404)
}
```

---

# **ЭТАП 10 — Финальная проверка (30–60 минут)**

Проверь:

✔ `/` открывается
✔ `/artist?id=X` работает
✔ кнопка event / action работает
✔ ошибки отлавливаются
✔ сервер не падает
✔ всё отображается

---

# 🎉 Итог

Этот roadmap:

### ⭐ Полный

### ⭐ Пошаговый

### ⭐ Ведёт от нуля до готового проекта

### ⭐ Гарантирует сдачу (всё по требованиям)

Примерное время выполнения:

👉 **8–14 часов суммарно**, если идти спокойно по шагам.

---

Если хочешь — можем идти **вместе по каждому этапу**, и я буду писать тебе точный код и проверять твоё выполнение.
