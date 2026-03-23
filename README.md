# Boot.dev Projects - Sergey D.

Репозиторий содержит мои проекты, выполненные в рамках обучения на платформе Boot.dev.

---

## 🐊 Project 1: RSS Gator (Current)
**Gator** — это CLI-агрегатор RSS-лент. Система позволяет подписываться на ресурсы, автоматически собирать посты в PostgreSQL и читать их в терминале.

### Как запустить:
1. **Перейдите в папку проекта**: `cd rss_gator`
2. **Установите зависимости**: `go mod tidy`
3. **Настройте БД**: 
   * Создайте базу `gator` в Postgres.
   * Выполните миграции: `goose -dir sql/schema postgres "postgres://USER:PASS@localhost:5432/gator" up`
4. **Соберите программу**: `go build -o gator`
5. **Использование**:
   * `./gator register <name>`
   * `./gator agg 1m` (запуск воркера)
   * `./gator browse 5` (чтение постов)

---

## 🎮 Project 2: Pokedex CLI
**Pokedex** — консольное приложение для работы с PokeAPI. Позволяет исследовать локации, ловить покемонов и просматривать их статы.

### Как запустить:
1. `cd pokedex-cli`
2. `go build -o pokedex`
3. `./pokedex`