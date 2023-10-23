package entity

// Подписчики
type Receiver struct {
	id          int64 `db:"receiver_id"`
	name        string
	email       string
	birthDate   string
	deviceToken string
}

// Отправитель пушки
type Sender struct {
}

// Создатель шаблонов
type Admin struct {
}
