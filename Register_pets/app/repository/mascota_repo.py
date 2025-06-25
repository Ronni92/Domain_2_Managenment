
from motor.motor_asyncio import AsyncIOMotorClient
from app.models.mascota import Mascota
import os
from dotenv import load_dotenv

load_dotenv()

MONGO_URI = os.getenv("MONGO_URI")
DB_NAME = os.getenv("DB_NAME")

client = AsyncIOMotorClient(MONGO_URI)
db = client[DB_NAME]
mascotas_collection = db["mascotas"]

async def crear_mascota(mascota: Mascota):
    mascota_dict = mascota.dict()
    result = await mascotas_collection.insert_one(mascota_dict)
    return str(result.inserted_id)
