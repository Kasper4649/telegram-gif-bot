from telegram import Update
from telegram.ext import Updater, CommandHandler, MessageHandler, CallbackContext, Filters
from config import TOKEN, WEBHOOK_HOST, PORT
import traceback
from loguru import logger
from moviepy.editor import VideoFileClip

def start_handler(update: Update, context: CallbackContext):
    update.message.reply_text("⛔")

def help_handler(update: Update, context: CallbackContext):
    update.message.reply_text("🚫")

def text_handler(update: Update, context: CallbackContext):
    update.message.reply_text("不陪聊。")

def photo_handler(update: Update, context: CallbackContext):
    update.message.reply_text("只接受 Animation。")

def animation_handler(update: Update, context: CallbackContext):
    print(update.message)
    context.bot.send_animation(update.effective_message.chat_id,
                               update.message.animation,
                               reply_to_message_id=update.message.message_id)

def handle_error(update, context):
    logger.error(f"Update: {update} caused error: {context.error}")
    traceback.print_exc()

def main():
    updater = Updater(TOKEN, use_context=True)
    updater.start_webhook(listen="0.0.0.0",
                          port=PORT,
                          url_path=TOKEN,
                          webhook_url=WEBHOOK_HOST + TOKEN)

    dispatcher = updater.dispatcher
    dispatcher.add_handler(CommandHandler("start", start_handler))
    dispatcher.add_handler(CommandHandler("help", help_handler))
    dispatcher.add_handler(MessageHandler(Filters.text, text_handler))
    dispatcher.add_handler(MessageHandler(~(Filters.text | Filters.animation), photo_handler))
    dispatcher.add_handler(MessageHandler(Filters.animation, animation_handler))
    dispatcher.add_error_handler(handle_error)

    updater.idle()

# def mp4_to_gif():
#     clip = VideoFileClip("a.mp4")
#     clip.write_gif("output.gif")

if __name__ == '__main__':
    main()