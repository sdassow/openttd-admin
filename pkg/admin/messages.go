package admin

import (
	"fmt"
	"time"

	"github.com/ghostiam/binstruct"
)

func convertDate(date int) time.Time {

	yr := 400 * (date / (DaysInYear*400 + 97))
	rem := date % (DaysInYear*400 + 97)

	if rem >= DaysInYear*100+25 {
		/* There are 25 leap years in the first 100 years after
		 * every 400th year, as every 400th year is a leap year */
		yr += 100
		rem -= DaysInYear*100 + 25

		/* There are 24 leap years in the next couple of 100 years */
		yr += 100 * (rem / (DaysInYear*100 + 24))
		rem = rem % (DaysInYear*100 + 24)
	}

	if !IsLeapYear(yr) && rem >= DaysInYear*4 {
		/* The first 4 year of the century are not always a leap year */
		yr += 4
		rem -= DaysInYear * 4
	}

	/* There is 1 leap year every 4 years */
	yr += 4 * (rem / (DaysInYear*4 + 1))
	rem = rem % (DaysInYear*4 + 1)

	/* The last (max 3) years to account for; the first one
	 * can be, but is not necessarily a leap year */
	//for rem >= (IsLeapYear(yr) ? DAYS_IN_LEAP_YEAR : DaysInYear) {
	for rem >= GetDaysInYear(yr) {
		//rem -= IsLeapYear(yr) ? DAYS_IN_LEAP_YEAR : DaysInYear
		rem -= GetDaysInYear(yr)
		yr++
	}

	/* Skip the 29th of February in non-leap years */
	//if !IsLeapYear(yr) && rem >= ACCUM_MAR - 1 {
	if !IsLeapYear(yr) && rem >= 2 {
		rem++
	}

	//ymd->year = yr;

	//log.Printf("date=%d, yr=%d, rem=%d", date, yr, rem)

	//t := time.Time{}
	//t.AddDate(yr, 0, rem)
	t := time.Date(yr, 1, rem+1, 0, 0, 0, 0, time.UTC)

	//x = _month_date_from_year_day[rem];
	//ymd->month = x >> 5;
	//ymd->day = x & 0x1F;

	return t
}

func IsLeapYear(yr int) bool {
	return yr%4 == 0 && (yr%100 != 0 || yr%400 == 0)
}

const DaysInYear int = 365

func GetDaysInYear(yr int) int {
	if IsLeapYear(yr) {
		return DaysInYear + 1
	}

	return DaysInYear
}

func NullString(r binstruct.Reader) (string, error) {
	var b []byte

	for {
		readByte, err := r.ReadByte()
		if binstruct.IsEOF(err) {
			break
		}
		if err != nil {
			return "", err
		}

		if readByte == 0x00 {
			break
		}

		b = append(b, readByte)
	}

	return string(b), nil
}

func GameDate(r binstruct.Reader) (time.Time, error) {
	date, err := r.ReadUint32()
	if err != nil {
		return time.Time{}, err
	}
	return convertDate(int(date)), nil
}

