
from fastapi import APIRouter, HTTPException
from app.models.mascota import Mascota
from app.repository.mascota_repo import crear_mascota

router = APIRouter()

@router.post("/mascotas")
async def registrar_mascota(mascota: Mascota):
    try:
        id = await crear_mascota(mascota)
        return {"mensaje": "Mascota registrada correctamente", "id": id}
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
