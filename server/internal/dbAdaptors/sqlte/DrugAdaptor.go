package sqlte

import (
	"security418/internal/drug"
)


// Add a new drug
func (this *adaptor) Add(drug *drug.Drug) error {
	tx, err := this.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO drugs (id, name, serial_no, tag_id, verdict, archived) VALUES (?, ?, ?, ?, ?, ?)",
		drug.ID, drug.Name, drug.SerialNo, drug.TagId, drug.Verdict, drug.Archived)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, convict := range drug.Convicts {
		_, err = tx.Exec("INSERT INTO convicts (drug_id, name, is_arrested) VALUES (?, ?, ?)",
			drug.ID, convict.Name, convict.IsArrested)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// GetAll drugs
func (this *adaptor) GetAll() ([]drug.Drug, error) {
	rows, err := this.db.Query("SELECT id, name, serial_no, tag_id, verdict, archived FROM drugs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var drugs []drug.Drug
	for rows.Next() {
		var drug drug.Drug
		err := rows.Scan(&drug.ID, &drug.Name, &drug.SerialNo, &drug.TagId, &drug.Verdict, &drug.Archived)
		if err != nil {
			return nil, err
		}
		drug.Convicts, err = this.getConvicts(drug.ID)
		if err != nil {
			return nil, err
		}
		drugs = append(drugs, drug)
	}
	return drugs, nil
}

// GetOne drug by ID
func (this *adaptor) GetOne(drugId int) (drug.Drug, error) {
	var drug drug.Drug
	err := this.db.QueryRow("SELECT id, name, serial_no, tag_id, verdict, archived FROM drugs WHERE id = ?", drugId).
		Scan(&drug.ID, &drug.Name, &drug.SerialNo, &drug.TagId, &drug.Verdict, &drug.Archived)
	if err != nil {
		return drug, err
	}
	drug.Convicts, err = this.getConvicts(drug.ID)
	return drug, err
}

// Update a drug by ID
func (this *adaptor) Update(drugId int, updatedDrug *drug.Drug) (drug.Drug, error) {
	tx, err := this.db.Begin()
	if err != nil {
		return drug.Drug{}, err
	}

	_, err = tx.Exec("UPDATE drugs SET name = ?, serial_no = ?, tag_id = ?, verdict = ?, archived = ? WHERE id = ?",
		updatedDrug.Name, updatedDrug.SerialNo, updatedDrug.TagId, updatedDrug.Verdict, updatedDrug.Archived, drugId)
	if err != nil {
		tx.Rollback()
		return drug.Drug{}, err
	}

	_, err = tx.Exec("DELETE FROM convicts WHERE drug_id = ?", drugId)
	if err != nil {
		tx.Rollback()
		return drug.Drug{}, err
	}

	for _, convict := range updatedDrug.Convicts {
		_, err = tx.Exec("INSERT INTO convicts (drug_id, name, is_arrested) VALUES (?, ?, ?)",
			drugId, convict.Name, convict.IsArrested)
		if err != nil {
			tx.Rollback()
			return drug.Drug{}, err
		}
	}

	tx.Commit()
	return *updatedDrug, nil
}

// Archive a drug by ID
func (this *adaptor) Archive(drugId int) error {
	_, err := this.db.Exec("UPDATE drugs SET archived = ? WHERE id = ?", true, drugId)
	return err
}

// getConvicts by drug ID
func (this *adaptor) getConvicts(drugId string) ([]drug.Convict, error) {
	rows, err := this.db.Query("SELECT name, is_arrested FROM convicts WHERE drug_id = ?", drugId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var convicts []drug.Convict
	for rows.Next() {
		var convict drug.Convict
		err := rows.Scan(&convict.Name, &convict.IsArrested)
		if err != nil {
			return nil, err
		}
		convicts = append(convicts, convict)
	}
	return convicts, nil
}

