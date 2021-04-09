from telegram import Update
from telegram.ext import Updater, CommandHandler, MessageHandler, CallbackContext, Filters
from config import TOKEN, WEBHOOK_HOST, PORT
import traceback
from loguru import logger
import random
import string
from moviepy.editor import VideoFileClip
from firebase import Firebase

firebase = Firebase()

def start_handler(update: Update, context: CallbackContext):
    update.message.reply_text("⛔")

def help_handler(update: Update, context: CallbackContext):
    update.message.reply_text("🚫")

def text_handler(update: Update, context: CallbackContext):
    update.message.reply_text("不陪聊。")

def other_handler(update: Update, context: CallbackContext):
    update.message.reply_text("只接受 Animation。")

def animation_handler(update: Update, context: CallbackContext):
    animation = update.message.animation
    # generate random file name
    file_name = ''.join(random.sample(string.ascii_letters + string.digits, 8)) + ".mp4"
    new_file_name = file_name.replace(".mp4", ".gif")

    context.bot.get_file(animation.file_id).download(file_name)
    mp4_to_gif(file_name, new_file_name)

    try:
        public_url = firebase.upload(new_file_name)
        logger.info(f"[firebase] uploaded {new_file_name}")
        update.message.reply_text(public_url,
                                  reply_to_message_id=update.message.message_id)
    except Exception as e:
        logger.error(f"[firebase] failed to upload: {e}")
        update.message.reply_text("unknown error",
                                  reply_to_message_id=update.message.message_id)


def error_handler(update, context):
    logger.error(f"[Update] {update} caused error: {context.error}")
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
    dispatcher.add_handler(MessageHandler(Filters.animation, animation_handler))
    dispatcher.add_handler(MessageHandler(~(Filters.text | Filters.animation), other_handler))
    dispatcher.add_error_handler(error_handler)

    updater.idle()

def mp4_to_gif(file_name, new_file_name):
    clip = VideoFileClip(file_name)
    clip.write_gif(new_file_name)

if __name__ == '__main__':
    main()