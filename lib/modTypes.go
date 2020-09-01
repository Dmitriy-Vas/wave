package lib

import (
	"github.com/Dmitriy-Vas/wave/lib/objects"
	"time"
)

type EmoticonRec struct {
	Name      []string
	Command   string
	Image     int
	IsDefault bool
	Animated  bool
	Unlocked  bool
}

type MapListRec struct {
	Name []string
}

type PlayerBuffsRec struct {
	Type     byte
	SpellNum int
	Timer    int
	Vital    int
}

type PlayerBuddiesRec struct {
	Name       string
	Sex        byte
	Classes    byte
	Sprite     int
	Hair       int
	FriendDate string
	Paperdoll  string
	HairTint   string
	Online     bool
}

type WorldRec struct {
	Name    string
	Sprite  int
	sizeX   int
	sizeY   int
	offsetX int
	offsetY int
	Files   byte
	Column  byte
	Map     []WorldMapRec
}

type WorldMapRec struct {
	Num    int
	Status byte
}

type FontDataRec struct {
	Name       string
	Font       []byte
	Sprite     int
	Separation int
	Height     int
}

type ServerOptionRec struct {
	Mode      []ServerModRec
	EventItem int
	Demo      bool
	Clone     bool
	Promotion PromotionRec
}

type PromotionRec struct {
	Type     byte
	Item     int
	GameCard []PromotionGameCardRec
}

type PromotionGameCardRec struct {
	Name  string
	Price string
	Bonus int
}

type ServerModRec struct {
	Value float64
	Timer int64
}

type ServerModeRec struct {
	Experience float64
	Drill      float64
	Drop       float64
	Damage     float64
	Heal       float64
	Junk       float64
	LostExp    float64
	Party      float64
	Profession float64
}

type PlayerCardRec struct {
	Level int
	Exp   int
}

type AwardsRec struct {
	Name       string
	ESName     string
	Desc       string
	Type       byte
	Level      []AwardsLevelRec
	LevelCount int
	Blocked    bool
}

type AwardsLevelRec struct {
	Required  int
	Icon      int
	AP        int
	Calaveras int
}

type PlayerProfessionRec struct {
	Level   byte
	Upgrade []byte
	Points  byte
	EXP     int
	Times   int64
}

type ProfessionRec struct {
	Icon    int
	Upgrade []ProfessionUpgradeRec
}

type ProfessionUpgradeRec struct {
	Name     []string
	Desc     []string
	MaxLevel byte
	Icon     int
}

type PlayerAwardsRec struct {
	Level   int
	Count   int
	GetDate string
}

type RegionRec struct {
	Name     []string
	Music    RegionMusicRec
	Fog      RegionFogRec
	Sunlight bool
	Night    RegionNightRec
}

type RegionNightRec struct {
	Alpha int
}

type RegionFogRec struct {
	Sprite int
	Alpha  byte
	Speed  byte
}

type RegionMusicRec struct {
	Name       string
	LoopState  bool
	StartPoint uint
	EndPoint   uint
}

type EventsRec struct {
	Name    []string
	Players []EventsRankRec
}

type EventsRankRec struct {
	Name    string
	Classes int
	Score   int
}

type OptionsInterfaceRec struct {
	Pos []objects.Vector2
}

type OptionsQuestLogRec struct {
	Player []OptionsPlayerQuestLogRec
}

type OptionsPlayerQuestLogRec struct {
	Name      string
	QuestList []int
}

type PlayerElementRec struct {
	Num  int
	Icon int
}

type PlayerComboRec struct {
	Num int
}

type MyMusicRec struct {
	Name   string
	Format string
	File   []byte
}

type createCharRec struct {
	ClassSelect    int
	SexSelect      byte
	HairSelect     int
	SpriteSelect   int
	HairTintSelect int
	Sprite         []createCharSpriteRec
	Hair           []createCharSpriteRec
	HairTint       []objects.Color
}

type createCharSpriteRec struct {
	Male   int
	Female int
	MFace  int
	FFace  int
}

