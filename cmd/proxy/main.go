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
		PacketInit:         wave.InitPacket,
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
	proxy.HookPacket(int64(lib.IncPing), true, func(conn *wave.Conn, packet wave.Packet) {
		log.Printf("Ping received")
	})
	// Edit packet value on-fly
	proxy.HookPacket(int64(lib.IncReceiveHour), false, func(conn *wave.Conn, packet wave.Packet) {
		receiveHourPacket := packet.(*incoming.ReceiveHourPacket)
		log.Printf("Changing hour from %d to 12", receiveHourPacket.Hour)
		receiveHourPacket.Hour = 12
	})
	// Manually create and send packet when player died
	proxy.HookPacket(int64(lib.IncPlayerDeath), false, func(conn *wave.Conn, packet wave.Packet) {
		p := &incoming.MapSoundPacket{
			DefaultPacket: &wave.DefaultPacket{
				ID:   int64(lib.IncMapSound),
				Send: true,
			},
			Num:   0,
			Sound: "buba_death.wav",
		}
		if err := conn.SendPacket(p, false); err != nil {
			log.Printf("Error sending packet: %v", err)
		}
	})
	// Ignore packet
	proxy.HookPacket(int64(lib.OutLogOut), true, func(conn *wave.Conn, packet wave.Packet) {
		packet.SetSend(false)
	})
}

// Example of using Incoming process
func IncomingProcess(buffer buffer.PacketBuffer, start bool) {
	raw, _ := hex.DecodeString("6a39570cc9de4ec71d64821894")
	cipher, _ := rc4.NewCipher(raw)
	cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
}

// Example of using Outgoing process
func OutgoingProcess(buffer buffer.PacketBuffer, start bool) {
	raw, _ := hex.DecodeString("c79332b197f92ba85ed281a023")
	cipher, _ := rc4.NewCipher(raw)
	cipher.XORKeyStream(buffer.Bytes()[5:], buffer.Bytes()[5:])
}

