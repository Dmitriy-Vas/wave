package lib

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

func PlayerInvData(b buffer.PacketBuffer, slot int) {
	Global.PlayerInventory[slot].Num = b.ReadInt(b.Bytes(), b.Index())
	Global.PlayerInventory[slot].Value = b.ReadLong(b.Bytes(), b.Index())
	Global.PlayerInventory[slot].Slot = b.ReadByte(b.Bytes(), b.Index())
	if GetItemInt(Global.PlayerInventory[slot].Num, ItemEnhancement) > 0 {
		Global.PlayerInventory[slot].Stat = make([]byte, 5)
		for i := range Global.PlayerInventory[slot].Stat {
			Global.PlayerInventory[slot].Stat[i] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
}

func GetItemInt(itemNum int32, property ItemPropertyType) int32 {
	if itemNum == 0 || itemNum > 1200 {
		return 0
	}
	if property < ItemLight || property > Itemexp_lost {
		return 0
	}
	return Item[itemNum].Int[property]
}

func SetPlayerCashInvItem(slot int, itemNum int32, itemValue int32) {
	Global.PlayerCashInventory[slot].Num = itemNum
	Global.PlayerCashInventory[slot].Value = int64(itemValue)
}

func ReadPlayerBankData(b buffer.PacketBuffer) InvItemRec {
	item := InvItemRec{
		Num:   b.ReadInt(b.Bytes(), b.Index()),
		Value: b.ReadLong(b.Bytes(), b.Index()),
		Slot:  b.ReadByte(b.Bytes(), b.Index()),
		Stat:  make([]byte, 6),
	}
	if true { //modTypes.Item[ptr.Num].Int[13] > 0; TODO
		for i := range item.Stat {
			item.Stat[i] = b.ReadByte(b.Bytes(), b.Index())
		}
	}
	return item
}

func WritePlayerBankData(b buffer.PacketBuffer, item InvItemRec) {
	b.WriteInt(b.Bytes(), item.Num, b.Index())
	b.WriteLong(b.Bytes(), item.Value, b.Index())
	b.WriteByte(b.Bytes(), item.Slot, b.Index())
	if true { //modTypes.Item[ptr.Num].Int[13] > 0; TODO
		for _, data := range item.Stat {
			b.WriteByte(b.Bytes(), data, b.Index())
		}
	}
}

func ReadCashShopData(b buffer.PacketBuffer) CashShopRec {
	item := CashShopRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Icon: b.ReadInt(b.Bytes(), b.Index()),
	}
	item.Slot = make([]TradeCashItemRec, b.ReadInt(b.Bytes(), b.Index()))
	for i := range item.Slot {
		item.Slot[i].Item = b.ReadInt(b.Bytes(), b.Index())
		item.Slot[i].Value = b.ReadLong(b.Bytes(), b.Index())
		item.Slot[i].Price = b.ReadLong(b.Bytes(), b.Index())
	}
	return item
}

func WriteCashShopData(b buffer.PacketBuffer, item CashShopRec) {
	b.WriteString(b.Bytes(), item.Name, b.Index())
	b.WriteInt(b.Bytes(), item.Icon, b.Index())
	b.WriteInt(b.Bytes(), int32(len(item.Slot)), b.Index())
	for _, slot := range item.Slot {
		b.WriteInt(b.Bytes(), slot.Item, b.Index())
		b.WriteLong(b.Bytes(), slot.Value, b.Index())
		b.WriteLong(b.Bytes(), slot.Price, b.Index())
	}
}

func ReadMapData(b buffer.PacketBuffer) MapRec {
	rec := MapRec{
		Name:         b.ReadString(b.Bytes(), b.Index(), 0),
		ESName:       b.ReadString(b.Bytes(), b.Index(), 0),
		Revision:     b.ReadInt(b.Bytes(), b.Index()),
		Moral:        b.ReadInt(b.Bytes(), b.Index()),
		MaxX:         b.ReadByte(b.Bytes(), b.Index()),
		MaxY:         b.ReadByte(b.Bytes(), b.Index()),
		Region:       b.ReadInt(b.Bytes(), b.Index()),
		Image:        b.ReadInt(b.Bytes(), b.Index()),
		Up:           b.ReadInt(b.Bytes(), b.Index()),
		Down:         b.ReadInt(b.Bytes(), b.Index()),
		Left:         b.ReadInt(b.Bytes(), b.Index()),
		Right:        b.ReadInt(b.Bytes(), b.Index()),
		BootMap:      b.ReadInt(b.Bytes(), b.Index()),
		BootX:        b.ReadInt(b.Bytes(), b.Index()),
		BootY:        b.ReadInt(b.Bytes(), b.Index()),
		Fog:          b.ReadInt(b.Bytes(), b.Index()),
		FogOpacity:   b.ReadByte(b.Bytes(), b.Index()),
		FogSpeed:     b.ReadInt(b.Bytes(), b.Index()),
		tileA:        b.ReadInt(b.Bytes(), b.Index()),
		tileB:        b.ReadInt(b.Bytes(), b.Index()),
		DayNight:     b.ReadInt(b.Bytes(), b.Index()),
		Parallax:     b.ReadInt(b.Bytes(), b.Index()),
		ParallaxType: b.ReadByte(b.Bytes(), b.Index()),
		Ambience:     b.ReadString(b.Bytes(), b.Index(), 0),
		NightAlpha:   b.ReadInt(b.Bytes(), b.Index()),
		InstanceMax:  b.ReadInt(b.Bytes(), b.Index()),
		MapCondition: b.ReadInt(b.Bytes(), b.Index()),
		Puzzle: objects.Vector4{
			A: b.ReadInt(b.Bytes(), b.Index()),
			B: b.ReadInt(b.Bytes(), b.Index()),
			C: b.ReadInt(b.Bytes(), b.Index()),
			D: b.ReadInt(b.Bytes(), b.Index()),
		},
		Drill: objects.Vector4{
			A: b.ReadInt(b.Bytes(), b.Index()),
			B: b.ReadInt(b.Bytes(), b.Index()),
			C: b.ReadInt(b.Bytes(), b.Index()),
			D: b.ReadInt(b.Bytes(), b.Index()),
		},
		tileOffset: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
	}
	// TODO finish map data
	return rec
}

func WriteMapData(b buffer.PacketBuffer, item MapRec) {

}

func ReadItemData(b buffer.PacketBuffer) ItemRec {
	item := ItemRec{
		Name:          b.ReadString(b.Bytes(), b.Index(), 0),
		ESName:        b.ReadString(b.Bytes(), b.Index(), 0),
		AddStat:       make([]int32, 7),  // TODO int to const
		VitalMode:     make([]int32, 3),  // TODO int to const
		ClassReq:      make([]int64, 6),  // TODO int to const
		StatReq:       make([]int32, 7),  // TODO int to const
		Int:           make([]int32, 21), // TODO int to const
		Bool:          make([]bool, 1),   // TODO int to const
		Element:       make([]int32, 4),  // TODO int to const
		ElementChance: make([]int32, 4),  // TODO int to const
	}
	for i := range item.AddStat {
		item.AddStat[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	item.Animation = b.ReadLong(b.Bytes(), b.Index())
	item.BindType = b.ReadByte(b.Bytes(), b.Index())
	for i := range item.VitalMode {
		item.VitalMode[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range item.ClassReq {
		item.ClassReq[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	item.Data1 = b.ReadInt(b.Bytes(), b.Index())
	item.Data2 = b.ReadInt(b.Bytes(), b.Index())
	item.Tool = b.ReadInt(b.Bytes(), b.Index())
	item.Handed = b.ReadLong(b.Bytes(), b.Index())
	item.LevelReq = b.ReadLong(b.Bytes(), b.Index())
	item.Mastery = b.ReadByte(b.Bytes(), b.Index())
	item.AccessReq = b.ReadLong(b.Bytes(), b.Index())
	item.Paperdoll = b.ReadLong(b.Bytes(), b.Index())
	item.Pic = b.ReadInt(b.Bytes(), b.Index())
	item.Price = b.ReadLong(b.Bytes(), b.Index())
	item.Rarity = b.ReadByte(b.Bytes(), b.Index())
	item.Speed = b.ReadLong(b.Bytes(), b.Index())
	for i := range item.StatReq {
		item.StatReq[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	item.Type = b.ReadByte(b.Bytes(), b.Index())
	item.Desc = b.ReadString(b.Bytes(), b.Index(), 0)
	item.Stackable = b.ReadBool(b.Bytes(), b.Index())
	item.Overol = b.ReadInt(b.Bytes(), b.Index())
	item.paperWidth = b.ReadInt(b.Bytes(), b.Index())
	item.ProjecTil = b.ReadInt(b.Bytes(), b.Index())
	item.MType = b.ReadInt(b.Bytes(), b.Index())
	item.HPBonus = b.ReadInt(b.Bytes(), b.Index())
	item.MPBonus = b.ReadInt(b.Bytes(), b.Index())
	item.WarpMap = b.ReadInt(b.Bytes(), b.Index())
	item.WarpX = b.ReadInt(b.Bytes(), b.Index())
	item.WarpY = b.ReadInt(b.Bytes(), b.Index())
	item.isHair = b.ReadInt(b.Bytes(), b.Index())
	item.BigPic = b.ReadInt(b.Bytes(), b.Index())
	item.comboSlot = b.ReadLong(b.Bytes(), b.Index())
	item.Recipe = b.ReadInt(b.Bytes(), b.Index())
	for i := range item.Int {
		item.Int[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range item.Bool {
		item.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	for i := range item.Element {
		item.Element[i] = b.ReadInt(b.Bytes(), b.Index())
		item.ElementChance[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	item.Revision = b.ReadInt(b.Bytes(), b.Index())
	return item
}

func ReadRandomItemData(b buffer.PacketBuffer) RandomItemRec {
	randomItem := RandomItemRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Item: make([]RandomItemListRec, 11), // TODO int to const
	}
	for i := range randomItem.Item {
		randomItem.Item[i].Num = b.ReadInt(b.Bytes(), b.Index())
		randomItem.Item[i].Value = b.ReadInt(b.Bytes(), b.Index())
		randomItem.Item[i].Luck = b.ReadByte(b.Bytes(), b.Index())
	}
	return randomItem
}

func ReadPuzzleData(b buffer.PacketBuffer) PuzzleRec {
	puzzle := PuzzleRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Map:  b.ReadInt(b.Bytes(), b.Index()),
		Point: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
		Size: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
		SwitchImage: b.ReadInt(b.Bytes(), b.Index()),
		SwitchPos: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
	}
	if puzzle.Size.X > 0 && puzzle.Size.Y > 0 {
		puzzle.Cube = make([][]PuzzleCubeRec, puzzle.Size.X)
		for i := range puzzle.Cube {
			puzzle.Cube[i] = make([]PuzzleCubeRec, puzzle.Size.Y)
		}
		for x := range puzzle.Cube {
			for y := range puzzle.Cube[x] {
				puzzle.Cube[x][y] = PuzzleCubeRec{
					Image:  b.ReadInt(b.Bytes(), b.Index()),
					Move:   b.ReadBool(b.Bytes(), b.Index()),
					Block:  b.ReadBool(b.Bytes(), b.Index()),
					Key:    b.ReadBool(b.Bytes(), b.Index()),
					Switch: b.ReadInt(b.Bytes(), b.Index()),
				}
			}
		}
	}
	return puzzle
}

func ReadAnimationData(b buffer.PacketBuffer) AnimationRec {
	animation := AnimationRec{
		Name:      b.ReadString(b.Bytes(), b.Index(), 0),
		Frames:    make([]int64, 2), // TODO int to const
		LoopCount: make([]int64, 2), // TODO int to const
		LoopTime:  make([]int64, 2), // TODO int to const
		Sprite:    make([]int64, 2), // TODO int to const
		XOffset:   make([]int64, 2), // TODO int to const
		YOffset:   make([]int64, 2), // TODO int to const
	}
	for i := range animation.Frames {
		animation.Frames[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range animation.LoopCount {
		animation.LoopCount[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range animation.LoopTime {
		animation.LoopTime[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range animation.Sprite {
		animation.Sprite[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range animation.XOffset {
		animation.XOffset[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range animation.YOffset {
		animation.YOffset[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	animation.Sound = b.ReadString(b.Bytes(), b.Index(), 0)
	animation.SoundValue = b.ReadInt(b.Bytes(), b.Index())
	return animation
}

func ReadNpcData(b buffer.PacketBuffer) NpcRec {
	npc := NpcRec{
		AttackSay:     b.ReadString(b.Bytes(), b.Index(), 0),
		Behaviour:     b.ReadByte(b.Bytes(), b.Index()),
		DropChance:    make([]int32, 16),          // TODO int to const
		DropItem:      make([]int32, 16),          // TODO int to const
		DropItemValue: make([]int32, 16),          // TODO int to const
		Spell:         make([]int32, 10),          // TODO int to const
		Element:       make([]int64, 4),           // TODO int to const
		Bool:          make([]bool, 26),           // TODO int to const
		Int:           make([]int32, 23),          // TODO int to const
		Vec:           make([]objects.Vector2, 1), // TODO int to const
		QuestList:     make([]int32, 6),           // TODO int to const
	}
	for i := range npc.DropChance {
		npc.DropChance[i] = b.ReadInt(b.Bytes(), b.Index())
		npc.DropItem[i] = b.ReadInt(b.Bytes(), b.Index())
		npc.DropItemValue[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	npc.Faction = b.ReadByte(b.Bytes(), b.Index())
	npc.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	npc.Range = b.ReadByte(b.Bytes(), b.Index())
	npc.RequireVar = b.ReadInt(b.Bytes(), b.Index())
	npc.AttackSound = b.ReadString(b.Bytes(), b.Index(), 0)
	npc.Light = b.ReadInt(b.Bytes(), b.Index())
	npc.CardDrop = b.ReadInt(b.Bytes(), b.Index())
	npc.CardNum = b.ReadInt(b.Bytes(), b.Index())
	for i := range npc.Spell {
		npc.Spell[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range npc.Element {
		npc.Element[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range npc.Bool {
		npc.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	for i := range npc.Int {
		npc.Int[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range npc.Vec {
		npc.Vec[i] = objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
	for i := range npc.QuestList {
		npc.QuestList[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	if npc.Bool[9] {
		npc.ScriptList = make([]NpcScriptListRec, 6) // TODO int to const
		for i := range npc.ScriptList {
			npc.ScriptList[i].Num = b.ReadInt(b.Bytes(), b.Index())
			npc.ScriptList[i].Interval = b.ReadInt(b.Bytes(), b.Index())
			npc.ScriptList[i].Variable = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	}
	npc.Offset = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	return npc
}

func ReadShopData(b buffer.PacketBuffer) ShopRec {
	shop := ShopRec{
		Name:      b.ReadString(b.Bytes(), b.Index(), 0),
		BuyRate:   b.ReadInt(b.Bytes(), b.Index()),
		priceItem: b.ReadInt(b.Bytes(), b.Index()),
		Order:     b.ReadByte(b.Bytes(), b.Index()),
		Type:      b.ReadByte(b.Bytes(), b.Index()),
		Pic:       b.ReadInt(b.Bytes(), b.Index()),
		TradeItem: make([]TradeItemRec, 37), // TODO int to const
	}
	for i := range shop.TradeItem {
		shop.TradeItem[i].CostItem = b.ReadInt(b.Bytes(), b.Index())
		shop.TradeItem[i].CostValue = b.ReadLong(b.Bytes(), b.Index())
		shop.TradeItem[i].Item = b.ReadInt(b.Bytes(), b.Index())
		shop.TradeItem[i].ItemValue = b.ReadLong(b.Bytes(), b.Index())
		shop.TradeItem[i].RequiredVar = b.ReadInt(b.Bytes(), b.Index())
	}
	return shop
}

func ReadSpellData(b buffer.PacketBuffer) SpellRec {
	spell := SpellRec{
		AccessReq:    b.ReadByte(b.Bytes(), b.Index()),
		AoE:          b.ReadLong(b.Bytes(), b.Index()),
		CastAnim:     b.ReadLong(b.Bytes(), b.Index()),
		CastTime:     b.ReadLong(b.Bytes(), b.Index()),
		CDTime:       b.ReadLong(b.Bytes(), b.Index()),
		ClassReq:     b.ReadInt(b.Bytes(), b.Index()),
		Dir:          b.ReadByte(b.Bytes(), b.Index()),
		Duration:     b.ReadLong(b.Bytes(), b.Index()),
		Icon:         b.ReadInt(b.Bytes(), b.Index()),
		Interval:     b.ReadLong(b.Bytes(), b.Index()),
		IsAoE:        b.ReadBool(b.Bytes(), b.Index()),
		LevelReq:     b.ReadInt(b.Bytes(), b.Index()),
		Map:          b.ReadInt(b.Bytes(), b.Index()),
		MPCost:       b.ReadInt(b.Bytes(), b.Index()),
		Name:         b.ReadString(b.Bytes(), b.Index(), 0),
		ESName:       b.ReadString(b.Bytes(), b.Index(), 0),
		Desc:         b.ReadString(b.Bytes(), b.Index(), 0),
		Range:        b.ReadLong(b.Bytes(), b.Index()),
		SpellAnim:    b.ReadLong(b.Bytes(), b.Index()),
		StunDuration: b.ReadLong(b.Bytes(), b.Index()),
		Type:         b.ReadByte(b.Bytes(), b.Index()),
		Vital:        b.ReadLong(b.Bytes(), b.Index()),
		X:            b.ReadInt(b.Bytes(), b.Index()),
		Y:            b.ReadInt(b.Bytes(), b.Index()),
		BuffType:     b.ReadLong(b.Bytes(), b.Index()),
		Frame:        b.ReadLong(b.Bytes(), b.Index()),
		NextRank:     b.ReadLong(b.Bytes(), b.Index()),
		NextUses:     b.ReadLong(b.Bytes(), b.Index()),
		MaxMob:       b.ReadLong(b.Bytes(), b.Index()),
		Stat:         b.ReadLong(b.Bytes(), b.Index()),
		Cases:        b.ReadByte(b.Bytes(), b.Index()),
		ElementBased: b.ReadBool(b.Bytes(), b.Index()),
		ClassBasic:   b.ReadBool(b.Bytes(), b.Index()),
		CastWalking:  b.ReadBool(b.Bytes(), b.Index()),
		HPCost:       b.ReadInt(b.Bytes(), b.Index()),
		Int:          make([]int32, 14), // TODO int to const
		Bool:         make([]bool, 3),   // TODO int to const
	}
	for i := range spell.Int {
		spell.Int[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range spell.Bool {
		spell.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	return spell
}

func ReadResourceData(b buffer.PacketBuffer) ResourceRec {
	resource := ResourceRec{
		Name:           b.ReadString(b.Bytes(), b.Index(), 0),
		Animation:      b.ReadInt(b.Bytes(), b.Index()),
		ESName:         b.ReadString(b.Bytes(), b.Index(), 0),
		SuccessMessage: b.ReadString(b.Bytes(), b.Index(), 0),
		ExhaustedImage: b.ReadInt(b.Bytes(), b.Index()),
		Health:         b.ReadLong(b.Bytes(), b.Index()),
		ItemReward:     make([]int32, 11), // TODO int to const
		ItemVal:        make([]int64, 11), // TODO int to const
		ItemLuck:       make([]int32, 11), // TODO int to const
		Int:            make([]int32, 3),  // TODO int to const
	}
	for i := range resource.ItemReward {
		resource.ItemReward[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range resource.ItemVal {
		resource.ItemVal[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	for i := range resource.ItemLuck {
		resource.ItemLuck[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	resource.ResourceImage = b.ReadInt(b.Bytes(), b.Index())
	resource.ResourceType = b.ReadByte(b.Bytes(), b.Index())
	resource.RespawnTime = b.ReadInt(b.Bytes(), b.Index())
	resource.ToolRequired = b.ReadByte(b.Bytes(), b.Index())
	resource.Walkthrough = b.ReadBool(b.Bytes(), b.Index())
	resource.NormalAnim = b.ReadBool(b.Bytes(), b.Index())
	resource.NormalRandom = b.ReadBool(b.Bytes(), b.Index())
	for i := range resource.Int {
		resource.Int[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	return resource
}

func ReadMoralData(b buffer.PacketBuffer) MoralRec {
	moral := MoralRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Int:  make([]int32, 2),
		Bool: make([]bool, 23),
	}
	for i := range moral.Bool {
		moral.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	for i := range moral.Int {
		moral.Int[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	return moral
}

func ReadCardData(b buffer.PacketBuffer) CardsRec {
	card := CardsRec{
		Name:   b.ReadString(b.Bytes(), b.Index(), 0),
		Desc:   b.ReadString(b.Bytes(), b.Index(), 0),
		Num:    b.ReadInt(b.Bytes(), b.Index()),
		Sprite: b.ReadInt(b.Bytes(), b.Index()),
		Card:   b.ReadInt(b.Bytes(), b.Index()),
		Offset: objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		},
		Icon: b.ReadInt(b.Bytes(), b.Index()),
	}
	return card
}

func ReadCraftData(b buffer.PacketBuffer) CraftingRec {
	craft := CraftingRec{
		Name:    b.ReadString(b.Bytes(), b.Index(), 0),
		Desc:    b.ReadString(b.Bytes(), b.Index(), 0),
		ItemReq: make([]int32, 3), // TODO int to const
		ItemVal: make([]int32, 3), // TODO int to const
	}
	for i := range craft.ItemReq {
		craft.ItemReq[i] = b.ReadInt(b.Bytes(), b.Index())
		craft.ItemVal[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	craft.Reward = b.ReadInt(b.Bytes(), b.Index())
	craft.RewardVal = b.ReadInt(b.Bytes(), b.Index())
	return craft
}

func ReadQuestData(b buffer.PacketBuffer) QuestRec {
	quest := QuestRec{
		Name:          b.ReadString(b.Bytes(), b.Index(), 0),
		ESName:        b.ReadString(b.Bytes(), b.Index(), 0),
		Repeat:        b.ReadBool(b.Bytes(), b.Index()),
		QuestLog:      b.ReadString(b.Bytes(), b.Index(), 0),
		Owner:         b.ReadInt(b.Bytes(), b.Index()),
		Type:          b.ReadInt(b.Bytes(), b.Index()),
		Status:        b.ReadBool(b.Bytes(), b.Index()),
		Item:          make([]QuestRewardItemRec, 10), // TODO int to const
		RequiredClass: make([]int32, 6),               // TODO int to const
		Task:          make([]TaskRec, 10),            // TODO int to const
		Story:         make([]string, 2),              // TODO int to const
		Bool:          make([]bool, 1),                // TODO int to const
	}
	for i := range quest.Item {
		quest.Item[i].Item = b.ReadInt(b.Bytes(), b.Index())
		quest.Item[i].Value = b.ReadLong(b.Bytes(), b.Index())
		quest.Item[i].Task = b.ReadInt(b.Bytes(), b.Index())
		quest.Item[i].Type = b.ReadInt(b.Bytes(), b.Index())
		quest.Item[i].Classes = b.ReadInt(b.Bytes(), b.Index())
	}
	quest.RequiredLevel = b.ReadInt(b.Bytes(), b.Index())
	quest.RequiredQuest = b.ReadInt(b.Bytes(), b.Index())
	quest.RequiredCompleteQuest = b.ReadBool(b.Bytes(), b.Index())
	for i := range quest.RequiredClass {
		quest.RequiredClass[i] = b.ReadInt(b.Bytes(), b.Index())
	}
	for i := range quest.Task {
		task := TaskRec{
			Amount:      b.ReadLong(b.Bytes(), b.Index()),
			Item:        b.ReadInt(b.Bytes(), b.Index()),
			Map:         b.ReadInt(b.Bytes(), b.Index()),
			NPC:         b.ReadInt(b.Bytes(), b.Index()),
			PartyQuest:  b.ReadInt(b.Bytes(), b.Index()),
			Order:       b.ReadInt(b.Bytes(), b.Index()),
			QuestEnd:    b.ReadBool(b.Bytes(), b.Index()),
			Resource:    b.ReadInt(b.Bytes(), b.Index()),
			Recipe:      b.ReadInt(b.Bytes(), b.Index()),
			Speech:      b.ReadString(b.Bytes(), b.Index(), 0),
			TaskLog:     b.ReadString(b.Bytes(), b.Index(), 0),
			RewardEXP:   b.ReadLong(b.Bytes(), b.Index()),
			SpellReward: make([]int32, 3),  // TODO int to const
			StartQuote:  make([]string, 2), // TODO int to const
			EndQuote:    make([]string, 2), // TODO int to const
			Sprite:      make([]int32, 2),  // TODO int to const
			Int:         make([]int32, 3),  // TODO int to const
		}
		for x := range task.SpellReward {
			task.SpellReward[x] = b.ReadInt(b.Bytes(), b.Index())
		}
		task.SpellUse = b.ReadInt(b.Bytes(), b.Index())
		task.Tutorial = b.ReadInt(b.Bytes(), b.Index())
		task.RewardRecipe = b.ReadInt(b.Bytes(), b.Index())
		task.RewardVariable = b.ReadInt(b.Bytes(), b.Index())
		task.RewardCalavera = b.ReadLong(b.Bytes(), b.Index())
		task.Image = b.ReadInt(b.Bytes(), b.Index())
		task.ImageVec = objects.Vector2{
			X: b.ReadInt(b.Bytes(), b.Index()),
			Y: b.ReadInt(b.Bytes(), b.Index()),
		}
		task.AutoComplete = b.ReadBool(b.Bytes(), b.Index())
		task.ElementReward = b.ReadInt(b.Bytes(), b.Index())
		for x := range task.StartQuote {
			task.StartQuote[x] = b.ReadString(b.Bytes(), b.Index(), 0)
			task.EndQuote[x] = b.ReadString(b.Bytes(), b.Index(), 0)
			task.Sprite[x] = b.ReadInt(b.Bytes(), b.Index())
		}
		for x := range task.Int {
			task.Int[x] = b.ReadInt(b.Bytes(), b.Index())
		}
		quest.Task[i] = task
	}
	for i := range quest.Story {
		quest.Story[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	quest.Promotion = b.ReadByte(b.Bytes(), b.Index())
	quest.QuestAtEnd = b.ReadInt(b.Bytes(), b.Index())
	for i := range quest.Bool {
		quest.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	return quest
}

func ReadConditionData(b buffer.PacketBuffer) ConditionRec {
	condition := ConditionRec{
		Name: b.ReadString(b.Bytes(), b.Index(), 0),
		Sprite: ConditionSpriteRec{
			pic:  b.ReadInt(b.Bytes(), b.Index()),
			Type: b.ReadInt(b.Bytes(), b.Index()),
			Mask: b.ReadInt(b.Bytes(), b.Index()),
			myFrame: objects.Vector2{
				X: b.ReadInt(b.Bytes(), b.Index()),
				Y: b.ReadInt(b.Bytes(), b.Index()),
			},
			Offset: objects.Vector2{
				X: b.ReadInt(b.Bytes(), b.Index()),
				Y: b.ReadInt(b.Bytes(), b.Index()),
			},
			Width:  b.ReadInt(b.Bytes(), b.Index()),
			Height: b.ReadInt(b.Bytes(), b.Index()),
		},
		ShowName: b.ReadBool(b.Bytes(), b.Index()),
		OnJoin:   b.ReadBool(b.Bytes(), b.Index()),
		MaxLines: b.ReadInt(b.Bytes(), b.Index()),
		Bool:     make([]bool, 1), // TODO int to const
	}
	condition.Line = make([]ActionRec, condition.MaxLines)
	for i := range condition.Line {
		action := ActionRec{
			Void:       b.ReadByte(b.Bytes(), b.Index()),
			NextLine:   b.ReadInt(b.Bytes(), b.Index()),
			ActionType: b.ReadByte(b.Bytes(), b.Index()),
			EndCode:    b.ReadBool(b.Bytes(), b.Index()),
			OrCode:     b.ReadBool(b.Bytes(), b.Index()),
			AndCode:    b.ReadBool(b.Bytes(), b.Index()),
			ElseCode:   b.ReadBool(b.Bytes(), b.Index()),
			isChoice:   b.ReadBool(b.Bytes(), b.Index()),
			Modifier:   b.ReadByte(b.Bytes(), b.Index()),
			ToEntity:   b.ReadByte(b.Bytes(), b.Index()),
			VarType:    make([]byte, 5),   // TODO int to const
			Var:        make([]string, 5), // TODO int to const
		}
		for x := range action.VarType {
			action.VarType[x] = b.ReadByte(b.Bytes(), b.Index())
			if Type := action.VarType[x]; Type > 0 && Type-1 <= 4 {
				action.Var[x] = b.ReadString(b.Bytes(), b.Index(), 0)
			}
		}
		condition.Line[i] = action
	}
	condition.MeetReq = b.ReadString(b.Bytes(), b.Index(), 0)
	condition.Faceset = b.ReadInt(b.Bytes(), b.Index())
	condition.Switch = b.ReadInt(b.Bytes(), b.Index())
	for i := range condition.Bool {
		condition.Bool[i] = b.ReadBool(b.Bytes(), b.Index())
	}
	return condition
}

func ReadEmoticonData(b buffer.PacketBuffer) EmoticonRec {
	emoticon := EmoticonRec{
		Name: make([]string, 3), // TODO int to const
	}
	for i := range emoticon.Name {
		emoticon.Name[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	emoticon.Command = b.ReadString(b.Bytes(), b.Index(), 0)
	emoticon.Image = b.ReadInt(b.Bytes(), b.Index())
	emoticon.IsDefault = b.ReadBool(b.Bytes(), b.Index())
	emoticon.Animated = b.ReadBool(b.Bytes(), b.Index())
	return emoticon
}

func ReadProfessionData(b buffer.PacketBuffer) []ProfessionRec {
	professions := make([]ProfessionRec, 4) // TODO int to const
	for i := range professions {
		profession := ProfessionRec{
			Icon:    b.ReadInt(b.Bytes(), b.Index()),
			Upgrade: make([]ProfessionUpgradeRec, 10), // TODO int to const
		}
		for x := range profession.Upgrade {
			profession.Upgrade[x].Name = make([]string, 3)
			profession.Upgrade[x].Desc = make([]string, 3)
			for f := range profession.Upgrade[x].Name {
				profession.Upgrade[x].Name[f] = b.ReadString(b.Bytes(), b.Index(), 0)
				profession.Upgrade[x].Desc[f] = b.ReadString(b.Bytes(), b.Index(), 0)
				profession.Upgrade[x].MaxLevel = b.ReadByte(b.Bytes(), b.Index())
				profession.Upgrade[x].Icon = b.ReadInt(b.Bytes(), b.Index())
			}
		}
		professions[i] = profession
	}
	return professions
}

func ReadNewspaperData(b buffer.PacketBuffer) NewspaperRec {
	news := NewspaperRec{
		Message: b.ReadString(b.Bytes(), b.Index(), 0),
		Url:     b.ReadString(b.Bytes(), b.Index(), 0),
		Poll:    make([]string, 4), // TODO int to const
	}
	if text := b.ReadString(b.Bytes(), b.Index(), 0); text != news.Url {
		news.Url = text
	}
	news.EventUrl = b.ReadString(b.Bytes(), b.Index(), 0)
	news.Title = b.ReadString(b.Bytes(), b.Index(), 0)
	news.ServerDay = b.ReadByte(b.Bytes(), b.Index())
	news.ServerMonth = b.ReadByte(b.Bytes(), b.Index())
	news.PollName = b.ReadString(b.Bytes(), b.Index(), 0)
	for i := range news.Poll {
		news.Poll[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	return news
}