type PlayPuzzleRec struct {
	Data   PuzzleRec
	Status int
}

type PuzzleRec struct {
	Name        string
	Map         int
	Point       objects.Vector2
	Size        objects.Vector2
	SwitchImage int
	SwitchPos   objects.Vector2
	Cube        []PuzzleCubeRec
}

type PuzzleCubeRec struct {
	Image        int
	Block        bool
	Move         bool
	Key          bool
	Switch       int
	isPlayerMove int
}

type AdminInfoRec struct {
	Map          []string
	Level        []int64
	PlayersOfMap []int
	Puzzle       []string
	PartyQuest   []string
	Elements     []string
	Update       bool
}

type ControlRec struct {
	ID int
}

type TextureBox struct {
	Sprite int
	Width  int
	Height int
	pos    objects.Vector2
	Color  objects.Color
}

type ChoiceRec struct {
	myChoice      int
	choiceText    int
	choiceSelText int
	genDialog     int
	Sprite        int
	DialogType    int
	choiceDialog  string
	Max           int
	ChoiceArray   []string
	Code          bool
	isCondition   bool
	ConditionLine objects.Vector2
}

type DrillModRec struct {
	DrillDepth int
	Active     bool
	Pause      bool
	SoundIndex int
	Dir        byte
	DirStar    int
}

type ConnectionRec struct {
	IP   string
	Port int
}

type OptionsRec struct {
	Username         string
	Music            bool
	Sound            bool
	VSync            bool
	FullScreen       bool
	Debug            bool
	CheckLang        bool
	WindowsCursor    bool
	MapClean         bool
	SaveID           bool
	Zoom             bool
	Lang             string
	Resolution       byte
	Launcher         byte
	MVolume          byte
	SVolume          byte
	ScreenShotFormat string
	Game             GameOptionsRec
	UI               UIRec
	MyJoystick       JoystickRec
}

type GameOptionsRec struct {
	PlayerName bool
	BadWords   bool
}

type JoystickRec struct {
	Active bool
	Analog bool
	ID     byte
}

type UIRec struct {
	BigEmoticon   bool
	Cooldown      bool
	ExpBarPercent bool
}

type newRECT struct {
	top    int
	left   int
	right  int
	bottom int
}

type PartyRec struct {
	Leader      int
	Member      []MemberPartyRec
	MemberCount byte
	Type        byte
	Level       int
	Experience  int64
	PartyQuest  InPartyQuestRec
	shareItems  bool
}

type InPartyQuestRec struct {
	Time int
}

type MemberPartyRec struct {
	Index    int
	Level    int
	Hair     int
	Sprite   int
	Classes  int
	Map      int
	Name     string
	HairTint string
}

type InvItemRec struct {
	Num   int
	Value int64
	Slot  byte
	Stat  []byte
}

type BankRec struct {
	Slot byte
	Item []InvItemRec
}

type PlayerRec struct {
	Name             string
	Classes          byte
	Sprite           int
	Sex              byte
	Access           byte
	Level            int
	EXP              int64
	PK               byte
	Honor            int
	HonorDate        time.Time
	Vital            []int64
	Stat             []int64
	POINTS           int
	Equipment        []InvItemRec
	CashEquipment    []InvItemRec
	Pin              int
	Map              int
	Pos              objects.Vector2
	Dir              byte
	MaxVital         []int64
	WalkingTimer     int64
	Offset           objects.Vector2
	ScreenPos        objects.Vector2
	Moving           bool
	Attacking        byte
	AttackTimer      int64
	MapGetTimer      int64
	Steps            byte
	Hair             int64
	HairTint         objects.Color
	Defensa          byte
	DefensaBool      bool
	shieldHP         int64
	CancelDefensa    byte
	SpellBuffer      PlayerSpellBufferRec
	ProjecTil        []ProjectilRec
	StartFlash       int64
	Flash            byte
	Hide             int64
	CurCard          int64
	AmountOfCrafting int64
	Ghost            byte
	AmountOfTitle    int64
	Title            []int
	curTitle         int64
	Language         int64
	AmountOfQuests   int
	DrillMod         DrillModRec
	Alpha            int64
	Temp             PlayerTempRec
	Reflection       ReflectionRec
	Interaction      int64
	PuzzleOffset     objects.Vector2
	pComboCount      int
	ElementSelect    []int
	Light            int
	AP               int
	moveSpeed        float64
	elapsedTime      int64
	AutoTarget       objects.Vector2
	IsFishing        int
	Polymorph        int
	DamageSkin       byte
	ProjectilSkin    byte
	BuffsEffect      PlayerBuffEffectRec
	serverDelay      int
	Trap             PlayerTrapRec
	Emoticon         PlayerEmoticonRec
	Trapped          TrappedRec
}

