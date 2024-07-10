// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"encoding/json"
	"fmt"
	"github.com/Deepbinder-main/cc-backend/internal/auth"
	"github.com/Deepbinder-main/cc-backend/internal/config"
	"github.com/Deepbinder-main/cc-backend/internal/graph"
	"github.com/Deepbinder-main/cc-backend/internal/repository"
	"github.com/Deepbinder-main/cc-backend/internal/util"
	"github.com/Deepbinder-main/cc-backend/pkg/log"
	"github.com/Deepbinder-main/cc-backend/pkg/schema"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

//	@title			ClusterCockpit REST API
//	@version		1.0.0
//	@description	API for batch job control.

//	@contact.name	ClusterCockpit Project
//	@contact.url	https://github.com/Deepbinder-main
//	@contact.email	support@clustercockpit.org

//	@license.name	MIT License
//	@license.url	https://opensource.org/licenses/MIT

//	@host		localhost:8080
//	@basePath	/api

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						X-Auth-Token

type RestApi struct {
	// JobRepository   *repository.JobRepository
	Resolver        *graph.Resolver
	Authentication  *auth.Authentication
	MachineStateDir string
	RepositoryMutex sync.Mutex
}

func (api *RestApi) MountRoutes(r *mux.Router) {
	r = r.PathPrefix("/api").Subrouter()
	r.StrictSlash(true)

	// r.HandleFunc("/jobs/start_job/", api.startJob).Methods(http.MethodPost, http.MethodPut)
	// r.HandleFunc("/jobs/stop_job/", api.stopJobByRequest).Methods(http.MethodPost, http.MethodPut)
	// r.HandleFunc("/jobs/stop_job/{id}", api.stopJobById).Methods(http.MethodPost, http.MethodPut)
	// // r.HandleFunc("/jobs/import/", api.importJob).Methods(http.MethodPost, http.MethodPut)

	// r.HandleFunc("/jobs/", api.getJobs).Methods(http.MethodGet)
	// r.HandleFunc("/jobs/{id}", api.getJobById).Methods(http.MethodPost)
	// r.HandleFunc("/jobs/tag_job/{id}", api.tagJob).Methods(http.MethodPost, http.MethodPatch)
	// r.HandleFunc("/jobs/metrics/{id}", api.getJobMetrics).Methods(http.MethodGet)
	// r.HandleFunc("/jobs/delete_job/", api.deleteJobByRequest).Methods(http.MethodDelete)
	// r.HandleFunc("/jobs/delete_job/{id}", api.deleteJobById).Methods(http.MethodDelete)
	// r.HandleFunc("/jobs/delete_job_before/{ts}", api.deleteJobBefore).Methods(http.MethodDelete)

	if api.MachineStateDir != "" {
		r.HandleFunc("/machine_state/{cluster}/{host}", api.getMachineState).Methods(http.MethodGet)
		r.HandleFunc("/machine_state/{cluster}/{host}", api.putMachineState).Methods(http.MethodPut, http.MethodPost)
	}

	if api.Authentication != nil {
		r.HandleFunc("/jwt/", api.getJWT).Methods(http.MethodGet)
		r.HandleFunc("/roles/", api.getRoles).Methods(http.MethodGet)
		r.HandleFunc("/users/", api.createUser).Methods(http.MethodPost, http.MethodPut)
		r.HandleFunc("/users/", api.getUsers).Methods(http.MethodGet)
		r.HandleFunc("/users/", api.deleteUser).Methods(http.MethodDelete)
		r.HandleFunc("/user/{id}", api.updateUser).Methods(http.MethodPost)
		r.HandleFunc("/configuration/", api.updateConfiguration).Methods(http.MethodPost)
		// Machine Configuration
		r.HandleFunc("/machine_conf", api.CreateMachineConf).Methods(http.MethodPost)
		r.HandleFunc("/machine_conf/{machine_id}", api.GetMachineConf).Methods(http.MethodGet)
		r.HandleFunc("/machine_conf/{id}", api.UpdateMachineConf).Methods(http.MethodPut)
		r.HandleFunc("/machine_conf/{id}", api.DeleteMachineConf).Methods(http.MethodDelete)
		// RabbitMQ Configuration
		r.HandleFunc("/rabbitmq_config", api.CreateRabbitMQConfig).Methods("POST")
		r.HandleFunc("/rabbitmq_config", api.GetRabbitMQConfig).Methods("GET")
		r.HandleFunc("/rabbitmq_config", api.UpdateRabbitMQConfig).Methods("PUT")
		r.HandleFunc("/rabbitmq_config", api.DeleteRabbitMQConfig).Methods("DELETE")
		// InfluxDB Configuration
		r.HandleFunc("/influxdb_config", api.CreateInfluxDBConfig).Methods("POST")
		r.HandleFunc("/influxdb_config", api.GetInfluxDBConfig).Methods("GET")
		r.HandleFunc("/influxdb_config", api.UpdateInfluxDBConfig).Methods("PUT")
		r.HandleFunc("/influxdb_config", api.DeleteInfluxDBConfig).Methods("DELETE")
		// File stash Configuration
		r.HandleFunc("/file_stash_url", api.CreateFileStashURL).Methods("POST")
		r.HandleFunc("/file_stash_url", api.GetFileStashURL).Methods("GET")
		r.HandleFunc("/file_stash_url", api.UpdateFileStashURL).Methods("PUT")
		r.HandleFunc("/file_stash_url", api.DeleteFileStashURL).Methods("DELETE")
		// Machine Configuration
		r.HandleFunc("/machine", api.CreateMachine).Methods("POST")
		r.HandleFunc("/machine/{machine_id}", api.GetMachine).Methods("GET")
		r.HandleFunc("/machine/{machine_id}", api.UpdateMachine).Methods("PUT")
		r.HandleFunc("/machine/{machine_id}", api.DeleteMachine).Methods("DELETE")
		r.HandleFunc("/machines", api.ListMachines).Methods("GET")
		// LV Storage Issuer routes
		r.HandleFunc("lv_storage_issuer", api.CreateLVStorageIssuer).Methods("POST")
		r.HandleFunc("lv_storage_issuers", api.GetLVStorageIssuers).Methods("GET")
		r.HandleFunc("lv_storage_issuer/{id}", api.UpdateLVStorageIssuer).Methods("PUT")
		r.HandleFunc("lv_storage_issuer/{id}", api.DeleteLVStorageIssuer).Methods("DELETE")
		// Physical Volume routes
		r.HandleFunc("physical_volumes", api.CreatePhysicalVolume).Methods("POST")
		r.HandleFunc("physical_volumes/{pv_id}", api.UpdatePhysicalVolume).Methods("PUT")
		r.HandleFunc("physical_volumes/{pv_id}", api.DeletePhysicalVolume).Methods("DELETE")

		// Notification routes
		r.HandleFunc("/notifications", api.CreateNotification).Methods("POST")
		r.HandleFunc("/notifications", api.GetNotifications).Methods("GET")
		r.HandleFunc("/notifications/{id}", api.DeleteNotification).Methods("DELETE")

		// Realtime Log routes
		r.HandleFunc("/realtime_logs", api.CreateRealtimeLog).Methods("POST")
		r.HandleFunc("/realtime_logs", api.GetRealtimeLogs).Methods("GET")
		r.HandleFunc("/realtime_logs/{id}", api.DeleteRealtimeLog).Methods("DELETE")

		// Volume Group routes
		r.HandleFunc("/volume_groups", api.CreateVolumeGroup).Methods("POST")
		r.HandleFunc("/volume_groups", api.GetVolumeGroups).Methods("GET")
		r.HandleFunc("/volume_groups/{id}", api.UpdateVolumeGroup).Methods("PUT")
		r.HandleFunc("/volume_groups/{id}", api.DeleteVolumeGroup).Methods("DELETE")

		// Logical Volume routes
		r.HandleFunc("/logical_volume", api.CreateLogicalVolume).Methods("POST")
		r.HandleFunc("/logical_volumes/{machine_id}", api.GetLogicalVolumes).Methods("GET")
		r.HandleFunc("/logical_volume/{lv_id}", api.UpdateLogicalVolume).Methods("PUT")
		r.HandleFunc("/logical_volume/{lv_id}", api.DeleteLogicalVolume).Methods("DELETE")

	}
}

