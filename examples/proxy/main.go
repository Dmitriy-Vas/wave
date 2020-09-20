package main

import (
	"crypto/rc4"
	"encoding/binary"
	"encoding/hex"
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"github.com/Dmitriy-Vas/wave/lib/packets/incoming"
	"github.com/Dmitriy-Vas/wave/lib/packets/outgoing"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"log"
	"net"
	"sync"
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", ":7999")
	remoteAddr, _ := net.ResolveTCPAddr("tcp", "74.91.123.86:7000")
	buf := &buffer.DefaultBuffer{
		PacketReader: &buffer.DefaultReader{Order: binary.LittleEndian},
		PacketWriter: &buffer.DefaultWriter{Order: binary.LittleEndian},
	}
	buf.SetInitLength(wave.Size64Bits)
	buf.SetMaxLength(0xfffff)

	config := wave.Config{
		PacketLengthSize:   wave.Size64Bits,
		PacketTypeSize:     wave.Size64Bits,
		RemoteAddress:      remoteAddr,
		LocalAddress:       localAddr,
		ConnectImmediately: true,
		Buffer:             &wrapper.Buffer{DefaultBuffer: buf},
		//OutgoingProcess:    OutgoingProcess,
		//IncomingProcess:    IncomingProcess,
	}
	proxy := wave.New(config)

	RegisterPackets(proxy)
	RegisterHooks(proxy)

	if err := proxy.Start(); err != nil {
		panic(err)
	}
	proxy.Close()
}

func RegisterHooks(proxy *wave.Proxy) {
	// Notifying when new ping request sent
	proxy.HookPacket(int64(lib.OutGetPing), true, func(conn *wave.Conn, packet wave.Packet) {
		log.Printf("Requesting ping from the server")
	})
	// Notifying when ping response received
	proxy.HookPacket(int64(lib.IncPing), false, func(conn *wave.Conn, packet wave.Packet) {
		log.Printf("Ping received")
	})
	// Edit packet value on-fly
	proxy.HookPacket(int64(lib.IncReceiveHour), false, func(conn *wave.Conn, packet wave.Packet) {
		receiveHourPacket := packet.(*incoming.ReceiveHourPacket)
		log.Printf("Changing hour from %d to 12", receiveHourPacket.Hour)
		receiveHourPacket.Hour = 12
	})
	// Manually create and send packet when player died
	var once = sync.Once{}
	var buf buffer.PacketBuffer = nil
	proxy.HookPacket(int64(lib.IncPlayerDeath), false, func(conn *wave.Conn, packet wave.Packet) {
		once.Do(func() { buf = conn.Buffer().Clone() })
		p := &incoming.MapSoundPacket{
			ID:    int64(lib.IncMapSound),
			Send:  true,
			Num:   0,
			Sound: "buba_death.wav",
		}
		if err := conn.SendPacket(buf, p, false); err != nil {
			log.Printf("Error sending packet: %v", err)
		}
	})
	// Ignore packet
	proxy.HookPacket(int64(lib.OutLogOut), true, func(conn *wave.Conn, packet wave.Packet) {
		packet.SetSend(false)
	})
}

// Example of using Incoming process
func IncomingProcess(buffer buffer.PacketBuffer, _ bool) {
	raw, _ := hex.DecodeString("6a39570cc9de4ec71d64821894")
	cipher, _ := rc4.NewCipher(raw)
	cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
}

// Example of using Outgoing process
func OutgoingProcess(buffer buffer.PacketBuffer, _ bool) {
	raw, _ := hex.DecodeString("c79332b197f92ba85ed281a023")
	cipher, _ := rc4.NewCipher(raw)
	cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
}

