package supplier

import(
	"net/http"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
	"thosai-chutney/core/misc"
	"thosai-chutney/core/auth"
)

func createSupplierFromSignup(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var supplier Supplier
		err := json.NewDecoder(request.Body).Decode(&supplier)
		utils.CheckError(err)

		var createResponse = misc.IdReturnStruct{}
		createResponse.Id = CreateSupplier(conn, &supplier)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(createResponse)
	}
}

func login(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var supplier misc.LoginRequest
		err := json.NewDecoder(request.Body).Decode(&supplier)
		utils.CheckError(err)

		searchResult := FindSupplier(conn, supplier.Id)
		if searchResult.Password == supplier.Password {
			token, err := auth.GenerateToken(supplier.Id)
			utils.CheckError(err)
			tokenResponse := misc.TokenReturnStruct{token}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			json.NewEncoder(writer).Encode(tokenResponse)
		} else {
			http.Error(writer, "Bad Request: Incorrect password", http.StatusBadRequest)
		}

	}
}

func SupplierRouter(conn *pgxpool.Pool) *http.ServeMux {
	supplierRouter := http.NewServeMux()
	supplierRouter.HandleFunc("/create", createSupplierFromSignup(conn))
	return supplierRouter
}
