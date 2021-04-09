import firebase_admin
from firebase_admin import storage, credentials
from config import FIREBASE_CREDENTIALS, BUCKET_NAME
import json

class Firebase(object):

    def __init__(self):
        cert = json.loads(FIREBASE_CREDENTIALS, strict=False)
        cred = credentials.Certificate(cert)
        firebase_admin.initialize_app(cred, {
            'storageBucket': BUCKET_NAME
        })
        self._bucket = storage.bucket()

    def upload(self, file_name):
        blob = self._bucket.blob(file_name)
        blob.upload_from_filename(file_name)
        blob.make_public()
        return blob.public_url

    def download(self, file_name):
        blob = self._bucket.blob(file_name)
        file = blob.download_as_bytes()
        return file