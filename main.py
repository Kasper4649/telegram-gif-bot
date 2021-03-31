from telegram import Update
from telegram.ext import Updater, CommandHandler, CallbackContext
from config import TOKEN, WEBHOOK_HOST, PORT

def start_handler(update: Update, context: CallbackContext):
    update.message.reply_text("⛔")

def help_handler(update: Update, context: CallbackContext):
    update.message.reply_text("🚫")


updater = Updater(TOKEN, use_context=True)

updater.start_webhook(listen="0.0.0.0",
                      port=PORT,
                      url_path=TOKEN,
                      webhook_url=WEBHOOK_HOST)
updater.dispatcher.add_handler(CommandHandler("start", start_handler))
updater.dispatcher.add_handler(CommandHandler("help", help_handler))

updater.start_polling()
updater.idle()