//じゃんけんの結果モデルの定義

package model

// じゃんけんの手をEnumで定義(rock, paper, scissors)
type Hand string

const (
	Rock     = Hand("rock")
	Paper    = Hand("paper")
	Scissors = Hand("scissors")
)

// じゃんけんの結果をEnumで定義(win, lose, draw)
type JankenResult string

const (
	Win  = JankenResult("win")
	Lose = JankenResult("lose")
	Draw = JankenResult("draw")
)

// じゃんけんの結果モデルの定義(Computerの手とUserの手と結果)
type JankenResponse struct {
	ComputerHand Hand         `json:"computer_hand"`
	UserHand     Hand         `json:"user_hand"`
	Result       JankenResult `json:"result"`
}
