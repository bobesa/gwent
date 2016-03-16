package gwent

import "testing"

func TestTemplateLeader(t *testing.T) {
	tmpl := CardTemplate{Type: "leader"}

	card := tmpl.Make()

	if card.Type() != TypeLeader {
		t.Error("Card should be a leader card")
	}
}
