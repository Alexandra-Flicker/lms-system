# 📚 Полное руководство по изучению LMS системы

## 🎯 Цель документа
Этот документ поможет новичку в программировании понять:
- Архитектуру современного Go приложения
- Принципы чистой архитектуры
- Основы языка Go
- Инструменты разработки (Docker, Keycloak)
- Как все компоненты работают вместе

---

## 🏗️ Что такое наш проект?

### Описание системы
**LMS (Learning Management System)** - это система управления обучением, которая позволяет:

**Для студентов:**
- Просматривать доступные курсы
- Покупать курсы
- Изучать уроки и главы
- Управлять своим профилем

**Для преподавателей:**
- Создавать курсы
- Добавлять главы и уроки
- Управлять контентом

**Для администраторов:**
- Управлять пользователями
- Контролировать всю систему
- Назначать роли

### Технологический стек
- **Backend**: Go (Golang)
- **База данных**: PostgreSQL
- **Аутентификация**: Keycloak
- **Контейнеризация**: Docker, Docker Compose
- **API**: REST API
- **Документация**: Swagger

---

## 🧩 Архитектура проекта

### Что такое "Чистая архитектура"?

```mermaid
graph TB
    subgraph "Внешний мир"
        UI["Web UI"]
        DB[("Database")]
        API["External APIs"]
    end
    
    subgraph "Наше приложение"
        subgraph "Delivery Layer"
            HTTP["HTTP Handlers"]
            MW["Middleware"]
        end
        
        subgraph "Business Layer"
            SRV["Services"]
            DOM["Domain Logic"]
        end
        
        subgraph "Infrastructure Layer"
            REPO["Repositories"]
            CLIENT["HTTP Clients"]
        end
    end
    
    UI --> HTTP
    HTTP --> SRV
    SRV --> REPO
    REPO --> DB
    SRV --> CLIENT
    CLIENT --> API
    
    classDef delivery fill:#e1f5fe
    classDef business fill:#f3e5f5
    classDef infra fill:#e8f5e8
    
    class HTTP,MW delivery
    class SRV,DOM business
    class REPO,CLIENT infra
```

Чистая архитектура - это способ организации кода, где:
1. **Бизнес-логика не зависит от внешних систем**
2. **Код легко тестировать**
3. **Легко заменять компоненты**
4. **Четкое разделение ответственности**

### Структура нашего проекта

```
📁 lms-system/
├── 🚀 cmd/lms/                    # Точка входа приложения
├── ⚙️  config/                     # Конфигурация
├── 🏠 internal/                   # Внутренний код приложения
│   ├── 🌐 delivery/               # Слой представления (HTTP)
│   │   └── http/
│   │       ├── 📋 handlers/       # Обработчики HTTP запросов
│   │       ├── 🛡️  middleware/     # Промежуточное ПО
│   │       ├── 🗺️  router.go       # Маршрутизация
│   │       └── 🖥️  server.go       # HTTP сервер
│   ├── 🧠 domain/                 # Бизнес логика (ядро)
│   │   ├── 📦 dto/               # Объекты передачи данных
│   │   ├── 🏗️  entity/            # Бизнес сущности
│   │   └── 🔌 *_interface.go     # Интерфейсы
│   ├── 🔧 infrastructure/         # Внешние зависимости
│   │   ├── 🌍 clients/           # Внешние API клиенты
│   │   └── 💾 repository/        # Работа с БД
│   └── ⚡ service/               # Бизнес логика
├── 🗃️  migrations/                 # Миграции БД
└── 🛠️  utils/                     # Утилиты
```

### Диаграмма зависимостей слоев