func RegisterPackets(proxy *wave.Proxy) {
	proxy.AddPacket(1, true, (*outgoing.NewAccountPacket)(nil))
	proxy.AddPacket(2, true, (*outgoing.LoginPacket)(nil))
	proxy.AddPacket(3, true, (*outgoing.AddCharPacket)(nil))
	proxy.AddPacket(5, true, (*outgoing.PlayerMsgPacket)(nil))
	proxy.AddPacket(6, true, (*outgoing.ProfessionEndPacket)(nil))
	proxy.AddPacket(7, true, (*outgoing.DragItemPacket)(nil))
	proxy.AddPacket(8, true, (*outgoing.SayMsgPacket)(nil))
	proxy.AddPacket(10, true, (*outgoing.PlayerDirPacket)(nil))
	proxy.AddPacket(11, true, (*outgoing.UseItemPacket)(nil))
	proxy.AddPacket(13, true, (*outgoing.HotkeysPacket)(nil))
	proxy.AddPacket(15, true, (*outgoing.WarpMeToPacket)(nil))
	proxy.AddPacket(16, true, (*outgoing.WarpToMePacket)(nil))
	proxy.AddPacket(17, true, (*outgoing.AdminTaskPacket)(nil))
	proxy.AddPacket(19, true, (*outgoing.UseMegaphonePacket)(nil))
	proxy.AddPacket(20, true, (*outgoing.ClientRevisionPacket)(nil))
	proxy.AddPacket(24, true, (*outgoing.DropItemPacket)(nil))
	proxy.AddPacket(25, true, (*outgoing.MapRespawnPacket)(nil))
	proxy.AddPacket(26, true, (*outgoing.ReportPlayerPacket)(nil))
	proxy.AddPacket(28, true, (*outgoing.EnhancementPacket)(nil))
	proxy.AddPacket(27, true, (*outgoing.BanListPacket)(nil))
	proxy.AddPacket(29, true, (*outgoing.BanPacket)(nil))
	proxy.AddPacket(34, true, (*outgoing.SetAccessPacket)(nil))
	proxy.AddPacket(35, true, (*outgoing.WhosOnlinePacket)(nil))
	proxy.AddPacket(36, true, (*outgoing.TryOpenUIPacket)(nil))
	proxy.AddPacket(37, true, (*outgoing.PlayerSearchPacket)(nil))
	proxy.AddPacket(38, true, (*outgoing.GetStatsPacket)(nil))
	proxy.AddPacket(39, true, (*outgoing.PlayerCommandPacket)(nil))
	proxy.AddPacket(40, true, (*outgoing.ButtonPacket)(nil))
	proxy.AddPacket(43, true, (*outgoing.TakeCapturePacket)(nil))
	proxy.AddPacket(44, true, (*outgoing.SwapSlotsPacket)(nil))
	proxy.AddPacket(46, true, (*outgoing.GetPingPacket)(nil))
	proxy.AddPacket(47, true, (*outgoing.UnequipPacket)(nil))
	proxy.AddPacket(48, true, (*outgoing.SpawnItemPacket)(nil))
	proxy.AddPacket(49, true, (*outgoing.TrainStatPacket)(nil))
	proxy.AddPacket(51, true, (*outgoing.RequestLevelUpPacket)(nil))
	proxy.AddPacket(53, true, (*outgoing.CloseShopPacket)(nil))
	proxy.AddPacket(54, true, (*outgoing.BuyItemPacket)(nil))
	proxy.AddPacket(55, true, (*outgoing.SellItemPacket)(nil))
	proxy.AddPacket(56, true, (*outgoing.ChangeBankSlotsPacket)(nil))
	proxy.AddPacket(57, true, (*outgoing.DepositItemPacket)(nil))
	proxy.AddPacket(58, true, (*outgoing.WithdrawItemPacket)(nil))
	proxy.AddPacket(59, true, (*outgoing.CloseBankPacket)(nil))
	proxy.AddPacket(60, true, (*outgoing.AdminWarpPacket)(nil))
	proxy.AddPacket(61, true, (*outgoing.TradeRequestPacket)(nil))
	proxy.AddPacket(62, true, (*outgoing.AcceptTradePacket)(nil))
	proxy.AddPacket(63, true, (*outgoing.DeclineTradePacket)(nil))
	proxy.AddPacket(64, true, (*outgoing.TradeItemPacket)(nil))
	proxy.AddPacket(65, true, (*outgoing.UntradeItemPacket)(nil))
	proxy.AddPacket(66, true, (*outgoing.RequestUseCharPacket)(nil))
	proxy.AddPacket(67, true, (*outgoing.RequestNewCharPacket)(nil))
	proxy.AddPacket(68, true, (*outgoing.RequestDelCharPacket)(nil))
	proxy.AddPacket(69, true, (*outgoing.StopDefencePacket)(nil))
	proxy.AddPacket(71, true, (*outgoing.HotbarChangePacket)(nil))
	proxy.AddPacket(72, true, (*outgoing.HotbarUsePacket)(nil))
	proxy.AddPacket(75, true, (*outgoing.WhisperPacket)(nil))
	proxy.AddPacket(76, true, (*outgoing.PartyRequestPacket)(nil))
	proxy.AddPacket(77, true, (*outgoing.ReplyPlayerInvitationPacket)(nil))
	proxy.AddPacket(78, true, (*outgoing.DeclinePartyPacket)(nil))
	proxy.AddPacket(79, true, (*outgoing.PartyLeavePacket)(nil))
	proxy.AddPacket(85, true, (*outgoing.RequestDataPacket)(nil))
	proxy.AddPacket(87, true, (*outgoing.CraftingPacket)(nil))
	proxy.AddPacket(89, true, (*outgoing.DeathTypePacket)(nil))
	proxy.AddPacket(90, true, (*outgoing.TitlePacket)(nil))
	proxy.AddPacket(96, true, (*outgoing.ChoicePacket)(nil))
	proxy.AddPacket(97, true, (*outgoing.InteractionPacket)(nil))
	proxy.AddPacket(98, true, (*outgoing.GetAdminInfoPacket)(nil))
	proxy.AddPacket(99, true, (*outgoing.BreakPointPacket)(nil))
	proxy.AddPacket(102, true, (*outgoing.SaveComboPacket)(nil))
	proxy.AddPacket(103, true, (*outgoing.LoginBackPacket)(nil))
	proxy.AddPacket(104, true, (*outgoing.WarpMapPacket)(nil))
	proxy.AddPacket(105, true, (*outgoing.LogOutPacket)(nil))
	proxy.AddPacket(106, true, (*outgoing.UseCashItemPacket)(nil))
	proxy.AddPacket(107, true, (*outgoing.GetPlayerInfoPacket)(nil))
	proxy.AddPacket(108, true, (*outgoing.UpdatePlayerHonorPacket)(nil))
	proxy.AddPacket(109, true, (*outgoing.UIBuddiesPacket)(nil))
	proxy.AddPacket(110, true, (*outgoing.DropCalaverasPacket)(nil))
	proxy.AddPacket(111, true, (*outgoing.CouponPacket)(nil))
	proxy.AddPacket(112, true, (*outgoing.CashShopUIPacket)(nil))
	proxy.AddPacket(113, true, (*outgoing.FavAwardsPacket)(nil))
	proxy.AddPacket(114, true, (*outgoing.SelectMinigamesPacket)(nil))
	proxy.AddPacket(115, true, (*outgoing.DemoValuePacket)(nil))
	proxy.AddPacket(116, true, (*outgoing.ForgeActionPacket)(nil))
	proxy.AddPacket(117, true, (*outgoing.CancelPacket)(nil))

	proxy.AddPacket(1, false, (*incoming.AlertMsgPacket)(nil))
	proxy.AddPacket(2, false, (*incoming.LoginOkPacket)(nil))
	proxy.AddPacket(3, false, (*incoming.ClassesDataPacket)(nil))
	proxy.AddPacket(4, false, (*incoming.InGamePacket)(nil))
	//proxy.AddPacket(5, false, (*incoming.PlayerInvUpdatePacket)(nil))
	proxy.AddPacket(8, false, (*incoming.CoinsPacket)(nil))
	proxy.AddPacket(9, false, (*incoming.SkullsPacket)(nil))
	proxy.AddPacket(10, false, (*incoming.PlayerStatsPacket)(nil))
	proxy.AddPacket(14, false, (*incoming.PlayerDirPacket)(nil))
	proxy.AddPacket(15, false, (*incoming.NpcDirPacket)(nil))
	proxy.AddPacket(16, false, (*incoming.PlayerXYPacket)(nil))
	proxy.AddPacket(17, false, (*incoming.AttackPacket)(nil))
	proxy.AddPacket(18, false, (*incoming.NpcAttackPacket)(nil))
	proxy.AddPacket(19, false, (*incoming.CheckMapPacket)(nil))
	proxy.AddPacket(21, false, (*incoming.MapItemDataPacket)(nil))
	proxy.AddPacket(22, false, (*incoming.MapNPCDataPacket)(nil))
	proxy.AddPacket(23, false, (*incoming.SteamAchievementPacket)(nil))
	proxy.AddPacket(24, false, (*incoming.GlobalMessagePacket)(nil))
	proxy.AddPacket(26, false, (*incoming.MapMessagePacket)(nil))
	proxy.AddPacket(27, false, (*incoming.SpawnItemPacket)(nil))
	proxy.AddPacket(29, false, (*incoming.SpawnNPCPacket)(nil))
	proxy.AddPacket(30, false, (*incoming.NpcDeadPacket)(nil))
	proxy.AddPacket(34, false, (*incoming.SpellsPacket)(nil))
	proxy.AddPacket(35, false, (*incoming.PlayerTimerPacket)(nil))
	proxy.AddPacket(36, false, (*incoming.ResourceCachePacket)(nil))
	proxy.AddPacket(38, false, (*incoming.PingPacket)(nil))
	proxy.AddPacket(39, false, (*incoming.ActionMessagePacket)(nil))
	proxy.AddPacket(41, false, (*incoming.PlayerSpritePacket)(nil))
	proxy.AddPacket(43, false, (*incoming.AnimationPacket)(nil))
	proxy.AddPacket(45, false, (*incoming.CooldownPacket)(nil))
	proxy.AddPacket(46, false, (*incoming.ClearSpellBufferPacket)(nil))
	proxy.AddPacket(47, false, (*incoming.SayMessagePacket)(nil))
	proxy.AddPacket(50, false, (*incoming.StunnedPacket)(nil))
	proxy.AddPacket(54, false, (*incoming.PlayerCurrentPacket)(nil))
	proxy.AddPacket(55, false, (*incoming.ClosePacket)(nil))
	proxy.AddPacket(57, false, (*incoming.TradeStatusPacket)(nil))
	//proxy.AddPacket(58, false, (*incoming.GameDataPacket)(nil))
	proxy.AddPacket(59, false, (*incoming.HighIndexPacket)(nil))
	proxy.AddPacket(60, false, (*incoming.TargetPacket)(nil))
	proxy.AddPacket(61, false, (*incoming.CharDataPacket)(nil))
	proxy.AddPacket(62, false, (*incoming.DefensaPacket)(nil))
	proxy.AddPacket(65, false, (*incoming.SendTimePacket)(nil))
	proxy.AddPacket(66, false, (*incoming.SolarEXPPacket)(nil))
	proxy.AddPacket(67, false, (*incoming.PlayerInvitePacket)(nil))
	proxy.AddPacket(68, false, (*incoming.PartyUpdatePacket)(nil))
	proxy.AddPacket(70, false, (*incoming.BossMsgPacket)(nil))
	proxy.AddPacket(71, false, (*incoming.PlayerXYMapPacket)(nil))
	proxy.AddPacket(72, false, (*incoming.ProjectilPacket)(nil))
	proxy.AddPacket(73, false, (*incoming.UpdateProjectilPacket)(nil))
	proxy.AddPacket(75, false, (*incoming.FlashPacket)(nil))
	proxy.AddPacket(76, false, (*incoming.RequestEditorPacket)(nil))
	proxy.AddPacket(78, false, (*incoming.PlayerCardsPacket)(nil))
	proxy.AddPacket(79, false, (*incoming.UpdateCraftingPacket)(nil))
	proxy.AddPacket(80, false, (*incoming.PlayerCraftingPacket)(nil))
	proxy.AddPacket(81, false, (*incoming.PlayerDeathPacket)(nil))
	proxy.AddPacket(82, false, (*incoming.UpdateTitlePacket)(nil))
	proxy.AddPacket(84, false, (*incoming.BuffInfoPacket)(nil))
	proxy.AddPacket(85, false, (*incoming.SpellAnimPacket)(nil))
	proxy.AddPacket(86, false, (*incoming.BattleMsgPacket)(nil))
	proxy.AddPacket(87, false, (*incoming.NPCProjectilPacket)(nil))
	proxy.AddPacket(88, false, (*incoming.MapSoundPacket)(nil))
	proxy.AddPacket(89, false, (*incoming.ArrowPacket)(nil))
	proxy.AddPacket(90, false, (*incoming.BanditPacket)(nil))
	proxy.AddPacket(92, false, (*incoming.PVPPacket)(nil))
	proxy.AddPacket(93, false, (*incoming.MapTopoDataPacket)(nil))
	proxy.AddPacket(94, false, (*incoming.TopoDrillPacket)(nil))
	proxy.AddPacket(96, false, (*incoming.ServerMsgPacket)(nil))
	proxy.AddPacket(97, false, (*incoming.DailyCheckPacket)(nil))
	proxy.AddPacket(98, false, (*incoming.SpawnMobItemPacket)(nil))
	proxy.AddPacket(99, false, (*incoming.ProjectilCrossPacket)(nil))
	proxy.AddPacket(100, false, (*incoming.ReceiveHourPacket)(nil))
	proxy.AddPacket(101, false, (*incoming.ReceiveBubblePacket)(nil))
	proxy.AddPacket(102, false, (*incoming.ChoicePacket)(nil))
	proxy.AddPacket(103, false, (*incoming.ClearChoicePacket)(nil))
	proxy.AddPacket(105, false, (*incoming.BreakPointPacket)(nil))
	proxy.AddPacket(107, false, (*incoming.SendConditionVarPacket)(nil))
	proxy.AddPacket(108, false, (*incoming.SendPlayerVarsPacket)(nil))
	proxy.AddPacket(109, false, (*incoming.QuestMsgPacket)(nil))
	proxy.AddPacket(112, false, (*incoming.UpdatePlayerListPacket)(nil))
	proxy.AddPacket(113, false, (*incoming.ComboPacket)(nil))
	proxy.AddPacket(114, false, (*incoming.TutorialPacket)(nil))
	proxy.AddPacket(115, false, (*incoming.PlayerElementPacket)(nil))
	proxy.AddPacket(117, false, (*incoming.EarthquakePacket)(nil))
	proxy.AddPacket(118, false, (*incoming.EventRankPacket)(nil))
	proxy.AddPacket(119, false, (*incoming.TravelInfoPacket)(nil))
	proxy.AddPacket(120, false, (*incoming.LoginBackOKPacket)(nil))
	proxy.AddPacket(121, false, (*incoming.PlayerAwardsPacket)(nil))
	proxy.AddPacket(122, false, (*incoming.AwardsPacket)(nil))
	proxy.AddPacket(124, false, (*incoming.OpenUIPacket)(nil))
	proxy.AddPacket(125, false, (*incoming.LogOutOKPacket)(nil))
	proxy.AddPacket(126, false, (*incoming.PlayerScorePacket)(nil))
	proxy.AddPacket(127, false, (*incoming.PartyQuestInfoPacket)(nil))
	proxy.AddPacket(128, false, (*incoming.DemoUpdateQuestPacket)(nil))
	proxy.AddPacket(129, false, (*incoming.CashInvUpdatePacket)(nil))
	proxy.AddPacket(130, false, (*incoming.PlayerInfoPacket)(nil))
	proxy.AddPacket(131, false, (*incoming.PlayerHonorPacket)(nil))
	proxy.AddPacket(135, false, (*incoming.DrillBreakPacket)(nil))
	proxy.AddPacket(136, false, (*incoming.ArenaBonusPacket)(nil))
	proxy.AddPacket(137, false, (*incoming.NpcBuffInfoPacket)(nil))
	proxy.AddPacket(138, false, (*incoming.AnnouncerPacket)(nil))
	proxy.AddPacket(139, false, (*incoming.OpenBackgroundPacket)(nil))
	proxy.AddPacket(140, false, (*incoming.CouponReadyPacket)(nil))
	proxy.AddPacket(141, false, (*incoming.ProfessionStartPacket)(nil))
	proxy.AddPacket(143, false, (*incoming.PlayerMasterPacket)(nil))
	proxy.AddPacket(144, false, (*incoming.UrlPacket)(nil))
	proxy.AddPacket(146, false, (*incoming.PlayerConnectedPacket)(nil))
	proxy.AddPacket(147, false, (*incoming.ServerStatsPacket)(nil))
	proxy.AddPacket(150, false, (*incoming.ShowEmoticonPacket)(nil))
	proxy.AddPacket(152, false, (*incoming.ServerVarsPacket)(nil))
	proxy.AddPacket(153, false, (*incoming.CashShopDataPacket)(nil))
	proxy.AddPacket(154, false, (*incoming.BlockListPacket)(nil))
	proxy.AddPacket(155, false, (*incoming.ConstantDataPacket)(nil))
}