type PlayerEmoticonRec struct {
	Num   int
	Timer int64
	isPin bool
}

type PlayerTrapRec struct {
	Spell    int
	X        int
	Y        int
	Duration int64
}

type PlayerBuffEffectRec struct {
	Poisoned bool
	Burned   bool
}

type PlayerTempRec struct {
	Invisible    bool
	PickUpItem   objects.Vector2
	PickUpTimer  int64
	PickUpPos    objects.Vector2
	MasterSprite int
	MasterTimer  int64
	MasterPos    objects.Vector2
	Mount        int
}

type PlayerSpellBufferRec struct {
	SpellAnimNum int
	SpellAnim    int
	SpellTimer   int64
	SpellBoom    bool
}

type ReflectionRec struct {
	pos objects.Vector2
	//rec  Rectangle
	Img  int
	Type int
}

type TileRec struct {
	Type      byte
	Data1     int
	Data2     int
	Data3     int
	Data4     string
	DirBlock  byte
	Topo      int
	DataTopo1 int
	DataTopo2 int
}

type MapRec struct {
	ID              int
	Name            string
	ESName          string
	Region          int
	Image           int
	Revision        int
	Moral           int
	Up              int
	Down            int
	Left            int
	Right           int
	BootMap         int
	BootX           int
	BootY           int
	MaxX            byte
	MaxY            byte
	Tile            []TileRec
	tileOffset      objects.Vector2
	Npc             []int
	FogOpacity      byte
	Fog             int
	FogSpeed        int
	tileA           int
	tileB           int
	DayNight        int
	Parallax        int
	ParallaxType    byte
	Ambience        string
	Lighting        bool
	NightAlpha      int
	InstanceMax     int
	MapCondition    int
	Puzzle          objects.Vector4
	Drill           objects.Vector4
	Temp            MapTempRec
	BackgroundColor string
}

type MapTempRec struct {
	EventCount int
	Events     []MapEventRec
}

type MapEventRec struct {
	Sprite  int
	Mode    byte
	Pos     objects.Vector2
	myFrame objects.Vector2
	Name    string
	OnJoin  bool
	Mask    int
	Width   int
	Height  int
	Offset  objects.Vector2
}

type spriteData struct {
	Flash bool
	Hide  bool
}

type ClassRec struct {
	Name     string
	Stat     []byte
	Sprite   []ClassSpriteRec
	Hair     []ClassSpriteRec
	HairTint []ClassHairTintRec
	Vital    []int64
}

type ClassHairTintRec struct {
	Name    string
	Color   objects.Color
	Premium bool
}

type ClassSpriteRec struct {
	Name    string
	Male    int
	Female  int
	MFace   int
	FFace   int
	Premium bool
}

type ItemRec struct {
	Name          string
	ESName        string
	Pic           int
	Revision      int
	Type          byte
	Data1         int
	Data2         int
	Tool          int
	ClassReq      []int64
	AccessReq     int64
	LevelReq      int64
	Mastery       byte
	Price         int64
	AddStat       []int
	Rarity        byte
	Speed         int64
	Handed        int64
	BindType      byte
	StatReq       []int
	Animation     int64
	Paperdoll     int64
	Desc          string
	Stackable     bool
	Overol        int
	paperWidth    int
	ProjecTil     int
	MType         int
	HPBonus       int
	MPBonus       int
	WarpMap       int
	WarpX         int
	WarpY         int
	isHair        int
	Element       []int
	ElementChance []int
	BigPic        int
	VitalMode     []int
	comboSlot     int64
	Recipe        int
	Int           []int
	Bool          []bool
}

