package localization

import "github.com/bwmarrin/discordgo"

func init() {
	data[discordgo.Russian] = map[string]string{
		"#admin":                              "администратор",
		"#admin.allow-role":                   "разрешить-роль",
		"#look":                               "смотреть",
		"#move":                               "двигаться",
		"#review":                             "обзор",
		"#admin.allow-role.value":             "ценить",
		"#.subject":                           "цель",
		"#look.info":                          "информация",
		"#look.reviews":                       "обзоры",
		"#admin.allow-role.Description":       "Отображение результатов в виде ролей. По умолчанию: Ложь",
		"#look.Description":                   "увидеть что-то",
		"#move.Description":                   "Переместить обзор на этот канал",
		"#review.Description":                 "Обзор пользователя",
		"#admin.allow-role.value.Description": "установленное значение",
		"#.subject.Description":               "Выберите цель",
		"#look.info.Description":              "Просмотр информации о пользователе",
		"#look.reviews.Description":           "Посмотрите отзывы, полученные субъектом",
		"#allow-role.NeedPermissions":         "У бота нет разрешений – Управление ролями",
		"#allow-role.InProgress":              "работа в процессе",
		"#allow-role.proc.Title":              "Изменить параметры",
		"#allow-role.proc.Description":        "Разрешить роли",
		"#allow-role.proc.InProgress":         "Судебное разбирательство",
		"#allow-role.proc.Done":               "полный",
		"#allow-role.Keep":                    "Настройки не изменены",
		"#look.info.IsNone":                   "Нет отзывов",
		"#look.reviews.IsNone":                "Нет отзывов",
		"#look.reviews.menu.Title":            "Отзывы для %s",
		"#look.reviews.menu.Page":             "%d страница",
		"#move.IsNone":                        "Для объекта не оставлено ни одного отзыва.",
		"#move.Move":                          "двигайтесь сюда",
		"#review.SelfReview":                  "Вы не можете проверить себя",
		"#review.modal.Title":                 "Обзор %s",
		"#review.lable.Score":                 "счет",
		"#review.lable.Title":                 "заголовок",
		"#review.lable.Content":               "деталь",
		"$review.IsEdited":                    "Этот отзыв был отредактирован",
		"$review.NoAuthor":                    "Удаленные отзывы невозможно восстановить, поскольку автора здесь не существует.",
		"$review.DM":                          "Написан новый отзыв",
    }
}
