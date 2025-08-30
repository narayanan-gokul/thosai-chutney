package consumer

import(
	"net/http"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
	"thosai-chutney/core/misc"
)

func createConsumerFromSignup(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var c Consumer
		err := json.NewDecoder(request.Body).Decode(&c)
		utils.CheckError(err)

		var createResponse = misc.IdReturnStruct{}
		createResponse.Id = CreateConsumer(conn, &c)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(createResponse)
	}
}

func ConsumerRouter(conn *pgxpool.Pool) *http.ServeMux {
	consumerRouter := http.NewServeMux()
	consumerRouter.HandleFunc("/create", createConsumerFromSignup(conn))
	return consumerRouter
}
