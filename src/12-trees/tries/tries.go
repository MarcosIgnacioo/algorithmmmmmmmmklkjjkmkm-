package tries

type TriesNode struct {
	IsWord     bool
	Characters *[26]rune
}

// La manera de crear caracteres va a ser con la formula del offset que enseño prime en el video
