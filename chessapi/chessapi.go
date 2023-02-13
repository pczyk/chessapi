package chessapi

type ChessAPI struct {
	baseUrl string
}

func NewChessAPI() ChessAPI {
	return ChessAPI{baseUrl: "https://api.chess.com/pub"}
}