func ParseMessage(p *Packet) (interface{}, error) {
	switch p.Type {
	case AdminPacketServerFull:
		return &ServerFull{}, nil

	case AdminPacketServerBanned:
		return &ServerBanned{}, nil

	case AdminPacketServerError:
		x := &ServerError{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerProtocol:
		x := &ServerProtocol{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerWelcome:
		x := &ServerWelcome{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerNewgame:
		return &ServerNewgame{}, nil

	case AdminPacketServerShutdown:
		return &ServerShutdown{}, nil

	case AdminPacketServerDate:
		x := &ServerDate{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerClientJoin:
		x := &ServerClientJoin{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerClientInfo:
		x := &ServerClientInfo{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerClientUpdate:
		x := &ServerClientUpdate{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerClientQuit:
		x := &ServerClientQuit{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerClientError:
		x := &ServerClientError{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyNew:
		x := &ServerCompanyNew{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyInfo:
		x := &ServerCompanyInfo{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyUpdate:
		x := &ServerCompanyUpdate{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyRemove:
		x := &ServerCompanyRemove{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyEconomy:
		x := &ServerCompanyEconomy{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCompanyStats:
		x := &ServerCompanyStats{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerChat:
		x := &ServerChat{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerRcon:
		x := &ServerRcon{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerConsole:
		x := &ServerConsole{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCmdNames:
		x := &ServerCmdNames{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerCmdLogging:
		x := &ServerCmdLogging{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerGamescript:
		x := &ServerGamescript{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerRconEnd:
		x := &ServerRconEnd{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	case AdminPacketServerPong:
		x := &ServerPong{}
		if err := binstruct.UnmarshalLE(p.Payload, x); err != nil {
			return nil, err
		}
		return x, nil

	default:
		return nil, fmt.Errorf("Unable to handle packet: %+v", p.Type)
	}

	return nil, nil
}

type ServerFull struct {
}

type ServerBanned struct {
}

type ServerError struct {
	Error   uint8
	KickMsg string `bin:"NullString"`
}

func (*ServerError) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerProtocol struct {
	Version     uint8
	Frequencies [AdminUpdateEnd]struct {
		Bool  bool
		Num   uint16
		Flags uint16
	}

	Bool bool
}

type ServerWelcome struct {
	ServerName     string `bin:"NullString"`
	Revision       string `bin:"NullString"`
	Dedicated      bool
	MapName        string `bin:"NullString"`
	GenerationSeed uint32
	Landscape      uint8
	Date           time.Time `bin:"GameDate"`
	MapSizeX       uint16
	MapSizeY       uint16
}

func (*ServerWelcome) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

func (*ServerWelcome) GameDate(r binstruct.Reader) (time.Time, error) {
	return GameDate(r)
}

type ServerNewgame struct {
}

type ServerShutdown struct {
}

type ServerDate struct {
	Date time.Time `bin:"GameDate"`
}

func (*ServerDate) GameDate(r binstruct.Reader) (time.Time, error) {
	return GameDate(r)
}

type ServerClientJoin struct {
	ClientId uint32
}

type ServerClientInfo struct {
	ClientId      uint32
	HostName      string `bin:"NullString"`
	ClientName    string `bin:"NullString"`
	ClientLang    NetworkLanguage
	JoinDate      time.Time `bin:"GameDate"`
	ClientPlayers uint8
}

func (*ServerClientInfo) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

func (*ServerClientInfo) GameDate(r binstruct.Reader) (time.Time, error) {
	return GameDate(r)
}

type ServerClientUpdate struct {
	ClientId      uint32
	ClientName    string `bin:"NullString"`
	ClientPlayers uint8
}

func (*ServerClientUpdate) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerClientQuit struct {
	ClientId uint32
}

type ServerClientError struct {
	ClientId uint32
	Code     NetworkErrorCode
}

type ServerCompanyNew struct {
	CompanyId uint32
}

const NumShareOwners int = 4

type ServerCompanyInfo struct {
	Index                uint8
	CompanyName          string `bin:"NullString"`
	ManagerName          string `bin:"NullString"`
	Colour               uint8
	IsPasswordProtected  bool
	InauguratedYear      uint32
	IsAI                 bool
	QuartersOfBankruptcy uint8
	ShareOwners          [NumShareOwners]uint8
}

func (*ServerCompanyInfo) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerCompanyUpdate struct {
	Index                uint8
	CompanyName          string `bin:"NullString"`
	ManagerName          string `bin:"NullString"`
	Colour               uint8
	IsPasswordProtected  bool
	QuartersOfBankruptcy uint8
	ShareOwners          [NumShareOwners]uint8
}

func (*ServerCompanyUpdate) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerCompanyRemove struct {
	CompanyId uint8
	Accr      uint8
}

const NumLastQuarters int = 2

type ServerCompanyEconomy struct {
	CompanyId      uint8
	Money          uint64
	CurrentLoan    uint64
	Income         uint64
	DeliveredCargo uint16
	LastQuarters   [NumLastQuarters]struct {
		CompanyValue       uint64
		PerformanceHistory uint16
		DeliveredCargo     uint16
	}
}

type ServerCompanyStats struct {
	CompanyId uint8
	Vehicles  [NetworkVehEnd]uint16
	Stations  [NetworkVehEnd]uint16
}

type ServerChat struct {
	Action   NetworkAction
	ClientId uint32
	SelfSend bool
	Message  string `bin:"NullString"`
	Data     uint64
}

func (*ServerChat) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerRcon struct {
	Colour Colour
	Result string `bin:"NullString"`
}

func (*ServerRcon) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerConsole struct {
	Origin  string `bin:"NullString"`
	Message string `bin:"NullString"`
}

func (*ServerConsole) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerCmdNames struct {
	Commands []struct {
		MoreData bool
		Num      uint16
		Name     string `bin:"NullString"`
	}
	EndOfPacket bool
}

func (*ServerCmdNames) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerCmdLogging struct {
	ClientId uint32
	Company  uint8
	Command  uint16
	P1       uint32
	P2       uint32
	Tile     uint32
	Text     string `bin:"NullString"`
	Frame    uint32
}

func (*ServerCmdLogging) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerGamescript struct {
	Text string `bin:"NullString"`
}

func (*ServerGamescript) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerRconEnd struct {
	Command string `bin:"NullString"`
}

func (*ServerRconEnd) NullString(r binstruct.Reader) (string, error) {
	return NullString(r)
}

type ServerPong struct {
	Payload uint32
}
