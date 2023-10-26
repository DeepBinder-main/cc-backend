// Copyright (C) 2023 NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scheduler

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ClusterCockpit/cc-backend/pkg/log"
)

type MetricConfig struct {
	Name string `json:"name"`
	Unit struct {
		Base string `json:"base"`
	} `json:"unit"`
	Scope       string  `json:"scope"`
	Aggregation string  `json:"aggregation"`
	Timestep    int     `json:"timestep"`
	Peak        float64 `json:"peak"`
	Normal      float64 `json:"normal"`
	Caution     float64 `json:"caution"`
	Alert       float64 `json:"alert"`
}
type SubCluster struct {
	Name           string `json:"name"`
	Nodes          string `json:"nodes"`
	ProcessorType  string `json:"processorType"`
	SocketsPerNode int    `json:"socketsPerNode"`
	CoresPerSocket int    `json:"coresPerSocket"`
	ThreadsPerCore int    `json:"threadsPerCore"`
	FlopRateScalar struct {
		Unit struct {
			Base   string `json:"base"`
			Prefix string `json:"prefix"`
		} `json:"unit"`
		Value float64 `json:"value"`
	} `json:"flopRateScalar"`
	FlopRateSimd struct {
		Unit struct {
			Base   string `json:"base"`
			Prefix string `json:"prefix"`
		} `json:"unit"`
		Value float64 `json:"value"`
	} `json:"flopRateSimd"`
	MemoryBandwidth struct {
		Unit struct {
			Base   string `json:"base"`
			Prefix string `json:"prefix"`
		} `json:"unit"`
		Value float64 `json:"value"`
	} `json:"memoryBandwidth"`
	Topology struct {
		Node         []int   `json:"node"`
		Socket       [][]int `json:"socket"`
		MemoryDomain [][]int `json:"memoryDomain"`
		Core         [][]int `json:"core"`
		Accelerators []struct {
			ID    string `json:"id"`
			Type  string `json:"type"`
			Model string `json:"model"`
		} `json:"accelerators"`
	} `json:"topology"`
}

type ClusterConfig struct {
	Name         string         `json:"name"`
	MetricConfig []MetricConfig `json:"metricConfig"`
	SubClusters  []SubCluster   `json:"subClusters"`
}

type Metadata struct {
	Plugin struct {
		Type string `json:"type"`
		Name string `json:"name"`
	} `json:"plugin"`
	Slurm struct {
		Version struct {
			Major int `json:"major"`
			Micro int `json:"micro"`
			Minor int `json:"minor"`
		} `json:"version"`
		Release string `json:"release"`
	} `json:"Slurm"`
}

type JobResource struct {
	Nodes          string          `json:"nodes"`
	AllocatedCores int             `json:"allocated_cores"`
	AllocatedHosts int             `json:"allocated_hosts"`
	AllocatedNodes []AllocatedNode `json:"allocated_nodes"`
}

type AllocatedNode struct {
	Sockets         map[string]Socket `json:"sockets"`
	Nodename        string            `json:"nodename"`
	CPUsUsed        *int              `json:"cpus_used"`
	MemoryUsed      *int              `json:"memory_used"`
	MemoryAllocated *int              `json:"memory_allocated"`
}

type Socket struct {
	Cores map[string]string `json:"cores"`
}