// // StartJobApiResponse model
// type StartJobApiResponse struct {
// 	// Database ID of new job
// 	DBID int64 `json:"id"`
// }

// // DeleteJobApiResponse model
// type DeleteJobApiResponse struct {
// 	Message string `json:"msg"`
// }

// // UpdateUserApiResponse model
// type UpdateUserApiResponse struct {
// 	Message string `json:"msg"`
// }

// // StopJobApiRequest model
// type StopJobApiRequest struct {
// 	// Stop Time of job as epoch
// 	StopTime  int64           `json:"stopTime" validate:"required" example:"1649763839"`
// 	State     schema.JobState `json:"jobState" validate:"required" example:"completed"` // Final job state
// 	JobId     *int64          `json:"jobId" example:"123000"`                           // Cluster Job ID of job
// 	Cluster   *string         `json:"cluster" example:"fritz"`                          // Cluster of job
// 	StartTime *int64          `json:"startTime" example:"1649723812"`                   // Start Time of job as epoch
// }

// // DeleteJobApiRequest model
// type DeleteJobApiRequest struct {
// 	JobId     *int64  `json:"jobId" validate:"required" example:"123000"` // Cluster Job ID of job
// 	Cluster   *string `json:"cluster" example:"fritz"`                    // Cluster of job
// 	StartTime *int64  `json:"startTime" example:"1649723812"`             // Start Time of job as epoch
// }

