package q4

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
		destinos := make(map[string]bool)

	for _, caminho := range caminhos {
		destinos[caminho[0]] = true
		delete(destinos, caminho[1])
	}

	if len(destinos) != 1 {
		return "", nil // ou retornar um erro, dependendo dos requisitos
	}

	var destino string
	for cidade := range destinos {
		destino = cidade
	}

	return destino, nil
}
