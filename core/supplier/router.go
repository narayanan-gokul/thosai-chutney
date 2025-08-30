package supplier

import(
	"net/http"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
	"thosai-chutney/core/misc"
)

func createSupplierFromSignup(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var c Supplier
		err := json.NewDecoder(request.Body).Decode(&c)
		utils.CheckError(err)

		var createResponse = misc.IdReturnStruct{}
		createResponse.Id = CreateSupplier(conn, &c)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(createResponse)
	}
}

func SupplierRouter(conn *pgxpool.Pool) *http.ServeMux {
	supplierRouter := http.NewServeMux()
	supplierRouter.HandleFunc("/create", createSupplierFromSignup(conn))
	return supplierRouter
}