type Job struct {
	Account                  string      `json:"account"`
	AccrueTime               int         `json:"accrue_time"`
	AdminComment             string      `json:"admin_comment"`
	ArrayJobID               int64       `json:"array_job_id"`
	ArrayTaskID              interface{} `json:"array_task_id"`
	ArrayMaxTasks            int         `json:"array_max_tasks"`
	ArrayTaskString          string      `json:"array_task_string"`
	AssociationID            int         `json:"association_id"`
	BatchFeatures            string      `json:"batch_features"`
	BatchFlag                bool        `json:"batch_flag"`
	BatchHost                string      `json:"batch_host"`
	Flags                    []string    `json:"flags"`
	BurstBuffer              string      `json:"burst_buffer"`
	BurstBufferState         string      `json:"burst_buffer_state"`
	Cluster                  string      `json:"cluster"`
	ClusterFeatures          string      `json:"cluster_features"`
	Command                  string      `json:"command"`
	Comment                  string      `json:"comment"`
	Container                string      `json:"container"`
	Contiguous               bool        `json:"contiguous"`
	CoreSpec                 interface{} `json:"core_spec"`
	ThreadSpec               interface{} `json:"thread_spec"`
	CoresPerSocket           interface{} `json:"cores_per_socket"`
	BillableTres             interface{} `json:"billable_tres"`
	CPUPerTask               interface{} `json:"cpus_per_task"`
	CPUFrequencyMinimum      interface{} `json:"cpu_frequency_minimum"`
	CPUFrequencyMaximum      interface{} `json:"cpu_frequency_maximum"`
	CPUFrequencyGovernor     interface{} `json:"cpu_frequency_governor"`
	CPUPerTres               string      `json:"cpus_per_tres"`
	Deadline                 int         `json:"deadline"`
	DelayBoot                int         `json:"delay_boot"`
	Dependency               string      `json:"dependency"`
	DerivedExitCode          int         `json:"derived_exit_code"`
	EligibleTime             int         `json:"eligible_time"`
	EndTime                  int64       `json:"end_time"`
	ExcludedNodes            string      `json:"excluded_nodes"`
	ExitCode                 int         `json:"exit_code"`
	Features                 string      `json:"features"`
	FederationOrigin         string      `json:"federation_origin"`
	FederationSiblingsActive string      `json:"federation_siblings_active"`
	FederationSiblingsViable string      `json:"federation_siblings_viable"`
	GresDetail               []string    `json:"gres_detail"`
	GroupID                  int         `json:"group_id"`
	GroupName                string      `json:"group_name"`
	JobID                    int64       `json:"job_id"`
	JobResources             JobResource `json:"job_resources"`
	JobState                 string      `json:"job_state"`
	LastSchedEvaluation      int         `json:"last_sched_evaluation"`
	Licenses                 string      `json:"licenses"`
	MaxCPUs                  int         `json:"max_cpus"`
	MaxNodes                 int         `json:"max_nodes"`
	MCSLabel                 string      `json:"mcs_label"`
	MemoryPerTres            string      `json:"memory_per_tres"`
	Name                     string      `json:"name"`
	Nodes                    string      `json:"nodes"`
	Nice                     interface{} `json:"nice"`
	TasksPerCore             interface{} `json:"tasks_per_core"`
	TasksPerNode             int         `json:"tasks_per_node"`
	TasksPerSocket           interface{} `json:"tasks_per_socket"`
	TasksPerBoard            int         `json:"tasks_per_board"`
	CPUs                     int32       `json:"cpus"`
	NodeCount                int32       `json:"node_count"`
	Tasks                    int         `json:"tasks"`
	HETJobID                 int         `json:"het_job_id"`
	HETJobIDSet              string      `json:"het_job_id_set"`
	HETJobOffset             int         `json:"het_job_offset"`
	Partition                string      `json:"partition"`
	MemoryPerNode            interface{} `json:"memory_per_node"`
	MemoryPerCPU             int         `json:"memory_per_cpu"`
	MinimumCPUsPerNode       int         `json:"minimum_cpus_per_node"`
	MinimumTmpDiskPerNode    int         `json:"minimum_tmp_disk_per_node"`
	PreemptTime              int         `json:"preempt_time"`
	PreSUSTime               int         `json:"pre_sus_time"`
	Priority                 int         `json:"priority"`
	Profile                  interface{} `json:"profile"`
	QoS                      string      `json:"qos"`
	Reboot                   bool        `json:"reboot"`
	RequiredNodes            string      `json:"required_nodes"`
	Requeue                  bool        `json:"requeue"`
	ResizeTime               int         `json:"resize_time"`
	RestartCnt               int         `json:"restart_cnt"`
	ResvName                 string      `json:"resv_name"`
	Shared                   *string     `json:"shared"`
	ShowFlags                []string    `json:"show_flags"`
	SocketsPerBoard          int         `json:"sockets_per_board"`
	SocketsPerNode           interface{} `json:"sockets_per_node"`
	StartTime                int64       `json:"start_time"`
	StateDescription         string      `json:"state_description"`
	StateReason              string      `json:"state_reason"`
	StandardError            string      `json:"standard_error"`
	StandardInput            string      `json:"standard_input"`
	StandardOutput           string      `json:"standard_output"`
	SubmitTime               int         `json:"submit_time"`
	SuspendTime              int         `json:"suspend_time"`
	SystemComment            string      `json:"system_comment"`
	TimeLimit                int         `json:"time_limit"`
	TimeMinimum              int         `json:"time_minimum"`
	ThreadsPerCore           interface{} `json:"threads_per_core"`
	TresBind                 string      `json:"tres_bind"`
	TresFreq                 string      `json:"tres_freq"`
	TresPerJob               string      `json:"tres_per_job"`
	TresPerNode              string      `json:"tres_per_node"`
	TresPerSocket            string      `json:"tres_per_socket"`
	TresPerTask              string      `json:"tres_per_task"`
	TresReqStr               string      `json:"tres_req_str"`
	TresAllocStr             string      `json:"tres_alloc_str"`
	UserID                   int         `json:"user_id"`
	UserName                 string      `json:"user_name"`
	Wckey                    string      `json:"wckey"`
	CurrentWorkingDirectory  string      `json:"current_working_directory"`
}