// // GetJobsApiResponse model
// type GetJobsApiResponse struct {
// 	Jobs  []*schema.JobMeta `json:"jobs"`  // Array of jobs
// 	Items int               `json:"items"` // Number of jobs returned
// 	Page  int               `json:"page"`  // Page id returned
// }

// ErrorResponse model
type ErrorResponse struct {
	// Statustext of Errorcode
	Status string `json:"status"`
	Error  string `json:"error"` // Error Message
}

// ApiTag model
type ApiTag struct {
	// Tag Type
	Type string `json:"type" example:"Debug"`
	Name string `json:"name" example:"Testjob"` // Tag Name
}

// type TagJobApiRequest []*ApiTag

// type GetJobApiRequest []string

// type GetJobApiResponse struct {
// 	Meta *schema.Job
// 	Data []*JobMetricWithName
// }

// type JobMetricWithName struct {
// 	Name   string             `json:"name"`
// 	Scope  schema.MetricScope `json:"scope"`
// 	Metric *schema.JobMetric  `json:"metric"`
// }

type ApiReturnedUser struct {
	Username string   `json:"username"`
	Name     string   `json:"name"`
	Roles    []string `json:"roles"`
	Email    string   `json:"email"`
	Projects []string `json:"projects"`
}

func handleError(err error, statusCode int, rw http.ResponseWriter) {
	log.Warnf("REST ERROR : %s", err.Error())
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	json.NewEncoder(rw).Encode(ErrorResponse{
		Status: http.StatusText(statusCode),
		Error:  err.Error(),
	})
}

func decode(r io.Reader, val interface{}) error {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	return dec.Decode(val)
}

func securedCheck(r *http.Request) error {
	user := repository.GetUserFromContext(r.Context())
	if user == nil {
		return fmt.Errorf("no user in context")
	}

	if user.AuthType == schema.AuthToken {
		// If nothing declared in config: deny all request to this endpoint
		if config.Keys.ApiAllowedIPs == nil || len(config.Keys.ApiAllowedIPs) == 0 {
			return fmt.Errorf("missing configuration key ApiAllowedIPs")
		}

		if config.Keys.ApiAllowedIPs[0] == "*" {
			return nil
		}

		// extract IP address
		IPAddress := r.Header.Get("X-Real-Ip")
		if IPAddress == "" {
			IPAddress = r.Header.Get("X-Forwarded-For")
		}
		if IPAddress == "" {
			IPAddress = r.RemoteAddr
		}

		if strings.Contains(IPAddress, ":") {
			IPAddress = strings.Split(IPAddress, ":")[0]
		}

		// check if IP is allowed
		if !util.Contains(config.Keys.ApiAllowedIPs, IPAddress) {
			return fmt.Errorf("unknown ip: %v", IPAddress)
		}
	}

	return nil
}