func RegisterPackets(proxy *wave.Proxy) {
	proxy.AddPacket(1, true, new(outgoing.NewAccountPacket))
	proxy.AddPacket(2, true, new(outgoing.LoginPacket))
	proxy.AddPacket(3, true, new(outgoing.AddCharPacket))
	proxy.AddPacket(5, true, new(outgoing.PlayerMsgPacket))
	proxy.AddPacket(6, true, new(outgoing.ProfessionEndPacket))
	proxy.AddPacket(7, true, new(outgoing.DragItemPacket))
	proxy.AddPacket(8, true, new(outgoing.SayMsgPacket))
	proxy.AddPacket(10, true, new(outgoing.PlayerDirPacket))
	proxy.AddPacket(11, true, new(outgoing.UseItemPacket))
	proxy.AddPacket(13, true, new(outgoing.HotkeysPacket))
	proxy.AddPacket(15, true, new(outgoing.WarpMeToPacket))
	proxy.AddPacket(16, true, new(outgoing.WarpToMePacket))
	proxy.AddPacket(17, true, new(outgoing.AdminTaskPacket))
	proxy.AddPacket(19, true, new(outgoing.UseMegaphonePacket))
	proxy.AddPacket(20, true, new(outgoing.ClientRevisionPacket))
	proxy.AddPacket(24, true, new(outgoing.DropItemPacket))
	proxy.AddPacket(25, true, new(outgoing.MapRespawnPacket))
	proxy.AddPacket(26, true, new(outgoing.ReportPlayerPacket))
	proxy.AddPacket(28, true, new(outgoing.EnhancementPacket))
	proxy.AddPacket(27, true, new(outgoing.BanListPacket))
	proxy.AddPacket(29, true, new(outgoing.BanPacket))
	proxy.AddPacket(34, true, new(outgoing.SetAccessPacket))
	proxy.AddPacket(35, true, new(outgoing.WhosOnlinePacket))
	proxy.AddPacket(36, true, new(outgoing.TryOpenUIPacket))
	proxy.AddPacket(37, true, new(outgoing.PlayerSearchPacket))
	proxy.AddPacket(38, true, new(outgoing.GetStatsPacket))
	proxy.AddPacket(39, true, new(outgoing.PlayerCommandPacket))
	proxy.AddPacket(40, true, new(outgoing.ButtonPacket))
	proxy.AddPacket(43, true, new(outgoing.TakeCapturePacket))
	proxy.AddPacket(44, true, new(outgoing.SwapSlotsPacket))
	proxy.AddPacket(46, true, new(outgoing.GetPingPacket))
	proxy.AddPacket(47, true, new(outgoing.UnequipPacket))
	proxy.AddPacket(48, true, new(outgoing.SpawnItemPacket))
	proxy.AddPacket(49, true, new(outgoing.TrainStatPacket))
	proxy.AddPacket(51, true, new(outgoing.RequestLevelUpPacket))
	proxy.AddPacket(53, true, new(outgoing.CloseShopPacket))
	proxy.AddPacket(54, true, new(outgoing.BuyItemPacket))
	proxy.AddPacket(55, true, new(outgoing.SellItemPacket))
	proxy.AddPacket(56, true, new(outgoing.ChangeBankSlotsPacket))
	proxy.AddPacket(57, true, new(outgoing.DepositItemPacket))
	proxy.AddPacket(58, true, new(outgoing.WithdrawItemPacket))
	proxy.AddPacket(59, true, new(outgoing.CloseBankPacket))
	proxy.AddPacket(60, true, new(outgoing.AdminWarpPacket))
	proxy.AddPacket(61, true, new(outgoing.TradeRequestPacket))
	proxy.AddPacket(62, true, new(outgoing.AcceptTradePacket))
	proxy.AddPacket(63, true, new(outgoing.DeclineTradePacket))
	proxy.AddPacket(64, true, new(outgoing.TradeItemPacket))
	proxy.AddPacket(65, true, new(outgoing.UntradeItemPacket))
	proxy.AddPacket(66, true, new(outgoing.RequestUseCharPacket))
	proxy.AddPacket(67, true, new(outgoing.RequestNewCharPacket))
	proxy.AddPacket(68, true, new(outgoing.RequestDelCharPacket))
	proxy.AddPacket(69, true, new(outgoing.StopDefencePacket))
	proxy.AddPacket(71, true, new(outgoing.HotbarChangePacket))
	proxy.AddPacket(72, true, new(outgoing.HotbarUsePacket))
	proxy.AddPacket(75, true, new(outgoing.WhisperPacket))
	proxy.AddPacket(76, true, new(outgoing.PartyRequestPacket))
	proxy.AddPacket(77, true, new(outgoing.ReplyPlayerInvitationPacket))
	proxy.AddPacket(78, true, new(outgoing.DeclinePartyPacket))
	proxy.AddPacket(79, true, new(outgoing.PartyLeavePacket))
	proxy.AddPacket(85, true, new(outgoing.RequestDataPacket))
	proxy.AddPacket(87, true, new(outgoing.CraftingPacket))
	proxy.AddPacket(89, true, new(outgoing.DeathTypePacket))
	proxy.AddPacket(90, true, new(outgoing.TitlePacket))
	proxy.AddPacket(96, true, new(outgoing.ChoicePacket))
	proxy.AddPacket(97, true, new(outgoing.InteractionPacket))
	proxy.AddPacket(98, true, new(outgoing.GetAdminInfoPacket))
	proxy.AddPacket(99, true, new(outgoing.BreakPointPacket))
	proxy.AddPacket(102, true, new(outgoing.SaveComboPacket))
	proxy.AddPacket(103, true, new(outgoing.LoginBackPacket))
	proxy.AddPacket(104, true, new(outgoing.WarpMapPacket))
	proxy.AddPacket(105, true, new(outgoing.LogOutPacket))
	proxy.AddPacket(106, true, new(outgoing.UseCashItemPacket))
	proxy.AddPacket(107, true, new(outgoing.GetPlayerInfoPacket))
	proxy.AddPacket(108, true, new(outgoing.UpdatePlayerHonorPacket))
	proxy.AddPacket(109, true, new(outgoing.UIBuddiesPacket))
	proxy.AddPacket(110, true, new(outgoing.DropCalaverasPacket))
	proxy.AddPacket(111, true, new(outgoing.CouponPacket))
	proxy.AddPacket(112, true, new(outgoing.CashShopUIPacket))
	proxy.AddPacket(113, true, new(outgoing.FavAwardsPacket))
	proxy.AddPacket(114, true, new(outgoing.SelectMinigamesPacket))
	proxy.AddPacket(115, true, new(outgoing.DemoValuePacket))
	proxy.AddPacket(116, true, new(outgoing.ForgeActionPacket))
	proxy.AddPacket(117, true, new(outgoing.CancelPacket))

	proxy.AddPacket(1, false, new(incoming.AlertMsgPacket))
	proxy.AddPacket(2, false, new(incoming.LoginOkPacket))
	proxy.AddPacket(4, false, new(incoming.InGamePacket))
	proxy.AddPacket(8, false, new(incoming.CoinsPacket))
	proxy.AddPacket(9, false, new(incoming.SkullsPacket))
	proxy.AddPacket(10, false, new(incoming.PlayerStatsPacket))
	proxy.AddPacket(14, false, new(incoming.PlayerDirPacket))
	proxy.AddPacket(15, false, new(incoming.NpcDirPacket))
	proxy.AddPacket(16, false, new(incoming.PlayerXYPacket))
	proxy.AddPacket(17, false, new(incoming.AttackPacket))
	proxy.AddPacket(18, false, new(incoming.NpcAttackPacket))
	proxy.AddPacket(19, false, new(incoming.CheckMapPacket))
	proxy.AddPacket(21, false, new(incoming.MapItemDataPacket))
	proxy.AddPacket(22, false, new(incoming.MapNPCDataPacket))
	proxy.AddPacket(23, false, new(incoming.SteamAchievementPacket))
	proxy.AddPacket(24, false, new(incoming.GlobalMessagePacket))
	proxy.AddPacket(26, false, new(incoming.MapMessagePacket))
	proxy.AddPacket(27, false, new(incoming.SpawnItemPacket))
	proxy.AddPacket(29, false, new(incoming.SpawnNPCPacket))
	proxy.AddPacket(30, false, new(incoming.NpcDeadPacket))
	proxy.AddPacket(34, false, new(incoming.SpellsPacket))
	proxy.AddPacket(35, false, new(incoming.PlayerTimerPacket))
	proxy.AddPacket(36, false, new(incoming.ResourceCachePacket))
	proxy.AddPacket(38, false, new(incoming.PingPacket))
	proxy.AddPacket(39, false, new(incoming.ActionMessagePacket))
	proxy.AddPacket(41, false, new(incoming.PlayerSpritePacket))
	proxy.AddPacket(43, false, new(incoming.AnimationPacket))
	proxy.AddPacket(45, false, new(incoming.CooldownPacket))
	proxy.AddPacket(46, false, new(incoming.ClearSpellBufferPacket))
	proxy.AddPacket(47, false, new(incoming.SayMessagePacket))
	proxy.AddPacket(50, false, new(incoming.StunnedPacket))
	proxy.AddPacket(54, false, new(incoming.PlayerCurrentPacket))
	proxy.AddPacket(55, false, new(incoming.ClosePacket))
	proxy.AddPacket(57, false, new(incoming.TradeStatusPacket))
	proxy.AddPacket(59, false, new(incoming.HighIndexPacket))
	proxy.AddPacket(60, false, new(incoming.TargetPacket))
	proxy.AddPacket(61, false, new(incoming.CharDataPacket))
	proxy.AddPacket(62, false, new(incoming.DefensaPacket))
	proxy.AddPacket(65, false, new(incoming.SendTimePacket))
	proxy.AddPacket(66, false, new(incoming.SolarEXPPacket))
	proxy.AddPacket(67, false, new(incoming.PlayerInvitePacket))
	proxy.AddPacket(68, false, new(incoming.PartyUpdatePacket))
	proxy.AddPacket(70, false, new(incoming.BossMsgPacket))
	proxy.AddPacket(71, false, new(incoming.PlayerXYMapPacket))
	proxy.AddPacket(72, false, new(incoming.ProjectilPacket))
	proxy.AddPacket(73, false, new(incoming.UpdateProjectilPacket))
	proxy.AddPacket(75, false, new(incoming.FlashPacket))
	proxy.AddPacket(76, false, new(incoming.RequestEditorPacket))
	proxy.AddPacket(78, false, new(incoming.PlayerCardsPacket))
	proxy.AddPacket(79, false, new(incoming.UpdateCraftingPacket))
	proxy.AddPacket(80, false, new(incoming.PlayerCraftingPacket))
	proxy.AddPacket(81, false, new(incoming.PlayerDeathPacket))
	proxy.AddPacket(82, false, new(incoming.UpdateTitlePacket))
	proxy.AddPacket(84, false, new(incoming.BuffInfoPacket))
	proxy.AddPacket(85, false, new(incoming.SpellAnimPacket))
	proxy.AddPacket(86, false, new(incoming.BattleMsgPacket))
	proxy.AddPacket(87, false, new(incoming.NPCProjectilPacket))
	proxy.AddPacket(88, false, new(incoming.MapSoundPacket))
	proxy.AddPacket(89, false, new(incoming.ArrowPacket))
	proxy.AddPacket(90, false, new(incoming.BanditPacket))
	proxy.AddPacket(92, false, new(incoming.PVPPacket))
	proxy.AddPacket(93, false, new(incoming.MapTopoDataPacket))
	proxy.AddPacket(94, false, new(incoming.TopoDrillPacket))
	proxy.AddPacket(96, false, new(incoming.ServerMsgPacket))
	proxy.AddPacket(97, false, new(incoming.DailyCheckPacket))
	proxy.AddPacket(98, false, new(incoming.SpawnMobItemPacket))
	proxy.AddPacket(99, false, new(incoming.ProjectilCrossPacket))
	proxy.AddPacket(100, false, new(incoming.ReceiveHourPacket))
	proxy.AddPacket(101, false, new(incoming.ReceiveBubblePacket))
	proxy.AddPacket(102, false, new(incoming.ChoicePacket))
	proxy.AddPacket(103, false, new(incoming.ClearChoicePacket))
	proxy.AddPacket(105, false, new(incoming.BreakPointPacket))
	proxy.AddPacket(107, false, new(incoming.SendConditionVarPacket))
	proxy.AddPacket(108, false, new(incoming.SendPlayerVarsPacket))
	proxy.AddPacket(109, false, new(incoming.QuestMsgPacket))
	proxy.AddPacket(112, false, new(incoming.UpdatePlayerListPacket))
	proxy.AddPacket(113, false, new(incoming.ComboPacket))
	proxy.AddPacket(114, false, new(incoming.TutorialPacket))
	proxy.AddPacket(115, false, new(incoming.PlayerElementPacket))
	proxy.AddPacket(117, false, new(incoming.EarthquakePacket))
	proxy.AddPacket(118, false, new(incoming.EventRankPacket))
	proxy.AddPacket(119, false, new(incoming.TravelInfoPacket))
	proxy.AddPacket(120, false, new(incoming.LoginBackOKPacket))
	proxy.AddPacket(121, false, new(incoming.PlayerAwardsPacket))
	proxy.AddPacket(122, false, new(incoming.AwardsPacket))
	proxy.AddPacket(124, false, new(incoming.OpenUIPacket))
	proxy.AddPacket(125, false, new(incoming.LogOutOKPacket))
	proxy.AddPacket(126, false, new(incoming.PlayerScorePacket))
	proxy.AddPacket(127, false, new(incoming.PartyQuestInfoPacket))
	proxy.AddPacket(128, false, new(incoming.DemoUpdateQuestPacket))
	proxy.AddPacket(129, false, new(incoming.CashInvUpdatePacket))
	proxy.AddPacket(130, false, new(incoming.PlayerInfoPacket))
	proxy.AddPacket(131, false, new(incoming.PlayerHonorPacket))
	proxy.AddPacket(135, false, new(incoming.DrillBreakPacket))
	proxy.AddPacket(136, false, new(incoming.ArenaBonusPacket))
	proxy.AddPacket(137, false, new(incoming.NpcBuffInfoPacket))
	proxy.AddPacket(138, false, new(incoming.AnnouncerPacket))
	proxy.AddPacket(139, false, new(incoming.OpenBackgroundPacket))
	proxy.AddPacket(140, false, new(incoming.CouponReadyPacket))
	proxy.AddPacket(141, false, new(incoming.ProfessionStartPacket))
	proxy.AddPacket(143, false, new(incoming.PlayerMasterPacket))
	proxy.AddPacket(144, false, new(incoming.UrlPacket))
	proxy.AddPacket(146, false, new(incoming.PlayerConnectedPacket))
	proxy.AddPacket(147, false, new(incoming.ServerStatsPacket))
	proxy.AddPacket(150, false, new(incoming.ShowEmoticonPacket))
	proxy.AddPacket(152, false, new(incoming.ServerVarsPacket))
	proxy.AddPacket(154, false, new(incoming.BlockListPacket))
	proxy.AddPacket(155, false, new(incoming.ConstantDataPacket))
}
