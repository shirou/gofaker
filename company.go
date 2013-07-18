package gofaker


type Company struct {
	Name string
	Phone string
}

func NewCompany(locale string) Company{
	b := Company{}

	b.Name = GetValue(locale, "company_names")
	b.Phone = GetValue(locale, "phone")

	return b
}