type MapItemRec struct {
	Item   InvItemRec
	Slot   byte
	Stat   []byte
	Frame  byte
	pos    objects.Vector2
	X2     int64
	Y2     int64
	XTmr   int64
	Timer  int64
	Up     bool
	YTmr   int64
	Jump   byte
	pIndex int
	pTimer int64
	Angle  int
}

type NpcRec struct {
	Name          string
	AttackSay     string
	Behaviour     byte
	Range         byte
	DropChance    []int
	DropItem      []int
	DropItemValue []int
	Int           []int
	Vec           []objects.Vector2
	Faction       byte
	CardDrop      int
	CardNum       int
	Spell         []int
	Element       []int64
	Bool          []bool
	RequireVar    int
	AttackSound   string
	Light         int
	QuestList     []int
	ScriptList    []NpcScriptListRec
	Offset        objects.Vector2
}

type NpcScriptListRec struct {
	Num      int
	Interval int
	Variable string
}

type MapNpcRec struct {
	Num         int
	Target      int
	targetType  byte
	Vital       []int64
	Map         int
	Pos         objects.Vector2
	Dir         byte
	HPSetTo     int64
	Offset      MapNpcOffsetRec
	XOffset     int
	YOffset     int
	Moving      int
	Attacking   int
	AttackTimer int64
	Steps       int64
	AttackNum   int64
	AnimTimer   int
	ProjecTil   []NPCProjectilRec
	AutoTarget  objects.Vector2
	Tile        []NpcTileRec
	questIcon   int
	Buffs       []int
	BuffsEffect NpcBuffEffectRec
	Trapped     TrappedRec
	Paperdoll   []int
}

type MapNpcOffsetRec struct {
	X      int
	Y      int
	Type   byte
	Toggle bool
}

type TrappedRec struct {
	Sprite   int
	Duration int64
}

type NpcBuffEffectRec struct {
	Burned   bool
	Freezing bool
	Poisoned bool
}

type NpcTileRec struct {
	Type    byte
	Timer   int64
	Pos     objects.Vector2
	Alpha   byte
	fadeOut bool
}

type TradeItemRec struct {
	Item        int
	CostItem    int
	RequiredVar int
	ItemValue   int64
	CostValue   int64
	Haircut     TradeItemHaircutRec
}

type TradeItemHaircutRec struct {
	Sprite int
	Color  objects.Color
}

type ShopRec struct {
	Name      string
	BuyRate   int
	TradeItem []TradeItemRec
	Pic       int
	priceItem int
	Order     byte
	Type      byte
}

type SpellRec struct {
	Name         string
	ESName       string
	Desc         string
	Dir          byte
	Type         byte
	AccessReq    byte
	MPCost       int
	LevelReq     int
	ClassReq     int
	CastTime     int64
	CDTime       int64
	Icon         int
	Map          int
	X            int
	Y            int
	Vital        int64
	Duration     int64
	Interval     int64
	Range        int64
	IsAoE        bool
	AoE          int64
	CastAnim     int64
	SpellAnim    int64
	StunDuration int64
	BuffType     int64
	Frame        int64
	NextRank     int64
	NextUses     int64
	MaxMob       int64
	Stat         int64
	Cases        byte
	ElementBased bool
	ClassBasic   bool
	CastWalking  bool
	HPCost       int
	Int          []int
	Bool         []bool
}

type MapResourceRec struct {
	X             int
	Y             int
	ResourceState byte
	Frame         byte
	Shadow        byte
	FallDir       byte
	FallAngle     int
	HitWood       float32
}

