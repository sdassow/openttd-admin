package admin

// echo $(grep ^type enums/enums.go | awk '{print$2}') | tr ' ' ','
//go:generate enumer -type=AdminCompanyRemoveReason,AdminUpdateFrequency,AdminUpdateType,Colour,DestType,Landscape,NetworkAction,NetworkErrorCode,NetworkLanguage,PacketType,PauseMode,VehicleType -json -text -output enumer.go

type AdminCompanyRemoveReason int

const (
	AdminCrrManual AdminCompanyRemoveReason = iota
	AdminCrrAutoclean
	AdminCrrBankrupt
)

type AdminUpdateFrequency uint16

const (
	AdminFrequencyPoll      AdminUpdateFrequency = 0x01
	AdminFrequencyDaily     AdminUpdateFrequency = 0x02
	AdminFrequencyWeekly    AdminUpdateFrequency = 0x04
	AdminFrequencyMonthly   AdminUpdateFrequency = 0x08
	AdminFrequencyQuarterly AdminUpdateFrequency = 0x10
	AdminFrequencyAnually   AdminUpdateFrequency = 0x20
	AdminFrequencyAutomatic AdminUpdateFrequency = 0x40
)

type AdminUpdateType uint16

const (
	AdminUpdateDate AdminUpdateType = iota
	AdminUpdateClientInfo
	AdminUpdateCompanyInfo
	AdminUpdateCompanyEconomy
	AdminUpdateCompanyStats
	AdminUpdateChat
	AdminUpdateConsole
	AdminUpdateCmdNames
	AdminUpdateCmdLogging
	AdminUpdateGamescript
	AdminUpdateEnd
)

type Colour uint16

const (
	DarkBlue Colour = iota
	PaleGreen
	Pink
	Yellow
	Red
	LightBlue
	Green
	DarkGreen
	Blue
	Cream
	Mauve
	Purple
	Orange
	Brown
	Grey
	White
	End
	Invalid Colour = 0xFF
)

type DestType int

const (
	DesttypeBroadcast DestType = iota
	DesttypeTeam
	DesttypeClient
)

type Landscape int

const (
	LandscapeTemperate Landscape = iota
	LandscapeArctic
	LandscapeTropic
	LandscapeToyland
	NumLandscape
)

type NetworkAction uint8

const (
	NetworkActionJoin NetworkAction = iota
	NetworkActionLeave
	NetworkActionServerMessage
	NetworkActionChat
	NetworkActionChatCompany
	NetworkActionChatClient
	NetworkActionGiveMoney
	NetworkActionNameChange
	NetworkActionCompanySpectator
	NetworkActionCompanyJoin
	NetworkActionCompanyNew
)

type NetworkErrorCode uint8

const (
	NetworkErrorGeneral NetworkErrorCode = iota // Try to use this one like never

	/* Signals from clients */
	NetworkErrorDesync
	NetworkErrorSavegameFailed
	NetworkErrorConnectionLost
	NetworkErrorIllegalPacket
	NetworkErrorNewgrfMismatch

	/* Signals from servers */
	NetworkErrorNotAuthorized
	NetworkErrorNotExpected
	NetworkErrorWrongRevision
	NetworkErrorNameInUse
	NetworkErrorWrongPassword
	NetworkErrorCompanyMismatch // Happens in ClientCommand
	NetworkErrorKicked
	NetworkErrorCheater
	NetworkErrorFull
)

type NetworkLanguage uint8

const (
	NetlangAny NetworkLanguage = iota
	NetlangEnglish
	NetlangGerman
	NetlangFrench
	NetlangBrazilian
	NetlangBulgarian
	NetlangChinese
	NetlangCzech
	NetlangDanish
	NetlangDutch
	NetlangEsperanto
	NetlangFinnish
	NetlangHungarian
	NetlangIcelandic
	NetlangItalian
	NetlangJapanese
	NetlangKorean
	NetlangLithuanian
	NetlangNorwegian
	NetlangPolish
	NetlangPortuguese
	NetlangRomanian
	NetlangRussian
	NetlangSlovak
	NetlangSlovenian
	NetlangSpanish
	NetlangSwedish
	NetlangTurkish
	NetlangUkrainian
	NetlangAfrikaans
	NetlangCroatian
	NetlangCatalan
	NetlangEstonian
	NetlangGalician
	NetlangGreek
	NetlangLatvian
	NetlangCount
)

type PacketType uint8

