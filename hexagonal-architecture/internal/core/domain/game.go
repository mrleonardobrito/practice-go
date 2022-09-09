package domain

const (
	GAME_STATE_NEW  = "new"
	GAME_STATE_WON  = "won"
	GAME_STATE_LOSE = "lose"
)

type Game struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	State         string        `json:"state"`
	BoardSettings BoardSettings `json:"board_settings"`
	Board         Board         `json:"board"`
}

func NewGame(id string, name string, size uint, bombs uint) Game {
	return Game{
		ID:    id,
		Name:  name,
		State: "new",
		BoardSettings: BoardSettings{
			Size:  size,
			Bombs: bombs,
		},
	}
}
