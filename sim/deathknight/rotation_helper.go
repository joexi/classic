package deathknight

import (
	"github.com/wowsims/wotlk/sim/core"
)

type RotationAction uint8

const (
	RotationAction_Skip RotationAction = iota
	RotationAction_CUSTOM1
	RotationAction_CUSTOM2
	RotationAction_CUSTOM3
	RotationAction_IT
	RotationAction_PS
	RotationAction_Obli
	RotationAction_BS
	RotationAction_BT
	RotationAction_UA
	RotationAction_RD
	RotationAction_Pesti
	RotationAction_FS
	RotationAction_HW
	RotationAction_ERW
	RotationAction_HB_Ghoul_RimeCheck
	RotationAction_PrioMode
	RotationAction_SS
	RotationAction_DND
	RotationAction_GF
	RotationAction_DC
	RotationAction_Garg
	RotationAction_AOTD
	RotationAction_BP
	RotationAction_FP
	RotationAction_UP
	RotationAction_RedoSequence
)

type Sequence struct {
	idx        int
	numActions int
	actions    []RotationAction
}

type DoRotationEvent func(sim *core.Simulation, target *core.Unit)
type RotationActionCustom func(sim *core.Simulation, target *core.Unit) bool

type RotationHelper struct {
	opener   *Sequence
	onOpener bool

	sequence *Sequence

	CastSuccessful     bool
	justCastPestilence bool

	DoRotationEvent DoRotationEvent
	ActionCustom1   RotationActionCustom
	ActionCustom2   RotationActionCustom
	ActionCustom3   RotationActionCustom
}

func TernaryRotationAction(condition bool, t RotationAction, f RotationAction) RotationAction {
	if condition {
		return t
	} else {
		return f
	}
}

func (r *RotationHelper) DefineOpener(actions []RotationAction) {
	r.opener = &Sequence{
		idx:        0,
		numActions: len(actions),
		actions:    actions,
	}
}

func (r *RotationHelper) PushSequence(actions []RotationAction) {
	if r.sequence == nil {
		r.sequence = &Sequence{
			idx:        0,
			numActions: len(actions),
			actions:    actions,
		}
	} else {
		panic("Tried to push sequence but sequence is currently Ongoing!")
	}
}

func (r *RotationHelper) RedoSequence(s *Sequence) {
	if r.sequence != nil {
		s.Reset()
		r.sequence = s
	} else {
		panic("Tried to redo sequence that wasn't ongoing!")
	}
}

func (r *RotationHelper) HasSequence() bool {
	return r.sequence != nil
}
