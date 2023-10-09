package main

import (
	"fmt"
	"log"
	"tg_bot_resender/internal/app/bot"
	"tg_bot_resender/internal/app/config"
)

func main() {
	_ = bot.New()
	conf := config.New()
	err := conf.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf.ConnString)
	// Замените "YOUR_BOT_TOKEN" на ваш токен, полученный от BotFather.
	// bot, err := tgbotapi.NewBotAPI("6472929625:AAFhTW07sz1dAtstC3UCLnVFnYOKRJq9SVg")
	// if err != nil {
	// 	log.Panic(err)
	// }

	// // Устанавливаем отладочный режим (выводим обновления в консоль).
	// bot.Debug = true

	// log.Printf("Авторизован как %s", bot.Self.UserName)

	// // Устанавливаем обработчик команды /start.
	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates, err := bot.GetUpdatesChan(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for update := range updates {
	// 	// if update.Message == nil {
	// 	// 	continue
	// 	// }

	// 	if update.Message != nil {
	// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	// 		if update.Message.IsCommand() {
	// 			switch update.Message.Command() {
	// 			case "start":
	// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я ваш бот.")
	// 				keyboard := tgbotapi.NewInlineKeyboardMarkup(
	// 					tgbotapi.NewInlineKeyboardRow(
	// 						tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "1"),
	// 						//tgbotapi.NewInlineKeyboardButtonWebApp("url", tgbotapi.WebAppInfo{URL: "https://t.me/promocodes_from_bot"}),
	// 						tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "2"),
	// 					),
	// 					tgbotapi.NewInlineKeyboardRow(
	// 						tgbotapi.NewInlineKeyboardButtonData("Кнопка 3", "3"),
	// 					),
	// 					// tgbotapi.NewKeyboardButtonRow(
	// 					// 	tgbotapi.NewKeyboardButton("Кнопка 1"),
	// 					// 	tgbotapi.NewKeyboardButton("Кнопка 2"),
	// 					// ),
	// 					// tgbotapi.NewKeyboardButtonRow(
	// 					// 	tgbotapi.NewKeyboardButton("Кнопка 3"),
	// 					// ),
	// 				)
	// 				msg.ReplyMarkup = keyboard
	// 				bot.Send(msg)
	// 			case "products":
	// 				// Здесь вы можете добавить логику для получения списка продукции из базы данных или другого источника.
	// 				// После получения списка продукции, создайте сообщение и отправьте его ботом.
	// 				products := []string{"Продукт 1", "Продукт 2", "Продукт 3"} // Замените этот список на ваш список продукции.

	// 				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список продукции:\n")
	// 				for _, product := range products {
	// 					msg.Text += product + "\n"
	// 				}

	// 				bot.Send(msg)

	// 			}
	// 		}
	// 	} else if update.CallbackQuery != nil {
	// 		callback := update.CallbackQuery

	// 		// Реагируем на нажатие кнопки с данным callback-данными.
	// 		if callback.Data == "1" || callback.Data == "2" || callback.Data == "3" {
	// 			responseMsg := tgbotapi.NewMessage(callback.Message.Chat.ID, "Вы нажали на инлайн-кнопку!")
	// 			bot.Send(responseMsg)
	// 		}
	// 	}
	// }
}