// createUser godoc
//
//	@summary		Adds a new user
//	@tags			User
//	@description	User specified in form data will be saved to database.
//	@description	Only accessible from IPs registered with apiAllowedIPs configuration option.
//	@accept			mpfd
//	@produce		plain
//	@param			username	formData	string	true	"Unique user ID"
//	@param			password	formData	string	true	"User password"
//	@param			role		formData	string	true	"User role"	Enums(admin, support, manager, user, api)
//	@param			project		formData	string	false	"Managed project, required for new manager role user"
//	@param			name		formData	string	false	"Users name"
//	@param			email		formData	string	false	"Users email"
//	@success		200			{string}	string	"Success Response"
//	@failure		400			{string}	string	"Bad Request"
//	@failure		401			{string}	string	"Unauthorized"
//	@failure		403			{string}	string	"Forbidden"
//	@failure		422			{string}	string	"Unprocessable Entity: creating user failed"
//	@failure		500			{string}	string	"Internal Server Error"
//	@security		ApiKeyAuth
//	@router			/users/ [post]
func (api *RestApi) createUser(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	rw.Header().Set("Content-Type", "text/plain")
	me := repository.GetUserFromContext(r.Context())
	if !me.HasRole(schema.RoleAdmin) {
		http.Error(rw, "Only admins are allowed to create new users", http.StatusForbidden)
		return
	}

	username, password, role, name, email, project := r.FormValue("username"),
		r.FormValue("password"), r.FormValue("role"), r.FormValue("name"),
		r.FormValue("email"), r.FormValue("project")

	if len(password) == 0 && role != schema.GetRoleString(schema.RoleApi) {
		http.Error(rw, "Only API users are allowed to have a blank password (login will be impossible)", http.StatusBadRequest)
		return
	}

	if len(project) != 0 && role != schema.GetRoleString(schema.RoleManager) {
		http.Error(rw, "only managers require a project (can be changed later)",
			http.StatusBadRequest)
		return
	} else if len(project) == 0 && role == schema.GetRoleString(schema.RoleManager) {
		http.Error(rw, "managers require a project to manage (can be changed later)",
			http.StatusBadRequest)
		return
	}

	if err := repository.GetUserRepository().AddUser(&schema.User{
		Username: username,
		Name:     name,
		Password: password,
		Email:    email,
		Projects: []string{project},
		Roles:    []string{role},
	}); err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	fmt.Fprintf(rw, "User %v successfully created!\n", username)
}

