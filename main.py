from telegram.ext import Updater, CommandHandler
from config import TOKEN, WEBHOOK_HOST, PORT


def start(update, context):
    update.message.reply_text("⛔")

def help(update, context):
    update.message.reply_text("🚫")


updater = Updater(TOKEN, use_context=True)
# add handlers
updater.start_webhook(listen="0.0.0.0",
                      port=PORT,
                      url_path=TOKEN,
                      webhook_url=WEBHOOK_HOST)
updater.dispatcher.add_handler(CommandHandler("start", start))
updater.dispatcher.add_handler(CommandHandler("help", help))
updater.idle()