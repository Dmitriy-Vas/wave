package lib

const (
	CashShopTop     int  = 106
	CashShopLeft    int  = 34
	CashShopOffsetY byte = 86
	CashShopOffsetX byte = 196
	CashShopColumns byte = 2
)
const (
	InvTop     byte = 39
	InvLeft    byte = 13
	InvOffsetY byte = 4
	InvOffsetX byte = 4
	InvColumns byte = 7
)
const (
	SpellTop     byte = 40
	SpellLeft    byte = 39
	SpellOffsetY byte = 4
	SpellOffsetX byte = 4
	SpellColumns byte = 6
)
const (
	EqTop     int64 = 95
	EqLeft    int64 = 122
	EqOffsetX int64 = 4
	EqOffsetY int64 = 4
	EqColumns byte  = 4
)
const (
	BIND_TYPE_NONE   byte = 0
	BIND_TYPE_PICKUP byte = 1
	BIND_TYPE_EQUIP  byte = 2
	BIND_TYPE_INV    byte = 3
)
const (
	TILE_TYPE_WALKABLE    byte = 0
	TILE_TYPE_BLOCKED     byte = 1
	TILE_TYPE_WARP        byte = 2
	TILE_TYPE_ITEM        byte = 3
	TILE_TYPE_NPCAVOID    byte = 4
	TILE_TYPE_KEY         byte = 5
	TILE_TYPE_KEYOPEN     byte = 6
	TILE_TYPE_RESOURCE    byte = 7
	TILE_TYPE_CONDITION   byte = 8
	TILE_TYPE_NPCSPAWN    byte = 9
	TILE_TYPE_SHOP        byte = 10
	TILE_TYPE_BANK        byte = 11
	TILE_TYPE_HEAL        byte = 12
	TILE_TYPE_TRAP        byte = 13
	TILE_TYPE_SLIDE       byte = 14
	TILE_TYPE_ENVINROMENT byte = 15
	TILE_TYPE_LIGHT       byte = 16
	TILE_TYPE_SOUND       byte = 17
	TILE_TYPE_OBJECT      byte = 18
	TILE_TYPE_RESTRICTION byte = 19
	TILE_TYPE_DIALOGUE    byte = 20
	TILE_TYPE_OX          byte = 21
	TILE_TYPE_PVPFREE     byte = 22
)
const (
	TOPO_TYPE_DRILL byte = 1
	TOPO_TYPE_WATER byte = 2
	TOPO_TYPE_COLOR byte = 3
)
const (
	ITEM_TYPE_NONE        byte = 0
	ITEM_TYPE_WEAPON      byte = 1
	ITEM_TYPE_ARMOR       byte = 2
	ITEM_TYPE_HELMET      byte = 3
	ITEM_TYPE_LEGS        byte = 4
	ITEM_TYPE_CAPE        byte = 5
	ITEM_TYPE_GLOVE       byte = 6
	ITEM_TYPE_RING        byte = 7
	ITEM_TYPE_MASK        byte = 8
	ITEM_TYPE_SHOULDER    byte = 9
	ITEM_TYPE_EARRING     byte = 10
	ITEM_TYPE_NECKLACE    byte = 11
	ITEM_TYPE_STRINGUER   byte = 12
	ITEM_TYPE_BELT        byte = 13
	ITEM_TYPE_MOUNT       byte = 14
	ITEM_TYPE_SHOES       byte = 15
	ITEM_TYPE_SHIELD      byte = 16
	ITEM_TYPE_CONSUME     byte = 17
	ITEM_TYPE_KEY         byte = 18
	ITEM_TYPE_SCRIPT      byte = 19
	ITEM_TYPE_SPELL       byte = 20
	ITEM_TYPE_RECIPE      byte = 21
	ITEM_TYPE_ARROW       byte = 22
	ITEM_TYPE_WARP        byte = 23
	ITEM_TYPE_PVP         byte = 24
	ITEM_TYPE_BUFF        byte = 25
	ITEM_TYPE_BAIT        byte = 26
	ITEM_TYPE_MASTER      byte = 27
	ITEM_TYPE_SPECIAL     byte = 28
	ITEM_TYPE_FORGE       byte = 29
	ITEM_TYPE_INSECT      byte = 30
	ITEM_TYPE_COMBO       byte = 31
	ITEM_TYPE_ENHANCEMENT byte = 32
	ITEM_TYPE_GIFTBOX     byte = 33
	ITEM_TYPE_COUNT       byte = 34
)
const (
	DIR_UP    byte = 0
	DIR_DOWN  byte = 1
	DIR_LEFT  byte = 2
	DIR_RIGHT byte = 3
)
const (
	PLAYER_NORMAL byte = 0
	PLAYER_MOD    byte = 1
	PLAYER_ADMIN  byte = 4
)
const (
	NPC_BEHAVIOUR_ATTACKONSIGHT      byte = 0
	NPC_BEHAVIOUR_ATTACKWHENATTACKED byte = 1
	NPC_BEHAVIOUR_FRIENDLY           byte = 2
	NPC_BEHAVIOUR_SHOPKEEPER         byte = 3
	NPC_BEHAVIOUR_GUARD              byte = 4
	NPC_BEHAVIOUR_FOLLOWER           byte = 5
	NPC_BEHAVIOUR_SCRIPT             byte = 6
	NPC_BEHAVIOUR_INSECT             byte = 7
)
const (
	HANDLE_MOUSE_DOWN       byte = 1
	HANDLE_MOUSE_UP         byte = 2
	HANDLE_MOUSE_DOBLECLICK byte = 3
)
const (
	VOID_CONDITION byte = 1
	VOID_FUNCTION  byte = 2
	VOID_REPEAT    byte = 3
	VOID_CHOICE    byte = 4
	VOID_CONTINUE  byte = 5
	VOID_TAG       byte = 6
	VOID_PUZZLE    byte = 7
)
const (
	EDITOR_ITEM         byte = 1
	EDITOR_NPC          byte = 2
	EDITOR_SPELL        byte = 3
	EDITOR_SHOP         byte = 4
	EDITOR_RESOURCE     byte = 5
	EDITOR_ANIMATION    byte = 6
	EDITOR_QUEST        byte = 7
	EDITOR_PROJECTIL    byte = 8
	EDITOR_MORAL        byte = 9
	EDITOR_FRAMES       byte = 10
	EDITOR_CARD         byte = 11
	EDITOR_MAP          byte = 12
	EDITOR_CRAFTING     byte = 13
	EDITOR_MAPTOPO      byte = 14
	EDITOR_GUI          byte = 15
	EDITOR_CONDITION    byte = 16
	EDITOR_QUESTTALK    byte = 17
	EDITOR_LANGUAGE     byte = 18
	EDITOR_MAPANIMATION byte = 19
	EDITOR_EVENT        byte = 20
	EDITOR_PUZZLE       byte = 21
	EDITOR_CLASS        byte = 22
	EDITOR_LEVEL        byte = 23
	EDITOR_TRAVEL       byte = 24
	EDITOR_COMBO        byte = 25
	EDITOR_ELEMENT      byte = 26
	EDITOR_PARTYQUEST   byte = 27
	EDITOR_AWARD        byte = 28
	EDITOR_CASHSHOP     byte = 29
	EDITOR_COUPON       byte = 30
	EDITOR_FORGE        byte = 31
	EDITOR_FONT         byte = 32
	EDITOR_WORLD        byte = 33
	EDITOR_DAILYCHECK   byte = 34
	EDITOR_EMOTICONS    byte = 35
	EDITOR_OXQUESTION   byte = 36
)
const (
	ACTIONMSG_STATIC int64 = 0
	ACTIONMSG_SCROLL int64 = 1
	ACTIONMSG_SCREEN int64 = 2
)
const (
	DIALOGUE_TYPE_NONE   byte = 0
	DIALOGUE_TYPE_TRADE  byte = 1
	DIALOGUE_TYPE_FORGET byte = 2
	DIALOGUE_TYPE_PARTY  byte = 3
	DIALOGUE_TYPE_BUDDY  byte = 4
)
const (
	ColorWhite       int64 = 1
	ColorBlack       int64 = 2
	ColorBrown       int64 = 3
	ColorCalmBrown   int64 = 4
	ColorYellow      int64 = 5
	ColorGray        int64 = 6
	ColorRed         int64 = 7
	ColorOrange      int64 = 8
	ColorPink        int64 = 9
	ColorSky         int64 = 10
	ColorLodo        int64 = 11
	ColorIce         int64 = 12
	ColorPurple      int64 = 13
	ColorGrayOp      int64 = 14
	ColorOrange2     int64 = 15
	ColorBrownOrange int64 = 16
	ColorStump       int64 = 17
	ColorYellowGray  int64 = 18
	ColorAtenuar     int64 = 19
	ColorCalmBrown2  int64 = 20
)
const (
	ATTACK_DAMAGE int64 = 1
	ATTACK_RANGE  int64 = 2
	ATTACK_DOBLE  int64 = 3
)
const (
	MaxPlayerArrow  byte = 3
	MaxPlayerBuffs  byte = 10
	MaxPlayerCombo  byte = 18
	MaxPlayerEmoji  byte = 50
	MaxPlayerMaster byte = 25
	MaxPlayerRank   byte = 10
	MaxPlayerSolar  byte = 2
)
const (
	MAX_ANIMATIONS         int     = 255
	MAX_ARROWS             byte    = 2
	MAX_BANK_PREMIUM_SLOTS byte    = 144
	MAX_BANK_SLOTS         byte    = 48
	MAX_BUFFS              int64   = 10
	MAX_BYTE               byte    = 255
	MAX_CARDS              int64   = 255
	MAX_CLASSREQ           byte    = 6
	MAX_CONDITIONS         int     = 400
	MAX_CPT                byte    = 6
	MAX_CRAFTING           int64   = 255
	MAX_CRAFTING_TASK      int64   = 3
	MAX_EMAIL              byte    = 50
	MAX_HOTBAR             int64   = 12
	MAX_ID                 byte    = 12
	MAX_ITEMS              int     = 1200
	MAX_LANGUAGE           byte    = 3
	MAX_LIGHTS             byte    = 25
	MAX_LOGMSG             byte    = 8
	MAX_MAPRANK            byte    = 4
	MAX_MAPS               int     = 200
	MAX_MAP_ANIMATIONS     byte    = 25
	MAX_MAP_ITEMS          byte    = 255
	MAX_MAP_NPCS           byte    = 30
	MAX_MAP_TOPO           int64   = 500
	MAX_MAP_TOPO_ITEMS     int64   = 10
	MAX_MINIMAPX           byte    = 31
	MAX_MINIMAPY           byte    = 25
	MAX_MORALS             byte    = 25
	MAX_NPC_DROPS          byte    = 16
	MAX_NPC_PROJECTILES    byte    = 5
	MAX_NPC_SPELLS         int64   = 10
	MAX_NUMBERS            int     = 300
	MAX_PARTYQUEST         int     = 10
	MAX_PARTYS             int64   = 100
	MAX_PARTY_MEMBERS      int64   = 4
	MAX_PLAYERVAR          int     = 100
	MAX_PLAYER_CHARACTERS  byte    = 3
	MAX_PLAYER_PROJECTILES byte    = 5
	MAX_PLAYER_SPELLS      byte    = 42
	MAX_PROJECTILES        int64   = 100
	MAX_PW                 byte    = 18
	MAX_RESOURCES          byte    = 100
	MAX_RESOURCE_REWARDS   byte    = 10
	MAX_SELECTED_COMBO     byte    = 3
	MAX_SHOPS              byte    = 200
	MAX_SKULL              int     = 5000
	MAX_SOUND              byte    = 50
	MAX_SPELLS             byte    = 255
	MAX_TALK               byte    = 25
	MAX_TEXT_CODE          byte    = 25
	MAX_TITLES             int64   = 255
	MAX_TRADES             byte    = 36
	MAX_WALK_SPEED         float64 = 3.2
)
const (
	MaxAwards          int  = 200
	MaxAwardsLevel     byte = 30
	MaxBuddies         byte = 25
	MaxButtons         int  = 500
	MaxEmoticons       int  = 255
	MaxFavAwards       byte = 6
	MaxInv             byte = 49
	MaxLangText        int  = 1000
	MaxMapRegion       byte = 70
	MaxNpcs            int  = 350
	MaxPVPRank         byte = 4
	MaxPlayers         int  = 250
	MaxPuzzles         int  = 50
	MaxQuestLogPlayers byte = 15
	MaxRestrictedWords byte = 100
	MaxTaskMsg         byte = 4
	MaxWalkSpeed       byte = 125
	MaxWorld           byte = 30
)
const (
	FishRodFrame           byte    = 13
	HOTBAR_TYPE_ITEM       byte    = 1
	HOTBAR_TYPE_SPELL      byte    = 2
	HatchetFrame           byte    = 14
	ITEM_SPAWN_TIME        int64   = 35000
	MASK_BELOW             byte    = 0
	MASK_PLAYER            byte    = 1
	MASK_UPPER             byte    = 2
	MAX_STACK              int     = 2000
	MinMapX                byte    = 31
	MinMapY                byte    = 23
	NAME_LENGTH            byte    = 20
	PIC_X                  int     = 32
	PIC_Y                  int     = 32
	PUZZLE_MSIZE           int     = 3
	Party_HPWidth          int64   = 124
	PickaxeFrame           byte    = 14
	SEX_FEMALE             byte    = 1
	SEX_MALE               byte    = 0
	SIZE_X                 int     = 32
	SIZE_Y                 int     = 32
	TRANSPARENT_BACKGROUND int     = 171
	TalkRange              byte    = 4
	WalkSpeed              float64 = 2.8
	YES                    byte    = 1
	NO                     byte    = 0
)
