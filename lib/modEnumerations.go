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

type Lang int32

const (
	LangCustom Lang = iota - 1
	LangSEquipItem
	LangSCannotEquipItem
	LangGDefenseStatInfo
	LangGAttackStatInfo
	LangGShopBuyThisFor
	LangShop_NotUnlocked
	LangNewspaper_BronzeDay
	LangGHowToMove
	LangGHowToAttack
	LangGPlayerBuyThisFor
	LangGHowToSkill
	LangQuestList
	LangGHowToCraft
	LangGHowToUseItem
	LangGHowToClass
	LangGHowToMonsterBook
	LangGHowToStats
	LangGSpellCooldown
	LangGSpellCastWalk
	LangGSpellTarget
	LangGSpellVitalRequired
	LangGArrowNeed
	LangGNotManaPoints
	LangGInventoryFull
	LangGShopNotEnough
	LangGGetExperience
	LangGGetCalaveras
	LangGGetItem
	LangGGetItemValue
	LangGRecipeRequire
	LangGAttackStat
	LangGDefenseStat
	LangStat_Spirit
	LangAgilityStat
	LangGIntelligenceStat
	LangSpiritStatInfo
	LangAgilityStatInfo
	LangGIntelligenceStatInfo
	LangGRecommended
	LangGSecondary
	LangEquipWeapon
	LangSellItem
	LangBuyItem
	LangShopDoesntWant
	LangAdjustKey
	LangAdjustKeyCannot
	LangItem_NotForYourClass
	LangQuest_GoTalk
	LangQuest_GoCraft
	LangQuest_GoGather
	LangQuestGoSlay
	LangQuestGoTrain
	LangQuestGoSpell
	LangQuestGoDrill
	LangQuestGoUseItem
	LangQuestGoGetCard
	LangQuestGoWalk
	LangQuestGoMap
	LangQuestGoUse
	LangLevel
	LangShop_NoValue
	LangChooseElement
	LangQuestTipSpell
	LangRecipeAlready
	LangRecipeMissIngredient
	LangNotInAParty
	LangHowToDrill
	LangPlayerKilledBy
	LangPlayerLostExp
	LangQuestOwner
	LangQuestObjetive
	LangQuestProgress
	LangQuestLocalization
	LangQuest_UnderDevelopment
	LangQuestUnavailable
	LangQuestReward
	LangQuestDetails
	LangQuestGoBuy
	LangQuestGoSell
	LangTitleInventory
	LangTitleConfig
	LangTitleParty
	LangTitleSelectChar
	LangTitleMonsterBook
	LangTitleRecipe
	LangTitleHelp
	LangTitleCharacter
	LangTitleRegister
	LangTitleKeyConfig
	LangTitleElement
	LangTitleAward
	LangTitleSpell
	LangTitleQuestBook
	LangTitleCombo
	LangTitleWorldmap
	LangOptMusic
	LangOptSound
	LangOptDebug
	LangOptFullscreen
	LangOptVSync
	LangOptJoystick
	LangOptKeyboard
	LangOptSize
	LangOptEmoticonSize
	LangTitleCashInventory
	LangPlayerStunned
	LangItemStatRequired
	LangItemLevelRequired
	LangDrillCannotHere
	LangTitleEquipment
	LangPlayerWasDied
	LangKillWaitSec
	LangKillRespawn
	LangInvitationSent
	LangPartyIsFull
	LangPartyYouAreNotLead
	LangTitlePartyInvitation
	LangPartyPlayerInviteYou
	LangPartyPlayerDecline
	LangPartyYouDecline
	LangItemCheckEquip
	LangItemCheckUnequip
	LangResourceWrongTool
	LangTitlePlayerInfo
	LangResourceNeedATool
	LangPlayerInfoStats
	LangPlayerInfoAwards
	LangTitleBug
	LangInvalidTarget
	LangPlayerNotSameMap
	LangHonorRequireLevel
	LangHonorWaitSec
	LangQuestIsCompleted
	LangQuestIsInProgress
	LangTitleEscapeMenu
	LangHonorPlayerTake
	LangHonorPlayerGive
	LangHonorYouGive
	LangHonorYouTake
	LangPlayerCannotRespawn
	LangInventoryHelpOne
	LangInventoryHelpTwo
	LangInventoryHelpThree
	LangItemDescNotBonus
	LangItemDescNotElemental
	LangItemDescSpeed
	LangItemDescCritical
	LangItemDescLuminosity
	LangItemDescNotStatBonus
	LangPlayerNotOnline
	LangBuddiesAlreadyAdded
	LangBuddiesAlreadyOnHisList
	LangBuddiesIsFull
	LangBuddiesTargetIsFull
	LangPlayerIsBusy
	LangBuddiesPlayerDecline
	LangBuddiesYouDecline
	LangTitleBuddiesInvitation
	LangBuddiesPlayerInviteYou
	LangTitleBuddies
	LangBuddiesPlayerAccept
	LangBuddiesYouAccept
	LangBuddiesIsOnline
	LangBankIsFull
	LangAlertClassNotAvailable
	LangAlertNameAtLeast
	LangAlertNameInvalid
	LangAlertCharExist
	LangAlertClassNotSelected
	LangAlertCantUseBody
	LangAlertCantUseHair
	LangAlertCantUseTint
	LangAlertNameIsUsed
	LangQuestGiveTo
	LangQuestGoLevel
	LangQuestCraftTo
	LangQuestGoHunt
	LangBackgroundClose
	LangMenuLoading
	LangQuestGoEvent
	LangLogQuestNew
	LangLogQuestComplete
	LangLogTaskComplete
	LangLogRecipeNew
	LangLogCardNew
	LangLogElementNew
	LangLogAwardUnlock
	LangLogSkillNew
	LangTitleCashShopPreview
	LangTitleCashShopWish
	LangTitleCashShopCoupon
	LangCouponEscape
	LangCouponWrong
	LangCouponWait
	LangCouponInvalid
	LangCouponCompleted
	LangCouponAlready
	LangTitleCouponReply
	LangTitleCashShopBuy
	LangCouponWannaBuy
	LangCouponInstructions
	LangTitleLogin
	LangTitle_ContentGuide
	LangServer_QuestMsgResearch
	LangHelp_Skulls
	LangHelp_WorldMap
	LangFavAwardAlready
	LangFavAwardUncompleted
	LangFavAwardFull
	LangFavAwardAddFavorite
	LangFavAwardDelFavorite
	LangAlertServerLost
	LangOptLanguages
	LangOptResolution
	LangCurrencyDrop
	LangCurrencyBuy
	LangCurrencySell
	LangCurrencyDeposit
	LangCurrencyWithdraw
	LangCurrencyCalaveras
	LangQuestLogCraft
	LangQuestTipSell
	LangQuestNeedLevel
	LangQuestNeedQuest
	LangQuestTipLevel
	LangFishingNeedBait
	LangServerPlayerJoin
	LangServerPlayerLeft
	LangQuestInventoryFull
	LangQuestTipFishing
	LangQuestTipTraining
	LangTitleMiniGames
	LangMinigamesTitleEvent
	LangMinigamesTitlePlayer
	LangMinigamesKrobiDesc
	LangMinigamesHalloweenDesc
	LangMinigamesPvpDesc
	LangNameKrobiIsland
	LangNameHalloweenTower
	LangNameBattleRoom
	LangQuestTipCatch
	LangLogMasterNew
	LangTitleMaster
	LangTitleElements
	LangComboSave
	LangMenuUsername
	LangMenuPassword
	LangMenuSaveID
	LangQuestLogUse
	LangQuestLogCard
	LangQuestLogKill
	LangQuestLogFish
	LangQuestLogCatch
	LangOnline
	LangOffline
	LangFriends
	LangTitleURL
	LangConditionIncorrectCode
	LangRegister_Username
	LangRegister_Email
	LangRegister_Password
	LangRegister_Repeat
	LangRegister_Captcha
	LangMenu_Copyright
	LangNewchar_Name
	LangNewChar_Appearance
	LangNewChar_Info
	LangEvent_ChestOpen
	LangQuest_TipKill
	LangItem_Sword
	LangItem_Bow
	LangItem_Staff
	LangItem_Dagger
	LangTitle_Shop
	LangTitle_Barber
	LangPlayer_TakeItem
	LangServer_AlreadyBuff
	LangGame_WrongWay
	LangPartyquest_Halloween1
	LangMenu_Clue
	LangPartyquest_Halloween2
	LangPartyquest_Halloween3
	LangPartyquest_Halloween4
	LangChat_SpamNotAllowed
	LangServer_FunctionNotAvailable
	LangPlayer_LevelUp
	LangTitle_Pins
	LangGame_PinUnlocked
	LangGame_PinAlready
	LangGame_PinName
	LangServer_PinUnlocked
	LangServer_PinEquip
	LangGame_PinDoesntHave
	LangCoupon_PromotionAlready
	LangGame_CutGrass
	LangGame_UsePickaxe
	LangGame_UseFishing
	LangGame_StoryQuest
	LangServer_FreeSkulls
	LangGame_UseWoodcut
	LangEvent_LevelRequired
	LangEvent_LeaderRequiredToStart
	LangEvent_PartyExceeded
	LangEvent_AliveRequired
	LangEvent_PlayerNotInMap
	LangEvent_NeedParty
	LangEvent_SomeoneIn
	LangEvent_LeaderRequiredToEnd
	LangWelcome_PlayersConnected
	LangTitle_Forge
	LangForge_MissingMaterials
	LangForge_Congratulations
	LangShop_Forge
	LangServer_DayTime
	LangServer_NightTime
	LangInsect_Caught
	LangInsect_Freed
	LangInsect_InventoryLiberated
	LangSpell_Damage
	LangSpell_Defense
	LangSpell_RegenMP
	LangSpell_HealHP
	LangSpell_HealMP
	LangSpell_Projectil
	LangSpell_Excavation
	LangSpell_Buff
	LangSpell_Debuff
	LangSpell_Warp
	LangServer_WorldBless
	LangItem_NotForYourSex
	LangItem_AlreadyPolyphorm
	LangItem_HolyRainAlreadyActive
	LangSettings_WindowsCursor
	LangSettings_CaptureQuality
	LangSettings_Enable
	LangSettings_Disable
	LangSettings_High
	LangSettings_Low
	LangQuest_Type
	LangQuest_Status
	LangQuest_TipDrill
	LangQuest_TipWalk
	LangQuest_TipTalk
	LangQuest_TypeGive
	LangQuest_TypeReach
	LangQuest_TypeWalk
	LangQuest_TypeTalk
	LangQuest_TypeLevel
	LangQuest_TypeCraft
	LangQuest_TypeSpell
	LangQuest_TypeDrill
	LangQuest_TypePartyLevel
	LangQuest_TypeCard
	LangQuest_TypeUseItem
	LangQuest_TypeInteract
	LangQuest_TypeFish
	LangQuest_TypeCatch
	LangQuest_TypeInvestigate
	LangQuest_TypeForge
	LangPlayer_KillBoss
	LangClass_KnightInfo
	LangClass_MageInfo
	LangClass_HunterInfo
	LangClass_DruidInfo
	LangClass_StalkerInfo
	LangClass_NecromancerInfo
	LangGame_Time
	LangParty_GetExperience
	LangTitle_TradeRequest
	LangTitle_TradePlayerRequestYou
	LangTrade_YouAcceptTargetRequest
	LangTrade_TargetAcceptYourRequest
	LangTrade_Completed
	LangTrade_InventoryFull
	LangTrade_YouDeclineTrade
	LangTrade_TargetDeclineTrade
	LangTrade_NotEnoughSkulls
	LangInvite_NoInvitation
	LangTrade_TargetInviteYou
	LangTrade_YouInviteTarget
	LangTrade_AlreadyOfferedItem
	LangTitle_Trade
	LangParty_PlayerAlreadyInParty
	LangParty_PlayerJoinParty
	LangParty_StartedNewParty
	LangParty_LevelUp
	LangCraft_RequireItem
	LangCraft_RequireItemValue
	LangCard_UnlockNewCard
	LangQuest_ReqInvSlot
	LangQuest_ReqItemToStart
	LangQuest_ReqItemToComplete
	LangQuest_YouAcceptTrade
	LangQuest_TargetAcceptTrade
	LangTrade_YouAlreadyAccepted
	LangDay_Monday
	LangDay_Tuesday
	LangDay_Wednesday
	LangDay_Thursday
	LangDay_Friday
	LangDay_Saturday
	LangDay_Sunday
	LangDay_Today
	LangKeyboard_Help
	LangTitle_MasterSkill
	LangLog_SkillMastery
	LangSpell_AlreadyMastery
	LangSpell_NotLearned
	LangServer_YouNeedMonsterCard
	LangServer_YouNeedAP
	LangSpell_NeedTheThreeStars
	LangTitle_DailyCheck
	LangDailycheck_DateToDate
	LangGame_OutdatedClient
	LangParty_ShareItems
	LangGame_ServerOffline
	LangServer_PlayerLeftMod
	LangMenu_OldPassword
	LangTitle_PasswordReset
	LangPlayer_YourBarIsFull
	LangPlayer_FillTheBarToDig
	LangServer_ItemNotMeetReq
	LangServer_CannotEquipTwoHandWeapon
	LangServer_CannotEquipAShield
	LangServer_CannotEquipOverall
	LangServer_CannotEquipPants
	LangWorldmap_LocationUnavailable
	LangServer_TermsOfService
	LangServer_YouHaveBeenDisconnected
	LangServer_ReadTermsOfService
	LangServer_MinLength
	LangServer_MaxLength
	LangServer_InvalidInfo
	LangChat_PressEnter
	LangReport_Player
	LangReport_Desc
	LangReport_ReasonNotDefined
	LangReport_Successful
	LangCurrency_Trade
	LangTitle_Joystick
	LangPlayer_Esia
	LangCharacter_Points
	LangPlayer_LeftToRespawn
	LangNewspaper_News
	LangNewspaper_Update
	LangNewspaper_Players
	LangNewspaper_Today
	LangNewspaper_Title
	LangNewspaper_Daily
	LangNewspaper_Check
	LangNewspaper_ArtBy
	LangNewspaper_EndDate
	LangNewspaper_Announcement
	LangNewspaper_Weekly
	LangNewspaper_Server
	LangHow_PlayerInformation
	LangHow_TurnAround
	LangHow_Skulls
	LangHow_WorldMap
	LangMegaphone_Item
	LangTitle_Megaphone
	LangDailycheck_AlreadyChecked
	LangEnhancement_Successfully
	LangUrl_GoTo
	LangGame_OutdatedRevision
	LangExplore_Npc
	LangExplore_Map
	LangInfo_Damage
	LangInfo_MagicDamage
	LangInfo_RangeDamage
	LangInfo_Critical
	LangInfo_Defense
	LangInfo_MagicDefense
	LangInfo_RangeDefense
	LangInfo_Speed
	LangCashshop_NotAvailable
	LangTitle_PhysicDamage
	LangTitle_MagicDamage
	LangTitle_RangeDamage
	LangTitle_CriticalDamage
	LangTitle_PhysicDefense
	LangTitle_MagicDefense
	LangTitle_RangeDefense
	LangTitle_Speed
	LangCurrency_InvalidAmount
	LangInventory_MissItem
	LangShop_MissItem
	LangTitle_Bank
	LangError_ShuttingDown
	LangError_AccountNotExist
	LangError_IncorrectPassword
	LangError_MultipleAccounts
	LangNewspaper_CommingSoon
	LangNewspaper_SocialMedia
	LangCashshop_Free
	LangAchievement_Add
	LangAchievement_Del
	LangCashinv_Font
	LangMinimap_Day
	LangMinimap_Night
	LangSpelldesc_Target
	LangSpelldesc_Cooldown
	LangNpc_ShamiDialogue
	LangName_Level
	LangQuest_Requirement
	LangQuest_AllClasses
	LangTutorial_HowToShop
	LangClass_CantLeave
	LangQuest_LogCompleted
	LangCraft_LevelRequire
	LangTutorial_HowToPotion
	LangTutorial_HowToTravel
	LangQuest_TipSpell
	LangTutorial_HowToNewSkill
	LangQuest_Available
	LangCoupon_Successful
	LangCashshop_Hot
	LangSettings_PlayerName
	LangEvent_LevelRequirement
	LangTitle_Information
	LangDojo_ContinueTrying
	LangDojo_YouHasBeenPromoted
	LangDojo_PlayerHasBeenPromoted
	LangDojo_IsFull
	LangEvent_HappyHalloween
	LangWorldboss_Summon
	LangCard_NpcRewards
	LangItem_Slots
	LangProfession_Digging
	LangProfession_Mining
	LangProfession_TreeCutting
	LangProfession_Fishing
	LangTitle_FishList
	LangTitle_Profession
	LangProfession_Times
	LangQuest_FindLocation
	LangLog_ProfessionLevelUp
	LangLog_ProfessionExp
	LangQuest_TipForge
	LangQuest_TipInvestigate
	LangQuest_TipBuy
	LangQuest_TipMannequin
	LangServer_BlockedTarget
	LangMenu_DelChar
	LangOpt_Cooldown
	LangCurrency_TradeSkulls
	LangCurrency_Stat
	LangServer_MoleHPEmpty
	LangTitle_HP
	LangTitle_MP
	LangInfo_HP
	LangInfo_MP
	LangTitle_Luminiscence
	LangInfo_Luminiscence
	LangTitle_DodgeChance
	LangInfo_DodgeChance
	LangTitle_Account
	LangTitle_GetUsername
	LangTitle_ResetPassword
	LangError_YouHavenFilled
	LangError_InvalidCaptcha
	LangError_NotSamePassword
	LangError_InvalidEmail
	LangError_AccountInUse
	LangServer_PasswordChangeSuccessfully
	LangNewspaper_FestivalDay
	LangNewspaper_DiggingDay
	LangNewspaper_TopPlayers
	LangNewspaper_TopPlayersDesc
	LangNewspaper_LatestAdventurer
	LangNewspaper_LatestAdventurerDesc
	LangTitle_KeySetting
	LangKeysetting_Description
	LangKeysetting_ClassicDesc
	LangKeysetting_NewDesc
	LangKeysetting_Classic
	LangKeysetting_New
	LangServer_ModExp
	LangServer_ModDrop
	LangServer_ModDig
	LangServer_ModDamage
	LangServer_ModHeal
	LangServer_ModJunk
	LangServer_ModLostExp
	LangServer_ModParty
	LangServer_ModProfession
	LangServer_ModCardDrop
	LangServer_ShutdownMin
	LangServer_ShutdownSec
	LangServer_PasswordResetSuccessfully
	LangServer_GetUsernameSuccessfully
	LangServer_AlreadySentRequest
	LangServer_RequestBusy
	LangMinigames_SpeedTest
	LangServer_CannotSpawnMob
	LangServer_GhostCannotEquip
	LangServer_UseProjectilSkin
	LangServer_UnlockSlot
	LangServer_AlreadyUnlockedSlot
	LangServer_AddedEmoji
	LangServer_AlreadyAddedEmoji
	LangTitle_HolyRainOwner
	LangServer_BanPlayer
	LangServer_FarAwayWarp
	LangUi_ComboDesc
	LangUi_ElementDesc
	LangUi_ArrowDesc
	LangUi_MasterDesc
	LangUi_elementstat
	LangUi_deck
	LangUi_special
	LangUi_trap
	LangTitle_projectiles
	LangUi_alreadybought
	LangUi_towhisper
	LangUi_fromwhisper
	LangUi_whisperalert
	LangUi_block
	LangUi_unblock
	LangUi_active
	LangUi_devicenotconnected
	LangUi_joystickdesc
	LangUi_screenshottaken
	LangUi_learnelement
	LangBlocklist_maxlimit
	LangBlocklist_delfrombuddies
	LangBlocklist_alreadyblocked
	LangBlocklist_noinyourlist
	LangQuest_rewardtakeitem
	LangCommand_invalid
	LangServer_notenoughstats
	LangHow_fishing
	LangHow_mining
	LangHow_woodcutting
	LangHow_enhancement
	LangHow_masteryskill
	LangHow_achievementpoints
	LangHow_emojis
	LangTitle_AP
	LangUi_enhancementstone
	LangSpell_cantride
	LangQuest_startitemreq
	LangUi_agimage
	LangServer_cantcastspell
	LangUi_DuelArena
	LangTitle_worldboss
	LangUi_cannotcreateparty
	LangUi_cannotjoinwithparty
	LangUi_cannotinvitetoparty
	LangSpell_cannotunequipweapon
	LangSpell_cannotstackbuff
	LangSetting_exppercentage
	LangSetting_badwords
	LangTip_botsnotallowed
	LangTip_joindiscord
	LangTip_visitwebsite
	LangTip_steamreview
	LangGamecard_description
	LangTitle_gamecard
	LangPremium_steamfailed
	LangPremium_steamcompleted
	LangPremium_steamqueuefull
	LangPremium_steamnotinitialized
	LangTitle_BotsNotAllowed
	LangName_TermsOfService
	LangError_notintrade
	LangTitle_TravelingEsia
	LangTip_TravelingEsia
	LangSpell_Mount
	LangSpell_Magnet
	LangSpell_Trap
	LangSpell_Activation
	LangSpell_None
	LangSpell_Single
	LangSpell_Self
	LangSpell_Range
	LangMastery_AvailableSkills
)
