package lib

type NPCType uint

const (
	NPCTypePhysical NPCType = iota
	NPCTypeMagical
	NPCTypeRanged
)

type ItemPropertyType uint

const (
	ItemLight ItemPropertyType = iota
	ItemSpeed
	ItemCritical
	ItemAchievement
	ItemSexReq
	ItemBuff
	ItemNpc
	ItemSpecial
	ItembuffDuration
	ItembuffInterval
	ItemHonorReq
	ItemHPRegen
	ItemMPRegen
	ItemEnhancement
	ItemDefense
	ItemMagicDefense
	ItemTexture
	ItemMaterial
	Itemanimated_paperdoll
	Itemluck
	Itemexp_lost
)

type ItemMaterialType uint

const (
	MaterialOre ItemMaterialType = iota + 1
	MaterialFish
	MaterialWood
	MaterialStone
)

type ProfessionType uint

const (
	ProfessionDigging ProfessionType = iota
	ProfessionMining
	ProfessionTreeCutting
	ProfessionFishing
)

type StatsType uint

const (
	StatsStrength StatsType = iota + 1
	StatsEndurance
	StatsSpirit
	Statswillpower
	StatsIntelligence
	StatsAgility
)

type VitalType uint

const (
	HP VitalType = iota + 1
	MP
	SP
)

type EquipmentType uint

const (
	EquipmentMask EquipmentType = iota + 1
	EquipmentCape
	EquipmentHelmet
	EquipmentShoulder
	EquipmentNecklace
	EquipmentWeapon
	EquipmentArmor
	EquipmentShield
	EquipmentEarring
	EquipmentRing
	EquipmentLegs
	EquipmentGlove
	EquipmentStringuer
	EquipmentBelt
	EquipmentShoes
	EquipmentMount
)

type CashEquipment uint

const (
	Cape CashEquipment = iota + 1
	Hat
	Neck
	Weapon
	Top
	Shield
	Ring
	Bottom
	Gloves
	Mount
	Shoes
	Face
)

type DefenseType uint

const (
	DefensePhisic DefenseType = iota + 1
	DefenseMagic
	DefenseRanged
)

type PlayerClassType uint

const (
	PlayerClassKnight PlayerClassType = iota + 1
	PlayerClassStalker
	PlayerClassMage
	PlayerClassHunter
	PlayerClassNecromancer
	PlayerClassDruid
)

type SpellType uint

const (
	SpellTypeDamageHP SpellType = iota
	SpellTypeDamageMP
	SpellTypeHealHP
	SpellTypeHealMP
	SpellTypeWarp
	SpellTypeAttack
	SpellTypeShield
	SpellTypeBuff
	SpellTypeResurrect
	SpellTypeProjectil
	SpellTypeActivation
	SpellTypeTopo
	SpellTypeSummon
	SpellTypeEarthquake
	SpellTypeDash
	SpellTypeTileburn
	SpellTypeMagnet
	SpellTypePolymorphism
	SpellTypePassive
	SpellTypeTrap
	SpellTypeMount
	SpellTypeCombo
)

type ChatMod uint

const (
	ChatMap ChatMod = iota
	ChatGame
	ChatParty
	ChatTrade
	ChatWhisper
)

type QuestType uint

const (
	QuestNormal QuestType = iota
	QuestStory
	QuestSpecial
)

type SkullType uint

const (
	SkullNormal SkullType = iota
	SkullBronze
)

type PlayerRace uint

const (
	RaceHuman PlayerRace = iota + 1
	RaceWolf
	RaceFox
)

type ToolType uint

const (
	ToolNone ToolType = iota
	ToolSword
	ToolBow
	ToolStaff
	ToolRod
	ToolHatchet
	ToolPickaxe
	ToolDagger
	ToolNet
)

type TargetType uint

const (
	TargetNone TargetType = iota
	TargetPlayer
	TargetNpc
	TargetEnvinroment
	TargetResource
)

type CurrenyOperation uint

const (
	CurrencyNone CurrenyOperation = iota
	CurrencyDrop
	CurrencySell
	CurrencyDeposit
	CurrencyWithdraw
	CurrencyBuy
	CurrencyDropCalaveras
	CurrencyTrade
	CurrencyTradeSkulls
	CurrencyStat
)

type ItemRarityType uint

const (
	RarityNormal ItemRarityType = iota
	RaritySpecial
	RarityPremium
	RarityArena
	RarityRare
	RarityEpic
	RarityLegendary
	RarityDojo
)
