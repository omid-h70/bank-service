package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/omid-h70/bank-service/internal/core/domain"
)

type CustomerRepositoryMySqlDB struct {
	client *sql.DB
}

func (c *CustomerRepositoryMySqlDB) FindMostActiveCustomersWithinTime(ctx context.Context, count int, time int) ([]domain.CustomerReportOut, error) {
	var (
		query string = fmt.Sprintf(`with top_txn_analysis
			as
			(
				select ctmr_txn_cnt.customer_id, txn.*, row_number() over (partition by ctmr_txn_cnt.customer_id order by txn.transaction_date desc) as row_num from transaction txn
				left join card crd on txn.card_id_from = crd.card_id or txn.card_id_to = crd.card_id
				left join account acnt on crd.account_id = acnt.account_id
				right join 
				(
					select ctmr_txn.customer_id from
					(
						select ctmr.customer_id, txn.* from transaction txn
						left join card crd on txn.card_id_from = crd.card_id or txn.card_id_to = crd.card_id
						left join account acnt on crd.account_id = acnt.account_id
						left join customer ctmr on acnt.customer_id = ctmr.customer_id
						WHERE txn.transaction_date >= (NOW() - INTERVAL 10 hour)
					) as ctmr_txn group by ctmr_txn.customer_id order by count(0) desc limit 3
				) as ctmr_txn_cnt on ctmr_txn_cnt.customer_id = acnt.customer_id
			) select * from top_txn_analysis where top_txn_analysis.row_num <= 10;
			`)
	)

	fmt.Println(query)
	return []domain.CustomerReportOut{}, nil
}

func NewCustomerRepositoryMySqlDB(clientIn *sql.DB) *CustomerRepositoryMySqlDB {
	return &CustomerRepositoryMySqlDB{clientIn}
}
