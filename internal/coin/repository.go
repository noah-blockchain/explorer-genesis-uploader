package coin

import (
	"sync"

	"github.com/go-pg/pg"
	"github.com/noah-blockchain/coinExplorer-tools/models"
)

type Repository struct {
	db       *pg.DB
	cache    *sync.Map
	invCache *sync.Map
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db:       db,
		cache:    new(sync.Map), //TODO: добавить реализацию очистки
		invCache: new(sync.Map), //TODO: добавить реализацию очистки
	}
}

// Find coin id by symbol
func (r *Repository) FindIdBySymbol(symbol string) (uint64, error) {
	//First look in the cache
	id, ok := r.cache.Load(symbol)
	if ok {
		return id.(uint64), nil
	}
	coin := new(models.Coin)
	err := r.db.Model(coin).
		Column("id").
		Where("symbol = ?", symbol).
		Select()

	if err != nil {
		return 0, err
	}
	r.cache.Store(symbol, coin.ID)
	return coin.ID, nil
}

// Find coin by symbol
func (r *Repository) FindCoinByID(id uint64) (*models.Coin, error) {
	coin := new(models.Coin)
	err := r.db.Model(coin).
		Column("id", "crr", "volume", "reserve_balance", "name", "price", "delegated", "updated_at", "capitalization").
		Where("id = ?", id).
		Select()

	if err != nil {
		return nil, err
	}
	return coin, nil
}

func (r *Repository) FindSymbolById(id uint64) (string, error) {
	//First look in the cache
	symbol, ok := r.invCache.Load(id)
	if ok {
		return symbol.(string), nil
	}
	coin := &models.Coin{ID: id}
	err := r.db.Model(coin).
		Where("id = ?", id).
		Limit(1).
		Select()

	if err != nil {
		return "", err
	}
	r.cache.Store(coin.Symbol, id)
	r.invCache.Store(id, coin.Symbol)
	return coin.Symbol, nil
}

func (r *Repository) Save(c *models.Coin) error {
	_, err := r.db.Model(c).
		Where("symbol = ?symbol").
		OnConflict("DO NOTHING"). //TODO: change to DO UPDATE
		SelectOrInsert()
	if err != nil {
		return err
	}
	r.cache.Store(c.Symbol, c.ID)
	return nil
}

func (r Repository) SaveAllIfNotExist(coins []*models.Coin) error {
	_, err := r.db.Model(&coins).OnConflict("(symbol) DO UPDATE").Insert()
	if err != nil {
		return err
	}
	for _, coin := range coins {
		r.cache.Store(coin.Symbol, coin.ID)
		r.invCache.Store(coin.ID, coin.Symbol)
	}
	return err
}

func (r *Repository) GetAllCoins() ([]*models.Coin, error) {
	var coins []*models.Coin
	err := r.db.Model(&coins).Order("symbol ASC").Select()
	return coins, err
}

func (r Repository) DeleteBySymbol(symbol string) error {
	coin := &models.Coin{Symbol: symbol}
	_, err := r.db.Model(coin).Where("symbol = ?symbol").Delete()
	return err
}

func (r *Repository) UpdateCoinOwner(symbol string, creationAddressID uint64) error {
	coin := models.Coin{CreationAddressID: &creationAddressID}
	_, err := r.db.Model(&coin).Column("creation_address_id").Where("symbol = ?", symbol).Update()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateCoinDelegation(id uint64, delegated uint64) error {
	coin := models.Coin{Delegated: delegated}
	_, err := r.db.Model(&coin).Column("delegated").Where("id = ?", id).Update()
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) ResetCoinDelegationNotInListIds(idList []uint64) error {
	if len(idList) > 0 {
		_, err := r.db.Query(nil, `update coins set delegated=0 where id not in (?) and delegated > 0;`, pg.In(idList))
		return err
	}
	return nil
}

func (r *Repository) UpdateCoinTransaction(symbol string, creationTransactionID uint64) error {
	coin := models.Coin{CreationTransactionID: &creationTransactionID}
	_, err := r.db.Model(&coin).Column("creation_transaction_id").Where("symbol = ?", symbol).Update()
	if err != nil {
		return err
	}
	return nil
}
