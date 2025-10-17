from pymongo import MongoClient
from bson.objectid import ObjectId
from dotenv import load_dotenv
import os
import time

load_dotenv()

env = os.getenv('NODE_ENV', 'development')
mongo_host = os.getenv('MONGO_HOST', 'localhost')
mongo_port = os.getenv('MONGO_PORT', '27017')
mongo_user = os.getenv('MONGO_USER')
mongo_password = os.getenv('MONGO_PASSWORD')

mongo_credential = f"{mongo_user}:{mongo_password}@" if mongo_user and mongo_password else ""
uri = f"mongodb://{mongo_credential}{mongo_host}:{mongo_port}"

def main():
    start_time = int(time.time() * 1000)
    print(f"🔹 Proceso iniciado en {start_time} ms")

    client = MongoClient(uri)
    print("✅ Conectado a MongoDB")

    try:
        db = client['sample_db']
        collection = db['mycollection']

        # Crear un documento
        doc = { "name": "Martín Rivas", "age": 42, "source": "python" }
        insert_result = collection.insert_one(doc)
        print(f"📝 Documento insertado con _id: {insert_result.inserted_id}")

        # Buscar por _id
        found_by_id = collection.find_one({ "_id": ObjectId(str(insert_result.inserted_id)) })
        print("🔍 Documento encontrado por ID:", found_by_id)

        # Buscar todos los documentos
        all_docs = list(collection.find({}))
        print(f"📄 Documentos encontrados ({len(all_docs)}):", all_docs)

    finally:
        client.close()
        print("🔒 Conexión cerrada")

        end_time = int(time.time() * 1000)
        print(f"🔹 Proceso finalizado en {end_time} ms")
        print(f"⏱️ Duración total: {end_time - start_time} ms")

if __name__ == "__main__":
    main()