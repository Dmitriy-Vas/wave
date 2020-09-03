package main

import (
	"github.com/Dmitriy-Vas/wave"
	. "github.com/Dmitriy-Vas/wave/lib/packets/incoming"
	"github.com/Dmitriy-Vas/wave/lib/packets/outgoing"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"log"
	"net"
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", ":7999")
	remoteAddr, _ := net.ResolveTCPAddr("tcp", "74.91.123.86:7000")
	config := wave.Config{
		PacketLengthSize:   8,
		PacketTypeSize:     8,
		RemoteAddress:      remoteAddr,
		LocalAddress:       localAddr,
		ConnectImmediately: true,
		BufferType:         (*wrapper.Buffer)(nil),
		ReaderType:         (*wrapper.Reader)(nil),
		WriterType:         (*wrapper.Writer)(nil),
		PacketInit:         wave.InitPacket,
		BufferInit:         wrapper.InitBuffer,
		ReaderInit:         wrapper.InitReader,
		WriterInit:         wrapper.InitWriter,
	}
	proxy := wave.New(config)

	RegisterPackets(proxy)
	RegisterHooks(proxy)

	if err := proxy.Start(); err != nil {
		panic(err)
	}
	proxy.Close()
}

func RegisterPackets(proxy *wave.Proxy) {
	proxy.AddPacket(2, true, new(outgoing.LoginPacket))
	proxy.AddPacket(20, true, new(outgoing.ClientRevisionPacket))

	proxy.AddPacket(1, false, new(AlertMsgPacket))
	proxy.AddPacket(2, false, new(LoginOkPacket))
	proxy.AddPacket(4, false, new(InGamePacket))
	proxy.AddPacket(8, false, new(CoinsPacket))
	proxy.AddPacket(9, false, new(SkullsPacket))
	proxy.AddPacket(10, false, new(PlayerStatsPacket))
	proxy.AddPacket(14, false, new(PlayerDirPacket))
	proxy.AddPacket(15, false, new(NpcDirPacket))
	proxy.AddPacket(16, false, new(PlayerXYPacket))
	proxy.AddPacket(17, false, new(AttackPacket))
	proxy.AddPacket(18, false, new(NpcAttackPacket))
	proxy.AddPacket(19, false, new(CheckMapPacket))
	proxy.AddPacket(21, false, new(MapItemDataPacket))
	proxy.AddPacket(22, false, new(MapNPCDataPacket))
	proxy.AddPacket(23, false, new(SteamAchievementPacket))
	proxy.AddPacket(24, false, new(GlobalMessagePacket))
	proxy.AddPacket(26, false, new(MapMessagePacket))
	proxy.AddPacket(27, false, new(SpawnItemPacket))
	proxy.AddPacket(29, false, new(SpawnNPCPacket))
	proxy.AddPacket(30, false, new(NpcDeadPacket))
	proxy.AddPacket(34, false, new(SpellsPacket))
	proxy.AddPacket(35, false, new(PlayerTimerPacket))
	proxy.AddPacket(36, false, new(ResourceCachePacket))
	proxy.AddPacket(38, false, new(PingPacket))
	proxy.AddPacket(39, false, new(ActionMessagePacket))
	proxy.AddPacket(41, false, new(PlayerSpritePacket))
	proxy.AddPacket(43, false, new(AnimationPacket))
	proxy.AddPacket(45, false, new(CooldownPacket))
	proxy.AddPacket(46, false, new(ClearSpellBufferPacket))
	proxy.AddPacket(47, false, new(SayMessagePacket))
	proxy.AddPacket(50, false, new(StunnedPacket))
	proxy.AddPacket(54, false, new(PlayerCurrentPacket))
	proxy.AddPacket(55, false, new(ClosePacket))
	proxy.AddPacket(57, false, new(TradeStatusPacket))
	proxy.AddPacket(59, false, new(HighIndexPacket))
	proxy.AddPacket(60, false, new(TargetPacket))
	proxy.AddPacket(61, false, new(CharDataPacket))
	proxy.AddPacket(62, false, new(DefensaPacket))
	proxy.AddPacket(65, false, new(SendTimePacket))
	proxy.AddPacket(66, false, new(SolarEXPPacket))
	proxy.AddPacket(67, false, new(PlayerInvitePacket))
	proxy.AddPacket(68, false, new(PartyUpdatePacket))
	proxy.AddPacket(70, false, new(BossMsgPacket))
	proxy.AddPacket(71, false, new(PlayerXYMapPacket))
	proxy.AddPacket(72, false, new(ProjectilPacket))
	proxy.AddPacket(73, false, new(UpdateProjectilPacket))
	proxy.AddPacket(75, false, new(FlashPacket))
	proxy.AddPacket(76, false, new(RequestEditorPacket))
	proxy.AddPacket(78, false, new(PlayerCardsPacket))
	proxy.AddPacket(79, false, new(UpdateCraftingPacket))
	proxy.AddPacket(80, false, new(PlayerCraftingPacket))
	proxy.AddPacket(81, false, new(PlayerDeathPacket))
	proxy.AddPacket(82, false, new(UpdateTitlePacket))
	proxy.AddPacket(84, false, new(BuffInfoPacket))
	proxy.AddPacket(85, false, new(SpellAnimPacket))
	proxy.AddPacket(86, false, new(BattleMsgPacket))
	proxy.AddPacket(87, false, new(NPCProjectilPacket))
	proxy.AddPacket(88, false, new(MapSoundPacket))
	proxy.AddPacket(89, false, new(ArrowPacket))
	proxy.AddPacket(90, false, new(BanditPacket))
	proxy.AddPacket(92, false, new(PVPPacket))
	proxy.AddPacket(93, false, new(MapTopoDataPacket))
	proxy.AddPacket(94, false, new(TopoDrillPacket))
	proxy.AddPacket(96, false, new(ServerMsgPacket))
	proxy.AddPacket(97, false, new(DailyCheckPacket))
	proxy.AddPacket(98, false, new(SpawnMobItemPacket))
	proxy.AddPacket(99, false, new(ProjectilCrossPacket))
	proxy.AddPacket(100, false, new(ReceiveHourPacket))
	proxy.AddPacket(101, false, new(ReceiveBubblePacket))
	proxy.AddPacket(102, false, new(ChoicePacket))
	proxy.AddPacket(103, false, new(ClearChoicePacket))
	proxy.AddPacket(105, false, new(BreakPointPacket))
	proxy.AddPacket(107, false, new(SendConditionVarPacket))
	proxy.AddPacket(108, false, new(SendPlayerVarsPacket))
	proxy.AddPacket(109, false, new(QuestMsgPacket))
	proxy.AddPacket(112, false, new(UpdatePlayerListPacket))
	proxy.AddPacket(113, false, new(ComboPacket))
	proxy.AddPacket(114, false, new(TutorialPacket))
	proxy.AddPacket(115, false, new(PlayerElementPacket))
	proxy.AddPacket(117, false, new(EarthquakePacket))
	proxy.AddPacket(118, false, new(EventRankPacket))
	proxy.AddPacket(119, false, new(TravelInfoPacket))
	proxy.AddPacket(120, false, new(LoginBackOKPacket))
	proxy.AddPacket(121, false, new(PlayerAwardsPacket))
	proxy.AddPacket(122, false, new(AwardsPacket))
	proxy.AddPacket(124, false, new(OpenUIPacket))
	proxy.AddPacket(125, false, new(LogOutOKPacket))
	proxy.AddPacket(126, false, new(PlayerScorePacket))
	proxy.AddPacket(127, false, new(PartyQuestInfoPacket))
	proxy.AddPacket(128, false, new(DemoUpdateQuestPacket))
	proxy.AddPacket(129, false, new(CashInvUpdatePacket))
	proxy.AddPacket(130, false, new(PlayerInfoPacket))
	proxy.AddPacket(131, false, new(PlayerHonorPacket))
	proxy.AddPacket(135, false, new(DrillBreakPacket))
	proxy.AddPacket(136, false, new(ArenaBonusPacket))
	proxy.AddPacket(137, false, new(NpcBuffInfoPacket))
	proxy.AddPacket(138, false, new(AnnouncerPacket))
	proxy.AddPacket(139, false, new(OpenBackgroundPacket))
	proxy.AddPacket(140, false, new(CouponReadyPacket))
	proxy.AddPacket(141, false, new(ProfessionStartPacket))
	proxy.AddPacket(143, false, new(PlayerMasterPacket))
	proxy.AddPacket(144, false, new(UrlPacket))
	proxy.AddPacket(146, false, new(PlayerConnectedPacket))
	proxy.AddPacket(147, false, new(ServerStatsPacket))
	proxy.AddPacket(150, false, new(ShowEmoticonPacket))
	proxy.AddPacket(152, false, new(ServerVarsPacket))
	proxy.AddPacket(154, false, new(BlockListPacket))
	proxy.AddPacket(155, false, new(ConstantDataPacket))
}

func RegisterHooks(proxy *wave.Proxy) {
	proxy.HookPacket(20, true, func(conn *wave.Conn, packet wave.Packet) {
		p := packet.(*outgoing.ClientRevisionPacket)
		log.Printf("%+v", *p)
	})
	proxy.HookPacket(38, false, PingHook)
	proxy.HookPacket(46, true, PongHook)
}

func PingHook(conn *wave.Conn, packet wave.Packet) {
	// TODO send pong to the server
}

func PongHook(conn *wave.Conn, packet wave.Packet) {
	// TODO ignore pong from client
}