type ResourceRec struct {
	Name           string
	ESName         string
	SuccessMessage string
	ResourceType   byte
	ResourceImage  int
	ExhaustedImage int
	ToolRequired   byte
	Health         int64
	RespawnTime    int
	Walkthrough    bool
	Animation      int
	ItemReward     []int
	ItemVal        []int64
	ItemLuck       []int
	NormalAnim     bool
	NormalRandom   bool
	Int            []int
}

type ActionMsgRec struct {
	Message string
	Created int64
	Type    int64
	Color   int64
	Scroll  int64
	X       int64
	Y       int64
	Timer   int64
}

type BloodRec struct {
	Sprite int64
	Timer  int64
	X      int64
	Y      int64
}

type AnimationRec struct {
	Name       string
	Sprite     []int64
	Frames     []int64
	LoopCount  []int64
	LoopTime   []int64
	XOffset    []int64
	YOffset    []int64
	Sound      string
	SoundValue int
}

type AnimInstanceRec struct {
	Animation    int64
	X            int64
	Y            int64
	Duration     int64
	lockindex    int
	LockType     byte
	Timer        []int64
	Used         []bool
	LoopIndex    []int64
	FrameIndex   []int64
	MasterSprite int
	MasterPos    objects.Vector2
	MasterTimer  int64
}

type ButtonRec struct {
	Num            int
	state          byte
	X              int64
	Y              int64
	Width          int
	Height         int
	Visible        bool
	PicNum         int64
	Name           string
	Sound          string
	Color          string
	FlipHorizontal bool
	flipV          bool
}

type GUIWindowRec struct {
	Name                  string
	pic                   int
	Type                  int64
	X                     int
	Y                     int
	X2                    int64
	Y2                    int64
	defaultX              int
	defaultY              int
	Width                 int
	Height                int
	isString              bool
	Max                   int64
	button                []int64
	label                 []string
	labelX                []int64
	labelY                []int64
	Drag                  bool
	Press                 bool
	Escape                bool
	Up                    bool
	Zoom                  bool
	IsVisible             bool
	CloseX                int
	CloseY                int
	ScreenPosition        int
	defaultScreenPosition int
	UIName                int
	Text                  string
	Visible               bool
	CanWrite              bool
	Resized               objects.Vector2
}

type ChatBubbleRec struct {
	Msg        string
	Colour     objects.Color
	Target     int64
	TargetType byte
	timer      int64
	active     bool
	pos        objects.Vector2
	tName      string
}

type CharDataRec struct {
	Name      string
	Level     int
	Classes   int
	Sprite    int
	Hair      int
	Sex       byte
	HairTint  objects.Color
	notHair   bool
	Equip     []int
	CashEquip []int
	Invisible []bool
}

type PlayerSpellRec struct {
	Spell  int
	Uses   int
	Master bool
}

type HotbarRec struct {
	Slot  int64
	sType byte
}

type MiniMapPlayerRec struct {
	X int64
	Y int64
}

type MiniMapNPCRec struct {
	X int64
	Y int64
}

type ServerAnnouncementRec struct {
	Message []string
	Timer   int64
}

type ProjectilRec struct {
	Name       string
	TravelTime int64
	Direction  int64
	X          int64
	Y          int64
	StateX     int64
	StateY     int64
	Pic        int
	Range      int64
	Damage     int64
	speed      int64
	timer      int64
	combo      int64
	Animation  int64
	Light      bool
	Int        []int
}

type MoralRec struct {
	Name string
	Int  []int
	Bool []bool
}

type StepsRec struct {
	Sprite int64
	timer  int64
	X      int64
	Y      int64
}

type FramesRec struct {
	Name          string
	Pic           int64
	PaperCheck    byte
	Frame1        int64
	Frame2        int64
	Frame3        int64
	Frame4        int64
	Timer1        int64
	Timer2        int64
	Timer3        int64
	Timer4        int64
	Paperdoll     int64
	PFrame1       int64
	PFrame2       int64
	PFrame3       int64
	PFrame4       int64
	Left1         int64
	Left2         int64
	Left3         int64
	Left4         int64
	KillPaperdoll byte
}

