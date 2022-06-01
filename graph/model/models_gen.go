// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewUser struct {
	ID                string    `json:"id"`
	Name              *string   `json:"name"`
	LastName          *string   `json:"lastName"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Admin             *bool     `json:"admin"`
	Root              *bool     `json:"root"`
	Verified          *bool     `json:"verified"`
	Reported          *bool     `json:"reported"`
	ReportReason      *string   `json:"reportReason"`
	ActiveContract    *bool     `json:"activeContract"`
	AdmissionDay      *string   `json:"admissionDay"`
	UnemploymentDay   *string   `json:"unemploymentDay"`
	WorkedHours       *int      `json:"workedHours"`
	CurrentBranch     *string   `json:"currentBranch"`
	OriginBranch      *string   `json:"originBranch"`
	MonetaryBonds     *float64  `json:"monetaryBonds"`
	MonetaryDiscounts *float64  `json:"monetaryDiscounts"`
	Mail              *string   `json:"mail"`
	AlternativeMails  []*string `json:"alternativeMails"`
	Phone             *string   `json:"phone"`
	AlternativePhones []*string `json:"alternativePhones"`
	Address           *string   `json:"address"`
	BornDay           *string   `json:"bornDay"`
	DegreeStudy       *string   `json:"degreeStudy"`
	RelationShip      *string   `json:"relationShip"`
	Curp              *string   `json:"curp"`
	CitizenID         *string   `json:"citizenId"`
	CredentialID      *string   `json:"credentialId"`
	OriginState       *string   `json:"originState"`
	Score             *string   `json:"score"`
	Qualities         *string   `json:"qualities"`
	Defects           *string   `json:"defects"`
}

type User struct {
	ID                string    `json:"id"`
	Name              *string   `json:"name"`
	LastName          *string   `json:"lastName"`
	Username          string    `json:"username"`
	Password          string    `json:"password"`
	Admin             *bool     `json:"admin"`
	Root              *bool     `json:"root"`
	Verified          *bool     `json:"verified"`
	Reported          *bool     `json:"reported"`
	ReportReason      *string   `json:"reportReason"`
	ActiveContract    *bool     `json:"activeContract"`
	AdmissionDay      *string   `json:"admissionDay"`
	UnemploymentDay   *string   `json:"unemploymentDay"`
	WorkedHours       *int      `json:"workedHours"`
	CurrentBranch     *string   `json:"currentBranch"`
	OriginBranch      *string   `json:"originBranch"`
	MonetaryBonds     *float64  `json:"monetaryBonds"`
	MonetaryDiscounts *float64  `json:"monetaryDiscounts"`
	Mail              *string   `json:"mail"`
	AlternativeMails  []*string `json:"alternativeMails"`
	Phone             *string   `json:"phone"`
	AlternativePhones []*string `json:"alternativePhones"`
	Address           *string   `json:"address"`
	BornDay           *string   `json:"bornDay"`
	DegreeStudy       *string   `json:"degreeStudy"`
	RelationShip      *string   `json:"relationShip"`
	Curp              *string   `json:"curp"`
	CitizenID         *string   `json:"citizenId"`
	CredentialID      *string   `json:"credentialId"`
	OriginState       *string   `json:"originState"`
	Score             *string   `json:"score"`
	Qualities         *string   `json:"qualities"`
	Defects           *string   `json:"defects"`
}
