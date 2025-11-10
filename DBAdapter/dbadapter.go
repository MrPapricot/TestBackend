package DBAdapter

import (
	"backend/DBAdapter/Models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func InitAdapter(host string, port string, user string, password string, dbname string) Adapter {
	query := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(query), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	adapter := Adapter{
		db: db,
	}
	adapter.migrate()
	return adapter
}

func (Adapter *Adapter) Blank() {

}

func (Adapter *Adapter) Close() {
	db, err := Adapter.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func (adapter *Adapter) FillTestData() {
	bacalavr := Models.QualificationLevel{ID: 1, Name: "Бакалавриат"}
	result := adapter.db.Create(&bacalavr)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании типов отношений: %v", result.Error)
	}
	log.Printf("Успешно создан Бакалавр")

	test := Models.TrainingDirection{ID: 1, Name: "Test", Duration: 8, QualificationLevelID: 1}

	result = adapter.db.Create(&test)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании направлений подготовки: %v", result.Error)
	}
	log.Printf("Успешно создано Тестовое направление подготовки")

	specializations := []Models.Specialization{
		{
			ID:                  29116,
			Name:                "Промышленная разработка программного обеспечения",
			TrainingDirectionID: 1,
		},
		{
			ID:                  29113,
			Name:                "Информационные технологии и интеллектуальный анализ данных",
			TrainingDirectionID: 1,
		},
	}

	result = adapter.db.Create(&specializations)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании специализаций: %v", result.Error)
	}
	log.Printf("Создано %d специализаций\n", result.RowsAffected)

	// Создание типов отношений с кодами модулей
	relationTypes := []Models.RelationType{
		// Блок 1: Дисциплины (модули)
		{Code: "Б1.БМ1", Name: "Базовая часть. Модуль базовой инженерной подготовки"},
		{Code: "Б1.БМ2", Name: "Базовая часть. Модуль направления подготовки"},
		{Code: "Б1.ВМ1", Name: "Вариативная часть. Модуль дополнительной специализации"},
		{Code: "Б1.ВМ2", Name: "Вариативная часть. Модуль специализации"},
		{Code: "Б1.ВМ3", Name: "Вариативная часть. Элективные дисциплины по физической культуре и спорту"},

		// Блок 2: Практики
		{Code: "Б2.В.1", Name: "Учебная практика"},
		{Code: "Б2.В.2", Name: "Производственная практика"},

		// Блок 3: ГИА
		{Code: "Б3.Б", Name: "Государственная итоговая аттестация"},

		// Факультативы
		{Code: "ФД.В", Name: "Факультативные дисциплины"},
	}

	// Добавление типов отношений в БД
	result = adapter.db.Create(&relationTypes)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании RelationTypes: %v", result.Error)
	}
	log.Printf("Успешно создано %d RelationTypes\n", result.RowsAffected)

	subjects := []Models.Subject{
		{Name: "История России"},
		{Name: "Основы российской государственности"},
		{Name: "Введение в проектную деятельность"},
		{Name: "Иностранный язык (английский)"},
		{Name: "Информатика 1.2"},
		{Name: "Математика 1.3"},
		{Name: "Введение в информационные технологии"},
		{Name: "Математическая логика и теория алгоритмов"},
		{Name: "Элективные дисциплины по физической культуре и спорту"},
		{Name: "Учебный проект"},
		{Name: "Начертательная геометрия и инженерная графика 1.4"},
		{Name: "Математика 2.3"},
		{Name: "Физика 1.2"},
		{Name: "Проектирование человеко-машинного интерфейса"},
		{Name: "Основы программирования на Python"},
		{Name: "Программирование на Java"},
		{Name: "Учебная практика по развитию цифровых компетенций"},
		{Name: "Физическая культура и спорт"},
		{Name: "Математика 3.3"},
		{Name: "Физика 2.2"},
		{Name: "Дискретная математика"},
		{Name: "Программирование мобильных приложений"},
		{Name: "Факультативные дисциплины по выбору студента"},
		{Name: "Математика 4.1"},
		{Name: "Физика 3.2"},
		{Name: "Безопасность жизнедеятельности"},
		{Name: "Комплексный экзамен по модулю базовой инженерной подготовки"},
		{Name: "Электротехника 1.3"},
		{Name: "Операционные системы и сети"},
		{Name: "Базы данных"},
		{Name: "Алгоритмы и структуры данных"},
		{Name: "Технологическая (проектно-технологическая) практика"},
		{Name: "Экономическая культура и финансовая грамотность"},
		{Name: "Основы права"},
		{Name: "Учебно-исследовательская работа студентов"},
		{Name: "Web-программирование"},
		{Name: "Архитектура вычислительных систем"},
		{Name: "Проектирование и архитектура программных систем"},
		{Name: "Дисциплины дополнительной специализации"},
		{Name: "Разработка и анализ требований"},
		{Name: "Конструирование программного обеспечения"},
		{Name: "Философия"},
		{Name: "Экономика и организация производства"},
		{Name: "Технология командной разработки программного обеспечения"},
		{Name: "Проектирование Интернет-приложений"},
		{Name: "Компьютерная графика"},
		{Name: "Компьютерное моделирование"},
		{Name: "Интеллектуальный анализ данных (дата-майнинг)"},
		{Name: "Методы и системы обработки данных"},
		{Name: "Междисциплинарный проект"},
		{Name: "Тестирование программного обеспечения"},
		{Name: "Технологии облачных вычислений в бизнесе"},
		{Name: "Управление программными проектами"},
		{Name: "Программная инженерия"},
		{Name: "Информационная безопасность и защита информации"},
		{Name: "Параллельное и распределённое программирование"},
		{Name: "Проектирование информационных систем"},
		{Name: "Современные концепции организации баз данных"},
		{Name: "Искусственный интеллект и логическое программирование"},
		{Name: "Преддипломная практика"},
		{Name: "Выпускная квалификационная работа бакалавра"},
		{Name: "Эргономика цифровой среды"},
		{Name: "Численное решение инженерных задач"},
		{Name: "Проектирование пользовательских интерфейсов"},
		{Name: "Информационные технологии в цифровой среде"},
		{Name: "Разработка Web-приложений"},
		{Name: "Операционные системы"},
		{Name: "Архитектура информационных систем"},
		{Name: "Управление данными"},
		{Name: "Интеллектуальные системы и технологии"},
		{Name: "Сбор, обработка и анализ данных"},
		{Name: "Английский язык в сфере информационных технологий"},
		{Name: "Системная инженерия"},
		{Name: "Методы и средства разработки информационных систем"},
		{Name: "Компьютерная геометрия и графика"},
		{Name: "Инфокоммуникационные системы и сети"},
		{Name: "Машинное обучение"},
		{Name: "Статистический анализ экспериментальных данных"},
		{Name: "Методы обработки изображений и распознавания образов"},
		{Name: "Введение в методы распознавания образов"},
		{Name: "Методы оптимизации"},
		{Name: "Инструментальные средства информационных систем"},
		{Name: "Безопасность информационных систем"},
		{Name: "Параллельные и распределенные вычисления в машинном обучении"},
		{Name: "Методы анализа текстов"},
		{Name: "Методика работы с текстовыми данными"},
	}

	result = adapter.db.Create(&subjects)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании предметов: %v", result.Error)
	}
	fmt.Printf("Создано %d предметов\n", result.RowsAffected)

	relations := []Models.SubjectSpecializationRelation{
		// Специализация 29116 - Промышленная разработка ПО
		{SubjectID: 1, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 2},
		{SubjectID: 2, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 2},
		{SubjectID: 3, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 1},
		{SubjectID: 4, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 3},
		{SubjectID: 5, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 3},
		{SubjectID: 6, SpecializationID: 29116, RelationTypeID: 1, Semester: 1, Credits: 8},
		{SubjectID: 7, SpecializationID: 29116, RelationTypeID: 2, Semester: 1, Credits: 3},
		{SubjectID: 8, SpecializationID: 29116, RelationTypeID: 2, Semester: 1, Credits: 5},
		{SubjectID: 9, SpecializationID: 29116, RelationTypeID: 5, Semester: 1, Credits: 0},

		{SubjectID: 1, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 2},
		{SubjectID: 4, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 3},
		{SubjectID: 10, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 1},
		{SubjectID: 11, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 3},
		{SubjectID: 12, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 6},
		{SubjectID: 13, SpecializationID: 29116, RelationTypeID: 1, Semester: 2, Credits: 4},
		{SubjectID: 14, SpecializationID: 29116, RelationTypeID: 2, Semester: 2, Credits: 2},
		{SubjectID: 15, SpecializationID: 29116, RelationTypeID: 2, Semester: 2, Credits: 2},
		{SubjectID: 16, SpecializationID: 29116, RelationTypeID: 2, Semester: 2, Credits: 4},
		{SubjectID: 9, SpecializationID: 29116, RelationTypeID: 5, Semester: 2, Credits: 0},
		{SubjectID: 17, SpecializationID: 29116, RelationTypeID: 6, Semester: 2, Credits: 6},

		{SubjectID: 10, SpecializationID: 29116, RelationTypeID: 1, Semester: 3, Credits: 1},
		{SubjectID: 18, SpecializationID: 29116, RelationTypeID: 1, Semester: 3, Credits: 2},
		{SubjectID: 19, SpecializationID: 29116, RelationTypeID: 1, Semester: 3, Credits: 6},
		{SubjectID: 20, SpecializationID: 29116, RelationTypeID: 1, Semester: 3, Credits: 4},
		{SubjectID: 16, SpecializationID: 29116, RelationTypeID: 2, Semester: 3, Credits: 4},
		{SubjectID: 21, SpecializationID: 29116, RelationTypeID: 2, Semester: 3, Credits: 6},
		{SubjectID: 22, SpecializationID: 29116, RelationTypeID: 2, Semester: 3, Credits: 4},
		{SubjectID: 9, SpecializationID: 29116, RelationTypeID: 5, Semester: 3, Credits: 0},
		{SubjectID: 23, SpecializationID: 29116, RelationTypeID: 9, Semester: 3, Credits: 2},

		{SubjectID: 10, SpecializationID: 29116, RelationTypeID: 1, Semester: 4, Credits: 1},
		{SubjectID: 24, SpecializationID: 29116, RelationTypeID: 1, Semester: 4, Credits: 3},
		{SubjectID: 25, SpecializationID: 29116, RelationTypeID: 1, Semester: 4, Credits: 4},
		{SubjectID: 26, SpecializationID: 29116, RelationTypeID: 1, Semester: 4, Credits: 3},
		{SubjectID: 27, SpecializationID: 29116, RelationTypeID: 1, Semester: 4, Credits: 0},
		{SubjectID: 28, SpecializationID: 29116, RelationTypeID: 2, Semester: 4, Credits: 3},
		{SubjectID: 29, SpecializationID: 29116, RelationTypeID: 2, Semester: 4, Credits: 3},
		{SubjectID: 30, SpecializationID: 29116, RelationTypeID: 2, Semester: 4, Credits: 6},
		{SubjectID: 31, SpecializationID: 29116, RelationTypeID: 2, Semester: 4, Credits: 4},
		{SubjectID: 9, SpecializationID: 29116, RelationTypeID: 5, Semester: 4, Credits: 0},
		{SubjectID: 32, SpecializationID: 29116, RelationTypeID: 6, Semester: 4, Credits: 6},
		{SubjectID: 23, SpecializationID: 29116, RelationTypeID: 9, Semester: 4, Credits: 2},

		{SubjectID: 33, SpecializationID: 29116, RelationTypeID: 1, Semester: 5, Credits: 2},
		{SubjectID: 34, SpecializationID: 29116, RelationTypeID: 1, Semester: 5, Credits: 2},
		{SubjectID: 35, SpecializationID: 29116, RelationTypeID: 2, Semester: 5, Credits: 1},
		{SubjectID: 36, SpecializationID: 29116, RelationTypeID: 2, Semester: 5, Credits: 6},
		{SubjectID: 37, SpecializationID: 29116, RelationTypeID: 2, Semester: 5, Credits: 5},
		{SubjectID: 38, SpecializationID: 29116, RelationTypeID: 2, Semester: 5, Credits: 4},
		{SubjectID: 39, SpecializationID: 29116, RelationTypeID: 3, Semester: 5, Credits: 3},
		{SubjectID: 40, SpecializationID: 29116, RelationTypeID: 4, Semester: 5, Credits: 4},
		{SubjectID: 23, SpecializationID: 29116, RelationTypeID: 9, Semester: 5, Credits: 2},

		{SubjectID: 42, SpecializationID: 29116, RelationTypeID: 1, Semester: 6, Credits: 2},
		{SubjectID: 43, SpecializationID: 29116, RelationTypeID: 1, Semester: 6, Credits: 3},
		{SubjectID: 35, SpecializationID: 29116, RelationTypeID: 2, Semester: 6, Credits: 1},
		{SubjectID: 44, SpecializationID: 29116, RelationTypeID: 2, Semester: 6, Credits: 5},
		{SubjectID: 45, SpecializationID: 29116, RelationTypeID: 2, Semester: 6, Credits: 4},
		{SubjectID: 46, SpecializationID: 29116, RelationTypeID: 2, Semester: 6, Credits: 4},
		{SubjectID: 39, SpecializationID: 29116, RelationTypeID: 3, Semester: 6, Credits: 3},
		{SubjectID: 41, SpecializationID: 29116, RelationTypeID: 4, Semester: 6, Credits: 5},
		{SubjectID: 32, SpecializationID: 29116, RelationTypeID: 7, Semester: 6, Credits: 6},
		{SubjectID: 23, SpecializationID: 29116, RelationTypeID: 9, Semester: 6, Credits: 2},

		{SubjectID: 35, SpecializationID: 29116, RelationTypeID: 2, Semester: 7, Credits: 1},
		{SubjectID: 55, SpecializationID: 29116, RelationTypeID: 2, Semester: 7, Credits: 3},
		{SubjectID: 39, SpecializationID: 29116, RelationTypeID: 3, Semester: 7, Credits: 3},
		{SubjectID: 47, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 48, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 49, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 50, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 51, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 52, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 4},
		{SubjectID: 53, SpecializationID: 29116, RelationTypeID: 4, Semester: 7, Credits: 4},
		{SubjectID: 23, SpecializationID: 29116, RelationTypeID: 9, Semester: 7, Credits: 2},

		{SubjectID: 35, SpecializationID: 29116, RelationTypeID: 2, Semester: 8, Credits: 1},
		{SubjectID: 56, SpecializationID: 29116, RelationTypeID: 2, Semester: 8, Credits: 5},
		{SubjectID: 50, SpecializationID: 29116, RelationTypeID: 4, Semester: 8, Credits: 3},
		{SubjectID: 54, SpecializationID: 29116, RelationTypeID: 4, Semester: 8, Credits: 3},
		{SubjectID: 57, SpecializationID: 29116, RelationTypeID: 4, Semester: 8, Credits: 3},
		{SubjectID: 58, SpecializationID: 29116, RelationTypeID: 4, Semester: 8, Credits: 3},
		{SubjectID: 59, SpecializationID: 29116, RelationTypeID: 4, Semester: 8, Credits: 3},
		{SubjectID: 60, SpecializationID: 29116, RelationTypeID: 7, Semester: 8, Credits: 6},
		{SubjectID: 61, SpecializationID: 29116, RelationTypeID: 8, Semester: 8, Credits: 9},

		// Специализация 29113 - Информационные технологии и интеллектуальный анализ данных
		{SubjectID: 1, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 2},
		{SubjectID: 2, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 2},
		{SubjectID: 3, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 1},
		{SubjectID: 4, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 3},
		{SubjectID: 5, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 3},
		{SubjectID: 6, SpecializationID: 29113, RelationTypeID: 1, Semester: 1, Credits: 8},
		{SubjectID: 15, SpecializationID: 29113, RelationTypeID: 2, Semester: 1, Credits: 4},
		{SubjectID: 7, SpecializationID: 29113, RelationTypeID: 2, Semester: 1, Credits: 4},
		{SubjectID: 9, SpecializationID: 29113, RelationTypeID: 5, Semester: 1, Credits: 0},

		// Продолжение для 29113...
		{SubjectID: 1, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 2},
		{SubjectID: 4, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 3},
		{SubjectID: 10, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 1},
		{SubjectID: 11, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 3},
		{SubjectID: 12, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 6},
		{SubjectID: 13, SpecializationID: 29113, RelationTypeID: 1, Semester: 2, Credits: 4},
		{SubjectID: 8, SpecializationID: 29113, RelationTypeID: 2, Semester: 2, Credits: 3},
		{SubjectID: 16, SpecializationID: 29113, RelationTypeID: 2, Semester: 2, Credits: 4},
		{SubjectID: 62, SpecializationID: 29113, RelationTypeID: 2, Semester: 2, Credits: 1},
		{SubjectID: 9, SpecializationID: 29113, RelationTypeID: 5, Semester: 2, Credits: 0},
		{SubjectID: 17, SpecializationID: 29113, RelationTypeID: 6, Semester: 2, Credits: 6},

		// Специализация 29113 - семестр 3
		{SubjectID: 10, SpecializationID: 29113, RelationTypeID: 1, Semester: 3, Credits: 1},
		{SubjectID: 18, SpecializationID: 29113, RelationTypeID: 1, Semester: 3, Credits: 2},
		{SubjectID: 19, SpecializationID: 29113, RelationTypeID: 1, Semester: 3, Credits: 6},
		{SubjectID: 20, SpecializationID: 29113, RelationTypeID: 1, Semester: 3, Credits: 4},
		{SubjectID: 63, SpecializationID: 29113, RelationTypeID: 2, Semester: 3, Credits: 4},
		{SubjectID: 30, SpecializationID: 29113, RelationTypeID: 2, Semester: 3, Credits: 4},
		{SubjectID: 64, SpecializationID: 29113, RelationTypeID: 2, Semester: 3, Credits: 3},
		{SubjectID: 65, SpecializationID: 29113, RelationTypeID: 2, Semester: 3, Credits: 3},
		{SubjectID: 9, SpecializationID: 29113, RelationTypeID: 5, Semester: 3, Credits: 0},
		{SubjectID: 23, SpecializationID: 29113, RelationTypeID: 9, Semester: 3, Credits: 2},

		// Специализация 29113 - семестр 4
		{SubjectID: 10, SpecializationID: 29113, RelationTypeID: 1, Semester: 4, Credits: 1},
		{SubjectID: 25, SpecializationID: 29113, RelationTypeID: 1, Semester: 4, Credits: 4},
		{SubjectID: 26, SpecializationID: 29113, RelationTypeID: 1, Semester: 4, Credits: 3},
		{SubjectID: 27, SpecializationID: 29113, RelationTypeID: 1, Semester: 4, Credits: 0},
		{SubjectID: 28, SpecializationID: 29113, RelationTypeID: 2, Semester: 4, Credits: 3},
		{SubjectID: 24, SpecializationID: 29113, RelationTypeID: 2, Semester: 4, Credits: 3},
		{SubjectID: 21, SpecializationID: 29113, RelationTypeID: 2, Semester: 4, Credits: 4},
		{SubjectID: 66, SpecializationID: 29113, RelationTypeID: 2, Semester: 4, Credits: 5},
		{SubjectID: 67, SpecializationID: 29113, RelationTypeID: 2, Semester: 4, Credits: 4},
		{SubjectID: 9, SpecializationID: 29113, RelationTypeID: 5, Semester: 4, Credits: 0},
		{SubjectID: 32, SpecializationID: 29113, RelationTypeID: 6, Semester: 4, Credits: 6},
		{SubjectID: 23, SpecializationID: 29113, RelationTypeID: 9, Semester: 4, Credits: 2},

		// Специализация 29113 - семестр 5
		{SubjectID: 33, SpecializationID: 29113, RelationTypeID: 1, Semester: 5, Credits: 2},
		{SubjectID: 34, SpecializationID: 29113, RelationTypeID: 1, Semester: 5, Credits: 2},
		{SubjectID: 68, SpecializationID: 29113, RelationTypeID: 2, Semester: 5, Credits: 6},
		{SubjectID: 69, SpecializationID: 29113, RelationTypeID: 2, Semester: 5, Credits: 4},
		{SubjectID: 70, SpecializationID: 29113, RelationTypeID: 2, Semester: 5, Credits: 4},
		{SubjectID: 39, SpecializationID: 29113, RelationTypeID: 3, Semester: 5, Credits: 3},
		{SubjectID: 71, SpecializationID: 29113, RelationTypeID: 4, Semester: 5, Credits: 5},
		{SubjectID: 35, SpecializationID: 29113, RelationTypeID: 4, Semester: 5, Credits: 1},
		{SubjectID: 23, SpecializationID: 29113, RelationTypeID: 9, Semester: 5, Credits: 2},

		// Специализация 29113 - семестр 6
		{SubjectID: 42, SpecializationID: 29113, RelationTypeID: 1, Semester: 6, Credits: 2},
		{SubjectID: 43, SpecializationID: 29113, RelationTypeID: 1, Semester: 6, Credits: 3},
		{SubjectID: 72, SpecializationID: 29113, RelationTypeID: 2, Semester: 6, Credits: 2},
		{SubjectID: 44, SpecializationID: 29113, RelationTypeID: 2, Semester: 6, Credits: 3},
		{SubjectID: 73, SpecializationID: 29113, RelationTypeID: 2, Semester: 6, Credits: 3},
		{SubjectID: 74, SpecializationID: 29113, RelationTypeID: 2, Semester: 6, Credits: 6},
		{SubjectID: 39, SpecializationID: 29113, RelationTypeID: 3, Semester: 6, Credits: 3},
		{SubjectID: 75, SpecializationID: 29113, RelationTypeID: 4, Semester: 6, Credits: 4},
		{SubjectID: 35, SpecializationID: 29113, RelationTypeID: 4, Semester: 6, Credits: 1},
		{SubjectID: 32, SpecializationID: 29113, RelationTypeID: 7, Semester: 6, Credits: 6},
		{SubjectID: 23, SpecializationID: 29113, RelationTypeID: 9, Semester: 6, Credits: 2},

		// Специализация 29113 - семестр 7
		{SubjectID: 72, SpecializationID: 29113, RelationTypeID: 2, Semester: 7, Credits: 2},
		{SubjectID: 76, SpecializationID: 29113, RelationTypeID: 2, Semester: 7, Credits: 4},
		{SubjectID: 77, SpecializationID: 29113, RelationTypeID: 2, Semester: 7, Credits: 5},
		{SubjectID: 39, SpecializationID: 29113, RelationTypeID: 3, Semester: 7, Credits: 3},
		{SubjectID: 78, SpecializationID: 29113, RelationTypeID: 4, Semester: 7, Credits: 3},
		{SubjectID: 50, SpecializationID: 29113, RelationTypeID: 4, Semester: 7, Credits: 4},
		{SubjectID: 79, SpecializationID: 29113, RelationTypeID: 4, Semester: 7, Credits: 5},
		{SubjectID: 80, SpecializationID: 29113, RelationTypeID: 4, Semester: 7, Credits: 5},
		{SubjectID: 35, SpecializationID: 29113, RelationTypeID: 4, Semester: 7, Credits: 1},
		{SubjectID: 23, SpecializationID: 29113, RelationTypeID: 9, Semester: 7, Credits: 2},

		// Специализация 29113 - семестр 8
		{SubjectID: 72, SpecializationID: 29113, RelationTypeID: 2, Semester: 8, Credits: 2},
		{SubjectID: 81, SpecializationID: 29113, RelationTypeID: 2, Semester: 8, Credits: 4},
		{SubjectID: 82, SpecializationID: 29113, RelationTypeID: 2, Semester: 8, Credits: 4},
		{SubjectID: 83, SpecializationID: 29113, RelationTypeID: 2, Semester: 8, Credits: 3},
		{SubjectID: 84, SpecializationID: 29113, RelationTypeID: 4, Semester: 8, Credits: 4},
		{SubjectID: 85, SpecializationID: 29113, RelationTypeID: 4, Semester: 8, Credits: 4},
		{SubjectID: 86, SpecializationID: 29113, RelationTypeID: 4, Semester: 8, Credits: 4},
		{SubjectID: 35, SpecializationID: 29113, RelationTypeID: 4, Semester: 8, Credits: 1},
		{SubjectID: 60, SpecializationID: 29113, RelationTypeID: 7, Semester: 8, Credits: 6},
		{SubjectID: 61, SpecializationID: 29113, RelationTypeID: 8, Semester: 8, Credits: 9},
	}

	result = adapter.db.Create(&relations)
	if result.Error != nil {
		log.Fatalf("Ошибка при создании связей предметов со специализациями: %v", result.Error)
	}
	log.Printf("Успешно создано %d связей предметов со специализациями\n", result.RowsAffected)
}

