from enum import Enum


class Alergeno(Enum):
    """Alérgenos de la taxonomía de Open Food Facts"""

    GLUTEN = "gluten"
    CRUSTACEOS = "crustáceos"
    HUEVOS = "huevos"
    PESCADO = "pescado"
    CACAHUETES = "cacahuetes"
    SOJA = "soja"
    LECHE = "leche"
    FRUTOS_DE_CASCARA = "frutos de cáscara"
    APIO = "apio"
    MOSTAZA = "mostaza"
    GRANOS_DE_SESAMO = "granos de sésamo"
    DIOXIDO_DE_AZUFRE_Y_SULFITOS = "dióxido de azufre y sulfitos"
    ALTRAMUCES = "altramuces"
    MOLUSCOS = "moluscos"
