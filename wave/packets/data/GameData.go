package data

type MyItemInt uint32

const (
	Light MyItemInt = iota
	Speed
	Critical
	Achievement
	SexReq
	Buff
	Npc
	Special
	buffDuration
	buffInterval
	HonorReq
	HPRegen
	MPRegen
	Enhancement
	Defense
	MagicDefense
	Texture
	Material
	animated_paperdoll
	luck
	exp_lost
	Count
)

func GetItemInt(item int32, itemInt MyItemInt) int32 {
	if item == 0 || item > 1200 {
		return 0
	}
	if itemInt < Light || itemInt > exp_lost {
		return 0
	} else {
		return 1 // TODO modTypes.Item[itemNum].Int[(int)Value];
	}
}