func (adapter *Adapter) migrate() {
	var err error
	err = Models.MigrateBaseTpuUsers(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating BaseTpuUsers\nError:\n%+v", err)
	}
	err = Models.MigrateEducationTypes(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating EducationTypes\nError:\n%+v", err)
	}
	err = Models.MigrateEngineeringSchools(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating EngineeringSchools\nError:\n%+v", err)
	}
	err = Models.MigrateDepartments(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Departments\nError:\n%+v", err)
	}
	err = Models.MigrateTutors(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Tutors\nError:\n%+v", err)
	}
	err = Models.MigrateRoles(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Roles\nError:\n%+v", err)
	}
	err = Models.MigratePlanners(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Planners\nError:\n%+v", err)
	}
	err = Models.MigrateQualificationLevels(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating QualificationLevels\nError:\n%+v", err)
	}
	err = Models.MigrateGroups(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Groups\nError:\n%+v", err)
	}
	err = Models.MigrateSpecializations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Specializations\nError:\n%+v", err)
	}
	err = Models.MigrateStudents(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Students\nError:\n%+v", err)
	}
	err = Models.MigrateSubjectSpecializationRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating SubjectSpecializationRelations\nError:\n%+v", err)
	}
	err = Models.MigrateSubjectTutorRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating SubjectTutorRelations\nError:\n%+v", err)
	}
	err = Models.MigrateSubjects(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Subjects\nError:\n%+v", err)
	}
	err = Models.MigrateTrainingDirections(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating TrainingDirections\nError:\n%+v", err)
	}
	err = Models.MigrateTutorDepartmentRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating TutorDepartmentRelations\nError:\n%+v", err)
	}
	err = Models.MigrateRelationTypes(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating RealtionTypes\nError:\n%+v", err)
	}

	log.Println("Successful Migration")
}