// deleteUser godoc
//
//	@summary		Deletes a user
//	@tags			User
//	@description	User defined by username in form data will be deleted from database.
//	@description	Only accessible from IPs registered with apiAllowedIPs configuration option.
//	@accept			mpfd
//	@produce		plain
//	@param			username	formData	string	true	"User ID to delete"
//	@success		200			"User deleted successfully"
//	@failure		400			{string}	string	"Bad Request"
//	@failure		401			{string}	string	"Unauthorized"
//	@failure		403			{string}	string	"Forbidden"
//	@failure		422			{string}	string	"Unprocessable Entity: deleting user failed"
//	@failure		500			{string}	string	"Internal Server Error"
//	@security		ApiKeyAuth
//	@router			/users/ [delete]
func (api *RestApi) deleteUser(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	if user := repository.GetUserFromContext(r.Context()); !user.HasRole(schema.RoleAdmin) {
		http.Error(rw, "Only admins are allowed to delete a user", http.StatusForbidden)
		return
	}

	username := r.FormValue("username")
	if err := repository.GetUserRepository().DelUser(username); err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

// getUsers godoc
//
//	@summary		Returns a list of users
//	@tags			User
//	@description	Returns a JSON-encoded list of users.
//	@description	Required query-parameter defines if all users or only users with additional special roles are returned.
//	@description	Only accessible from IPs registered with apiAllowedIPs configuration option.
//	@produce		json
//	@param			not-just-user	query		bool				true	"If returned list should contain all users or only users with additional special roles"
//	@success		200				{array}		api.ApiReturnedUser	"List of users returned successfully"
//	@failure		400				{string}	string				"Bad Request"
//	@failure		401				{string}	string				"Unauthorized"
//	@failure		403				{string}	string				"Forbidden"
//	@failure		500				{string}	string				"Internal Server Error"
//	@security		ApiKeyAuth
//	@router			/users/ [get]
func (api *RestApi) getUsers(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	if user := repository.GetUserFromContext(r.Context()); !user.HasRole(schema.RoleAdmin) {
		http.Error(rw, "Only admins are allowed to fetch a list of users", http.StatusForbidden)
		return
	}

	users, err := repository.GetUserRepository().ListUsers(r.URL.Query().Get("not-just-user") == "true")
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(users)
}

// updateUser godoc
//
//	@summary		Updates an existing user
//	@tags			User
//	@description	Modifies user defined by username (id) in one of four possible ways.
//	@description	If more than one formValue is set then only the highest priority field is used.
//	@description	Only accessible from IPs registered with apiAllowedIPs configuration option.
//	@accept			mpfd
//	@produce		plain
//	@param			id				path		string	true	"Database ID of User"
//	@param			add-role		formData	string	false	"Priority 1: Role to add"		Enums(admin, support, manager, user, api)
//	@param			remove-role		formData	string	false	"Priority 2: Role to remove"	Enums(admin, support, manager, user, api)
//	@param			add-project		formData	string	false	"Priority 3: Project to add"
//	@param			remove-project	formData	string	false	"Priority 4: Project to remove"
//	@success		200				{string}	string	"Success Response Message"
//	@failure		400				{string}	string	"Bad Request"
//	@failure		401				{string}	string	"Unauthorized"
//	@failure		403				{string}	string	"Forbidden"
//	@failure		422				{string}	string	"Unprocessable Entity: The user could not be updated"
//	@failure		500				{string}	string	"Internal Server Error"
//	@security		ApiKeyAuth
//	@router			/user/{id} [post]
func (api *RestApi) updateUser(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	if user := repository.GetUserFromContext(r.Context()); !user.HasRole(schema.RoleAdmin) {
		http.Error(rw, "Only admins are allowed to update a user", http.StatusForbidden)
		return
	}

	// Get Values
	newrole := r.FormValue("add-role")
	delrole := r.FormValue("remove-role")
	newproj := r.FormValue("add-project")
	delproj := r.FormValue("remove-project")

	// TODO: Handle anything but roles...
	if newrole != "" {
		if err := repository.GetUserRepository().AddRole(r.Context(), mux.Vars(r)["id"], newrole); err != nil {
			http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		rw.Write([]byte("Add Role Success"))
	} else if delrole != "" {
		if err := repository.GetUserRepository().RemoveRole(r.Context(), mux.Vars(r)["id"], delrole); err != nil {
			http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		rw.Write([]byte("Remove Role Success"))
	} else if newproj != "" {
		if err := repository.GetUserRepository().AddProject(r.Context(), mux.Vars(r)["id"], newproj); err != nil {
			http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		rw.Write([]byte("Add Project Success"))
	} else if delproj != "" {
		if err := repository.GetUserRepository().RemoveProject(r.Context(), mux.Vars(r)["id"], delproj); err != nil {
			http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		rw.Write([]byte("Remove Project Success"))
	} else {
		http.Error(rw, "Not Add or Del [role|project]?", http.StatusInternalServerError)
	}
}

func (api *RestApi) getJWT(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	rw.Header().Set("Content-Type", "text/plain")
	username := r.FormValue("username")
	me := repository.GetUserFromContext(r.Context())
	if !me.HasRole(schema.RoleAdmin) {
		if username != me.Username {
			http.Error(rw, "Only admins are allowed to sign JWTs not for themselves",
				http.StatusForbidden)
			return
		}
	}

	user, err := repository.GetUserRepository().GetUser(username)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	jwt, err := api.Authentication.JwtAuth.ProvideJWT(user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(jwt))
}

func (api *RestApi) getRoles(rw http.ResponseWriter, r *http.Request) {
	err := securedCheck(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusForbidden)
		return
	}

	user := repository.GetUserFromContext(r.Context())
	if !user.HasRole(schema.RoleAdmin) {
		http.Error(rw, "only admins are allowed to fetch a list of roles", http.StatusForbidden)
		return
	}

	roles, err := schema.GetValidRoles(user)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(roles)
}

func (api *RestApi) updateConfiguration(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/plain")
	key, value := r.FormValue("key"), r.FormValue("value")

	fmt.Printf("REST > KEY: %#v\nVALUE: %#v\n", key, value)

	if err := repository.GetUserCfgRepo().UpdateConfig(key, value, repository.GetUserFromContext(r.Context())); err != nil {
		http.Error(rw, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	rw.Write([]byte("success"))
}

func (api *RestApi) putMachineState(rw http.ResponseWriter, r *http.Request) {
	if api.MachineStateDir == "" {
		http.Error(rw, "REST > machine state not enabled", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	cluster := vars["cluster"]
	host := vars["host"]
	dir := filepath.Join(api.MachineStateDir, cluster)
	if err := os.MkdirAll(dir, 0755); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	filename := filepath.Join(dir, fmt.Sprintf("%s.json", host))
	f, err := os.Create(filename)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	if _, err := io.Copy(f, r.Body); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (api *RestApi) getMachineState(rw http.ResponseWriter, r *http.Request) {
	if api.MachineStateDir == "" {
		http.Error(rw, "REST > machine state not enabled", http.StatusNotFound)
		return
	}

	vars := mux.Vars(r)
	filename := filepath.Join(api.MachineStateDir, vars["cluster"], fmt.Sprintf("%s.json", vars["host"]))

	// Sets the content-type and 'Last-Modified' Header and so on automatically
	http.ServeFile(rw, r, filename)
}
