package shared

type FactionsEnum string

const (
	AlexandrianMacedonian FactionsEnum = "AlexandrianMacedonian"
	AncientBritish                     = "AncientBritish"
)

type PhasesEnum string

const (
	PHASE_1_VICTORY_CHECK         PhasesEnum = "PHASE_1_VICTORY_CHECK"
	PHASE_2_ADVANCE                          = "PHASE_2_ADVANCE"
	PHASE_3_FLANK_ATTACKS                    = "PHASE_3_FLANK_ATTACKS"
	PHASE_4_PLAYER_ACTIONS                   = "PHASE_4_PLAYER_ACTIONS"
	PHASE_4_DEFENDER_ACTIONS                 = "PHASE_4_DEFENDER_ACTIONS"
	PHASE_5_DRAW_CARDS                       = "PHASE_5_DRAW_CARDS"
	GAME_OVER                                = "GAME_OVER"
	GAME_START_STATE                         = "GAME_START_STATE"
	SELECT_CARDS_STATE                       = "SELECT_CARDS_STATE"
	SELECT_COLUMN_TO_PLAY_STATE              = "SELECT_COLUMN_TO_PLAY_STATE"
	DISCARD_CARDS_FROM_HAND_STATE            = "DISCARD_CARDS_FROM_HAND_STATE"
)
