package matematica

import "testing"

const erroPadrao = "Valor esperado %v mas o resultado encontrato foi %v."

func TesteMedia(t *testing.T) {
	t.Parallel()
	valorEperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEperado {
		t.Errorf(erroPadrao, valorEperado, valor)
	}
}
