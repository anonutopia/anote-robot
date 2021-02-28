package main

import ui18n "github.com/unknwon/i18n"

func trUser(user *User, message string) string {
	if len(user.Language) > 0 {
		return ui18n.Tr(user.Language, message)
	} else {
		return ui18n.Tr(LANG, message)
	}
}

func tr(chatId int, message string) string {
	var lng string
	u := &User{TelegramID: chatId}
	db.First(u, u)
	if u.ID != 0 {
		if len(u.Language) > 0 {
			return ui18n.Tr(u.Language, message)
		} else {
			return ui18n.Tr(LANG, message)
		}
	} else {
		if chatId == T_ANON_BALKAN {
			lng = LANG_HR
		} else {
			lng = LANG
		}
	}
	return ui18n.Tr(lng, message)
}