const (
	AdminPacketAdminJoin            PacketType = 0 ///< The admin announces and authenticates itself to the server.
	AdminPacketAdminQuit            PacketType = 1 ///< The admin tells the server that it is quitting.
	AdminPacketAdminUpdateFrequency PacketType = 2 ///< The admin tells the server the update frequency of a particular piece of information.
	AdminPacketAdminPoll            PacketType = 3 ///< The admin explicitly polls for a piece of information.
	AdminPacketAdminChat            PacketType = 4 ///< The admin sends a chat message to be distributed.
	AdminPacketAdminRcon            PacketType = 5 ///< The admin sends a remote console command.
	AdminPacketAdminGamescript      PacketType = 6 ///< The admin sends a Json string for the GameScript.
	AdminPacketAdminPing            PacketType = 7 ///< The admin sends a ping to the server, expecting a ping-reply (Pong) packet.

	AdminPacketServerFull           PacketType = 100 ///< The server tells the admin it cannot accept the admin.
	AdminPacketServerBanned         PacketType = 101 ///< The server tells the admin it is banned.
	AdminPacketServerError          PacketType = 102 ///< The server tells the admin an error has occurred.
	AdminPacketServerProtocol       PacketType = 103 ///< The server tells the admin its protocol version.
	AdminPacketServerWelcome        PacketType = 104 ///< The server welcomes the admin to a game.
	AdminPacketServerNewgame        PacketType = 105 ///< The server tells the admin its going to start a new game.
	AdminPacketServerShutdown       PacketType = 106 ///< The server tells the admin its shutting down.
	AdminPacketServerDate           PacketType = 107 ///< The server tells the admin what the current game date is.
	AdminPacketServerClientJoin     PacketType = 108 ///< The server tells the admin that a client has joined.
	AdminPacketServerClientInfo     PacketType = 109 ///< The server gives the admin information about a client.
	AdminPacketServerClientUpdate   PacketType = 110 ///< The server gives the admin an information update on a client.
	AdminPacketServerClientQuit     PacketType = 111 ///< The server tells the admin that a client quit.
	AdminPacketServerClientError    PacketType = 112 ///< The server tells the admin that a client caused an error.
	AdminPacketServerCompanyNew     PacketType = 113 ///< The server tells the admin that a new company has started.
	AdminPacketServerCompanyInfo    PacketType = 114 ///< The server gives the admin information about a company.
	AdminPacketServerCompanyUpdate  PacketType = 115 ///< The server gives the admin an information update on a company.
	AdminPacketServerCompanyRemove  PacketType = 116 ///< The server tells the admin that a company was removed.
	AdminPacketServerCompanyEconomy PacketType = 117 ///< The server gives the admin some economy related company information.
	AdminPacketServerCompanyStats   PacketType = 118 ///< The server gives the admin some statistics about a company.
	AdminPacketServerChat           PacketType = 119 ///< The server received a chat message and relays it.
	AdminPacketServerRcon           PacketType = 120 ///< The server's reply to a remove console command.
	AdminPacketServerConsole        PacketType = 121 ///< The server gives the admin the data that got printed to its console.
	AdminPacketServerCmdNames       PacketType = 122 ///< The server gives the admin names of all DoCommands.
	AdminPacketServerCmdLogging     PacketType = 123 ///< The server gives the admin DoCommand information (for logging purposes only).
	AdminPacketServerGamescript     PacketType = 124 ///< The server gives the admin information from the GameScript in Json.
	AdminPacketServerRconEnd        PacketType = 125 ///< The server indicates that the remote console command has completed.
	AdminPacketServerPong           PacketType = 126 ///< The server replies to a ping request from the admin.
	AdminPacketServerEnd            PacketType = 127 ///< For internal reference only, mark the end.

	InvalidAdminPacket PacketType = 0xFF ///< An invalid marker for admin packets.
)

type PauseMode int

const (
	PmUnpaused            PauseMode = 0  ///< A normal unpaused game
	PmPausedNormal        PauseMode = 1  ///< A game normally paused
	PmPausedSaveload      PauseMode = 2  ///< A game paused for saving/loading
	PmPausedJoin          PauseMode = 4  ///< A game paused for 'pause_on_join'
	PmPausedError         PauseMode = 8  ///< A game paused because a (critical) error
	PmPausedActiveClients PauseMode = 16 ///< A game paused for 'min_active_clients'
	PmPausedGameScript    PauseMode = 32 ///< A game paused by a game script

	/** Pause mode bits when paused for network reasons. */
	PmbPausedNetwork PauseMode = PmPausedActiveClients | PmPausedJoin
)

type VehicleType int

const (
	NetworkVehTrain VehicleType = iota
	NetworkVehLorry
	NetworkVehBus
	NetworkVehPlane
	NetworkVehShip
	NetworkVehEnd
)
