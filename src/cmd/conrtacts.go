package cmd

type UserInputKey string

const (
	FullName          UserInputKey = "fullName"
	ShortName         UserInputKey = "shortName"
	Birthday          UserInputKey = "birthday"
	PlaceOfBirth      UserInputKey = "placeOfBirth"
	OfficialAddress   UserInputKey = "officialAddress"
	ActualAddress     UserInputKey = "actualAddress"
	Identifier        UserInputKey = "identifier"
	Occupation        UserInputKey = "occupation"
	NumberOfProtocol  UserInputKey = "numberOfProtocol"
	DateOfProtocol    UserInputKey = "dateOfProtocol"
	DateOfAccident    UserInputKey = "dateOfAccident"
	TimeOfAccident    UserInputKey = "timeOfAccident"
	AddressOfAccident UserInputKey = "addressOfAccident"
	NumberOfOrdinance UserInputKey = "numberOfOrdinance"
	DateOfOrdinance   UserInputKey = "dateOfOrdinance"
	TimeOfOrdinance   UserInputKey = "timeOfOrdinance"
	Decision          UserInputKey = "decision"
	DateOfEnactment   UserInputKey = "dateOfEnactment"
)

type UserInputsMap map[UserInputKey]string
