package lib

// Some parameters may vary per connection
// Recommended to create new libGlobals instead of reusing
type libGlobals struct {
	PlayerCoins         int32
	Hotbar              []HotbarRec
	Hotkeys             []int32
	PlayerInventory     []InvItemRec
	PlayerCashInventory []InvItemRec
	PlayerSpells        []PlayerSpellRec
	PlayerVars          []bool
	TNL                 int64
	PlayerQuests        []PlayerQuestRec
	PlayerCraft         []int32
	PlayerCards         []PlayerCardRec
	PlayerAwards        []PlayerAwardsRec
	PlayerCalaveras     []int64
	PlayerPin           []bool
	PlayerProfession    []PlayerProfessionRec
}

func NewGlobals() *libGlobals {
	return &libGlobals{
		Hotbar:              make([]HotbarRec, GlobalHotbarAmount),
		Hotkeys:             make([]int32, GlobalHotkeysAmount),
		PlayerInventory:     make([]InvItemRec, GlobalPlayerInventoryAmount),
		PlayerCashInventory: make([]InvItemRec, GlobalPlayerCashInventoryAmount),
		PlayerSpells:        make([]PlayerSpellRec, GlobalPlayerSpellAmount),
		PlayerVars:          make([]bool, GlobalPlayerVarAmount),
		PlayerQuests:        make([]PlayerQuestRec, GlobalPlayerQuestAmount),
		PlayerCraft:         make([]int32, GlobalPlayerCraftAmount),
		PlayerCards:         make([]PlayerCardRec, GlobalPlayerCardAmount),
		PlayerAwards:        make([]PlayerAwardsRec, GlobalPlayerAwardCount),
		PlayerCalaveras:     make([]int64, GlobalPlayerCalaverasAmount),
		PlayerPin:           make([]bool, GlobalPlayerPinAmount),
		PlayerProfession:    make([]PlayerProfessionRec, GlobalPlayerProfessionAmount),
	}
}

const (
	GlobalHotbarAmount              = 12
	GlobalHotkeysAmount             = 59
	GlobalPlayerInventoryAmount     = 50
	GlobalPlayerCashInventoryAmount = 241
	GlobalPlayerSpellAmount         = 42
	GlobalPlayerVarAmount           = 100
	GlobalPlayerQuestAmount         = 201
	GlobalPlayerCraftAmount         = 255
	GlobalPlayerCardAmount          = 255
	GlobalPlayerAwardCount          = 200
	GlobalPlayerCalaverasAmount     = 1
	GlobalPlayerPinAmount           = 50
	GlobalPlayerProfessionAmount    = 3
)

var Global = NewGlobals()

var Item = make([]ItemRec, 1201)  // TODO fill
var Quest = make([]QuestRec, 201) // TODO fill
var MyIndex = 0
