package distributor

import(
	"net/http"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
	"thosai-chutney/core/misc"
)

func createDistributorFromSignup(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var c Distributor
		err := json.NewDecoder(request.Body).Decode(&c)
		utils.CheckError(err)

		var createResponse = misc.IdReturnStruct{}
		createResponse.Id = CreateDistributor(conn, &c)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(createResponse)
	}
}

func DistributorRouter(conn *pgxpool.Pool) *http.ServeMux {
	distributorRouter := http.NewServeMux()
	distributorRouter.HandleFunc("/create", createDistributorFromSignup(conn))
	return distributorRouter
}