type SlurmPayload struct {
	Meta   Metadata      `json:"meta"`
	Errors []interface{} `json:"errors"`
	Jobs   []Job         `json:"jobs"`
}

type DumpedComment struct {
	Administrator interface{} `json:"administrator"`
	Job           interface{} `json:"job"`
	System        interface{} `json:"system"`
}

type MaxLimits struct {
	Running struct {
		Tasks int `json:"tasks"`
	} `json:"max"`
}

type ArrayInfo struct {
	JobID  int         `json:"job_id"`
	Limits MaxLimits   `json:"limits"`
	Task   interface{} `json:"task"`
	TaskID interface{} `json:"task_id"`
}

type Association struct {
	Account   string      `json:"account"`
	Cluster   string      `json:"cluster"`
	Partition interface{} `json:"partition"`
	User      string      `json:"user"`
}

type TimeInfo struct {
	Elapsed    int64 `json:"elapsed"`
	Eligible   int64 `json:"eligible"`
	End        int64 `json:"end"`
	Start      int64 `json:"start"`
	Submission int64 `json:"submission"`
	Suspended  int64 `json:"suspended"`
	System     struct {
		Seconds      int `json:"seconds"`
		Microseconds int `json:"microseconds"`
	} `json:"system"`
	Limit int `json:"limit"`
	Total struct {
		Seconds      int `json:"seconds"`
		Microseconds int `json:"microseconds"`
	} `json:"total"`
	User struct {
		Seconds      int `json:"seconds"`
		Microseconds int `json:"microseconds"`
	} `json:"user"`
}

type ExitCode struct {
	Status     string `json:"status"`
	ReturnCode int    `json:"return_code"`
}

