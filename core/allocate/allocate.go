package allocate

import (
	"context"
	"fmt"
	"net/http"
	"thosai-chutney/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Supply struct {
	PostCode int
	ItemId int
	Quantity int
}

type Demand struct {
	PostCode int
	ItemId int
	Quantity int
}

type Allocation struct {
	PostCode int
	ItemId int
	Fulfilled bool
}

func allocate(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		tx, err := conn.Begin(context.Background())
		utils.CheckError(err)
		tx.Rollback(context.Background())

		var supply Supply
		var totalSupply []Supply
		var demand Demand
		var totalDemand []Demand
		var allocation Allocation
		var totalAllocation []Allocation

		total_supply_rows, err := tx.Query(
			context.Background(),
			"SELECT postcode, item.item_id, SUM(quantity) AS quantity_sum FROM item JOIN shipment ON item.item_id = shipment.item_id JOIN supplier ON supplier.supp_id = shipment.supp_id GROUP BY postcode, item.item_id ORDER BY supplier.postcode ASC, item.item_id ASC, quantity_sum DESC",
		)
		utils.CheckError(err)
		defer total_supply_rows.Close()

		total_demand_rows, err := tx.Query(
			context.Background(),
			"SELECT distributor.postcode, item.item_id, SUM(quantity) AS quantity_sum FROM item JOIN cart ON cart.item_id = item.item_id JOIN consumer ON (consumer.cons_id IS NOT NULL AND consumer.cons_id = cart.cons_id) JOIN distributor ON cart.dist_id = distributor.dist_id GROUP BY distributor.postcode, item.item_id ORDER BY distributor.postcode ASC, item.item_id ASC, quantity_sum DESC",
		)
		utils.CheckError(err)
		defer total_demand_rows.Close()

		for total_supply_rows.Next() {
			err = total_supply_rows.Scan(&supply.PostCode, &supply.ItemId, &supply.Quantity)
			utils.CheckError(err)
			totalSupply = append(totalSupply, supply)
		}

		for total_demand_rows.Next() {
			err = total_demand_rows.Scan(&demand.PostCode, &demand.ItemId, &demand.Quantity)
			utils.CheckError(err)
			totalDemand = append(totalDemand, demand)
		}

		i := 0
		j := 0

		for i < len(totalSupply) {
			for j < len(totalDemand) {
				supply := totalSupply[i]
				demand := totalDemand[i]
				supplyPostCode := supply.PostCode
				demandPostCode := demand.PostCode
				if supplyPostCode > demandPostCode {
					var allocation Allocation
					allocation.PostCode = demandPostCode
					allocation.Fulfilled = false
					allocation.ItemId = totalDemand[j].ItemId
					totalAllocation = append(totalAllocation, allocation)
					j++
				} else if supplyPostCode < demandPostCode {
					i++
				} else {
					if supply.ItemId == demand.ItemId {
						allocation.PostCode = supply.PostCode
						allocation.Fulfilled = supply.Quantity >= demand.Quantity
						allocation.ItemId = supply.ItemId
						totalAllocation = append(totalAllocation, allocation)
						j++
						i++
					} else if supply.ItemId < demand.ItemId {
						i++
					} else {
						allocation.PostCode = supply.PostCode
						allocation.ItemId = demand.ItemId
						allocation.Fulfilled = false
						totalAllocation = append(totalAllocation, allocation)
						j++
					}
				}
			}
		}
		fmt.Printf("%+v\n",totalAllocation)
	}
}

func AllocationRouter(conn *pgxpool.Pool) *http.ServeMux {
	allocationRouter := http.NewServeMux()
	allocationRouter.HandleFunc("/", allocate(conn))
	return allocationRouter
}
