package drug

type Service interface {
  AddDrug(d *Drug) error
  GetAll() ([]Drug, error)
  UpdateDrug(int,  *Drug) (Drug, error)
  Archive(int) (error)
}


