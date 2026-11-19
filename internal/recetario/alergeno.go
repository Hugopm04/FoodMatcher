package recetario

type Alergeno int

const (
	AlergenoGluten Alergeno = iota + 1
	AlergenoCrustaceos
	AlergenoHuevos
	AlergenoPescado
	AlergenoCacahuetes
	AlergenoSoja
	AlergenoLeche
	AlergenoFrutosDeCascara
	AlergenoApio
	AlergenoMostaza
	AlergenoGranosDeSesamo
	AlergenoDioxidoDeAzufreYSulfitos
	AlergenoAltramuces
	AlergenoMoluscos
)
