import firebase_admin
from firebase_admin import storage, credentials
from config import FIREBASE_CREDENTIALS, BUCKET_NAME
import json
from loguru import logger

class Firebase(object):

    def __init__(self):
        cert = json.loads(FIREBASE_CREDENTIALS, strict=False)
        cred = credentials.Certificate(cert)
        firebase_admin.initialize_app(cred, {
            'storageBucket': BUCKET_NAME
        })
        self._bucket = storage.bucket()

    def upload(self, file, file_name):
        blob = self._bucket.blob(file_name)
        try:
            blob.upload_from_file(file)
        except Exception as e:
            logger.error(f"failed to upload: {e}")

    def download(self, file_name):
        file = ""
        blob = self._bucket.blob(file_name)
        try:
            file = blob.download_as_bytes()
        except Exception as e:
            logger.error(f"failed to download: {e}")

        return file