type DumpedJob struct {
	Account         string        `json:"account"`
	Comment         DumpedComment `json:"comment"`
	AllocationNodes int           `json:"allocation_nodes"`
	Array           ArrayInfo     `json:"array"`
	Association     Association   `json:"association"`
	Cluster         string        `json:"cluster"`
	Constraints     string        `json:"constraints"`
	Container       interface{}   `json:"container"`
	DerivedExitCode ExitCode      `json:"derived_exit_code"`
	Time            TimeInfo      `json:"time"`
	ExitCode        ExitCode      `json:"exit_code"`
	Flags           []string      `json:"flags"`
	Group           string        `json:"group"`
	Het             struct {
		JobID     int         `json:"job_id"`
		JobOffset interface{} `json:"job_offset"`
	} `json:"het"`
	JobID int64  `json:"job_id"`
	Name  string `json:"name"`
	MCS   struct {
		Label string `json:"label"`
	} `json:"mcs"`
	Nodes     string `json:"nodes"`
	Partition string `json:"partition"`
	Priority  int    `json:"priority"`
	QoS       string `json:"qos"`
	Required  struct {
		CPUs   int `json:"CPUs"`
		Memory int `json:"memory"`
	} `json:"required"`
	KillRequestUser interface{} `json:"kill_request_user"`
	Reservation     struct {
		ID   int `json:"id"`
		Name int `json:"name"`
	} `json:"reservation"`
	State struct {
		Current string `json:"current"`
		Reason  string `json:"reason"`
	} `json:"state"`
	Steps []struct {
		Nodes struct {
			List  []string `json:"list"`
			Count int      `json:"count"`
			Range string   `json:"range"`
		} `json:"nodes"`
		Tres struct {
			Requested struct {
				Max     []interface{} `json:"max"`
				Min     []interface{} `json:"min"`
				Average []interface{} `json:"average"`
				Total   []interface{} `json:"total"`
			} `json:"requested"`
			Consumed struct {
				Max     []interface{} `json:"max"`
				Min     []interface{} `json:"min"`
				Average []interface{} `json:"average"`
				Total   []interface{} `json:"total"`
			} `json:"consumed"`
			Allocated []struct {
				Type  string      `json:"type"`
				Name  interface{} `json:"name"`
				ID    int         `json:"id"`
				Count int         `json:"count"`
			} `json:"allocated"`
		} `json:"tres"`
		Time     TimeInfo `json:"time"`
		ExitCode ExitCode `json:"exit_code"`
		Tasks    struct {
			Count int `json:"count"`
		} `json:"tasks"`
		PID interface{} `json:"pid"`
		CPU struct {
			RequestedFrequency struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"requested_frequency"`
			Governor []interface{} `json:"governor"`
		} `json:"CPU"`
		KillRequestUser interface{} `json:"kill_request_user"`
		State           string      `json:"state"`
		Statistics      struct {
			CPU struct {
				ActualFrequency int `json:"actual_frequency"`
			} `json:"CPU"`
			Energy struct {
				Consumed int `json:"consumed"`
			} `json:"energy"`
		} `json:"statistics"`
		Step struct {
			JobID int `json:"job_id"`
			Het   struct {
				Component interface{} `json:"component"`
			} `json:"het"`
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"step"`
		Task struct {
			Distribution string `json:"distribution"`
		} `json:"task"`
	} `json:"steps"`
	Tres struct {
		Allocated []struct {
			Type  string      `json:"type"`
			Name  interface{} `json:"name"`
			ID    int         `json:"id"`
			Count int         `json:"count"`
		} `json:"allocated"`
		Requested []struct {
			Type  string      `json:"type"`
			Name  interface{} `json:"name"`
			ID    int         `json:"id"`
			Count int         `json:"count"`
		} `json:"requested"`
	} `json:"tres"`
	User  string `json:"user"`
	Wckey struct {
		Wckey string   `json:"wckey"`
		Flags []string `json:"flags"`
	} `json:"wckey"`
	WorkingDirectory string `json:"working_directory"`
}

type SlurmDBPayload struct {
	Meta   Metadata    `json:"meta"`
	Errors []string    `json:"errors"`
	Jobs   []DumpedJob `json:"jobs"`
}

func DecodeClusterConfig(filename string) (ClusterConfig, error) {
	var clusterConfig ClusterConfig

	file, err := os.Open(filename)
	if err != nil {
		log.Errorf("Cluster config file not found. No cores/GPU ids available.")
		return clusterConfig, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&clusterConfig)
	if err != nil {
		log.Errorf("Error decoding cluster config file: %v", err)
	}

	log.Printf("Name: %s\n", clusterConfig.Name)
	log.Printf("MetricConfig:\n")
	for _, metric := range clusterConfig.MetricConfig {
		log.Printf("  Name: %s\n", metric.Name)
		log.Printf("  Unit Base: %s\n", metric.Unit.Base)
		log.Printf("  Scope: %s\n", metric.Scope)
		log.Printf("  Aggregation: %s\n", metric.Aggregation)
		log.Printf("  Timestep: %d\n", metric.Timestep)
		log.Printf("  Peak: %f\n", metric.Peak)
		log.Printf("  Normal: %f\n", metric.Normal)
		log.Printf("  Caution: %f\n", metric.Caution)
		log.Printf("  Alert: %f\n", metric.Alert)
	}
	log.Printf("SubClusters:\n")
	for _, subCluster := range clusterConfig.SubClusters {
		log.Printf("  Name: %s\n", subCluster.Name)
		log.Printf("  Nodes: %s\n", subCluster.Nodes)
		log.Printf("  Processor Type: %s\n", subCluster.ProcessorType)
		log.Printf("  Sockets Per Node: %d\n", subCluster.SocketsPerNode)
		log.Printf("  Cores Per Socket: %d\n", subCluster.CoresPerSocket)
		log.Printf("  Threads Per Core: %d\n", subCluster.ThreadsPerCore)
		log.Printf("  Flop Rate Scalar Unit Base: %s\n", subCluster.FlopRateScalar.Unit.Base)
		log.Printf("  Flop Rate Scalar Unit Prefix: %s\n", subCluster.FlopRateScalar.Unit.Prefix)
		log.Printf("  Flop Rate Scalar Value: %f\n", subCluster.FlopRateScalar.Value)
		log.Printf("  Flop Rate Simd Unit Base: %s\n", subCluster.FlopRateSimd.Unit.Base)
		log.Printf("  Flop Rate Simd Unit Prefix: %s\n", subCluster.FlopRateSimd.Unit.Prefix)
		log.Printf("  Flop Rate Simd Value: %f\n", subCluster.FlopRateSimd.Value)
		log.Printf("  Memory Bandwidth Unit Base: %s\n", subCluster.MemoryBandwidth.Unit.Base)
		log.Printf("  Memory Bandwidth Unit Prefix: %s\n", subCluster.MemoryBandwidth.Unit.Prefix)
		log.Printf("  Memory Bandwidth Value: %f\n", subCluster.MemoryBandwidth.Value)
		log.Printf("  Topology Node: %v\n", subCluster.Topology.Node)
		log.Printf("  Topology Socket: %v\n", subCluster.Topology.Socket)
		log.Printf("  Topology Memory Domain: %v\n", subCluster.Topology.MemoryDomain)
		log.Printf("  Topology Core: %v\n", subCluster.Topology.Core)
		log.Printf("  Topology Accelerators:\n")
		for _, accelerator := range subCluster.Topology.Accelerators {
			log.Printf("    ID: %s\n", accelerator.ID)
			log.Printf("    Type: %s\n", accelerator.Type)
			log.Printf("    Model: %s\n", accelerator.Model)
		}
	}

	return clusterConfig, nil
}

func UnmarshalSlurmPayload(jsonPayload string) (SlurmPayload, error) {
	var slurmData SlurmPayload
	err := json.Unmarshal([]byte(jsonPayload), &slurmData)
	if err != nil {
		return slurmData, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}
	return slurmData, nil
}