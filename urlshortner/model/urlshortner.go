package model


func GetAllUrlshortners() ([]Urlshortner, error) {
	var urlshortners []Urlshortner

	tx := db.Find(&urlshortners)
	if tx.Error != nil {
		return []Urlshortner{}, tx.Error
	}

	return urlshortners, nil
}

func GetUrlshortner(id uint64) (Urlshortner, error) {
	var urlshortner Urlshortner

	tx := db.Where("id = ?", id).First(&urlshortner)

	if tx.Error != nil {
		return Urlshortner{}, tx.Error
	}

	return urlshortner, nil
}

func CreateUrlshortner(urlshortner Urlshortner) error {
	tx := db.Create(&urlshortner)
	return tx.Error
}

func UpdateUrlshortner(urlshortner Urlshortner) error {

	tx := db.Save(&urlshortner)
	return tx.Error
}

func DeleteUrlshortner(id uint64) error {

	tx := db.Unscoped().Delete(&Urlshortner{}, id)
	return tx.Error
}

func FindByUrlshortnerUrl(url string) (Urlshortner, error) {
	var urlshortner Urlshortner
	tx := db.Where("urlshortner = ?", url).First(&urlshortner)
	return urlshortner, tx.Error
}