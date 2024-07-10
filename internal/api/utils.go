// Add these imports if not already present
package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	// "github.com/Deepbinder-main/cc-backend/internal/repository"
	sqlcdb "github.com/Deepbinder-main/cc-backend/internal/repository/sqlc/db"
	// "github.com/Deepbinder-main/cc-backend/internal/repository"

	"github.com/gorilla/mux"
	"github.com/oklog/ulid/v2"
)

// Add these methods to the Service struct

// CreateMachine godoc
//
//	@summary    Creates a new machine record
//	@tags       Machine
//	@accept     mpfd
//	@produce    json
//	@param      machine_id  formData    string          true    "Machine ID"
//	@param      hostname    formData    string          true    "Hostname"
//	@param      os_version  formData    string          true    "OS Version"
//	@param      ip_address  formData    string          true    "IP Address"
//	@success    201         {object}    Machine         "Created machine"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine [post]
func (api *Service) CreateMachine(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreateMachineParams{
		MachineID: ulid.Make().String(),
		Hostname:  r.FormValue("hostname"),
		OsVersion: r.FormValue("os_version"),
		IpAddress: r.FormValue("ip_address"),
	}

	if params.MachineID == "" || params.Hostname == "" || params.OsVersion == "" || params.IpAddress == "" {
		log.Printf("all fields are required")
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreateMachine(r.Context(), params)
	if err != nil {
		log.Printf("error creating machine: %v", err)
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetMachine godoc
//
//	@summary    Retrieves a machine record
//	@tags       Machine
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@success    200         {object}    Machine         "Retrieved machine"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine/{machine_id} [get]
func (api *Service) GetMachine(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	machine, err := api.r.GetMachine(r.Context(), machineID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(machine)
}

// UpdateMachine godoc
//
//	@summary    Updates a machine record
//	@tags       Machine
//	@accept     mpfd
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@param      hostname    formData    string          true    "Hostname"
//	@param      os_version  formData    string          true    "OS Version"
//	@param      ip_address  formData    string          true    "IP Address"
//	@success    200         {object}    Machine         "Updated machine"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine/{machine_id} [put]
func (api *Service) UpdateMachine(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	params := sqlcdb.UpdateMachineParams{
		MachineID: machineID,
		Hostname:  r.FormValue("hostname"),
		OsVersion: r.FormValue("os_version"),
		IpAddress: r.FormValue("ip_address"),
	}

	if params.Hostname == "" || params.OsVersion == "" || params.IpAddress == "" {
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.UpdateMachine(r.Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeleteMachine godoc
//
//	@summary    Deletes a machine record
//	@tags       Machine
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@success    204         "No Content"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine/{machine_id} [delete]
func (api *Service) DeleteMachine(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	err = api.r.DeleteMachine(r.Context(), machineID)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// TODO : ListMachines still need to develop

// ListMachines godoc
//
//	@summary    Lists all machines
//	@tags       Machine
//	@produce    json
//	@success    200         {array}     Machine         "List of machines"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /machines [get]
func (api *Service) ListMachines(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	// Assuming you have a ListMachines query in your SQLC api.r
	// machines, err := api.r.ListMachines(r.Context())
	// return still development

	json.NewEncoder(rw).Encode("still development")
}

// Add these methods to the Service struct

// CreateMachineConf godoc
//
//	@summary    Creates a new machine configuration
//	@tags       MachineConf
//	@accept     mpfd
//	@produce    json
//	@param      machine_id   formData    string          true    "Machine ID"
//	@param      hostname     formData    string          true    "Hostname"
//	@param      username     formData    string          true    "Username"
//	@param      passphrase   formData    string          false   "Passphrase"
//	@param      port_number  formData    int             true    "Port Number"
//	@param      password     formData    string          false   "Password"
//	@param      host_key     formData    string          false   "Host Key"
//	@param      folder_path  formData    string          false   "Folder Path"
//	@success    201          {object}    MachineConf     "Created machine configuration"
//	@failure    400          {object}    ErrorResponse   "Bad Request"
//	@failure    500          {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine_conf [post]
func (api *Service) CreateMachineConf(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	portNumber, err := strconv.Atoi(r.FormValue("port_number"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.CreateMachineConfParams{
		MachineID:  r.FormValue("machine_id"),
		Hostname:   r.FormValue("hostname"),
		Username:   r.FormValue("username"),
		Passphrase: sql.NullString{String: r.FormValue("passphrase"), Valid: r.FormValue("passphrase") != ""},
		PortNumber: int32(portNumber),
		Password:   sql.NullString{String: r.FormValue("password"), Valid: r.FormValue("password") != ""},
		HostKey:    sql.NullString{String: r.FormValue("host_key"), Valid: r.FormValue("host_key") != ""},
		FolderPath: sql.NullString{String: r.FormValue("folder_path"), Valid: r.FormValue("folder_path") != ""},
	}

	err = api.r.CreateMachineConf(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetMachineConf godoc
//
//	@summary    Retrieves a machine configuration
//	@tags       MachineConf
//	@produce    json
//	@param      machine_id   path        string          true    "Machine ID"
//	@success    200          {object}    MachineConf     "Retrieved machine configuration"
//	@failure    404          {object}    ErrorResponse   "Not Found"
//	@failure    500          {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine_conf/{machine_id} [get]
func (api *Service) GetMachineConf(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	machineConf, err := api.r.GetMachineConf(r.Context(), machineID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(machineConf)
}

// UpdateMachineConf godoc
//
//	@summary    Updates a machine configuration
//	@tags       MachineConf
//	@accept     mpfd
//	@produce    json
//	@param      id           path        int             true    "Machine Configuration ID"
//	@param      hostname     formData    string          true    "Hostname"
//	@param      username     formData    string          true    "Username"
//	@param      passphrase   formData    string          false   "Passphrase"
//	@param      port_number  formData    int             true    "Port Number"
//	@param      password     formData    string          false   "Password"
//	@param      host_key     formData    string          false   "Host Key"
//	@param      folder_path  formData    string          false   "Folder Path"
//	@success    200          {object}    MachineConf     "Updated machine configuration"
//	@failure    400          {object}    ErrorResponse   "Bad Request"
//	@failure    404          {object}    ErrorResponse   "Not Found"
//	@failure    500          {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine_conf/{id} [put]
func (api *Service) UpdateMachineConf(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	portNumber, err := strconv.Atoi(r.FormValue("port_number"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.UpdateMachineConfParams{
		ID:         int32(id),
		Hostname:   r.FormValue("hostname"),
		Username:   r.FormValue("username"),
		Passphrase: sql.NullString{String: r.FormValue("passphrase"), Valid: r.FormValue("passphrase") != ""},
		PortNumber: int32(portNumber),
		Password:   sql.NullString{String: r.FormValue("password"), Valid: r.FormValue("password") != ""},
		HostKey:    sql.NullString{String: r.FormValue("host_key"), Valid: r.FormValue("host_key") != ""},
		FolderPath: sql.NullString{String: r.FormValue("folder_path"), Valid: r.FormValue("folder_path") != ""},
	}

	err = api.r.UpdateMachineConf(r.Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeleteMachineConf godoc
//
//	@summary    Deletes a machine configuration
//	@tags       MachineConf
//	@produce    json
//	@param      id           path        int             true    "Machine Configuration ID"
//	@success    204          "No Content"
//	@failure    400          {object}    ErrorResponse   "Bad Request"
//	@failure    500          {object}    ErrorResponse   "Internal Server Error"
//	@router     /machine_conf/{id} [delete]
func (api *Service) DeleteMachineConf(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeleteMachineConf(r.Context(), int32(id))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Add these methods to the Service struct

// CreateRabbitMQConfig godoc
//
//	@summary    Creates or updates RabbitMQ configuration
//	@tags       RabbitMQ
//	@accept     mpfd
//	@produce    json
//	@param      conn_url    formData    string          true    "Connection URL"
//	@param      username    formData    string          true    "Username"
//	@param      password    formData    string          true    "Password"
//	@success    200         {object}    RabbitMqConfig  "Created/Updated RabbitMQ configuration"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /rabbitmq_config [post]
func (api *Service) CreateRabbitMQConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreateRabbitMQConfigParams{
		ConnUrl:  r.FormValue("conn_url"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	err = api.r.CreateRabbitMQConfig(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(params)
}

// GetRabbitMQConfig godoc
//
//	@summary    Retrieves the RabbitMQ configuration
//	@tags       RabbitMQ
//	@produce    json
//	@success    200         {object}    RabbitMqConfig  "Retrieved RabbitMQ configuration"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /rabbitmq_config [get]
func (api *Service) GetRabbitMQConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	config, err := api.r.GetRabbitMQConfig(r.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(config)
}

// UpdateRabbitMQConfig godoc
//
//	@summary    Updates the RabbitMQ configuration
//	@tags       RabbitMQ
//	@accept     mpfd
//	@produce    json
//	@param      conn_url    formData    string          true    "Connection URL"
//	@param      username    formData    string          true    "Username"
//	@param      password    formData    string          true    "Password"
//	@success    200         {object}    RabbitMqConfig  "Updated RabbitMQ configuration"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /rabbitmq_config [put]
func (api *Service) UpdateRabbitMQConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.UpdateRabbitMQConfigParams{
		ConnUrl:  r.FormValue("conn_url"),
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	err = api.r.UpdateRabbitMQConfig(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(params)
}

// DeleteRabbitMQConfig godoc
//
//	@summary    Deletes the RabbitMQ configuration
//	@tags       RabbitMQ
//	@produce    json
//	@success    204         "No Content"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /rabbitmq_config [delete]
func (api *Service) DeleteRabbitMQConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	err = api.r.DeleteRabbitMQConfig(r.Context())
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Add these methods to the Service struct

// CreateInfluxDBConfig godoc
//
//	@summary    Creates or updates InfluxDB configuration
//	@tags       InfluxDB
//	@accept     mpfd
//	@produce    json
//	@param      type                   formData    string          true    "Type"
//	@param      database_name          formData    string          true    "Database Name"
//	@param      host                   formData    string          true    "Host"
//	@param      port                   formData    int             true    "Port"
//	@param      user                   formData    string          true    "User"
//	@param      password               formData    string          true    "Password"
//	@param      organization           formData    string          true    "Organization"
//	@param      ssl_enabled            formData    bool            true    "SSL Enabled"
//	@param      batch_size             formData    int             true    "Batch Size"
//	@param      retry_interval         formData    string          true    "Retry Interval"
//	@param      retry_exponential_base formData    int             true    "Retry Exponential Base"
//	@param      max_retries            formData    int             true    "Max Retries"
//	@param      max_retry_time         formData    string          true    "Max Retry Time"
//	@param      meta_as_tags           formData    string          false   "Meta as Tags"
//	@success    200                    {object}    InfluxdbConfiguration  "Created/Updated InfluxDB configuration"
//	@failure    400                    {object}    ErrorResponse   "Bad Request"
//	@failure    500                    {object}    ErrorResponse   "Internal Server Error"
//	@router     /influxdb_config [post]
func (api *Service) CreateInfluxDBConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	port, err := strconv.Atoi(r.FormValue("port"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	batchSize, err := strconv.Atoi(r.FormValue("batch_size"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	retryExponentialBase, err := strconv.Atoi(r.FormValue("retry_exponential_base"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	maxRetries, err := strconv.Atoi(r.FormValue("max_retries"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	sslEnabled, err := strconv.ParseBool(r.FormValue("ssl_enabled"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.CreateInfluxDBConfigurationParams{
		Type:                 r.FormValue("type"),
		DatabaseName:         r.FormValue("database_name"),
		Host:                 r.FormValue("host"),
		Port:                 int32(port),
		User:                 r.FormValue("user"),
		Password:             r.FormValue("password"),
		Organization:         r.FormValue("organization"),
		SslEnabled:           sslEnabled,
		BatchSize:            int32(batchSize),
		RetryInterval:        r.FormValue("retry_interval"),
		RetryExponentialBase: int32(retryExponentialBase),
		MaxRetries:           int32(maxRetries),
		MaxRetryTime:         r.FormValue("max_retry_time"),
		MetaAsTags:           sql.NullString{String: r.FormValue("meta_as_tags"), Valid: r.FormValue("meta_as_tags") != ""},
	}

	err = api.r.CreateInfluxDBConfiguration(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(params)
}

// GetInfluxDBConfig godoc
//
//	@summary    Retrieves the InfluxDB configuration
//	@tags       InfluxDB
//	@produce    json
//	@success    200         {object}    InfluxdbConfiguration  "Retrieved InfluxDB configuration"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /influxdb_config [get]
func (api *Service) GetInfluxDBConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	config, err := api.r.GetInfluxDBConfiguration(r.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(config)
}

// UpdateInfluxDBConfig godoc
//
//	@summary    Updates the InfluxDB configuration
//	@tags       InfluxDB
//	@accept     mpfd
//	@produce    json
//	@param      type                   formData    string          true    "Type"
//	@param      database_name          formData    string          true    "Database Name"
//	@param      host                   formData    string          true    "Host"
//	@param      port                   formData    int             true    "Port"
//	@param      user                   formData    string          true    "User"
//	@param      password               formData    string          true    "Password"
//	@param      organization           formData    string          true    "Organization"
//	@param      ssl_enabled            formData    bool            true    "SSL Enabled"
//	@param      batch_size             formData    int             true    "Batch Size"
//	@param      retry_interval         formData    string          true    "Retry Interval"
//	@param      retry_exponential_base formData    int             true    "Retry Exponential Base"
//	@param      max_retries            formData    int             true    "Max Retries"
//	@param      max_retry_time         formData    string          true    "Max Retry Time"
//	@param      meta_as_tags           formData    string          false   "Meta as Tags"
//	@success    200                    {object}    InfluxdbConfiguration  "Updated InfluxDB configuration"
//	@failure    400                    {object}    ErrorResponse   "Bad Request"
//	@failure    500                    {object}    ErrorResponse   "Internal Server Error"
//	@router     /influxdb_config [put]
func (api *Service) UpdateInfluxDBConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	port, err := strconv.Atoi(r.FormValue("port"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	batchSize, err := strconv.Atoi(r.FormValue("batch_size"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	retryExponentialBase, err := strconv.Atoi(r.FormValue("retry_exponential_base"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	maxRetries, err := strconv.Atoi(r.FormValue("max_retries"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	sslEnabled, err := strconv.ParseBool(r.FormValue("ssl_enabled"))
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.UpdateInfluxDBConfigurationParams{
		Type:                 r.FormValue("type"),
		DatabaseName:         r.FormValue("database_name"),
		Host:                 r.FormValue("host"),
		Port:                 int32(port),
		User:                 r.FormValue("user"),
		Password:             r.FormValue("password"),
		Organization:         r.FormValue("organization"),
		SslEnabled:           sslEnabled,
		BatchSize:            int32(batchSize),
		RetryInterval:        r.FormValue("retry_interval"),
		RetryExponentialBase: int32(retryExponentialBase),
		MaxRetries:           int32(maxRetries),
		MaxRetryTime:         r.FormValue("max_retry_time"),
		MetaAsTags:           sql.NullString{String: r.FormValue("meta_as_tags"), Valid: r.FormValue("meta_as_tags") != ""},
	}

	err = api.r.UpdateInfluxDBConfiguration(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(params)
}

// DeleteInfluxDBConfig godoc
//
//	@summary    Deletes the InfluxDB configuration
//	@tags       InfluxDB
//	@produce    json
//	@success    204         "No Content"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /influxdb_config [delete]
func (api *Service) DeleteInfluxDBConfig(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	err = api.r.DeleteInfluxDBConfiguration(r.Context(), 1) // Assuming the ID is always 1 for the single row
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Add these methods to the Service struct

// CreateFileStashURL godoc
//
//	@summary    Creates or updates File Stash URL
//	@tags       FileStash
//	@accept     mpfd
//	@produce    json
//	@param      url         formData    string          true    "File Stash URL"
//	@success    200         {object}    FileStashUrl    "Created/Updated File Stash URL"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /file_stash_url [post]
func (api *Service) CreateFileStashURL(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		handleError(errors.New("URL is required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreateFileStashURL(r.Context(), url)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	fileStashURL := sqlcdb.FileStashUrl{
		Url: url,
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(fileStashURL)
}

// GetFileStashURL godoc
//
//	@summary    Retrieves the File Stash URL
//	@tags       FileStash
//	@produce    json
//	@success    200         {object}    FileStashUrl    "Retrieved File Stash URL"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /file_stash_url [get]
func (api *Service) GetFileStashURL(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	fileStashURL, err := api.r.GetFileStashURL(r.Context())
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(fileStashURL)
}

// UpdateFileStashURL godoc
//
//	@summary    Updates the File Stash URL
//	@tags       FileStash
//	@accept     mpfd
//	@produce    json
//	@param      url         formData    string          true    "File Stash URL"
//	@success    200         {object}    FileStashUrl    "Updated File Stash URL"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /file_stash_url [put]
func (api *Service) UpdateFileStashURL(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		handleError(errors.New("URL is required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.UpdateFileStashURL(r.Context(), url)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	fileStashURL := sqlcdb.FileStashUrl{
		Url: url,
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(fileStashURL)
}

// DeleteFileStashURL godoc
//
//	@summary    Deletes the File Stash URL
//	@tags       FileStash
//	@produce    json
//	@success    204         "No Content"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /file_stash_url [delete]
func (api *Service) DeleteFileStashURL(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	err = api.r.DeleteFileStashURL(r.Context(), 1) // Assuming the ID is always 1 for the single row
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// CreateLVStorageIssuer godoc
//
//	@summary    Creates a new LV Storage Issuer
//	@tags       LVStorageIssuer
//	@accept     mpfd
//	@produce    json
//	@param      machine_id            formData    string  true    "Machine ID"
//	@param      inc_buffer            formData    int     false   "Increment Buffer"
//	@param      dec_buffer            formData    int     false   "Decrement Buffer"
//	@param      hostname              formData    string  true    "Hostname"
//	@param      username              formData    string  true    "Username"
//	@param      minAvailableSpaceGB   formData    float64 true    "Minimum Available Space in GB"
//	@param      maxAvailableSpaceGB   formData    float64 true    "Maximum Available Space in GB"
//	@success    201         {object}  LvStorageIssuer     "Created LV Storage Issuer"
//	@failure    400         {object}  ErrorResponse       "Bad Request"
//	@failure    500         {object}  ErrorResponse       "Internal Server Error"
//	@router     /lv_storage_issuer [post]
func (api *Service) CreateLVStorageIssuer(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	incBuffer, _ := strconv.Atoi(r.FormValue("inc_buffer"))
	decBuffer, _ := strconv.Atoi(r.FormValue("dec_buffer"))
	minAvailableSpaceGB, _ := strconv.ParseFloat(r.FormValue("minAvailableSpaceGB"), 64)
	maxAvailableSpaceGB, _ := strconv.ParseFloat(r.FormValue("maxAvailableSpaceGB"), 64)

	params := sqlcdb.CreateLVStorageIssuerParams{
		MachineID:           r.FormValue("machine_id"),
		IncBuffer:           sql.NullInt32{Int32: int32(incBuffer), Valid: r.FormValue("inc_buffer") != ""},
		DecBuffer:           sql.NullInt32{Int32: int32(decBuffer), Valid: r.FormValue("dec_buffer") != ""},
		Hostname:            r.FormValue("hostname"),
		Username:            r.FormValue("username"),
		Minavailablespacegb: minAvailableSpaceGB,
		Maxavailablespacegb: maxAvailableSpaceGB,
	}

	err = api.r.CreateLVStorageIssuer(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetLVStorageIssuers godoc
//
//	@summary    Retrieves all LV Storage Issuers
//	@tags       LVStorageIssuer
//	@produce    json
//	@success    200         {array}   LvStorageIssuer     "Retrieved LV Storage Issuers"
//	@failure    500         {object}  ErrorResponse       "Internal Server Error"
//	@router     /lv_storage_issuers [get]
func (api *Service) GetLVStorageIssuers(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	issuers, err := api.r.GetLVStorageIssuers(r.Context())
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	json.NewEncoder(rw).Encode(issuers)
}

// UpdateLVStorageIssuer godoc
//
//	@summary    Updates an LV Storage Issuer
//	@tags       LVStorageIssuer
//	@accept     mpfd
//	@produce    json
//	@param      id                    path        int     true    "LV Storage Issuer ID"
//	@param      inc_buffer            formData    int     false   "Increment Buffer"
//	@param      dec_buffer            formData    int     false   "Decrement Buffer"
//	@param      hostname              formData    string  true    "Hostname"
//	@param      username              formData    string  true    "Username"
//	@param      minAvailableSpaceGB   formData    float64 true    "Minimum Available Space in GB"
//	@param      maxAvailableSpaceGB   formData    float64 true    "Maximum Available Space in GB"
//	@success    200         {object}  LvStorageIssuer     "Updated LV Storage Issuer"
//	@failure    400         {object}  ErrorResponse       "Bad Request"
//	@failure    404         {object}  ErrorResponse       "Not Found"
//	@failure    500         {object}  ErrorResponse       "Internal Server Error"
//	@router     /lv_storage_issuer/{id} [put]
func (api *Service) UpdateLVStorageIssuer(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	incBuffer, _ := strconv.Atoi(r.FormValue("inc_buffer"))
	decBuffer, _ := strconv.Atoi(r.FormValue("dec_buffer"))
	minAvailableSpaceGB, _ := strconv.ParseFloat(r.FormValue("minAvailableSpaceGB"), 64)
	maxAvailableSpaceGB, _ := strconv.ParseFloat(r.FormValue("maxAvailableSpaceGB"), 64)

	params := sqlcdb.UpdateLVStorageIssuerParams{
		ID:                  int32(id),
		IncBuffer:           sql.NullInt32{Int32: int32(incBuffer), Valid: r.FormValue("inc_buffer") != ""},
		DecBuffer:           sql.NullInt32{Int32: int32(decBuffer), Valid: r.FormValue("dec_buffer") != ""},
		Hostname:            r.FormValue("hostname"),
		Username:            r.FormValue("username"),
		Minavailablespacegb: minAvailableSpaceGB,
		Maxavailablespacegb: maxAvailableSpaceGB,
	}

	err = api.r.UpdateLVStorageIssuer(r.Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeleteLVStorageIssuer godoc
//
//	@summary    Deletes an LV Storage Issuer
//	@tags       LVStorageIssuer
//	@produce    json
//	@param      id          path        int     true    "LV Storage Issuer ID"
//	@success    204         "No Content"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /lv_storage_issuer/{id} [delete]
func (api *Service) DeleteLVStorageIssuer(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeleteLVStorageIssuer(r.Context(), int32(id))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Notifications

// CreateNotification godoc
//
//	@summary    Creates a new notification
//	@tags       Notifications
//	@accept     mpfd
//	@produce    json
//	@param      message     formData    string          true    "Notification message"
//	@success    201         {object}    Notification    "Created notification"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /notifications [post]
func (api *Service) CreateNotification(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	message := r.FormValue("message")
	if message == "" {
		handleError(errors.New("message is required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreateNotification(r.Context(), message)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]string{"message": message})
}

// GetNotifications godoc
//
//	@summary    Retrieves notifications
//	@tags       Notifications
//	@produce    json
//	@param      limit       query       int             false   "Limit the number of notifications"
//	@success    200         {array}     Notification    "Retrieved notifications"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /notifications [get]
func (api *Service) GetNotifications(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10 // Default limit
	}

	notifications, err := api.r.GetNotifications(r.Context(), int32(limit))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	json.NewEncoder(rw).Encode(notifications)
}

// DeleteNotification godoc
//
//	@summary    Deletes a notification
//	@tags       Notifications
//	@produce    json
//	@param      id          path        int             true    "Notification ID"
//	@success    204         "No Content"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /notifications/{id} [delete]
func (api *Service) DeleteNotification(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeleteNotification(r.Context(), int32(id))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Realtime Logs

// CreateRealtimeLog godoc
//
//	@summary    Creates a new realtime log
//	@tags       RealtimeLogs
//	@accept     mpfd
//	@produce    json
//	@param      log_message formData    string          true    "Log message"
//	@param      machine_id  formData    string          true    "Machine ID"
//	@success    201         {object}    RealtimeLog     "Created realtime log"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /realtime_logs [post]
func (api *Service) CreateRealtimeLog(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreateRealtimeLogParams{
		LogMessage: r.FormValue("log_message"),
		MachineID:  r.FormValue("machine_id"),
	}

	if params.LogMessage == "" || params.MachineID == "" {
		handleError(errors.New("log_message and machine_id are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreateRealtimeLog(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetRealtimeLogs godoc
//
//	@summary    Retrieves realtime logs for a machine
//	@tags       RealtimeLogs
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@param      limit       query       int             false   "Limit the number of logs"
//	@success    200         {array}     RealtimeLog     "Retrieved realtime logs"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /realtime_logs/{machine_id} [get]
func (api *Service) GetRealtimeLogs(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit <= 0 {
		limit = 10 // Default limit
	}

	logs, err := api.r.GetRealtimeLogs(r.Context(), sqlcdb.GetRealtimeLogsParams{
		MachineID: machineID,
		Limit:     int32(limit),
	})
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	json.NewEncoder(rw).Encode(logs)
}

// DeleteRealtimeLog godoc
//
//	@summary    Deletes a realtime log
//	@tags       RealtimeLogs
//	@produce    json
//	@param      id          path        int             true    "Realtime Log ID"
//	@success    204         "No Content"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /realtime_logs/{id} [delete]
func (api *Service) DeleteRealtimeLog(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeleteRealtimeLog(r.Context(), int32(id))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// Volume Groups

// CreateVolumeGroup godoc
//
//	@summary    Creates a new volume group
//	@tags       VolumeGroups
//	@accept     mpfd
//	@produce    json
//	@param      machine_id  formData    string          true    "Machine ID"
//	@param      vg_name     formData    string          true    "Volume Group Name"
//	@param      pv_count    formData    string          true    "PV Count"
//	@param      lv_count    formData    string          true    "LV Count"
//	@param      snap_count  formData    string          true    "Snap Count"
//	@param      vg_attr     formData    string          true    "VG Attributes"
//	@param      vg_size     formData    string          true    "VG Size"
//	@param      vg_free     formData    string          true    "VG Free"
//	@success    201         {object}    VolumeGroup     "Created volume group"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /volume_groups [post]
func (api *Service) CreateVolumeGroup(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreateVolumeGroupParams{
		MachineID: r.FormValue("machine_id"),
		VgName:    r.FormValue("vg_name"),
		PvCount:   r.FormValue("pv_count"),
		LvCount:   r.FormValue("lv_count"),
		SnapCount: r.FormValue("snap_count"),
		VgAttr:    r.FormValue("vg_attr"),
		VgSize:    r.FormValue("vg_size"),
		VgFree:    r.FormValue("vg_free"),
	}

	err = api.r.CreateVolumeGroup(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetVolumeGroups godoc
//
//	@summary    Retrieves volume groups for a machine
//	@tags       VolumeGroups
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@success    200         {array}     VolumeGroup     "Retrieved volume groups"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /volume_groups/{machine_id} [get]
func (api *Service) GetVolumeGroups(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	volumeGroups, err := api.r.GetVolumeGroups(r.Context(), machineID)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	json.NewEncoder(rw).Encode(volumeGroups)
}

// UpdateVolumeGroup godoc
//
//	@summary    Updates a volume group
//	@tags       VolumeGroups
//	@accept     mpfd
//	@produce    json
//	@param      vg_id       path        int             true    "Volume Group ID"
//	@param      vg_name     formData    string          true    "Volume Group Name"
//	@param      pv_count    formData    string          true    "PV Count"
//	@param      lv_count    formData    string          true    "LV Count"
//	@param      snap_count  formData    string          true    "Snap Count"
//	@param      vg_attr     formData    string          true    "VG Attributes"
//	@param      vg_size     formData    string          true    "VG Size"
//	@param      vg_free     formData    string          true    "VG Free"
//	@success    200         {object}    VolumeGroup     "Updated volume group"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /volume_groups/{vg_id} [put]
func (api *Service) UpdateVolumeGroup(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	vgID, err := strconv.Atoi(mux.Vars(r)["vg_id"])
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.UpdateVolumeGroupParams{
		VgID:      int32(vgID),
		VgName:    r.FormValue("vg_name"),
		PvCount:   r.FormValue("pv_count"),
		LvCount:   r.FormValue("lv_count"),
		SnapCount: r.FormValue("snap_count"),
		VgAttr:    r.FormValue("vg_attr"),
		VgSize:    r.FormValue("vg_size"),
		VgFree:    r.FormValue("vg_free"),
	}

	err = api.r.UpdateVolumeGroup(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeleteVolumeGroup godoc
//
//	@summary    Deletes a volume group
//	@tags       VolumeGroup
//	@produce    json
//	@param      group_id    path        string          true    "Volume Group ID"
//	@success    204         "No Content"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /volume_group/{group_id} [delete]
func (api *Service) DeleteVolumeGroup(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupIDStr := vars["group_id"]

	// Convert group_id to int32
	groupIDInt, err := strconv.Atoi(groupIDStr)
	if err != nil {
		http.Error(rw, "Invalid group ID", http.StatusBadRequest)
		return
	}
	groupID := int32(groupIDInt)

	err = api.r.DeleteVolumeGroup(r.Context(), groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(rw, "Volume Group not found", http.StatusNotFound)
		} else {
			http.Error(rw, "Failed to delete Volume Group", http.StatusInternalServerError)
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// CreatePhysicalVolume godoc
//
//	@summary    Creates a new physical volume record
//	@tags       PhysicalVolume
//	@accept     mpfd
//	@produce    json
//	@param      machine_id  formData    string  true    "Machine ID"
//	@param      pv_name     formData    string  true    "Physical Volume Name"
//	@param      vg_name     formData    string  true    "Volume Group Name"
//	@param      pv_fmt      formData    string  true    "Physical Volume Format"
//	@param      pv_attr     formData    string  true    "Physical Volume Attributes"
//	@param      pv_size     formData    string  true    "Physical Volume Size"
//	@param      pv_free     formData    string  true    "Physical Volume Free Space"
//	@success    201         {object}    PhysicalVolume  "Created physical volume"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /physical_volume [post]
func (api *Service) CreatePhysicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreatePhysicalVolumeParams{
		MachineID: r.FormValue("machine_id"),
		PvName:    r.FormValue("pv_name"),
		VgName:    r.FormValue("vg_name"),
		PvFmt:     r.FormValue("pv_fmt"),
		PvAttr:    r.FormValue("pv_attr"),
		PvSize:    r.FormValue("pv_size"),
		PvFree:    r.FormValue("pv_free"),
	}

	if params.MachineID == "" || params.PvName == "" || params.VgName == "" || params.PvFmt == "" || params.PvAttr == "" || params.PvSize == "" || params.PvFree == "" {
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreatePhysicalVolume(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetPhysicalVolumes godoc
//
//	@summary    Retrieves physical volume records for a machine
//	@tags       PhysicalVolume
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@success    200         {array}     PhysicalVolume  "Retrieved physical volumes"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /physical_volumes/{machine_id} [get]
func (api *Service) GetPhysicalVolumes(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	physicalVolumes, err := api.r.GetPhysicalVolumes(r.Context(), machineID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(physicalVolumes)
}

// UpdatePhysicalVolume godoc
//
//	@summary    Updates a physical volume record
//	@tags       PhysicalVolume
//	@accept     mpfd
//	@produce    json
//	@param      pv_id       path        int     true    "Physical Volume ID"
//	@param      pv_name     formData    string  true    "Physical Volume Name"
//	@param      vg_name     formData    string  true    "Volume Group Name"
//	@param      pv_fmt      formData    string  true    "Physical Volume Format"
//	@param      pv_attr     formData    string  true    "Physical Volume Attributes"
//	@param      pv_size     formData    string  true    "Physical Volume Size"
//	@param      pv_free     formData    string  true    "Physical Volume Free Space"
//	@success    200         {object}    PhysicalVolume  "Updated physical volume"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /physical_volume/{pv_id} [put]
func (api *Service) UpdatePhysicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	pvIDStr := mux.Vars(r)["pv_id"]
	pvID, err := strconv.Atoi(pvIDStr)
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.UpdatePhysicalVolumeParams{
		PvID:   int32(pvID),
		PvName: r.FormValue("pv_name"),
		VgName: r.FormValue("vg_name"),
		PvFmt:  r.FormValue("pv_fmt"),
		PvAttr: r.FormValue("pv_attr"),
		PvSize: r.FormValue("pv_size"),
		PvFree: r.FormValue("pv_free"),
	}

	if params.PvName == "" || params.VgName == "" || params.PvFmt == "" || params.PvAttr == "" || params.PvSize == "" || params.PvFree == "" {
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.UpdatePhysicalVolume(r.Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeletePhysicalVolume godoc
//
//	@summary    Deletes a physical volume record
//	@tags       PhysicalVolume
//	@produce    json
//	@param      pv_id   path        int             true    "Physical Volume ID"
//	@success    204     "No Content"
//	@failure    400     {object}    ErrorResponse   "Bad Request"
//	@failure    500     {object}    ErrorResponse   "Internal Server Error"
//	@router     /physical_volume/{pv_id} [delete]
func (api *Service) DeletePhysicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	pvIDStr := mux.Vars(r)["pv_id"]
	pvID, err := strconv.Atoi(pvIDStr)
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeletePhysicalVolume(r.Context(), int32(pvID))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}



// CreateLogicalVolume godoc
//
//	@summary    Creates a new logical volume record
//	@tags       LogicalVolume
//	@accept     mpfd
//	@produce    json
//	@param      machine_id  formData    string          true    "Machine ID"
//	@param      lv_name     formData    string          true    "Logical Volume Name"
//	@param      vg_name     formData    string          true    "Volume Group Name"
//	@param      lv_attr     formData    string          true    "Logical Volume Attributes"
//	@param      lv_size     formData    string          true    "Logical Volume Size"
//	@success    201         {object}    LogicalVolume   "Created logical volume"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /logical_volume [post]
func (api *Service) CreateLogicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	params := sqlcdb.CreateLogicalVolumeParams{
		MachineID: r.FormValue("machine_id"),
		LvName:    r.FormValue("lv_name"),
		VgName:    r.FormValue("vg_name"),
		LvAttr:    r.FormValue("lv_attr"),
		LvSize:    r.FormValue("lv_size"),
	}

	if params.MachineID == "" || params.LvName == "" || params.VgName == "" || params.LvAttr == "" || params.LvSize == "" {
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.CreateLogicalVolume(r.Context(), params)
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(params)
}

// GetLogicalVolumes godoc
//
//	@summary    Retrieves logical volume records for a machine
//	@tags       LogicalVolume
//	@produce    json
//	@param      machine_id  path        string          true    "Machine ID"
//	@success    200         {array}     LogicalVolume   "Retrieved logical volumes"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /logical_volumes/{machine_id} [get]
func (api *Service) GetLogicalVolumes(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	machineID := mux.Vars(r)["machine_id"]

	logicalVolumes, err := api.r.GetLogicalVolumes(r.Context(), machineID)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(logicalVolumes)
}

// UpdateLogicalVolume godoc
//
//	@summary    Updates a logical volume record
//	@tags       LogicalVolume
//	@accept     mpfd
//	@produce    json
//	@param      lv_id       path        int             true    "Logical Volume ID"
//	@param      lv_name     formData    string          true    "Logical Volume Name"
//	@param      vg_name     formData    string          true    "Volume Group Name"
//	@param      lv_attr     formData    string          true    "Logical Volume Attributes"
//	@param      lv_size     formData    string          true    "Logical Volume Size"
//	@success    200         {object}    LogicalVolume   "Updated logical volume"
//	@failure    400         {object}    ErrorResponse   "Bad Request"
//	@failure    404         {object}    ErrorResponse   "Not Found"
//	@failure    500         {object}    ErrorResponse   "Internal Server Error"
//	@router     /logical_volume/{lv_id} [put]
func (api *Service) UpdateLogicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	lvIDStr := mux.Vars(r)["lv_id"]
	lvID, err := strconv.Atoi(lvIDStr)
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	params := sqlcdb.UpdateLogicalVolumeParams{
		LvID:   int32(lvID),
		LvName: r.FormValue("lv_name"),
		VgName: r.FormValue("vg_name"),
		LvAttr: r.FormValue("lv_attr"),
		LvSize: r.FormValue("lv_size"),
	}

	if params.LvName == "" || params.VgName == "" || params.LvAttr == "" || params.LvSize == "" {
		handleError(errors.New("all fields are required"), http.StatusBadRequest, rw)
		return
	}

	err = api.r.UpdateLogicalVolume(r.Context(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			handleError(err, http.StatusNotFound, rw)
		} else {
			handleError(err, http.StatusInternalServerError, rw)
		}
		return
	}

	json.NewEncoder(rw).Encode(params)
}

// DeleteLogicalVolume godoc
//
//	@summary    Deletes a logical volume record
//	@tags       LogicalVolume
//	@produce    json
//	@param      lv_id   path        int             true    "Logical Volume ID"
//	@success    204     "No Content"
//	@failure    400     {object}    ErrorResponse   "Bad Request"
//	@failure    500     {object}    ErrorResponse   "Internal Server Error"
//	@router     /logical_volume/{lv_id} [delete]
func (api *Service) DeleteLogicalVolume(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	lvIDStr := mux.Vars(r)["lv_id"]
	lvID, err := strconv.Atoi(lvIDStr)
	if err != nil {
		handleError(err, http.StatusBadRequest, rw)
		return
	}

	err = api.r.DeleteLogicalVolume(r.Context(), int32(lvID))
	if err != nil {
		handleError(err, http.StatusInternalServerError, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
    
