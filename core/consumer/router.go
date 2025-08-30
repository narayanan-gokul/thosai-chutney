package consumer

import(
	"net/http"
	"encoding/json"

	"github.com/jackc/pgx/v5/pgxpool"
	"thosai-chutney/utils"
	"thosai-chutney/core/misc"
	"thosai-chutney/core/auth"
)

func createConsumerFromSignup(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var consumer Consumer
		err := json.NewDecoder(request.Body).Decode(&consumer)
		utils.CheckError(err)
		defer request.Body.Close()

		var createResponse = misc.IdReturnStruct{}
		createResponse.Id = CreateConsumer(conn, &consumer)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
		json.NewEncoder(writer).Encode(createResponse)
	}
}

func login(conn *pgxpool.Pool) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var consumer misc.LoginRequest
		err := json.NewDecoder(request.Body).Decode(&consumer)
		utils.CheckError(err)

		searchResult := FindConsumer(conn, consumer.Id)
		if searchResult.Password == consumer.Password {
			token, err := auth.GenerateToken(consumer.Id)
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

func ConsumerRouter(conn *pgxpool.Pool) *http.ServeMux {
	consumerRouter := http.NewServeMux()
	consumerRouter.HandleFunc("/create", createConsumerFromSignup(conn))
	consumerRouter.HandleFunc("/login", login(conn))
	return consumerRouter
}
