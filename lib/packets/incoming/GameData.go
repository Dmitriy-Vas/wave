package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (d *GameDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GameDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GameDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GameDataPacket) SetSend(value bool) {
	d.Send = value
}

type GameDataPacket struct {
	ID   int64
	Send bool

	Demo              bool
	ItemsNum          []int32
	Items             []lib.ItemRec
	RandomItemsNum    []int32
	RandomItems       []lib.RandomItemRec
	PuzzlesNum        []int32
	Puzzles           []lib.PuzzleRec
	AnimationsNum     []int32
	Animations        []lib.AnimationRec
	NpcsNum           []int32
	Npcs              []lib.NpcRec
	ShopsNum          []int32
	Shops             []lib.ShopRec
	SpellsNum         []int32
	Spells            []lib.SpellRec
	ResourcesNum      []int32
	Resources         []lib.ResourceRec
	MoralsNum         []int32
	Morals            []lib.MoralRec
	ProjectilsNum     []int32
	Projectils        []lib.ProjectilRec
	CardsNum          []int32
	Cards             []lib.CardsRec
	CraftsNum         []int32
	Crafts            []lib.CraftingRec
	QuestsNum         []int32
	Quests            []lib.QuestRec
	CashShopsNum      []int32
	CashShops         []lib.CashShopRec
	TopItems          []int32
	HotItems          []lib.HotItemRec
	Variable1         byte
	ConditionsNum     []int32
	Conditions        []lib.ConditionRec
	EmoticonsNum      []int32
	Emoticons         []lib.EmoticonRec
	DailyCheck        []lib.InvItemRec
	MapList           []lib.MapListRec
	Pins              []lib.PinsRec
	Professions       []lib.ProfessionRec
	Newspaper         lib.NewspaperRec
	EventTime         int32
	WorldBoss         []time.Time
	Promotion         lib.PromotionRec
	ServerMods        []lib.ServerModRec
	PlayerBlockedList string
	ServerGlobalVar   int32
	DailyCheckDate    time.Time
	WorldBless        lib.WorldBlessRec
	Speed             lib.SpeedFormulaRec
	Variable2         byte
	Selected          byte
	RestrictedWords   []string

	Data []byte
}

func (packet *GameDataPacket) Read(b buffer.PacketBuffer) {
	packet.Data = b.(*wrapper.Buffer).DefaultBuffer.ReadBytes(b.Bytes(), b.Index(), b.Len()-16)
}

func (packet *GameDataPacket) Write(b buffer.PacketBuffer) {
	b.(*wrapper.Buffer).DefaultBuffer.WriteBytes(b.Bytes(), packet.Data, b.Index())
}

