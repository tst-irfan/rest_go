// path db/query.go

package db

type QueryHelper[T any] struct {
	Model T
}

type QueryHelperInterface[T any] interface {
	FindAll() ([]T, error)
	FindAllWithPagination(page int, size int) ([]T, error)
	FindByID(ID uint) (*T, error)
	Create(data T) (*T, error)
	Update(data T) (*T, error)
	DeleteByID(ID uint) error
	FindOneByColumn(column string, value interface{}) (*T, error)
	FindManyByColumn(column string, value interface{}) ([]T, error)
	Where(query interface{}, args ...interface{}) ([]T, error)
	FirstWhere(query interface{}, args ...interface{}) (*T, error)
	DeleteWhere(query interface{}, args ...interface{}) error
	Count() (int, error)
}

func (qh *QueryHelper[T]) FindAll() ([]T, error) {
	var results []T
	err := DB.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (qh *QueryHelper[T]) FindAllWithPagination(page int, size int) ([]T, error) {
	var results []T
	err := DB.Limit(size).Offset((page - 1) * size).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (qh *QueryHelper[T]) FindByID(ID uint) (*T, error) {
	var result T
	err := DB.Where("id = ?", ID).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (qh *QueryHelper[T]) Create(data T) (*T, error) {
	err := DB.Create(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (qh *QueryHelper[T]) Update(data T) (*T, error) {
	err := DB.Save(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (qh *QueryHelper[T]) DeleteByID(ID uint) error {
	err := DB.Where("id = ?", ID).Delete(&qh.Model).Error
	if err != nil {
		return err
	}
	return nil
}

func (qh *QueryHelper[T]) FindOneByColumn(column string, value interface{}) (*T, error) {
	var result T
	err := DB.Where(column+" = ?", value).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (qh *QueryHelper[T]) FindManyByColumn(column string, value interface{}) ([]T, error) {
	var results []T
	err := DB.Where(column+" = ?", value).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (qh *QueryHelper[T]) Where(query interface{}, args ...interface{}) ([]T, error) {
	var results []T
	err := DB.Where(query, args...).Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (qh *QueryHelper[T]) FirstWhere(query interface{}, args ...interface{}) (*T, error) {
	var result T
	err := DB.Where(query, args...).First(&result).Error
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (qh *QueryHelper[T]) DeleteWhere(query interface{}, args ...interface{}) error {
	err := DB.Where(query, args...).Delete(&qh.Model).Error
	if err != nil {
		return err
	}
	return nil
}

func (qh *QueryHelper[T]) Count() (int, error) {
	var count int64
	err := DB.Model(&qh.Model).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}
