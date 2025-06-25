
from pydantic import BaseModel, EmailStr
from typing import Optional

class Mascota(BaseModel):
    nombre: str
    especie: str
    raza: Optional[str]
    edad: Optional[int]
    peso: Optional[float]
    duenio_email: EmailStr
