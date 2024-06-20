package drug

type drugService struct {
  repo DrugRepo
}

func NewDrugService(drugRepo DrugRepo) (*drugService) {
  return &drugService{ 
    repo: drugRepo,
  }
}

func (this *drugService) AddDrug(drug *Drug) error {
  return this.repo.Add(drug)
}

func (this *drugService)  GetAll() ([]Drug, error){
  return this.repo.GetAll()
}

  // GetOne(drugId string) (Drug error)
func (this *drugService) GetOne(drugId int) (Drug, error) {
  return this.repo.GetOne(drugId)
}

func (this *drugService) UpdateDrug(drugId int, drugUpdate *Drug) (Drug, error){
  return this.repo.Update(drugId, drugUpdate)
}

func (this *drugService) Archive(drugId int) (error) {
  return this.repo.Archive(drugId)
}
