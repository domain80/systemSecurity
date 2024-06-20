package drug

type Drug struct {
  ID string
  Name string
  SerialNo string
  TagId string
  Convicts []Convict
  Verdict string
  Archived bool
}

type Convict struct {
  Name string
  IsArrested bool
}
