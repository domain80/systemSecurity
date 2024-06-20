package drug

type DrugRepo interface {
  Add(drug *Drug) error
  GetAll() ([]Drug, error)
  GetOne(drugId int) (Drug, error)
  Update(drugId int, updatedDrug *Drug) (Drug, error)
  Archive(drugId int) (error)
}