type CardsRec struct {
	Name   string
	Desc   string
	Num    int
	Card   int
	Sprite int
	Icon   int
	Offset objects.Vector2
}

type CraftingRec struct {
	Name      string
	Desc      string
	ItemReq   []int
	ItemVal   []int
	Reward    int
	RewardVal int
}

type NumberRec struct {
	Num     []string
	Width   []int64
	Pos     objects.Vector2
	Scroll  int64
	Created int64
	nColor  objects.Color
	aMsg    string
	Skin    byte
}

type MapSoundRec struct {
	pos         objects.Vector2
	SoundHandle string
	getTick     int64
	timer       int64
	Repeat      int
	SoundIndex  int
	isLoop      bool
	Destroy     bool
	Playing     bool
	delay       int64
}

type MapAnimationRec struct {
	Anim         []EachMapAnimationRec
	WeatherType  byte
	WeatherSpeed int64
	WeatherRate  int64
	EarthType    int64
	EarthSprite1 int64
	EarthSprite2 int64
	EarthSprite3 int64
	World        byte
}

type EachMapAnimationRec struct {
	Sprite     int
	X          int
	Y          int
	Size       byte
	Frames     byte
	Layer      byte
	Looop      int64
	XOffset    int64
	YOffset    int64
	Transport  int64
	CustomSize objects.Vector2
}

type FontRec struct {
	Ascii  []string
	Width  []int64
	X      int64
	Y      int64
	Scroll int64
	Active bool
	Leng   int64
}

type QuestTalkRec struct {
	Name string
	Lang []QuestLangTalkRec
}

type QuestLangTalkRec struct {
	Message []string
	Sprite  []int
}

type LanguageRec struct {
	Text []LanguageTextRec
}

type LanguageTextRec struct {
	Message  string
	Tutorial int
}

type GameFontsRec struct {
	Name string
}

type TitleRec struct {
	Name string
	Desc string
	Item int64
}

type TaskMsgRec struct {
	Msg       string
	FontColor objects.Color
	Created   int64
	dY        int
}

type NPCProjectilRec struct {
	num        int64
	TravelTime int64
	timer      int64
	Direction  int64
	X          int64
	Y          int64
	stateX     int64
	stateY     int64
	Launched   bool
}

type LogMsgRec struct {
	Msg         string
	nColor      objects.Color
	Created     int64
	Transparent int
	dY          int64
}

type MapTopoRec struct {
	Name    string
	Depth   int64
	Item    []int64
	ItemVal []int64
	Luck    []int64
}

type ParticleRec struct {
	Type     int64
	X        int64
	Y        int64
	Velocity int64
	InUse    int64
	Sprite   int
	timer    int64
	alpha    int64
	rotation int64
	size     int64
	Movement int64
}

type LightRec struct {
	Name       string
	OffSet     objects.Vector2
	Color      objects.Color
	Ratio      int
	BlinkRatio float64
}

type ConditionRec struct {
	Name     string
	Line     []ActionRec
	MaxLines int
	Sprite   ConditionSpriteRec
	ShowName bool
	OnJoin   bool
	MeetReq  string
	Faceset  int
	Switch   int
	Bool     []bool
}

type ConditionSpriteRec struct {
	pic     int
	Type    int
	myFrame objects.Vector2
	Mask    int
	Width   int
	Height  int
	Offset  objects.Vector2
}

type ActionRec struct {
	Void       byte
	NextLine   int
	ActionType byte
	EndCode    bool
	OrCode     bool
	AndCode    bool
	Modifier   byte
	ToEntity   byte
	ElseCode   bool
	isChoice   bool
	Var        []string
	VarType    []byte
	BreakPoint bool
}

type tmrEnvinromentRec struct {
	X   int64
	Y   int64
	tmr int64
	rnd int64
}

type MapEditorDataRec struct {
	Data1 int
	Data2 int
	Data3 int
	Data4 string
}