func (packet *GameDataPacket) RealRead(b buffer.PacketBuffer) {
	// TODO decompress
	packet.Demo = b.ReadBool(b.Bytes(), b.Index())
	itemsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ItemsNum = make([]int32, itemsAmount)
	packet.Items = make([]lib.ItemRec, itemsAmount)
	for i, _ := range packet.Items {
		packet.ItemsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Items[i] = lib.ReadItemData(b)
	}
	randomItemsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.RandomItemsNum = make([]int32, randomItemsAmount)
	packet.RandomItems = make([]lib.RandomItemRec, randomItemsAmount)
	for i, _ := range packet.RandomItems {
		packet.RandomItemsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.RandomItems[i] = lib.ReadRandomItemData(b)
	}
	puzzlesAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.PuzzlesNum = make([]int32, puzzlesAmount)
	packet.Puzzles = make([]lib.PuzzleRec, puzzlesAmount)
	for i, _ := range packet.Puzzles {
		packet.PuzzlesNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Puzzles[i] = lib.ReadPuzzleData(b)
	}
	animationAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.AnimationsNum = make([]int32, animationAmount)
	packet.Animations = make([]lib.AnimationRec, animationAmount)
	for i, _ := range packet.Animations {
		packet.AnimationsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Animations[i] = lib.ReadAnimationData(b)
	}
	npcsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.NpcsNum = make([]int32, npcsAmount)
	packet.Npcs = make([]lib.NpcRec, npcsAmount)
	for i, _ := range packet.Npcs {
		packet.NpcsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Npcs[i] = lib.ReadNpcData(b)
	}
	shopsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ShopsNum = make([]int32, shopsAmount)
	packet.Shops = make([]lib.ShopRec, shopsAmount)
	for i, _ := range packet.Shops {
		packet.ShopsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Shops[i] = lib.ReadShopData(b)
	}
	spellsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.SpellsNum = make([]int32, spellsAmount)
	packet.Spells = make([]lib.SpellRec, spellsAmount)
	for i, _ := range packet.Spells {
		packet.SpellsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Spells[i] = lib.ReadSpellData(b)
	}
	resourcesAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ResourcesNum = make([]int32, resourcesAmount)
	packet.Resources = make([]lib.ResourceRec, resourcesAmount)
	for i, _ := range packet.Resources {
		packet.ResourcesNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Resources[i] = lib.ReadResourceData(b)
	}
	moralsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.MoralsNum = make([]int32, moralsAmount)
	packet.Morals = make([]lib.MoralRec, moralsAmount)
	for i, _ := range packet.Morals {
		packet.MoralsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Morals[i] = lib.ReadMoralData(b)
	}
	projectilsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.ProjectilsNum = make([]int32, projectilsAmount)
	packet.Projectils = make([]lib.ProjectilRec, projectilsAmount)
	for i, _ := range packet.Projectils {
		packet.ProjectilsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		projectil := lib.ProjectilRec{
			Name:      b.ReadString(b.Bytes(), b.Index(), 0),
			Damage:    int64(b.ReadInt(b.Bytes(), b.Index())),
			Pic:       b.ReadInt(b.Bytes(), b.Index()),
			Range:     int64(b.ReadInt(b.Bytes(), b.Index())),
			Speed:     int64(b.ReadInt(b.Bytes(), b.Index())),
			Animation: int64(b.ReadInt(b.Bytes(), b.Index())),
			Light:     b.ReadBool(b.Bytes(), b.Index()),
			Int:       make([]int32, 3), // TODO int to const
		}
		for x, _ := range projectil.Int {
			projectil.Int[x] = b.ReadInt(b.Bytes(), b.Index())
		}
	}
	cardsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CardsNum = make([]int32, cardsAmount)
	packet.Cards = make([]lib.CardsRec, cardsAmount)
	for i, _ := range packet.Cards {
		packet.CardsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Cards[i] = lib.ReadCardData(b)
	}
	craftsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CraftsNum = make([]int32, craftsAmount)
	packet.Crafts = make([]lib.CraftingRec, craftsAmount)
	for i, _ := range packet.Cards {
		packet.CraftsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Crafts[i] = lib.ReadCraftData(b)
	}
	questsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.QuestsNum = make([]int32, questsAmount)
	packet.Quests = make([]lib.QuestRec, questsAmount)
	for i, _ := range packet.Quests {
		packet.QuestsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Quests[i] = lib.ReadQuestData(b)
	}
	cashShopsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.CashShopsNum = make([]int32, cashShopsAmount)
	packet.CashShops = make([]lib.CashShopRec, cashShopsAmount)
	for i, _ := range packet.CashShops {
		packet.CashShopsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.CashShops[i] = lib.ReadCashShopData(b)
	}
	packet.TopItems = make([]int32, 8) // TODO int to const
	for i, _ := range packet.TopItems {
		packet.TopItems[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	packet.HotItems = make([]lib.HotItemRec, 24) // TODO int to const
	for i, _ := range packet.HotItems {
		packet.HotItems[i] = lib.HotItemRec{
			Category:    b.ReadInt(b.Bytes(), b.Index()),
			SubCategory: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	if packet.Variable1 == 4 {
		conditionsAmount := b.ReadInt(b.Bytes(), b.Index())
		packet.ConditionsNum = make([]int32, conditionsAmount)
		packet.Conditions = make([]lib.ConditionRec, conditionsAmount)
		for i, _ := range packet.Conditions {
			packet.ConditionsNum[i] = b.ReadInt(b.Bytes(), b.Index())
			packet.Conditions[i] = lib.ReadConditionData(b)
		}
	}
	emoticonsAmount := b.ReadInt(b.Bytes(), b.Index())
	packet.EmoticonsNum = make([]int32, emoticonsAmount)
	packet.Emoticons = make([]lib.EmoticonRec, emoticonsAmount)
	for i, _ := range packet.Emoticons {
		packet.EmoticonsNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Emoticons[i] = lib.ReadEmoticonData(b)
	}
	packet.DailyCheck = make([]lib.InvItemRec, 28) // TODO int to const
	for i, _ := range packet.DailyCheck {
		packet.DailyCheck[i] = lib.InvItemRec{
			Num:   b.ReadInt(b.Bytes(), b.Index()),
			Value: b.ReadLong(b.Bytes(), b.Index()),
		}
	}
	packet.MapList = make([]lib.MapListRec, 200) // TODO int to const
	for i, _ := range packet.MapList {
		packet.MapList[i].Name = make([]string, 2) // TODO int to const
		for x, _ := range packet.MapList[i].Name {
			packet.MapList[i].Name[x] = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	}
	packet.Pins = make([]lib.PinsRec, 51)
	for i, _ := range packet.Pins {
		packet.Pins[i] = lib.PinsRec{
			Item: b.ReadInt(b.Bytes(), b.Index()),
			Desc: b.ReadString(b.Bytes(), b.Index(), 0),
		}
	}
	packet.Professions = lib.ReadProfessionData(b)
	packet.Newspaper = lib.ReadNewspaperData(b)
	packet.EventTime = b.ReadInt(b.Bytes(), b.Index())
	packet.WorldBoss = make([]time.Time, 2) // TODO int to const
	for i, _ := range packet.WorldBoss {
		packet.WorldBoss[i] = wrapper.ReadDate(b)
	}
	packet.Promotion = lib.PromotionRec{
		Type: b.ReadByte(b.Bytes(), b.Index()),
		Item: b.ReadInt(b.Bytes(), b.Index()),
		GameCard: make([]lib.PromotionGameCardRec,
			b.ReadByte(b.Bytes(), b.Index())),
	}
	for i, _ := range packet.Promotion.GameCard {
		card := lib.PromotionGameCardRec{
			Name:  b.ReadString(b.Bytes(), b.Index(), 0),
			Price: b.ReadString(b.Bytes(), b.Index(), 0),
			Bonus: b.ReadInt(b.Bytes(), b.Index()),
		}
		packet.Promotion.GameCard[i] = card
	}
	packet.ServerMods = make([]lib.ServerModRec, lib.ServerModQuestExp+1)
	for i := lib.ServerModExperience; i <= lib.ServerModQuestExp; i++ {
		packet.ServerMods[i] = lib.ServerModRec{
			Value: b.ReadDouble(b.Bytes(), b.Index()),
			Timer: int64(b.ReadInt(b.Bytes(), b.Index())),
		}
	}
	packet.PlayerBlockedList = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.WorldBless = lib.WorldBlessRec{Active: b.ReadBool(b.Bytes(), b.Index())}
	if packet.WorldBless.Active {
		packet.WorldBless.Owner = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Sprite = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Hair = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Paperdoll = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.WorldBless.Time = b.ReadInt(b.Bytes(), b.Index())
		packet.WorldBless.Value = b.ReadDouble(b.Bytes(), b.Index())
	}
	packet.Speed = lib.SpeedFormulaRec{
		Normal: b.ReadDouble(b.Bytes(), b.Index()),
		Buff:   b.ReadDouble(b.Bytes(), b.Index()),
		Mount:  b.ReadDouble(b.Bytes(), b.Index()),
	}
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Selected = b.ReadByte(b.Bytes(), b.Index())
	packet.RestrictedWords = make([]string, 101) // TODO int to const
	for i, _ := range packet.RestrictedWords {
		packet.RestrictedWords[i] = b.ReadString(b.Bytes(), b.Index(), 0)
		if packet.RestrictedWords[i] == "" {
			break
		}
	}
}

func (packet *GameDataPacket) RealWrite(b buffer.PacketBuffer) {
	// TODO compress
}