```mermaid
graph TD
    subgraph "🌐 Delivery Layer"
        A["HTTP Handlers"]
        B["Middleware"]
        C["Router"]
    end
    
    subgraph "⚡ Service Layer"
        D["Auth Service"]
        E["LMS Service"]
    end
    
    subgraph "🧠 Domain Layer"
        F["Entities"]
        G["DTOs"]
        H["Interfaces"]
    end
    
    subgraph "🔧 Infrastructure Layer"
        I["PostgreSQL Repository"]
        J["Keycloak Client"]
    end
    
    A --> D
    A --> E
    D --> J
    E --> I
    D -.-> H
    E -.-> H
    I -.-> F
    J -.-> G
    
    classDef delivery fill:#ffebcd
    classDef service fill:#e6f3ff
    classDef domain fill:#f0fff0
    classDef infra fill:#fff0f5
    
    class A,B,C delivery
    class D,E service
    class F,G,H domain
    class I,J infra
```

### Принцип слоев

#### 1. **Domain (Домен)** - Ядро системы
```go
// Пример сущности Course
type Course struct {
    ID          uint   `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description"`
    Price       int    `json:"price"`
}
```

**Что здесь:**
- Бизнес сущности (Course, User, Lesson)
- Интерфейсы для репозиториев и сервисов
- DTO (Data Transfer Objects) для передачи данных
- Бизнес правила

#### 2. **Service (Сервисы)** - Бизнес логика
```go
func (s *Service) CreateCourse(ctx context.Context, course entity.Course) (uint, error) {
    // Валидация бизнес правил
    if course.Title == "" {
        return 0, errors.New("title is required")
    }
    
    // Сохранение через репозиторий
    return s.repo.CreateCourse(ctx, course)
}
```

**Что здесь:**
- Обработка бизнес логики
- Валидация данных
- Координация между компонентами

#### 3. **Infrastructure (Инфраструктура)** - Внешние системы
```go
func (r *Repository) CreateCourse(ctx context.Context, course entity.Course) (uint, error) {
    var id uint
    err := r.db.QueryRow(
        "INSERT INTO courses (title, description, price) VALUES ($1, $2, $3) RETURNING id",
        course.Title, course.Description, course.Price,
    ).Scan(&id)
    return id, err
}
```

**Что здесь:**
- Работа с базой данных
- HTTP клиенты для внешних API
- Файловая система

#### 4. **Delivery (Доставка)** - Интерфейс пользователя
```go
func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
    var course entity.Course
    json.NewDecoder(r.Body).Decode(&course)
    
    id, err := h.service.CreateCourse(r.Context(), course)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    
    json.NewEncoder(w).Encode(map[string]uint{"id": id})
}
```

**Что здесь:**
- HTTP обработчики
- Сериализация/десериализация JSON
- Обработка ошибок

---

## 🔤 Основы языка Go

### 1. Структуры (Structs)
```go
// Определение структуры
type User struct {
    ID       uint   `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// Создание экземпляра
user := User{
    ID:       1,
    Username: "john",
    Email:    "john@example.com",
}
```

### 2. Интерфейсы (Interfaces)
```go
// Определение интерфейса
type UserRepository interface {
    CreateUser(ctx context.Context, user User) error
    GetUserByID(ctx context.Context, id uint) (*User, error)
}

// Реализация интерфейса
type PostgresUserRepository struct {
    db *sql.DB
}

func (r *PostgresUserRepository) CreateUser(ctx context.Context, user User) error {
    // Реализация
    return nil
}
```

### 3. Указатели (Pointers)
```go
// Функция принимает указатель
func UpdateUser(user *User) {
    user.Username = "new_username" // Изменяет оригинальный объект
}

// Функция принимает значение
func PrintUser(user User) {
    fmt.Println(user.Username) // Получает копию
}
```

### 4. Обработка ошибок
```go
func DoSomething() error {
    // Что-то пошло не так
    if someCondition {
        return errors.New("something went wrong")
    }
    return nil // Успех
}

// Использование
if err := DoSomething(); err != nil {
    log.Fatal(err) // Обработка ошибки
}
```

### 5. Goroutines и каналы (Concurrency)
```go
// Goroutine - легковесный поток
go func() {
    fmt.Println("Выполняется асинхронно")
}()

// Каналы для коммуникации
ch := make(chan string)
go func() {
    ch <- "Привет из goroutine"
}()
message := <-ch // Получение из канала
```

### 6. Контекст (Context)
```go
// Контекст для отмены операций и таймаутов
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

// Передача контекста в функции
func DoWork(ctx context.Context) error {
    select {
    case <-ctx.Done():
        return ctx.Err() // Операция отменена
    case <-time.After(1 * time.Second):
        // Работа выполнена
        return nil
    }
}
```

---

## 🌐 HTTP и REST API

### Что такое REST?
REST (Representational State Transfer) - архитектурный стиль для веб-сервисов:

- **GET** - получение данных
- **POST** - создание новых данных
- **PUT** - обновление существующих данных
- **DELETE** - удаление данных

### API Карта нашего проекта

```mermaid
graph LR
    subgraph "🔓 Public API"
        A["GET /public/courses"]
        B["GET /public/courses/id"]
    end
    
    subgraph "🔐 Auth API"
        C["POST /auth/login"]
        D["POST /auth/refresh"]
    end
    
    subgraph "👤 User API"
        E["PUT /user/profile"]
        F["PUT /user/change-password"]
        G["POST /user/courses/buy"]
        H["GET /user/courses"]
    end
    
    subgraph "👑 Admin API"
        I["POST /admin/users/register"]
        J["POST /admin/courses"]
        K["PUT /admin/courses/id"]
        L["DELETE /admin/courses/id"]
    end
    
    classDef public fill:#e8f5e8
    classDef auth fill:#fff3cd
    classDef user fill:#cce5ff
    classDef admin fill:#ffcccc
    
    class A,B public
    class C,D auth
    class E,F,G,H user
    class I,J,K,L admin
```

### REST Endpoints таблица

| Метод | Endpoint | Доступ | Описание |
|-------|----------|--------|----------|
| 🟢 GET | `/public/courses` | Все | Получить все курсы |
| 🟢 GET | `/public/courses/{id}` | Все | Получить курс по ID |
| 🟡 POST | `/auth/login` | Все | Авторизация |
| 🟡 POST | `/auth/refresh` | Все | Обновить токен |
| 🔵 PUT | `/user/profile` | Пользователь | Обновить профиль |
| 🔵 PUT | `/user/change-password` | Пользователь | Сменить пароль |
| 🔵 POST | `/user/courses/buy` | Пользователь | Купить курс |
| 🔴 POST | `/admin/users/register` | Админ | Создать пользователя |
| 🔴 POST | `/admin/courses` | Админ | Создать курс |
| 🔴 PUT | `/admin/courses/{id}` | Админ | Обновить курс |
| 🔴 DELETE | `/admin/courses/{id}` | Админ | Удалить курс |

### HTTP Status коды

| Код | Название | Когда используется | Пример |
|-----|----------|-------------------|---------|
| 🟢 200 | OK | Успешный запрос | Получение списка курсов |
| 🟢 201 | Created | Ресурс создан | Новый курс создан |
| 🟡 400 | Bad Request | Неверный запрос | Пустое поле "title" |
| 🔴 401 | Unauthorized | Не авторизован | Нет токена |
| 🔴 403 | Forbidden | Нет прав доступа | Обычный пользователь пытается создать курс |
| 🔴 404 | Not Found | Ресурс не найден | Курс с ID=999 не существует |
| 🔴 500 | Server Error | Ошибка сервера | База данных недоступна |

---

## 🛡️ Аутентификация и авторизация

### Что такое JWT (JSON Web Token)?
JWT - это токен, который содержит информацию о пользователе:

```
Header.Payload.Signature
```

**Пример:**
```json
{
  "sub": "user123",
  "roles": ["ROLE_USER"],
  "exp": 1640995200
}
```

### Процесс аутентификации

```mermaid
sequenceDiagram
    participant C as 👤 Client
    participant A as 🖥️ App
    participant K as 🔐 Keycloak
    
    Note over C,K: 1. Логин пользователя
    C->>A: POST /auth/login {username, password}
    A->>K: POST /token {username, password, client_id}
    K-->>A: {access_token, refresh_token}
    A-->>C: {access_token, refresh_token}
    
    Note over C,K: 2. Защищенный запрос
    C->>A: GET /user/profile<br/>Authorization: Bearer token
    A->>A: Проверить токен
    alt Токен валиден
        A-->>C: Profile data
    else Токен невалиден
        A-->>C: 401 Unauthorized
    end
    
    Note over C,K: 3. Обновление токена
    C->>A: POST /auth/refresh {refresh_token}
    A->>K: POST /token {refresh_token, grant_type}
    K-->>A: {new_access_token, new_refresh_token}
    A-->>C: {new_access_token, new_refresh_token}
```

### Пример работы с токенами

```bash
# 1. Логин
curl -X POST http://localhost:8082/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "user1",
    "password": "password123"
  }'

# Ответ:
{
  "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9...",
  "refresh_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9..."
}

# 2. Использование токена
curl -X GET http://localhost:8082/api/v1/user/profile \
  -H "Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9..."
```

### Роли в системе

```mermaid
graph TD
    subgraph "🏛️ Система ролей"
        ADMIN["👑 ROLE_ADMIN<br/>Полный доступ"]
        TEACHER["🎓 ROLE_TEACHER<br/>Создание контента"]
        USER["👤 ROLE_USER<br/>Базовый доступ"]
    end
    
    subgraph "📋 Права доступа"
        A1["Управление пользователями"]
        A2["Все административные функции"]
        T1["Создание курсов"]
        T2["Редактирование контента"]
        U1["Просмотр курсов"]
        U2["Покупка курсов"]
        U3["Управление профилем"]
    end
    
    ADMIN --> A1
    ADMIN --> A2
    ADMIN --> T1
    ADMIN --> T2
    ADMIN --> U1
    ADMIN --> U2
    ADMIN --> U3
    
    TEACHER --> T1
    TEACHER --> T2
    TEACHER --> U1
    TEACHER --> U2
    TEACHER --> U3
    
    USER --> U1
    USER --> U2
    USER --> U3
    
    classDef admin fill:#ffcccc
    classDef teacher fill:#ffffcc
    classDef user fill:#ccffcc
    classDef permission fill:#e6f3ff
    
    class ADMIN admin
    class TEACHER teacher
    class USER user
    class A1,A2,T1,T2,U1,U2,U3 permission
```

---

## 🔧 Инструменты разработки

### Docker - Контейнеризация

**Что такое Docker?**
Docker позволяет упаковать приложение со всеми зависимостями в контейнер, который может работать везде одинаково.

**Dockerfile - инструкция для создания образа:**
```dockerfile
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/lms/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

**Основные команды:**
```bash
docker build -t lms-app .          # Создать образ
docker run -p 8080:8080 lms-app    # Запустить контейнер
docker ps                          # Список запущенных контейнеров
docker logs container_id           # Посмотреть логи
```

### Docker Compose - Оркестрация контейнеров

**Что такое Docker Compose?**
Инструмент для управления несколькими контейнерами как единой системой.

#### Архитектура нашей системы

```mermaid
graph TB
    subgraph "🐳 Docker Compose Environment"
        subgraph "🖥️ Application"
            APP["lms-app:8082<br/>Go Application"]
        end
        
        subgraph "💾 Database"
            DB[("PostgreSQL:5432<br/>Database")]
        end
        
        subgraph "🔐 Authentication"
            KC["Keycloak:8080<br/>Auth Server"]
        end
        
        subgraph "🌐 Network"
            NET["lms-network<br/>Internal Docker Network"]
        end
    end
    
    subgraph "🌍 External"
        CLIENT["👤 Client<br/>Browser/Mobile"]
        ADMIN["👑 Admin<br/>Keycloak UI"]
    end
    
    CLIENT --> APP
    ADMIN --> KC
    APP --> DB
    APP --> KC
    
    APP -.-> NET
    DB -.-> NET
    KC -.-> NET
    
    classDef app fill:#e1f5fe
    classDef db fill:#f3e5f5
    classDef auth fill:#fff3e0
    classDef network fill:#f1f8e9
    classDef external fill:#fce4ec
    
    class APP app
    class DB db
    class KC auth
    class NET network
    class CLIENT,ADMIN external
```

**docker-compose.yml - конфигурация:**
```yaml
version: '3.8'
services:
  # 🖥️ Наше приложение
  app:
    build: .
    ports:
      - "8082:8082"  # Внешний:Внутренний порт
    depends_on:
      - db
      - keycloak
    environment:
      - DB_HOST=db   # Имя сервиса = hostname
      
  # 💾 База данных
  db:
    image: postgres:13
    environment:
      POSTGRES_DB: lms_system
      POSTGRES_USER: lms
      POSTGRES_PASSWORD: lms
    volumes:
      - postgres_data:/var/lib/postgresql/data
      
  # 🔐 Сервер аутентификации
  keycloak:
    image: quay.io/keycloak/keycloak:20.0
    environment:
      KEYCLOAK_ADMIN: admin
      KEYCLOAK_ADMIN_PASSWORD: admin
    ports:
      - "8080:8080"

volumes:
  postgres_data:  # Постоянное хранилище для БД
```

**Основные команды:**
```bash
docker-compose up -d        # Запустить все сервисы
docker-compose down         # Остановить все сервисы
docker-compose logs app     # Посмотреть логи приложения
docker-compose ps           # Статус сервисов
```

### Keycloak - Управление аутентификацией

**Что такое Keycloak?**
Keycloak - это сервер аутентификации и авторизации, который предоставляет:
- Управление пользователями
- Единый вход (SSO)
- Управление ролями
- OAuth 2.0 / OpenID Connect
- Социальные логины

**Основные концепции:**
- **Realm** - изолированная область для пользователей
- **Client** - приложение, которое использует Keycloak
- **User** - пользователь системы
- **Role** - роль пользователя
- **Token** - JWT токен с информацией о пользователе

**Как мы используем Keycloak:**
1. Создаем пользователей через Admin API
2. Аутентифицируем пользователей
3. Получаем токены доступа
4. Проверяем права доступа

---

## 📊 База данных и миграции

### PostgreSQL
Реляционная база данных, которую мы используем для хранения:
- Курсов (courses)
- Глав (chapters)
- Уроков (lessons)
- Доступов пользователей (user_course_access)

### Пример таблицы:
```sql
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### GORM - ORM для Go
```go
type Course struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"not null"`
    Description string
    Price       int       `gorm:"default:0"`
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Автоматическая миграция
db.AutoMigrate(&Course{})

// Создание записи
db.Create(&Course{Title: "Go Course", Price: 1000})

// Поиск
var course Course
db.First(&course, 1) // Найти по ID
```

---

## 🔄 Жизненный цикл HTTP запроса

Давайте проследим, что происходит при запросе `GET /api/v1/public/courses`:

```mermaid
sequenceDiagram
    participant C as 👤 Client
    participant R as 🗺️ Router
    participant H as 📋 Handler
    participant S as ⚡ Service
    participant Repo as 💾 Repository
    participant DB as 🗄️ Database
    
    Note over C,DB: Получение списка курсов
    
    C->>R: GET /api/v1/public/courses
    Note right of R: 1. Маршрутизация
    R->>H: courseHandler.GetAllCourses()
    
    Note right of H: 2. HTTP Layer
    H->>S: service.GetAllCourses(ctx)
    
    Note right of S: 3. Business Logic
    S->>Repo: repository.GetAllCourses(ctx)
    
    Note right of Repo: 4. Data Access
    Repo->>DB: SELECT * FROM courses
    DB-->>Repo: [Course Records]
    
    Note right of Repo: 5. Маппинг в Entity
    Repo-->>S: []entity.Course
    
    Note right of S: 6. Бизнес-обработка
    S-->>H: []entity.Course, nil
    
    Note right of H: 7. HTTP Response
    H-->>R: JSON Response
    R-->>C: HTTP 200 + Course List
```

### Детальный разбор каждого шага

#### 🗺️ 1. Router (Маршрутизация)
```go
// internal/delivery/http/router.go
r.Get("/courses", courseHandler.GetAllCourses)
```
**Задача**: Направить запрос к правильному обработчику

#### 📋 2. Handler (HTTP слой)
```go
// internal/delivery/http/handlers/course/get_all_courses.go
func (h *Handler) GetAllCourses(w http.ResponseWriter, r *http.Request) {
    courses, err := h.service.GetAllCourses(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(courses)
}
```
**Задача**: HTTP протокол, сериализация JSON, обработка ошибок

#### ⚡ 3. Service (Бизнес логика)
```go
// internal/service/lms/get_all_courses.go
func (s *Service) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
    // Здесь может быть бизнес-логика:
    // - кеширование
    // - валидация
    // - логирование
    
    return s.courseRepo.GetAllCourses(ctx)
}
```
**Задача**: Бизнес правила, координация

#### 💾 4. Repository (Доступ к данным)
```go
// internal/infrastructure/repository/postgres/course/repository.go
func (r *Repository) GetAllCourses(ctx context.Context) ([]entity.Course, error) {
    var courses []entity.Course
    err := r.db.WithContext(ctx).Find(&courses).Error
    return courses, err
}
```
**Задача**: Работа с базой данных

#### 🗄️ 5. Database (Хранение)
```sql
SELECT id, title, description, price, created_at, updated_at 
FROM courses 
ORDER BY created_at DESC;
```
**Задача**: Хранение и извлечение данных

---

## 🧪 Тестирование

### Unit тесты
```go
func TestCreateCourse(t *testing.T) {
    // Arrange - подготовка
    mockRepo := &MockCourseRepository{}
    service := NewService(mockRepo, logger)
    course := entity.Course{Title: "Test Course"}
    
    // Act - выполнение
    id, err := service.CreateCourse(context.Background(), course)
    
    // Assert - проверка
    assert.NoError(t, err)
    assert.Greater(t, id, uint(0))
}
```

### Integration тесты
Тестирование всей цепочки: HTTP → Service → Database

---

## 🚀 Запуск и разработка

### Схема развертывания

```mermaid
flowchart TD
    subgraph "🔧 Development Flow"
        A["📥 git clone"] --> B["⚙️ Configure .env"]
        B --> C{"🤔 Development Mode?"}
        
        C -->|Yes| D["🐳 docker-compose up -d db keycloak"]
        D --> E["🏃 go run cmd/lms/main.go"]
        
        C -->|No| F["🐳 docker-compose up -d"]
        
        E --> G["🌐 localhost:8082"]
        F --> G
    end
    
    subgraph "🔍 Check Services"
        G --> H["✅ App: http://localhost:8082"]
        G --> I["🔐 Keycloak: http://localhost:8080"]
        G --> J["💾 PostgreSQL: localhost:5432"]
    end
    
    classDef dev fill:#e3f2fd
    classDef service fill:#f3e5f5
    classDef check fill:#e8f5e8
    
    class A,B,C,D,E,F dev
    class G service
    class H,I,J check
```

### 🏗️ Локальная разработка (рекомендуется)
```bash
# 1. Клонировать репозиторий
git clone <repository>
cd lms-system

# 2. Настроить переменные окружения
cp .env.example .env
# Отредактировать .env под ваши нужды

# 3. Запустить только зависимости
docker-compose up -d db keycloak

# 4. Подождать готовности сервисов (30-60 сек)
docker-compose logs keycloak  # Проверить готовность

# 5. Запустить приложение локально
go run cmd/lms/main.go

# 6. Проверить работоспособность
curl http://localhost:8082/api/v1/public/courses
```

### 🐳 Полный запуск в Docker
```bash
# Запустить всю систему
docker-compose up -d

# Проверить статус
docker-compose ps

# Посмотреть логи
docker-compose logs -f app
```

### 🔧 Полезные команды для разработки

| Команда | Описание | Когда использовать |
|---------|----------|--------------------|
| `go run cmd/lms/main.go` | Запуск без компиляции | Разработка |
| `go build -o lms cmd/lms/main.go` | Компиляция | Проверка сборки |
| `docker-compose up -d db` | Только БД | Тестирование |
| `docker-compose logs app` | Логи приложения | Отладка |
| `docker-compose down` | Остановить все | Очистка |

---

## 📈 Мониторинг и логирование

### Логирование с Logrus:
```go
logger := logrus.New()
logger.WithFields(logrus.Fields{
    "user_id": userID,
    "action":  "create_course",
}).Info("Course created successfully")
```

### Структурированные логи:
```json
{
  "level": "info",
  "msg": "Course created successfully",
  "user_id": 123,
  "action": "create_course",
  "time": "2023-12-01T10:00:00Z"
}
```

---

## 🔐 Безопасность

### 1. Валидация входных данных
```go
if course.Title == "" {
    return errors.New("title is required")
}
```

### 2. Авторизация на уровне маршрутов
```go
r.Use(middleware.AuthMiddleware)
r.Use(middleware.OnlyRoles(common.RoleAdmin))
```

### 3. Безопасные пароли
- Минимальная длина
- Хеширование в Keycloak

### 4. HTTPS в продакшне
- TLS сертификаты
- Безопасная передача токенов

---

## 📚 План изучения (Road Map)

```mermaid
graph TD
    subgraph "🌱 Beginner Level (1-2 месяца)"
        A1["📖 Go Tour & Syntax"]
        A2["🌐 HTTP & REST Basics"]
        A3["📄 JSON & APIs"]
        A4["🗄️ SQL Basics"]
        A5["🔧 Git & GitHub"]
    end
    
    subgraph "📈 Intermediate Level (2-3 месяца)"
        B1["🏗️ Clean Architecture"]
        B2["🎯 Design Patterns"]
        B3["🧪 Testing"]
        B4["🐳 Docker Basics"]
        B5["🔐 Authentication & JWT"]
    end
    
    subgraph "🚀 Advanced Level (3-6 месяцев)"
        C1["🏢 Microservices"]
        C2["☸️ Kubernetes"]
        C3["📊 Monitoring & Metrics"]
        C4["🔄 CI/CD Pipelines"]
        C5["☁️ Cloud Platforms"]
    end
    
    A1 --> A2 --> A3 --> A4 --> A5
    A5 --> B1 --> B2 --> B3 --> B4 --> B5
    B5 --> C1 --> C2 --> C3 --> C4 --> C5
    
    classDef beginner fill:#e8f5e8
    classDef intermediate fill:#fff3cd
    classDef advanced fill:#ffebcd
    
    class A1,A2,A3,A4,A5 beginner
    class B1,B2,B3,B4,B5 intermediate
    class C1,C2,C3,C4,C5 advanced
```

### 📋 Детальный чеклист

#### 🌱 Уровень 1: Основы (1-2 месяца)
- [ ] **Go Basics** 
  - [ ] Переменные, типы данных, функции
  - [ ] Структуры и методы
  - [ ] Интерфейсы и полиморфизм
  - [ ] Обработка ошибок
  - [ ] Goroutines и каналы
  
- [ ] **Web Development**
  - [ ] HTTP протокол
  - [ ] REST API принципы
  - [ ] JSON сериализация
  - [ ] HTTP клиенты

- [ ] **Database**
  - [ ] SQL basics (SELECT, INSERT, UPDATE, DELETE)
  - [ ] Связи между таблицами
  - [ ] Индексы и производительность

#### 📈 Уровень 2: Архитектура (2-3 месяца)
- [ ] **Clean Architecture**
  - [ ] Слои архитектуры
  - [ ] Dependency Injection
  - [ ] Repository Pattern
  - [ ] Service Layer

- [ ] **Testing**
  - [ ] Unit тесты
  - [ ] Integration тесты
  - [ ] Mocking
  - [ ] Test coverage

- [ ] **Tools**
  - [ ] Docker контейнеры
  - [ ] Docker Compose
  - [ ] Environment variables
  - [ ] Logging

#### 🚀 Уровень 3: Production (3-6 месяцев)
- [ ] **Scalability**
  - [ ] Микросервисы
  - [ ] Load balancing
  - [ ] Caching (Redis)
  - [ ] Message queues

- [ ] **DevOps**
  - [ ] CI/CD pipelines
  - [ ] Kubernetes
  - [ ] Monitoring (Prometheus)
  - [ ] Cloud deployment

---

## 🎯 Следующие шаги

### Карта развития проекта

```mermaid
mindmap
  root((LMS Project))
    Backend Improvements
      Testing
        Unit Tests
        Integration Tests
        E2E Tests
      Performance
        Caching (Redis)
        Database Optimization
        Load Balancing
      Features
        File Upload
        Search (Elasticsearch)
        Notifications
        Analytics
    
    Frontend
      Web UI
        React/Vue.js
        Admin Panel
        Student Dashboard
      Mobile
        React Native
        Flutter
        Progressive Web App
    
    Infrastructure
      Monitoring
        Prometheus
        Grafana
        Alerting
      DevOps
        CI/CD
        Kubernetes
        Cloud Deploy
      Security
        Rate Limiting
        Input Validation
        HTTPS/SSL
```

### 🔧 Технические улучшения

| Приоритет | Область | Улучшение | Описание |
|-----------|---------|-----------|----------|
| 🔴 Высокий | Testing | Unit тесты | Покрытие всех сервисов тестами |
| 🔴 Высокий | Security | Rate Limiting | Защита от брутфорса |
| 🟡 Средний | Performance | Redis Cache | Кеширование популярных запросов |
| 🟡 Средний | Features | File Upload | Загрузка материалов курсов |
| 🟢 Низкий | UX | Search | Поиск по курсам |
| 🟢 Низкий | Analytics | Metrics | Дашборды и аналитика |

### 🌐 Фронтенд разработка

```mermaid
graph LR
    subgraph "🖥️ Web Application"
        A["React.js App"]
        B["Admin Panel"]
        C["Student Portal"]
        D["Teacher Dashboard"]
    end
    
    subgraph "📱 Mobile Application"
        E["React Native"]
        F["Flutter"]
        G["PWA"]
    end
    
    subgraph "🔗 API Integration"
        H["Auth Module"]
        I["Course Module"]
        J["User Module"]
    end
    
    A --> H
    B --> H
    C --> H
    D --> H
    
    A --> I
    B --> I
    C --> I
    
    E --> H
    F --> H
    G --> H
    
    classDef web fill:#e3f2fd
    classDef mobile fill:#f3e5f5
    classDef api fill:#e8f5e8
    
    class A,B,C,D web
    class E,F,G mobile
    class H,I,J api
```

---

## 📖 Полезные ресурсы

### Документация:
- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Docker Documentation](https://docs.docker.com/)
- [Keycloak Documentation](https://www.keycloak.org/documentation)

### Книги:
- "Clean Architecture" - Robert Martin
- "Go in Action" - William Kennedy
- "Building Microservices" - Sam Newman

### Онлайн курсы:
- Go Tour (tour.golang.org)
- Docker Getting Started
- PostgreSQL Tutorial

---

## 🤔 Часто задаваемые вопросы

### Q: Зачем нужна чистая архитектура?
**A:** Она делает код более понятным, тестируемым и гибким. Можно легко заменить базу данных или добавить новый способ аутентификации.

### Q: Почему Go для backend разработки?
**A:** Go быстрый, простой, хорошо подходит для сетевого программирования и имеет отличную поддержку concurrency.

### Q: Зачем нужен Keycloak?
**A:** Он решает сложные задачи аутентификации и авторизации, предоставляя готовое решение с веб-интерфейсом для управления пользователями.

### Q: Что такое middleware?
**A:** Это функции, которые выполняются до или после основного обработчика запроса. Например, проверка авторизации или логирование.

---

**Удачи в изучении! 🚀**

*Этот документ - ваш путеводитель. Не бойтесь экспериментировать, задавать вопросы и изучать код. Программирование - это практика, так что пишите код каждый день!*